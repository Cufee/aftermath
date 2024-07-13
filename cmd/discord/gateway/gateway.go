package gateway

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var _ Client = &gatewayClient{}

type Client interface {
	Connect() error
	Handler(fn interface{}) func()
	SetStatus(status status, text string, emoji *discordgo.Emoji) error
}

type gatewayClient struct {
	session *discordgo.Session
}

func NewClient(token string, intent discordgo.Intent) (*gatewayClient, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.Identify.Intents = intent
	return &gatewayClient{
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
