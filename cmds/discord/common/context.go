package common

import (
	"context"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/disgoorg/disgo/discord"
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
	User database.User

	Member      discord.User
	Interaction discord.Interaction

	Locale   language.Tag
	Localize localization.Printer

	Core core.Client
}

func NewContext(ctx context.Context, client core.Client) *Context {
	c := &Context{Context: ctx, Core: client, Locale: language.English}

	c.User, _ = ctx.Value(ContextKeyUser).(database.User)
	c.Member, _ = ctx.Value(ContextKeyMember).(discord.User)
	c.Interaction, _ = ctx.Value(ContextKeyInteraction).(discord.Interaction)

	// Use the user locale selection by default with fallback to guild settings
	if c.Interaction.Locale() != "" {
		c.Locale = LocaleToLanguageTag(c.Interaction.Locale())
	} else if locale := c.Interaction.GuildLocale(); locale != nil {
		c.Locale = LocaleToLanguageTag(*locale)
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
