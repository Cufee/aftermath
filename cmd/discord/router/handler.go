package router

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"slices"
	"time"

	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/retry"
)

var supportedInteractionTypes = []discordgo.InteractionType{discordgo.InteractionApplicationCommand, discordgo.InteractionMessageComponent, discordgo.InteractionApplicationCommandAutocomplete, discordgo.InteractionModalSubmit}

// https://github.com/bsdlp/discord-interactions-go/blob/main/interactions/verify.go
func verifyPublicKey(r *http.Request, key ed25519.PublicKey) bool {
	var msg bytes.Buffer

	signature := r.Header.Get("X-Signature-Ed25519")
	if signature == "" {
		return false
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	if len(sig) != ed25519.SignatureSize || sig[63]&224 != 0 {
		return false
	}

	timestamp := r.Header.Get("X-Signature-Timestamp")
	if timestamp == "" {
		return false
	}

	msg.WriteString(timestamp)

	defer r.Body.Close()
	var body bytes.Buffer

	// at the end of the function, copy the original body back into the request
	defer func() {
		r.Body = io.NopCloser(&body)
	}()

	// copy body into buffers
	_, err = io.Copy(&msg, io.TeeReader(r.Body, &body))
	if err != nil {
		return false
	}

	return ed25519.Verify(key, msg.Bytes(), sig)
}

/*
Returns a handler for the current router
*/
func (router *Router) HTTPHandler() (http.HandlerFunc, error) {
	if router.publicKey == "" {
		return nil, errors.New("missing publicKey")
	}
	hexDecodedKey, err := hex.DecodeString(router.publicKey)
	if err != nil {
		return nil, errors.New("invalid public key")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if ok := verifyPublicKey(r, hexDecodedKey); !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Debug().Msg("invalid request received on callback handler")
			return
		}

		defer r.Body.Close()
		log.Debug().Msg("received a discord callback")

		var data discordgo.Interaction
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Err(err).Msg("failed to decode callback body")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if data.Type == discordgo.InteractionPing {
			sendPingReply(w)
			return
		}

		cmd, err := router.routeInteraction(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Err(err).Str("id", data.ID).Msg("failed to route interaction")
			return
		}

		// ack the interaction proactively
		payload, needsAck, err := deferredInteractionResponsePayload(data.Type, cmd.Ephemeral)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Err(err).Str("id", data.ID).Msg("failed to make an ack response payload")
			return
		}
		if !needsAck {
			router.handleInteraction(data, cmd) // we wait for the handler to reply
			return
		}

		res := retry.Retry(
			func() (struct{}, error) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
				defer cancel()
				err := router.restClient.SendInteractionResponse(ctx, data.ID, data.Token, payload, nil)
				return struct{}{}, err
			},
			3,
			time.Millisecond*50,
			// break if the error means we were able to ack on the last request
			func(err error) bool { return errors.Is(err, rest.ErrInteractionAlreadyAcked) })
		if res.Err != nil && !errors.Is(res.Err, rest.ErrInteractionAlreadyAcked) {
			log.Err(res.Err).Str("id", data.ID).Msg("failed to ack an interaction")
			// cross our fingers and hope discord registered one of those requests, or will propagate the ack from the response body
		}

		// write the ack into response, this typically does nothing/is flaky
		body, err := json.Marshal(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Err(err).Str("id", data.ID).Msg("failed to encode interaction ack payload")
			return
		}
		_, err = w.Write(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Err(err).Str("id", data.ID).Msg("failed to write interaction ack payload")
			return
		}

		// run the interaction handler
		go router.handleInteraction(data, cmd)
	}, nil
}

func (r *Router) routeInteraction(interaction discordgo.Interaction) (builder.Command, error) {
	var matchKey string

	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		data, _ := interaction.Data.(discordgo.ApplicationCommandInteractionData)
		matchKey = data.Name
	case discordgo.InteractionMessageComponent:
		data, _ := interaction.Data.(discordgo.MessageComponentInteractionData)
		matchKey = data.CustomID
	case discordgo.InteractionApplicationCommandAutocomplete:
		data, _ := interaction.Data.(discordgo.ApplicationCommandInteractionData)
		matchKey = "autocomplete_" + data.Name
	}

	if matchKey == "" {
		return builder.Command{}, errors.New("match key not found")
	}

	for _, cmd := range r.commands {
		if cmd.Match(matchKey) {
			return cmd, nil
		}
	}
	return builder.Command{}, fmt.Errorf("failed to match %s to a command handler", matchKey)
}

func (r *Router) handleInteraction(interaction discordgo.Interaction, command builder.Command) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cCtx, err := common.NewContext(ctx, interaction, r.restClient, r.core)
	if err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		log.Err(err).Msg("failed to create a common.Context for a handler")
		r.sendInteractionReply(ctx, interaction, discordgo.InteractionResponseData{Content: "Something unexpected happened and your command failed. Please try again in a few seconds."})
		return
	}

	handler := command.Handler
	for i := len(command.Middleware) - 1; i >= 0; i-- {
		handler = command.Middleware[i](cCtx, handler)
	}

	func() {
		defer func() {
			if rec := recover(); rec != nil {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				log.Error().Str("stack", string(debug.Stack())).Msg("panic in interaction handler")
				r.sendInteractionReply(ctx, interaction, discordgo.InteractionResponseData{Content: cCtx.Localize("common_error_unhandled_not_reported")})
			}
		}()
		err = handler(cCtx)
		if err != nil {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			log.Err(err).Msg("handler returned an error")
			r.sendInteractionReply(ctx, interaction, discordgo.InteractionResponseData{Content: cCtx.Localize("common_error_unhandled_not_reported")})
			return
		}
	}()
}

func sendPingReply(w http.ResponseWriter) {
	_, err := w.Write([]byte(fmt.Sprintf(`{"type": %d}`, discordgo.InteractionPing)))
	if err != nil {
		log.Err(err).Msg("failed to reply to a discord PING")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (r *Router) sendInteractionReply(ctx context.Context, interaction discordgo.Interaction, data discordgo.InteractionResponseData) {
	var handler func() error

	if slices.Contains(supportedInteractionTypes, interaction.Type) {
		handler = func() error {
			return r.restClient.UpdateOrSendInteractionResponse(ctx, interaction.AppID, interaction.ID, interaction.Token, discordgo.InteractionResponse{Data: &data, Type: discordgo.InteractionResponseChannelMessageWithSource}, nil)
		}
	} else {
		log.Error().Stack().Any("data", data).Str("id", interaction.ID).Msg("unknown interaction type received")
		handler = func() error {
			return r.restClient.UpdateOrSendInteractionResponse(ctx, interaction.AppID, interaction.ID, interaction.Token, discordgo.InteractionResponse{Data: &discordgo.InteractionResponseData{Content: "Something unexpected happened and your command failed."}, Type: discordgo.InteractionResponseChannelMessageWithSource}, nil)
		}
	}

	err := handler()
	if err != nil {
		log.Err(err).Stack().Any("data", data).Str("id", interaction.ID).Msg("failed to send an interaction response")
		return
	}
}

func deferredInteractionResponsePayload(t discordgo.InteractionType, ephemeral bool) (discordgo.InteractionResponse, bool, error) {
	var response discordgo.InteractionResponse
	if ephemeral {
		if response.Data == nil {
			response.Data = &discordgo.InteractionResponseData{}
		}
		response.Data.Flags = discordgo.MessageFlagsEphemeral
	}

	switch t {
	case discordgo.InteractionApplicationCommand, discordgo.InteractionMessageComponent, discordgo.InteractionModalSubmit:
		response.Type = discordgo.InteractionResponseDeferredChannelMessageWithSource
	case discordgo.InteractionApplicationCommandAutocomplete:
		return response, false, nil

	default:
		return response, false, fmt.Errorf("interaction type %s not supported", t.String())
	}

	return response, true, nil
}
