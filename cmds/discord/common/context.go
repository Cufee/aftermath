package common

import (
	"context"

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
	c      context.Context
	User   database.User
	Member discordgo.User

	Locale   language.Tag
	Localize localization.Printer

	Core core.Client

	rest        *rest.Client
	interaction discordgo.Interaction
	respondCh   chan<- discordgo.InteractionResponseData
}

func NewContext(ctx context.Context, interaction discordgo.Interaction, respondCh chan<- discordgo.InteractionResponseData, rest *rest.Client, client core.Client) *Context {
	c := &Context{c: ctx, Locale: language.English, Core: client, rest: rest, interaction: interaction, respondCh: respondCh}

	if interaction.User != nil {
		c.Member = *interaction.User
	}
	if interaction.Member != nil {
		c.Member = *interaction.Member.User
	}

	c.User, _ = ctx.Value(ContextKeyUser).(database.User)

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
	return c
}

func (c *Context) Reply(message string) error {
	c.respondCh <- discordgo.InteractionResponseData{Content: message}
	return nil
}
