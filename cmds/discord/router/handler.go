package router

import (
	"context"
	"reflect"

	"github.com/cufee/aftermath/cmds/discord/commands"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/interactions"
	"github.com/cufee/aftermath/internal/database"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/rs/zerolog/log"
)

/*
Creates a new context and runs the middleware chain
*/
func (r *Router) handleEvent(event bot.Event) {
	nextHandler, exists := r.getEventHandler(event)
	if !exists {
		return
	}

	// Check if the event can be cast to an interaction, we need this to pull some data from it later
	_, ok := event.(discord.Interaction)
	if !ok {
		log.Error().Msg("bad interaction, failed to cast")
		return
	}

	ctx := context.Background()
	for i := len(r.middleware) - 1; i >= 0; i-- {
		nextHandler = r.middleware[i](ctx, nextHandler)
	}
	nextHandler(ctx, event)
}

func (r *Router) getEventHandler(event bot.Event) (func(ctx context.Context, event bot.Event), bool) {
	switch reflect.TypeOf(event) {
	case reflect.TypeOf(&events.ApplicationCommandInteractionCreate{}):
		command := event.(*events.ApplicationCommandInteractionCreate)
		if handler, ok := r.commandHandlers[command.Data.CommandName()]; ok {
			return func(ctx context.Context, e bot.Event) {
				user, _ := ctx.Value(common.ContextKeyUser).(database.User)
				err := handler(commands.ContextFrom(ctx, r.core, command, user))
				if err != nil {
					log.Err(err).Str("command", command.Data.CommandName()).Msg("error while handling a command")
				}
			}, true
		}

	case reflect.TypeOf(&events.ComponentInteractionCreate{}):
		interaction := event.(*events.ComponentInteractionCreate)
		if handler, ok := r.interactionHandlers[interaction.Data.CustomID()]; ok {
			return func(ctx context.Context, event bot.Event) {
				user, _ := ctx.Value(common.ContextKeyUser).(database.User)
				err := handler(interactions.ContextFrom(ctx, r.core, interaction, user))
				if err != nil {
					log.Err(err).Str("customId", interaction.Data.CustomID()).Msg("error while handling an interaction")
				}
			}, true
		}

	case reflect.TypeOf(&events.InteractionCreate{}):
		// For some reason, events come twice.
		// one of the events will have this generic payload, while the others have a specific one
		// we just ignore this event
		return nil, false
	}

	log.Info().Type("eventType", event).Msg("received an unhandled interaction type")
	return nil, false
}
