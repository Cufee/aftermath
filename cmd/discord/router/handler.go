package router

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/json"

	"github.com/pkg/errors"
	"golang.org/x/text/language"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/retry"
)

var supportedInteractionTypes = []discordgo.InteractionType{discordgo.InteractionApplicationCommand, discordgo.InteractionMessageComponent, discordgo.InteractionApplicationCommandAutocomplete, discordgo.InteractionModalSubmit}

/*
Returns a handler for the current router
*/
func (router *router) HTTPHandler() (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		if ok, _ := router.validator.Validate(r); !ok {
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

		// we validated the request and are working on it
		w.WriteHeader(http.StatusProcessing)

		cmd, err := router.routeInteraction(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Err(err).Str("id", data.ID).Msg("failed to route interaction")
			return
		}

		if !shouldAckInteraction(data.Type) {
			// discord requires a response within 3 seconds, allow 250ms on top of that to handle an error or timeout
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*2750)
			defer cancel()

			router.handleInteraction(ctx, data, cmd) // wait for the handler to finish
			return
		}

		res := retry.Retry(
			func() (struct{}, error) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
				defer cancel()

				err := router.restClient.AckInteractionResponse(ctx, data.ID, data.Token, cmd.Ephemeral)
				if errors.Is(err, rest.ErrInteractionAlreadyAcked) {
					err = nil
				}
				return struct{}{}, err
			},
			5,
			time.Millisecond*250)
		if res.Err != nil {
			log.Warn().Err(res.Err).Str("id", data.ID).Msg("failed to ack an interaction")
			// cross our fingers and hope discord registered one of those requests, or will propagate the ack from the response body
		}

		w.WriteHeader(http.StatusAccepted)

		// run the interaction handler
		go func() {
			// while the limit for a deferred reply is on a scale of minutes, we probably do not want a user to wait that long
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			router.handleInteraction(ctx, data, cmd)
		}()
	}, nil
}

func (r *router) routeInteraction(interaction discordgo.Interaction) (builder.Command, error) {
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
		key, ok := fullAutocompleteName(data.Options)
		if ok {
			matchKey = "autocomplete_" + data.Name + "_" + key
		}
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

func (r *router) handleInteraction(ctx context.Context, interaction discordgo.Interaction, command builder.Command) {
	cCtx, err := newContext(ctx, interaction, r.restClient, r.core)
	if err != nil {
		log.Err(err).Msg("failed to create a common.Context for a handler")
		r.sendInteractionReply(interaction, discordgo.InteractionResponseData{
			Content: "Something unexpected happened and your command failed. Please try again in a few seconds." + "\n-# " + interaction.ID,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						common.ButtonJoinPrimaryGuild("Need Help?"),
					}},
			},
		})
		return
	}

	defer func() {
		if rec := recover(); rec != nil {
			log.Error().Str("stack", string(debug.Stack())).Msg("panic in interaction handler")
			r.sendInteractionReply(interaction, discordgo.InteractionResponseData{
				Content: cCtx.Localize("common_error_unhandled_not_reported") + "\n-# " + interaction.ID,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							common.ButtonJoinPrimaryGuild(cCtx.Localize("buttons_join_primary_guild")),
						}},
				},
			})
		}
	}()

	handler := command.Handler
	for i := len(r.middleware) - 1; i >= 0; i-- {
		handler = r.middleware[i](cCtx, handler)
	}
	for i := len(command.Middleware) - 1; i >= 0; i-- {
		handler = command.Middleware[i](cCtx, handler)
	}

	err = handler(cCtx)
	if err != nil {
		log.Err(err).Msg("handler returned an error")
		r.sendInteractionReply(interaction, discordgo.InteractionResponseData{Content: cCtx.Localize("common_error_unhandled_not_reported") + "\n-# " + interaction.ID,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						common.ButtonJoinPrimaryGuild(cCtx.Localize("buttons_join_primary_guild")),
					}},
			},
		})
		return
	}

}

func sendPingReply(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf(`{"type": %d}`, discordgo.InteractionPing)))
	if err != nil {
		log.Err(err).Msg("failed to reply to a discord PING")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (r *router) sendInteractionReply(interaction discordgo.Interaction, data discordgo.InteractionResponseData) {
	// autocomplete interactions require a reply within 3 seconds, but since the UI error state is decent, we can just skip the reply
	if interaction.Type == discordgo.InteractionApplicationCommandAutocomplete {
		return
	}

	// if the interaction type is unknown, make a new error response
	if !slices.Contains(supportedInteractionTypes, interaction.Type) {
		printer, err := localization.NewPrinter("discord", language.English)
		if err != nil {
			log.Err(err).Msg("failed to create a new locale printer")
			printer = func(_ string) string { return "Join Aftermath Official" } // we only need it for this one button
		}

		data = discordgo.InteractionResponseData{
			Content: "Something unexpected happened and your command failed.",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						common.ButtonJoinPrimaryGuild(printer("buttons_join_primary_guild")),
					}},
			},
		}
		log.Error().Stack().Any("data", data).Str("id", interaction.ID).Msg("unknown interaction type received")
	}

	res := retry.Retry(func() (struct{}, error) {
		// use a background context for this
		// since the command is already handled, there is not much point in using the route handler context timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := r.restClient.UpdateInteractionResponse(ctx, interaction.AppID, interaction.Token, data, nil)
		return struct{}{}, err
	}, 5, time.Second) // better late than never
	if res.Err != nil {
		log.Err(res.Err).Stack().Any("data", data).Str("id", interaction.ID).Msg("failed to send an interaction response")
		return
	}
}

func shouldAckInteraction(t discordgo.InteractionType) bool {
	switch t {
	case discordgo.InteractionApplicationCommand, discordgo.InteractionMessageComponent, discordgo.InteractionModalSubmit:
		return true
	default:
		return false
	}
}

func fullAutocompleteName(options []*discordgo.ApplicationCommandInteractionDataOption) (string, bool) {
	for _, opt := range options {
		if opt.Type == discordgo.ApplicationCommandOptionSubCommandGroup || opt.Type == discordgo.ApplicationCommandOptionSubCommand {
			n, ok := fullAutocompleteName(opt.Options)
			return opt.Name + "_" + n, ok
		}
		if opt.Focused {
			return opt.Name, true
		}
	}
	return "", false
}
