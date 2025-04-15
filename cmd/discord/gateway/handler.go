package gateway

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/log"
)

func ToGatewayHandler[T any](client *gatewayClient, handler common.EventHandler[T]) func(s *discordgo.Session, e *T) {
	return func(s *discordgo.Session, e *T) {
		if !handler.Match(client.core.Database(), s, e) {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		cCtx, err := newContext(ctx, client, e)
		if errors.Is(err, ErrBotUserInitiated) {
			return
		}
		if errors.Is(err, ErrUnsupportedEvent) {
			log.Debug().Str("type", fmt.Sprintf("%T", e)).Msg("unsupported gateway event type")
			return
		}
		if err != nil {
			log.Err(err).Msg("failed to create a context for gateway event")
			return
		}

		chain := func(ctx common.Context) error {
			return handler.Handle(ctx, e)
		}
		for i := len(client.middleware) - 1; i >= 0; i-- {
			chain = client.middleware[i](cCtx, chain)
		}
		err = chain(cCtx)
		if err != nil {
			log.Err(err).Str("command", cCtx.ID()).Str("interactionId", cCtx.InteractionID()).Msg("gateway event handler returned an error")
			return
		}
	}
}
