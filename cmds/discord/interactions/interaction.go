package interactions

import (
	c "context"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/internal/database"
	"github.com/disgoorg/disgo/events"
)

type Handler func(*context) error

type InteractionOptions struct {
	Name string
}

type Interaction struct {
	InteractionOptions
	Handler
}

type context struct {
	*common.Context
	event *events.ComponentInteractionCreate
}

func ContextFrom(ctx c.Context, client core.Client, event *events.ComponentInteractionCreate, user database.User) *context {
	return &context{
		common.NewContext(ctx, client, user, event.Locale().String()),
		event,
	}
}
