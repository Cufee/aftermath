package gateway

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/log"
	"github.com/lucsky/cuid"
	"golang.org/x/text/language"
)

type eventContext struct {
	id string

	context.Context

	user models.User

	locale   language.Tag
	localize localization.Printer
	gw       *gatewayClient

	userID    string
	guildID   string
	messageID string
	channelID string
	eventType string
}

var ErrUnsupportedEvent = errors.New("event type not supported")
var ErrInvalidEvent = errors.New("event type missing required fields")
var ErrBotUserInitiated = errors.New("event initiated by a bot user")

func newContext(ctx context.Context, gw *gatewayClient, event any) (common.Context, error) {
	var c *eventContext
	switch cast := event.(type) {
	case *discordgo.MessageReactionAdd:
		locale := language.English
		if cast.Member != nil && cast.Member.User != nil {
			locale = common.LocaleToLanguageTag(discordgo.Locale(cast.Member.User.Locale))

			if cast.Member.User.Bot || cast.UserID == constants.DiscordBotUserID {
				return nil, ErrBotUserInitiated
			}
		}
		c = &eventContext{id: cuid.New(), Context: ctx, gw: gw, userID: cast.UserID, guildID: cast.GuildID, channelID: cast.ChannelID, messageID: cast.MessageID, locale: locale}

	case *discordgo.MessageCreate:
		if cast.Author == nil {
			return nil, ErrInvalidEvent
		}
		if cast.Author.Bot || cast.Author.ID == constants.DiscordBotUserID {
			return nil, ErrBotUserInitiated
		}
		c = &eventContext{id: cuid.New(), Context: ctx, gw: gw, userID: cast.Author.ID, guildID: cast.GuildID, channelID: cast.ChannelID, messageID: cast.ID, locale: common.LocaleToLanguageTag(discordgo.Locale(cast.Author.Locale))}

	default:
		return nil, ErrUnsupportedEvent
	}

	if c.userID == "" {
		return nil, ErrInvalidEvent
	}
	c.eventType = fmt.Sprintf("%T", event)

	user, err := gw.core.Database().GetOrCreateUserByID(ctx, c.userID, database.WithConnections(), database.WithSubscriptions(), database.WithContent())
	if err != nil {
		return nil, err
	}
	c.user = user

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
	meta["eventType"] = c.eventType
	meta["eventMessageId"] = c.messageID

	data := models.DiscordInteraction{
		Snowflake: c.InteractionID(),
		EventID:   c.eventType,
		MessageID: msg.ID,
		Locale:    c.locale,
		UserID:    c.user.ID,
		GuildID:   c.guildID,
		ChannelID: c.channelID,
		Type:      models.InteractionTypeGatewayEvent,
		Meta:      meta,
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
	return c.CreateMessage(c.Context, c.channelID, reply)
}

func (c *eventContext) InteractionFollowUp(reply common.Reply) (discordgo.Message, error) {
	return c.CreateMessage(c.Context, c.channelID, reply)
}

func (c *eventContext) Ctx() context.Context {
	return c.Context
}

func (c *eventContext) User() models.User {
	return c.user
}

func (c *eventContext) Member() discordgo.User {
	return discordgo.User{ID: c.userID}
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
	return discordgo.Interaction{}
}

func (c *eventContext) InteractionID() string {
	return c.id
}

func (c *eventContext) Reply() common.Reply {
	return common.ContextReply(c)
}

func (c *eventContext) Err(err error) error {
	log.Err(err).Str("interactionId", c.id).Msg("error while handling an interaction")
	button := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			common.ButtonJoinPrimaryGuild(c.localize("buttons_have_a_question_question")),
		}}
	return c.Reply().Hint(c.id).Component(button).Send("common_error_unhandled_reported")
}

func (c *eventContext) Error(message string) error {
	return c.Err(errors.New(message))
}

func (c *eventContext) ID() string {
	return strings.TrimPrefix(c.eventType, "*discordgo.")
}

func (c *eventContext) Options() common.Options {
	return common.Options{}
}

func (c *eventContext) CommandData() (discordgo.ApplicationCommandInteractionData, bool) {
	return discordgo.ApplicationCommandInteractionData{}, false
}

func (c *eventContext) ComponentData() (discordgo.MessageComponentInteractionData, bool) {
	return discordgo.MessageComponentInteractionData{}, false
}

func (c *eventContext) AutocompleteData() (discordgo.ApplicationCommandInteractionData, bool) {
	return discordgo.ApplicationCommandInteractionData{}, false
}

func (c *eventContext) DeleteResponse(ctx context.Context) error {
	return nil
}
func (c *eventContext) CreateMessage(ctx context.Context, channelID string, reply common.Reply) (discordgo.Message, error) {
	data, files := reply.Peek().Build(c.localize)
	msg, err := c.gw.rest.CreateMessage(ctx, channelID, discordgo.MessageSend{
		Content:    data.Content,
		Components: data.Components,
		Embeds:     data.Embeds,
		Reference:  reply.Peek().Reference,
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
