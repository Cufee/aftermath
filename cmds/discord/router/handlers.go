package router

import (
	"context"

	"github.com/cufee/aftermath/cmds/discord/commands"
	"github.com/cufee/aftermath/cmds/discord/interactions"
	"github.com/cufee/aftermath/internal/database"
	"github.com/disgoorg/disgo/events"
	"github.com/rs/zerolog/log"
)

func (r *Router) commandHandler(ctx context.Context, event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()

	if handler, ok := r.commandHandlers[data.CommandName()]; ok {
		user, _ := ctx.Value("user").(database.User)
		user.ID = event.User().ID.String()

		err := handler(commands.ContextFrom(ctx, event, user))
		if err != nil {
			// TODO: handle
			_ = err
		}
		return
	}
	log.Error().Str("command", data.CommandName()).Msg("received an unexpected command")
}

func (r *Router) interactionHandler(ctx context.Context, event *events.ComponentInteractionCreate) {
	if handler, ok := r.interactionHandlers[event.Data.CustomID()]; ok {
		user, _ := ctx.Value("user").(database.User)
		user.ID = event.User().ID.String()

		err := handler(interactions.ContextFrom(ctx, event, user))
		if err != nil {
			// TODO: handle
			_ = err
		}
		return
	}
	log.Error().Str("customId", event.Data.CustomID()).Msg("received an unexpected component interaction")
}
