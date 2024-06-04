package router

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

// we need to respond within 15 minutes of the initial interaction
var interactionTimeout = time.Minute*14 + time.Second*45

type interactionState struct {
	lock sync.RWMutex

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

	// handler := httpserver.HandleInteraction(hexDecodedKey, client.Logger(), handlers.DefaultHTTPServerEventHandlerFunc(client))

	return func(w http.ResponseWriter, r *http.Request) {
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

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*4)
		defer cancel()

		w.Header().Set("Content-Type", "application/json")
		err = router.routeInteraction(ctx, w, data)
		if err != nil {
			log.Err(err).Any("body", data).Msg("failed to route an interaction")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}, nil
}

func (router *Router) routeInteraction(_ context.Context, w http.ResponseWriter, interaction discordgo.Interaction) error {
	if interaction.Type == discordgo.InteractionPing {
		_, err := w.Write([]byte(fmt.Sprintf(`{"type": %d}`, discordgo.InteractionPing)))
		return err
	}

	responseCh := make(chan discordgo.InteractionResponseData, 1)
	doneCh := make(chan struct{}, 1)
	var state interactionState

	state.lock.Lock()
	defer func() {
		state.acked = true
		state.lock.Unlock()
	}()

	go func() {
		defer close(doneCh)
		defer close(responseCh)

		// create a timer for the interaction response
		responseTimer := time.NewTimer(interactionTimeout)
		defer responseTimer.Stop()

		// create a new context and cancel it on reply
		workerCtx, cancelWorker := context.WithCancel(context.Background())
		defer cancelWorker()

		// handle the interaction
		go router.handleInteraction(workerCtx, interaction, responseCh, doneCh)

		for {
			select {
			case <-responseTimer.C:
				state.lock.Lock()
				// discord will handle the user-facing error, so we just log it
				log.Error().Any("interaction", interaction).Msg("interaction update timed out")
				defer state.lock.Unlock()
				return

			case <-doneCh:
				// we are done, there is nothing else we should do here
				state.lock.Lock()
				// wait in case responseCh is still busy sending some data over
				defer state.lock.Unlock()
				return

			case data := <-responseCh:
				state.lock.Lock()
				router.sendInteractionReply(interaction, &state, data)
				state.lock.Unlock()

			}
		}
	}()

	return sendDeferredInteractionResponse(w, interaction.Type)
}

func (r *Router) sendInteractionReply(interaction discordgo.Interaction, state *interactionState, data discordgo.InteractionResponseData) {
	payload := discordgo.InteractionResponse{Data: &data}

	var err error
	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		if state.replied {
			payload.Type = discordgo.InteractionResponseUpdateMessage
			err = r.restClient.UpdateInteractionMessage(interaction.Token, payload)
		} else if state.acked {
			payload.Type = discordgo.InteractionResponseChannelMessageWithSource
			err = r.restClient.SendInteractionResponse(interaction.ID, interaction.Token, payload)
		} else {
			payload.Type = discordgo.InteractionResponseUpdateMessage
			err = r.restClient.SendInteractionResponse(interaction.ID, interaction.Token, payload)
		}
	}
	if err != nil {
		log.Err(err).Msg("failed to send an interaction response")
	}
}

func sendDeferredInteractionResponse(w http.ResponseWriter, t discordgo.InteractionType) error {
	switch t {
	case discordgo.InteractionApplicationCommand:
		_, err := w.Write([]byte(fmt.Sprintf(`{"type": %d}`, discordgo.InteractionResponseDeferredChannelMessageWithSource)))
		return err
	case discordgo.InteractionMessageComponent:
		_, err := w.Write([]byte(fmt.Sprintf(`{"type": %d}`, discordgo.InteractionResponseDeferredMessageUpdate)))
		return err
	default:
		log.Error().Str("type", t.String()).Msg("received an unsupported interaction type")
		http.Error(w, "unsupported interaction type", http.StatusInternalServerError)
		return nil
	}
}
