package middleware

import (
	"context"

	"github.com/cufee/aftermath/internal/database"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
)

type MiddlewareFunc func(context.Context, func(context.Context, bot.Event)) func(context.Context, bot.Event)

func FetchUser(key any, client database.Client) MiddlewareFunc {
	return func(_ context.Context, f func(context.Context, bot.Event)) func(context.Context, bot.Event) {
		return func(ctx context.Context, e bot.Event) {
			interaction, ok := e.(discord.Interaction)
			if !ok {
				return
			}
			f(context.WithValue(ctx, key, database.User{ID: interaction.User().ID.String()}), e)
		}
	}
}
