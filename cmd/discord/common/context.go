package common

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord/rest"

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
}

func NewContext(ctx context.Context, interaction discordgo.Interaction, rest *rest.Client, client core.Client) (*Context, error) {
	c := &Context{Context: ctx, Locale: language.English, Core: client, rest: rest, interaction: interaction}

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

func (c *Context) respond(data discordgo.InteractionResponse, files []rest.File) error {
	select {
	case <-c.Context.Done():
		return c.Context.Err()
	default:
		ctx, cancel := context.WithTimeout(c.Context, time.Millisecond*3000)
		defer cancel()
		err := c.rest.UpdateOrSendInteractionResponse(ctx, c.interaction.AppID, c.interaction.ID, c.interaction.Token, data, files)
		if os.IsTimeout(err) {
			// request timed out, most likely a discord issue
			log.Error().Str("id", c.ID()).Str("interactionId", c.interaction.ID).Msg("discord api: request timed out while responding to an interaction")
			ctx, cancel := context.WithTimeout(c.Context, time.Millisecond*1000)
			defer cancel()
			return c.rest.UpdateOrSendInteractionResponse(ctx, c.interaction.AppID, c.interaction.ID, c.interaction.Token, discordgo.InteractionResponse{Data: &discordgo.InteractionResponseData{Content: c.Localize("common_error_discord_outage")}, Type: discordgo.InteractionResponseChannelMessageWithSource}, nil)
		}
		return err
	}
}

func (c *Context) InteractionID() string {
	return c.interaction.ID
}

func (c *Context) Reply() reply {
	return reply{ctx: c}
}

func (c *Context) Err(err error) error {
	log.Err(err).Str("interactionId", c.interaction.ID).Msg("error while handling an interaction")
	return c.Reply().Send("common_error_unhandled_reported")
}

func (c *Context) Error(message string) error {
	log.Error().Str("message", message).Str("interactionId", c.interaction.ID).Msg("error while handling an interaction")
	return c.Reply().Send("common_error_unhandled_reported")
}

func (c *Context) isCommand() bool {
	return c.interaction.Type == discordgo.InteractionApplicationCommand
}

func (c *Context) isComponentInteraction() bool {
	return c.interaction.Type == discordgo.InteractionMessageComponent
}

func (c *Context) isAutocompleteInteraction() bool {
	return c.interaction.Type == discordgo.InteractionApplicationCommandAutocomplete
}

func (c *Context) isModalSubmit() bool {
	return c.interaction.Type == discordgo.InteractionModalSubmit
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
	if data, ok := c.interaction.Data.(discordgo.ApplicationCommandInteractionData); ok {
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

func (c *Context) AutocompleteData() (discordgo.ApplicationCommandInteractionData, bool) {
	if !c.isAutocompleteInteraction() {
		return discordgo.ApplicationCommandInteractionData{}, false
	}
	data, ok := c.interaction.Data.(discordgo.ApplicationCommandInteractionData)
	return data, ok
}
