package common

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/discord/rest"

	"github.com/cufee/aftermath/internal/database"
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
	User   database.User
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

	user, err := client.DB.GetOrCreateUserByID(ctx, c.Member.ID, database.WithConnections(), database.WithSubscriptions())
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

func (c *Context) Reply(key string) error {
	c.respondCh <- discordgo.InteractionResponseData{Content: c.Localize(key)}
	return nil
}

func (c *Context) Err(err error) error {
	log.Err(err).Str("interactionId", c.interaction.ID).Msg("error while handling an interaction")
	return c.Reply(err.Error())
}

func (c *Context) Error(message string) error {
	log.Error().Str("message", message).Str("interactionId", c.interaction.ID).Msg("error while handling an interaction")
	return c.Reply(message)
}

func (c *Context) ReplyFmt(key string, args ...any) error {
	c.respondCh <- discordgo.InteractionResponseData{Content: fmt.Sprintf(c.Localize(key), args...)}
	return nil
}

func (c *Context) File(r io.Reader, name string) error {
	c.respondCh <- discordgo.InteractionResponseData{Files: []*discordgo.File{{Reader: r, Name: name}}}
	return nil
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

func (c *Context) Option(name string) any {
	if data, ok := c.CommandData(); ok {
		for _, opt := range data.Options {
			fmt.Printf("%+v", opt)
			if opt.Name == name {
				return opt.Value
			}
		}
	}
	return nil
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
