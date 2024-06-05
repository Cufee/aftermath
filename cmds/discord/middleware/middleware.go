package middleware

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/internal/database"
)

type MiddlewareFunc func(context.Context, func(context.Context, *discordgo.InteractionCreate)) func(context.Context, *discordgo.InteractionCreate)

func FetchUser(key any, client database.Client) MiddlewareFunc {
	return func(_ context.Context, f func(context.Context, *discordgo.InteractionCreate)) func(context.Context, *discordgo.InteractionCreate) {
		return func(ctx context.Context, i *discordgo.InteractionCreate) {
			user := database.User{ID: i.User.ID}

			f(context.WithValue(ctx, key, user), i)
		}
	}
}
