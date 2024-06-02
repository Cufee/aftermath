package commands

import (
	c "context"

	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/internal/database"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type Handler func(*context) error

type Command struct {
	discord.ApplicationCommandCreate
	Handler
}

type context struct {
	*common.Context
	event *events.ApplicationCommandInteractionCreate
}

func ContextFrom(ctx c.Context, event *events.ApplicationCommandInteractionCreate, user database.User) *context {
	return &context{
		common.NewContext(ctx, user, event.Locale().String()),
		event,
	}
}
