package gateway

import (
	"context"
	"os"
	"runtime/debug"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/log"
)

/*
*
Register a command router handler for all gateway commands
*/
func (gw *gatewayClient) RouterHandler() {
	gw.Handler(func(s *discordgo.Session, e *discordgo.InteractionCreate) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		if e == nil {
			log.Error().Msg("received a nil interaction event in gateway command router")
			return
		}
		ctx, err := newContext(c, gw, *e.Interaction)
		if err != nil {
			log.Error().Str("interaction", e.ID).Str("type", e.Type.String()).Msg("failed to create an event context")
			sendRawError(gw.rest, e.Interaction, "Something unexpected happened and your command failed.\n*This error was reported automatically, you can also reach out to our team on Aftermath Official*", false)
			return
		}

		defer func() {
			if rec := recover(); rec != nil {
				log.Error().Str("stack", string(debug.Stack())).Msg("panic in interaction handler")
				sendError(ctx, ctx.Localize("common_error_unhandled_reported")+"\n-# "+e.ID)
			}
		}()

		var matchKey string

		switch e.Type {
		case discordgo.InteractionApplicationCommand:
			data, _ := e.Data.(discordgo.ApplicationCommandInteractionData)
			matchKey = data.Name
		case discordgo.InteractionMessageComponent:
			data, _ := e.Data.(discordgo.MessageComponentInteractionData)
			matchKey = data.CustomID
		case discordgo.InteractionApplicationCommandAutocomplete:
			data, _ := e.Data.(discordgo.ApplicationCommandInteractionData)
			key, ok := fullAutocompleteName(data.Options)
			if ok {
				matchKey = "autocomplete_" + data.Name + "_" + key
			}
		}

		if matchKey == "" {
			log.Error().Str("interaction", e.ID).Str("type", e.Type.String()).Msg("match key not found for interaction")
			sendError(ctx, ctx.Localize("common_error_unhandled_reported"))
			return
		}

		for _, cmd := range gw.commands {
			if cmd.Match(matchKey) {
				if e.Type != discordgo.InteractionApplicationCommandAutocomplete {
					// Ack the interaction
					payload := discordgo.InteractionResponseData{}
					if cmd.Ephemeral {
						payload.Flags = discordgo.MessageFlagsEphemeral
					}
					err = s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
						Data: &payload,
					})
					if err != nil {
						log.Err(err).Str("interaction", e.ID).Str("type", e.Type.String()).Msg("interaction handler failed")
						sendError(ctx, ctx.Localize("common_error_unhandled_reported")+"\n-# "+e.ID)
						return
					}
				}

				// Run command handler
				err := cmd.Handler(ctx)
				if err != nil {
					log.Err(err).Str("interaction", e.ID).Str("type", e.Type.String()).Msg("interaction handler failed")
					if os.IsTimeout(err) {
						sendRawError(gw.rest, e.Interaction, ctx.Localize("common_error_unhandled_reported")+"\n-# "+e.ID, true)
						return
					}
					sendError(ctx, ctx.Localize("common_error_unhandled_reported")+"\n-# "+e.ID)
					return
				}
				return
			}
		}

		// Error
		log.Error().Str("interaction", e.ID).Str("type", e.Type.String()).Str("key", matchKey).Msg("failed to match to a command handler")
		sendError(ctx, ctx.Localize("common_error_unhandled_reported"))
	})
}

func sendError(ctx common.Context, message string) {
	err := ctx.
		Reply().
		Component(discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				common.ButtonJoinPrimaryGuild(ctx.Localize("buttons_join_primary_guild")),
			}}).
		Send(message)
	if err != nil {
		log.Err(err).Msg("failed to send an interaction response")
	}
}

func sendRawError(rest *rest.Client, interaction *discordgo.Interaction, message string, acked bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var err error
	if acked {
		_, err = rest.SendInteractionResponse(ctx, interaction.ID, interaction.Token, discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: message + "\n-# " + interaction.ID,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							common.ButtonJoinPrimaryGuild("Need Help?"),
						}},
				},
			}}, nil)
	} else {
		_, err = rest.UpdateInteractionResponse(ctx, interaction.AppID, interaction.Token, discordgo.InteractionResponseData{
			Content: message + "\n-# " + interaction.ID,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						common.ButtonJoinPrimaryGuild("Need Help?"),
					}},
			}}, nil)
	}

	if err != nil {
		log.Err(err).Msg("failed to send an interaction response")
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
