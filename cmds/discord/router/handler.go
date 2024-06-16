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
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/rs/zerolog/log"
)

type interactionState struct {
	mx      *sync.Mutex
	acked   bool
	replied bool
}

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

		command, err := router.routeInteraction(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Err(err).Str("id", data.ID).Msg("failed to route interaction")
			return
		}

		// ack the interaction
		err = writeDeferredInteractionResponseAck(w, data.Type, command.Ephemeral)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Err(err).Str("id", data.ID).Msg("failed to ack an interaction")
			return
		}

		// route the interaction to a proper handler for a later reply
		state := &interactionState{mx: &sync.Mutex{}}
		state.mx.Lock()
		go func() {
			// unlock once this context is done and the ack is delivered
			<-r.Context().Done()
			log.Debug().Str("id", data.ID).Msg("sent an interaction ack")

			state.acked = true
			state.mx.Unlock()
		}()

		router.startInteractionHandlerAsync(data, state, command)
	}, nil
}

func (r *Router) routeInteraction(interaction discordgo.Interaction) (*builder.Command, error) {
	var matchKey string

	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		data, _ := interaction.Data.(discordgo.ApplicationCommandInteractionData)
		matchKey = data.Name
	case discordgo.InteractionMessageComponent:
		data, _ := interaction.Data.(discordgo.MessageComponentInteractionData)
		matchKey = data.CustomID
	}

	if matchKey == "" {
		return nil, errors.New("match key not found")
	}

	for _, cmd := range r.commands {
		if cmd.Match(matchKey) {
			return &cmd, nil
		}
	}
	return nil, fmt.Errorf("failed to match %s to a command handler", matchKey)
}

func (r *Router) handleInteraction(ctx context.Context, cancel context.CancelFunc, interaction discordgo.Interaction, command *builder.Command, reply chan<- discordgo.InteractionResponseData) {
	defer cancel()

	cCtx, err := common.NewContext(ctx, interaction, reply, r.restClient, r.core)
	if err != nil {
		log.Err(err).Msg("failed to create a common.Context for a handler")
		select {
		case <-ctx.Done():
			return
		default:
			reply <- discordgo.InteractionResponseData{Content: cCtx.Localize("common_error_unhandled_not_reported")}
		}
		return
	}

	handler := command.Handler
	for i := len(command.Middleware) - 1; i >= 0; i-- {
		handler = command.Middleware[i](cCtx, handler)
	}

	err = handler(cCtx)
	if err != nil {
		log.Err(err).Msg("handler returned an error")
		select {
		case <-ctx.Done():
			return
		default:
			reply <- discordgo.InteractionResponseData{Content: cCtx.Localize("common_error_unhandled_not_reported")}
		}
		return
	}
}

func sendPingReply(w http.ResponseWriter) {
	_, err := w.Write([]byte(fmt.Sprintf(`{"type": %d}`, discordgo.InteractionPing)))
	if err != nil {
		log.Err(err).Msg("failed to reply to a discord PING")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (router *Router) startInteractionHandlerAsync(interaction discordgo.Interaction, state *interactionState, command *builder.Command) {
	log.Info().Str("id", interaction.ID).Msg("started handling an interaction")

	// create a new context and cancel it on reply
	// we need to respond within 15 minutes, but it's safe to assume something failed if we did not reply quickly
	workerCtx, cancelWorker := context.WithTimeout(context.Background(), time.Second*10)

	// handle the interaction
	responseCh := make(chan discordgo.InteractionResponseData)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Error().Stack().Any("recover", r).Msg("panic in interaction handler")
				state.mx.Lock()
				router.sendInteractionReply(interaction, state, discordgo.InteractionResponseData{Content: "Something unexpected happened and your command failed. Please try again in a few seconds."})
				state.mx.Unlock()
			}
		}()
		router.handleInteraction(workerCtx, cancelWorker, interaction, command, responseCh)
	}()

	go func() {
		defer close(responseCh)
		defer cancelWorker()

		for {
			select {
			case <-workerCtx.Done():
				state.mx.Lock()
				defer state.mx.Unlock()
				// we are done, there is nothing else we should do here
				// lock in case responseCh is still busy sending some data over
				log.Info().Str("id", interaction.ID).Msg("finished handling an interaction")
				return

			case data := <-responseCh:
				state.mx.Lock()
				log.Debug().Str("id", interaction.ID).Msg("sending an interaction reply")
				router.sendInteractionReply(interaction, state, data)
				state.mx.Unlock()
			}
		}
	}()
}

func (r *Router) sendInteractionReply(interaction discordgo.Interaction, state *interactionState, data discordgo.InteractionResponseData) {
	var handler func() error

	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		if state.replied || state.acked {
			// We already replied to this interaction - edit the message
			handler = func() error {
				return r.restClient.UpdateInteractionResponse(interaction.AppID, interaction.Token, data)
			}
		} else {
			// We never replied to this message, create a new reply
			handler = func() error {
				payload := discordgo.InteractionResponse{Data: &data, Type: discordgo.InteractionResponseChannelMessageWithSource}
				return r.restClient.SendInteractionResponse(interaction.ID, interaction.Token, payload)
			}
		}
	}

	err := handler()
	if err != nil {
		log.Err(err).Stack().Any("state", state).Any("data", data).Str("id", interaction.ID).Msg("failed to send an interaction response")
		return
	}
	state.acked = true
	state.replied = true
}

func writeDeferredInteractionResponseAck(w http.ResponseWriter, t discordgo.InteractionType, ephemeral bool) error {
	var response discordgo.InteractionResponse
	if ephemeral {
		response.Data.Flags = discordgo.MessageFlagsEphemeral
	}

	switch t {
	case discordgo.InteractionApplicationCommand:
		response.Type = discordgo.InteractionResponseDeferredChannelMessageWithSource
	case discordgo.InteractionMessageComponent:
		response.Type = discordgo.InteractionResponseDeferredMessageUpdate

	default:
		return fmt.Errorf("interaction type %s not supported", t.String())
	}

	body, _ := json.Marshal(response)
	_, err := w.Write(body)
	return err
}
