package common

import (
	"context"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/discord/rest"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/localization"

	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

type contextKey int

const (
	ContextKeyUser contextKey = iota
	ContextKeyMember
	ContextKeyInteraction
)

type Context struct {
	context.Context
	User   models.User
	Member discordgo.User

	Locale   language.Tag
	Localize localization.Printer

	Core core.Client

	rest        *rest.Client
	interaction discordgo.Interaction
	respondCh   chan<- discordgo.InteractionResponseData
}

func NewContext(ctx context.Context, interaction discordgo.Interaction, respondCh chan<- discordgo.InteractionResponseData, rest *rest.Client, client core.Client) (*Context, error) {
	c := &Context{Context: ctx, Locale: language.English, Core: client, rest: rest, interaction: interaction, respondCh: respondCh}

	if interaction.User != nil {
		c.Member = *interaction.User
	}
	if interaction.Member != nil {
		c.Member = *interaction.Member.User
	}

	if c.Member.ID == "" {
		return nil, errors.New("failed to get a valid discord user id")
	}

	user, err := client.Database().GetOrCreateUserByID(ctx, c.Member.ID, database.WithConnections(), database.WithSubscriptions())
	if err != nil {
		return nil, err
	}
	c.User = user

	// Use the user locale selection by default with fallback to guild settings
	if c.interaction.Locale != "" {
		c.Locale = LocaleToLanguageTag(c.interaction.Locale)
	}

	printer, err := localization.NewPrinter("discord", c.Locale)
	if err != nil {
		log.Err(err).Msg("failed to get a localization printer for context")
		c.Localize = func(s string) string { return s }
	} else {
		c.Localize = printer
	}
	return c, nil
}

func (c *Context) respond(data discordgo.InteractionResponseData) error {
	select {
	case <-c.Context.Done():
		return c.Context.Err()
	default:
		c.respondCh <- data
	}
	return nil
}

func (c *Context) Message(message string) error {
	if message == "" {
		return errors.New("bad reply call with blank message")
	}
	return c.respond(discordgo.InteractionResponseData{Content: message})
}

func (c *Context) Reply(key string) error {
	return c.Message(c.Localize(key))
}

func (c *Context) Err(err error) error {
	log.Err(err).Str("interactionId", c.interaction.ID).Msg("error while handling an interaction")
	return c.Reply("common_error_unhandled_not_reported")
}

func (c *Context) Error(message string) error {
	log.Error().Str("message", message).Str("interactionId", c.interaction.ID).Msg("error while handling an interaction")
	return c.Reply("common_error_unhandled_not_reported")
}

func (c *Context) ReplyFmt(key string, args ...any) error {
	return c.Message(fmt.Sprintf(c.Localize(key), args...))
}

func (c *Context) File(r io.Reader, name string) error {
	if r == nil {
		return errors.New("bad Context#File call with nil io.Reader")
	}
	return c.respond(discordgo.InteractionResponseData{Files: []*discordgo.File{{Reader: r, Name: name}}})
}

func (c *Context) isCommand() bool {
	return c.interaction.Type == discordgo.InteractionApplicationCommand
}

func (c *Context) isComponentInteraction() bool {
	return c.interaction.Type == discordgo.InteractionMessageComponent
}

func (c *Context) ID() string {
	if c.isCommand() {
		d, _ := c.CommandData()
		return d.Name
	}
	if c.isComponentInteraction() {
		d, _ := c.ComponentData()
		return d.CustomID
	}
	return ""
}

func (c *Context) Options() options {
	if data, ok := c.CommandData(); ok {
		return data.Options
	}
	return options{}
}

type options []*discordgo.ApplicationCommandInteractionDataOption

func (o options) Value(name string) any {
	for _, opt := range o {
		if opt.Name == name {
			return opt.Value
		}
	}
	return nil
}

func (o options) Subcommand() (string, options, bool) {
	for _, opt := range o {
		if opt.Type == discordgo.ApplicationCommandOptionSubCommandGroup {
			name, opts, ok := options(opt.Options).Subcommand()
			return opt.Name + "_" + name, opts, ok
		}
		if opt.Type == discordgo.ApplicationCommandOptionSubCommand {
			return opt.Name, opt.Options, true
		}
	}
	return "", options{}, false
}

func GetOption[T any](c *Context, name string) (T, bool) {
	var v T
	if data, ok := c.CommandData(); ok {
		for _, opt := range data.Options {
			if opt.Name == name {
				v, _ = opt.Value.(T)

				return v, true
			}
		}
	}
	return v, false
}

func (c *Context) CommandData() (discordgo.ApplicationCommandInteractionData, bool) {
	if !c.isCommand() {
		return discordgo.ApplicationCommandInteractionData{}, false
	}
	data, ok := c.interaction.Data.(discordgo.ApplicationCommandInteractionData)
	return data, ok
}

func (c *Context) ComponentData() (discordgo.MessageComponentInteractionData, bool) {
	if !c.isComponentInteraction() {
		return discordgo.MessageComponentInteractionData{}, false
	}
	data, ok := c.interaction.Data.(discordgo.MessageComponentInteractionData)
	return data, ok
}
