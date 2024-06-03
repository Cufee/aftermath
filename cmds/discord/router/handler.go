package router

import (
	"context"
	"reflect"

	"github.com/cufee/aftermath/cmds/discord/commands"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/interactions"
	"github.com/cufee/aftermath/cmds/discord/middleware"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/rs/zerolog/log"
)

/*
Creates a new context and runs the middleware chain
*/
func (r *Router) handleEvent(event bot.Event) {
	nextHandler, middleware, exists := r.getEventHandler(event)
	if !exists {
		return
	}

	// Check if the event can be cast to an interaction, we need this to pull some data from it later
	interaction, ok := event.(discord.Interaction)
	if !ok {
		log.Error().Msg("bad interaction, failed to cast")
		return
	}

	withMember := context.WithValue(context.Background(), common.ContextKeyMember, interaction.User())
	withInteraction := context.WithValue(withMember, common.ContextKeyInteraction, interaction)

	chain := append(middleware, r.middleware...)
	for i := len(chain) - 1; i >= 0; i-- {
		nextHandler = chain[i](withInteraction, nextHandler)
	}
	nextHandler(withInteraction, event)
}

func (r *Router) getEventHandler(event bot.Event) (func(ctx context.Context, event bot.Event), []middleware.MiddlewareFunc, bool) {
	switch reflect.TypeOf(event) {
	case reflect.TypeOf(&events.ApplicationCommandInteractionCreate{}):
		event := event.(*events.ApplicationCommandInteractionCreate)
		if command, ok := r.commands[event.Data.CommandName()]; ok {
			return func(ctx context.Context, _ bot.Event) {
				err := command.Handler(commands.ContextFrom(ctx, r.core, event))
				if err != nil {
					log.Err(err).Str("command", event.Data.CommandName()).Msg("error while handling a command")
				}
			}, command.Middleware, true
		}

	case reflect.TypeOf(&events.ComponentInteractionCreate{}):
		event := event.(*events.ComponentInteractionCreate)
		if interaction, ok := r.interactions[event.Data.CustomID()]; ok {
			return func(ctx context.Context, _ bot.Event) {
				err := interaction.Handler(interactions.ContextFrom(ctx, r.core, event))
				if err != nil {
					log.Err(err).Str("customId", event.Data.CustomID()).Msg("error while handling an interaction")
				}
			}, interaction.Middleware, true
		}

	case reflect.TypeOf(&events.InteractionCreate{}):
		// For some reason, events come twice.
		// one of the events will have this generic payload, while the others have a specific one
		// we just ignore this event
		return nil, nil, false
	}

	log.Info().Type("eventType", event).Msg("received an unhandled interaction type")
	return nil, nil, false
}
