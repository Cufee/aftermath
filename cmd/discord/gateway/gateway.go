package gateway

import (
	"context"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/logic"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/pkg/errors"
	"github.com/servusdei2018/shards/v2"
)

var _ Client = &gatewayClient{}

type Client interface {
	discord.Commander

	Session(guildID string) *discordgo.Session

	Connect() error
	Disconnect() error
	Handler(fn interface{})
	SetStatus(status status, text string, emoji *discordgo.Emoji) error
}

type gatewayClient struct {
	core core.Client

	manager *shards.Manager

	rest       *rest.Client
	commands   []builder.Command
	middleware []middleware.MiddlewareFunc
}

func NewClient(core core.Client, token string, intent discordgo.Intent) (*gatewayClient, error) {
	mgr, err := shards.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	mgr.RegisterIntent(intent)

	rest, err := rest.NewClient(token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a new rest client")
	}

	return &gatewayClient{
		manager: mgr,
		core:    core,
		rest:    rest,
	}, nil
}

func (c *gatewayClient) Connect() error {
	if c.manager.Gateway.DataReady {
		return errors.New("already connected")
	}
	return c.manager.Start()
}

func (c *gatewayClient) Disconnect() error {
	return c.manager.Shutdown()
}

func (c *gatewayClient) Handler(fn interface{}) {
	c.manager.AddHandler(fn)
}

func (c *gatewayClient) Session(guildID string) *discordgo.Session {
	if guildID == "" {
		return c.manager.SessionForDM()
	}
	id, _ := strconv.Atoi(guildID)
	return c.manager.SessionForGuild(int64(id))
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
	session := c.manager.SessionForDM()

	switch status {
	default:
		return errors.New("invalid status provided")

	case StatusListening:
		return session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusOnline), Activities: activity(text, emoji, discordgo.ActivityTypeListening)})
	case StatusWatching:
		return session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusOnline), Activities: activity(text, emoji, discordgo.ActivityTypeWatching)})

	case StatusYellow:
		return session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusIdle), Activities: activity(text, emoji, discordgo.ActivityTypeCustom)})
	case StatusGreen:
		return session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusOnline), Activities: activity(text, emoji, discordgo.ActivityTypeCustom)})
	case StatusRed:
		return session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusDoNotDisturb), Activities: activity(text, emoji, discordgo.ActivityTypeCustom)})

	case StatusCustom:
		return session.UpdateStatusComplex(discordgo.UpdateStatusData{Status: string(discordgo.StatusOnline), Activities: activity(text, emoji, discordgo.ActivityTypeCustom)})
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
