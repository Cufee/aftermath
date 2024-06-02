package commands

import (
	c "context"

	"github.com/cufee/aftermath/cmds/core"
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

func ContextFrom(ctx c.Context, client core.Client, event *events.ApplicationCommandInteractionCreate, user database.User) *context {
	return &context{
		common.NewContext(ctx, client, user, event.Locale().String()),
		event,
	}
}
