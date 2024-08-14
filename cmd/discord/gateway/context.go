package gateway

import (
	"context"
	"errors"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/log"
	"golang.org/x/text/language"
)

type eventContext struct {
	context.Context

	user   models.User
	member discordgo.User

	locale   language.Tag
	localize localization.Printer
	gw       *gatewayClient

	interaction discordgo.Interaction
}

func newContext(ctx context.Context, gw *gatewayClient, interaction discordgo.Interaction) (common.Context, error) {
	c := &eventContext{Context: ctx, locale: language.English, interaction: interaction, gw: gw}

	if interaction.User != nil {
		c.member = *interaction.User
	}
	if interaction.Member != nil {
		c.member = *interaction.Member.User
	}

	if c.member.ID == "" {
		return nil, errors.New("failed to get a valid discord user id")
	}

	user, err := gw.core.Database().GetOrCreateUserByID(ctx, c.member.ID, database.WithConnections(), database.WithSubscriptions(), database.WithContent())
	if err != nil {
		return nil, err
	}
	c.user = user

	// Use the user locale selection by default with fallback to guild settings
	if c.interaction.Locale != "" {
		c.locale = common.LocaleToLanguageTag(c.interaction.Locale)
	}

	printer, err := localization.NewPrinterWithFallback("discord", c.locale)
	if err != nil {
		log.Err(err).Msg("failed to get a localization printer for context")
		c.localize = func(s string) string { return s }
	} else {
		c.localize = printer
	}
	return c, nil
}

func (c *eventContext) saveInteractionEvent(msg discordgo.Message, msgErr error, reply common.Reply) {
	meta := reply.Metadata()
	i := c.interaction
	i.Token = "..."
	meta["interaction"] = i

	data := models.DiscordInteraction{
		Snowflake: c.InteractionID(),
		ID:        c.interaction.ID,
		EventID:   c.ID(),
		MessageID: msg.ID,
		Locale:    c.locale,
		UserID:    c.user.ID,
		GuildID:   c.interaction.GuildID,
		ChannelID: c.interaction.ChannelID,
		Type:      models.InteractionTypeUnknown,
		Meta:      meta,
	}
	if c.isCommand() {
		data.Type = models.InteractionTypeCommand
	}
	if c.isComponentInteraction() {
		data.Type = models.InteractionTypeComponent
	}
	if c.isAutocompleteInteraction() {
		data.Type = models.InteractionTypeAutocomplete
	}
	if msgErr != nil {
		data.Result = "error"
		data.Meta["error"] = msgErr.Error()
	} else {
		data.Result = "ok"
	}

	// save event to db
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	data, err := c.gw.core.Database().CreateDiscordInteraction(ctx, data)
	if err != nil {
		log.Err(err).Msg("failed to save interaction event")
	}
}
func (c *eventContext) InteractionResponse(reply common.Reply) (discordgo.Message, error) {
	data, files := reply.Peek().Build(c.localize)
	select {
	case <-c.Context.Done():
		go c.saveInteractionEvent(discordgo.Message{}, c.Context.Err(), reply)
		return discordgo.Message{}, c.Context.Err()
	default:
		msg, err := common.WithRetry(func() (discordgo.Message, error) {
			// since we already finished handling the interaction, there is no need to use the handler context
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()

			if c.interaction.Type == discordgo.InteractionApplicationCommandAutocomplete {
				msg, err := c.gw.rest.SendInteractionResponse(ctx, c.interaction.ID, c.interaction.Token, discordgo.InteractionResponse{Type: discordgo.InteractionApplicationCommandAutocompleteResult, Data: &data}, nil)
				if errors.Is(err, rest.ErrInteractionAlreadyAcked) {
					err = nil
				}
				return msg, err
			}
			return c.gw.rest.UpdateInteractionResponse(ctx, c.interaction.AppID, c.interaction.Token, data, files)
		})

		go c.saveInteractionEvent(msg, err, reply)
		return msg, err
	}
}

func (c *eventContext) InteractionFollowUp(reply common.Reply) (discordgo.Message, error) {
	data, files := reply.Peek().Build(c.localize)
	select {
	case <-c.Context.Done():
		go c.saveInteractionEvent(discordgo.Message{}, c.Context.Err(), reply)
		return discordgo.Message{}, c.Context.Err()
	default:
		return common.WithRetry(func() (discordgo.Message, error) {
			// since we already finished handling the interaction, there is no need to use the handler context
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			msg, err := c.gw.rest.SendInteractionFollowup(ctx, c.interaction.AppID, c.interaction.Token, data, files)

			go c.saveInteractionEvent(msg, err, reply)
			return msg, err
		})
	}
}

func (c *eventContext) Ctx() context.Context {
	return c.Context
}

func (c *eventContext) User() models.User {
	return c.user
}

func (c *eventContext) Member() discordgo.User {
	return c.member
}

func (c *eventContext) Locale() language.Tag {
	return c.locale
}

func (c *eventContext) Core() core.Client {
	return c.gw.core
}

func (c *eventContext) Localize(key string) string {
	return c.localize(key)
}

func (c *eventContext) Rest() *rest.Client {
	return c.gw.rest
}

func (c *eventContext) RawInteraction() discordgo.Interaction {
	return c.interaction
}

func (c *eventContext) InteractionID() string {
	return c.interaction.ID
}

func (c *eventContext) Reply() common.Reply {
	return common.ContextReply(c)
}

func (c *eventContext) Err(err error) error {
	log.Err(err).Str("interactionId", c.interaction.ID).Msg("error while handling an interaction")
	button := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			common.ButtonJoinPrimaryGuild(c.localize("buttons_have_a_question_question")),
		}}
	return c.Reply().Hint(c.interaction.ID).Component(button).Send("common_error_unhandled_reported")
}

func (c *eventContext) Error(message string) error {
	return c.Err(errors.New(message))
}

func (c *eventContext) isCommand() bool {
	return c.interaction.Type == discordgo.InteractionApplicationCommand
}

func (c *eventContext) isComponentInteraction() bool {
	return c.interaction.Type == discordgo.InteractionMessageComponent
}

func (c *eventContext) isAutocompleteInteraction() bool {
	return c.interaction.Type == discordgo.InteractionApplicationCommandAutocomplete
}

func (c *eventContext) ID() string {
	var id string
	if c.isCommand() {
		d, _ := c.CommandData()
		id = d.Name
	}
	if c.isComponentInteraction() {
		d, _ := c.ComponentData()
		id = d.CustomID
	}
	if c.isAutocompleteInteraction() {
		d, _ := c.AutocompleteData()
		id = d.Name
	}
	if id != "" {
		return id
	}
	return "unknown"
}

func (c *eventContext) Options() common.Options {
	if data, ok := c.interaction.Data.(discordgo.ApplicationCommandInteractionData); ok {
		var o common.Options = data.Options
		return o.Deep()
	}
	return common.Options{}
}

func (c *eventContext) CommandData() (discordgo.ApplicationCommandInteractionData, bool) {
	if !c.isCommand() {
		return discordgo.ApplicationCommandInteractionData{}, false
	}
	data, ok := c.interaction.Data.(discordgo.ApplicationCommandInteractionData)
	return data, ok
}

func (c *eventContext) ComponentData() (discordgo.MessageComponentInteractionData, bool) {
	if !c.isComponentInteraction() {
		return discordgo.MessageComponentInteractionData{}, false
	}
	data, ok := c.interaction.Data.(discordgo.MessageComponentInteractionData)
	return data, ok
}

func (c *eventContext) AutocompleteData() (discordgo.ApplicationCommandInteractionData, bool) {
	if !c.isAutocompleteInteraction() {
		return discordgo.ApplicationCommandInteractionData{}, false
	}
	data, ok := c.interaction.Data.(discordgo.ApplicationCommandInteractionData)
	return data, ok
}

func (c *eventContext) DeleteResponse(ctx context.Context) error {
	return c.gw.rest.DeleteInteractionResponse(ctx, c.interaction.AppID, c.interaction.Token)
}
func (c *eventContext) CreateMessage(ctx context.Context, channelID string, reply common.Reply) (discordgo.Message, error) {
	data, files := reply.Peek().Build(c.localize)
	msg, err := c.gw.rest.CreateMessage(ctx, channelID, discordgo.MessageSend{
		Content:    data.Content,
		Components: data.Components,
		Embeds:     data.Embeds,
	}, files)

	go c.saveInteractionEvent(msg, err, reply)
	return msg, err
}
func (c *eventContext) UpdateMessage(ctx context.Context, channelID string, messageID string, reply common.Reply) (discordgo.Message, error) {
	data, files := reply.Peek().Build(c.localize)
	edit := discordgo.MessageEdit{
		Attachments: data.Attachments,
	}
	if data.Content != "" {
		edit.Content = &data.Content
	}
	if data.Components != nil {
		edit.Components = &data.Components
	}
	if data.Embeds != nil {
		edit.Embeds = &data.Embeds
	}
	msg, err := c.gw.rest.UpdateMessage(ctx, channelID, messageID, edit, files)

	go c.saveInteractionEvent(msg, err, reply)
	return msg, err
}
func (c *eventContext) CreateDMChannel(ctx context.Context, userID string) (discordgo.Channel, error) {
	return c.gw.rest.CreateDMChannel(ctx, userID)
}
