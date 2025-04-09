package common

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/guregu/null/v6"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord/rest"

	"github.com/cufee/aftermath/internal/database/models"

	"golang.org/x/text/language"
)

type Options []*discordgo.ApplicationCommandInteractionDataOption

type Error null.Value[byte]

var (
	UnknownError     Error = Error(null.ValueFrom[byte](0))
	ApplicationError Error = Error(null.ValueFrom[byte](1))
	UserError        Error = Error(null.ValueFrom[byte](2))
)

type Context interface {
	Ctx() context.Context
	Core() core.Client
	Rest() *rest.Client

	ID() string
	User() models.User
	Member() discordgo.User
	ChannelID() string
	GuildID() null.String

	Locale() language.Tag
	Localize(string) string

	InteractionID() string
	RawInteraction() discordgo.Interaction
	CommandData() (discordgo.ApplicationCommandInteractionData, bool)
	ComponentData() (discordgo.MessageComponentInteractionData, bool)
	AutocompleteData() (discordgo.ApplicationCommandInteractionData, bool)

	Reply() Reply
	Err(err error, kind Error) error
	Error(message string, kind Error) error
	Options() Options

	SetError(Error)
	HasError() bool
	ErrorType() Error

	WithAds(v bool)
	ShowAds() bool

	DeleteResponse(ctx context.Context) error
	CreateMessage(ctx context.Context, channelID string, reply Reply) (discordgo.Message, error)
	UpdateMessage(ctx context.Context, channelID string, messageID string, reply Reply) (discordgo.Message, error)

	CreateDMChannel(ctx context.Context, userID string) (discordgo.Channel, error)

	InteractionResponse(reply Reply) (discordgo.Message, error)
	InteractionFollowUp(ctx context.Context, reply Reply) (discordgo.Message, error)
}

func (o Options) Value(name string) any {
	for _, opt := range o {
		if opt.Name == name {
			return opt.Value
		}
	}
	return nil
}

func (o Options) Deep() Options {
	for _, opt := range o {
		var opts Options = opt.Options
		o = append(o, opts.Deep()...)
	}
	return o
}

func (o Options) Subcommand() (string, Options, bool) {
	for _, opt := range o {
		if opt.Type == discordgo.ApplicationCommandOptionSubCommandGroup {
			name, opts, ok := Options(opt.Options).Subcommand()
			return opt.Name + "_" + name, opts, ok
		}
		if opt.Type == discordgo.ApplicationCommandOptionSubCommand {
			return opt.Name, opt.Options, true
		}
	}
	return "", Options{}, false
}

func GetOption[T any](data []*discordgo.ApplicationCommandInteractionDataOption, name string) (T, bool) {
	var v T
	for _, opt := range data {
		if opt.Name == name {
			v, _ = opt.Value.(T)
			return v, true
		}
	}
	return v, false
}
