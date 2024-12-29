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

		err = handler.Handle(cCtx, e)
		if err != nil {
			log.Err(err).Msg("gateway event handler returned an error")
			return
		}
	}
}
