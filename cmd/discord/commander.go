package discord

import (
	"context"

	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/middleware"
)

type Commander interface {
	LoadCommands(commands ...builder.Command)
	UpdateLoadedCommands(ctx context.Context) error

	LoadMiddleware(middleware ...middleware.MiddlewareFunc)
}
