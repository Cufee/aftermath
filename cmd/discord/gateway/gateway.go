package gateway

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/logic"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/pkg/errors"
)

var _ Client = &gatewayClient{}

type Client interface {
	Connect() error
	Handler(fn interface{}) func()
	SetStatus(status status, text string, emoji *discordgo.Emoji) error
}

type gatewayClient struct {
	core core.Client

	rest       *rest.Client
	session    *discordgo.Session
	commands   []builder.Command
	middleware []middleware.MiddlewareFunc
}

func NewClient(core core.Client, token string, intent discordgo.Intent) (*gatewayClient, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	rest, err := rest.NewClient(token)
	if err != nil {
		return nil, errors.Errorf("failed to create a new rest client :%w", err)
	}

	dg.Identify.Intents = intent
	return &gatewayClient{
		core:    core,
		rest:    rest,
		session: dg,
	}, nil
}

func (c *gatewayClient) Connect() error {
	if c.session.DataReady {
		return errors.New("already connected")
	}
	return c.session.Open()
}

func (c *gatewayClient) Handler(fn interface{}) func() {
	return c.session.AddHandler(fn)
}

func (c *gatewayClient) Session() *discordgo.Session {
	return c.session
}

/*
Loads commands into the router, does not update bot commands through Discord API
*/
func (c *gatewayClient) LoadCommands(commands ...builder.Command) {
	c.commands = append(c.commands, commands...)
}

/*
Loads interactions into the router
*/
func (c *gatewayClient) LoadMiddleware(middleware ...middleware.MiddlewareFunc) {
	c.middleware = append(c.middleware, middleware...)
}

/*
Loads interactions into the router
*/
func (c *gatewayClient) UpdateLoadedCommands(ctx context.Context) error {
	return logic.UpdateCommands(ctx, c.core.Database(), c.rest, c.commands)
}

type status int

const (
	StatusListening = iota
	StatusWatching  = iota

	StatusCustom = iota
	StatusYellow = iota
	StatusGreen  = iota
	StatusRed    = iota
)

func (c *gatewayClient) SetStatus(status status, text string, emoji *discordgo.Emoji) error {
	switch status {
	default:
		return errors.New("invalid status provided")

	case StatusListening:
		return c.session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusOnline), Activities: activity(text, emoji, discordgo.ActivityTypeListening)})
	case StatusWatching:
		return c.session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusOnline), Activities: activity(text, emoji, discordgo.ActivityTypeWatching)})

	case StatusYellow:
		return c.session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusIdle), Activities: activity(text, emoji, discordgo.ActivityTypeCustom)})
	case StatusGreen:
		return c.session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusOnline), Activities: activity(text, emoji, discordgo.ActivityTypeCustom)})
	case StatusRed:
		return c.session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusDoNotDisturb), Activities: activity(text, emoji, discordgo.ActivityTypeCustom)})

	case StatusCustom:
		return c.session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusOnline), Activities: activity(text, emoji, discordgo.ActivityTypeCustom)})
	}
}

func activity(text string, emoji *discordgo.Emoji, at discordgo.ActivityType) []*discordgo.Activity {
	a := discordgo.Activity{
		Name:  text,
		State: text,
		Type:  at,
	}
	if emoji != nil {
		a.Emoji = *emoji
	}
	return []*discordgo.Activity{&a}
}
