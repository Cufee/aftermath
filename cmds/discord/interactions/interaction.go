package interactions

import (
	c "context"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/middleware"
	"github.com/disgoorg/disgo/events"
)

type Handler func(*context) error

type InteractionOptions struct {
	Name string
}

type Interaction struct {
	InteractionOptions
	Handler
	Middleware []middleware.MiddlewareFunc
}

type context struct {
	*common.Context
	event *events.ComponentInteractionCreate
}

func ContextFrom(ctx c.Context, client core.Client, event *events.ComponentInteractionCreate) *context {
	return &context{
		common.NewContext(ctx, client),
		event,
	}
}
