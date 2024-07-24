package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/discordinteraction"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

func (c client) CreateDiscordInteraction(ctx context.Context, data models.DiscordInteraction) (models.DiscordInteraction, error) {
	user, err := c.db.User.Get(ctx, data.UserID)
	if err != nil {
		return models.DiscordInteraction{}, errors.Wrap(err, "failed to get user")
	}

	record, err := c.db.DiscordInteraction.Create().
		SetChannelID(data.ChannelID).
		SetEventID(data.EventID).
		SetGuildID(data.GuildID).
		SetLocale(data.Locale.String()).
		SetMessageID(data.MessageID).
		SetMetadata(data.Meta).
		SetResult(data.Result).
		SetType(data.Type).
		SetUser(user).Save(ctx)
	if err != nil {
		return models.DiscordInteraction{}, err
	}

	return toDiscordInteraction(record), nil
}

func (c client) GetDiscordInteraction(ctx context.Context, id string) (models.DiscordInteraction, error) {
	record, err := c.db.DiscordInteraction.Get(ctx, id)
	if err != nil {
		return models.DiscordInteraction{}, err
	}
	return toDiscordInteraction(record), nil
}

func toDiscordInteraction(record *db.DiscordInteraction) models.DiscordInteraction {
	locale, err := language.Parse(record.Locale)
	if err != nil {
		locale = language.English
	}
	i := models.DiscordInteraction{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,

		Result:    record.Result,
		UserID:    record.UserID,
		GuildID:   record.GuildID,
		ChannelID: record.ChannelID,
		MessageID: record.MessageID,

		Locale:  locale,
		Type:    record.Type,
		EventID: record.EventID,
		Meta:    record.Metadata,
	}
	if i.Meta == nil {
		i.Meta = make(map[string]any)
	}
	return i
}

type interactionQuery struct {
	id           []string
	userID       []string
	guildID      []string
	channelID    []string
	messageID    []string
	eventID      []string
	kind         []models.DiscordInteractionType
	createdAfter *time.Time
	limit        int
}

func (q *interactionQuery) build() []predicate.DiscordInteraction {
	if q.limit == 0 {
		q.limit = 10
	}

	var where []predicate.DiscordInteraction
	// ID
	if q.id != nil {
		if len(q.id) == 1 {
			where = append(where, discordinteraction.ID(q.id[0]))
		} else {
			where = append(where, discordinteraction.IDIn(q.id...))
		}
	}
	// UserID
	if q.userID != nil {
		if len(q.userID) == 1 {
			where = append(where, discordinteraction.UserID(q.userID[0]))
		} else {
			where = append(where, discordinteraction.UserIDIn(q.userID...))
		}
	}
	// GuildID
	if q.guildID != nil {
		if len(q.guildID) == 1 {
			where = append(where, discordinteraction.GuildID(q.guildID[0]))
		} else {
			where = append(where, discordinteraction.GuildIDIn(q.guildID...))
		}
	}
	// ChannelID
	if q.channelID != nil {
		if len(q.channelID) == 1 {
			where = append(where, discordinteraction.ChannelID(q.channelID[0]))
		} else {
			where = append(where, discordinteraction.ChannelIDIn(q.channelID...))
		}
	}
	// MessageID
	if q.messageID != nil {
		if len(q.messageID) == 1 {
			where = append(where, discordinteraction.MessageID(q.messageID[0]))
		} else {
			where = append(where, discordinteraction.MessageIDIn(q.messageID...))
		}
	}
	// EventID
	if q.eventID != nil {
		if len(q.eventID) == 1 {
			where = append(where, discordinteraction.EventID(q.eventID[0]))
		} else {
			where = append(where, discordinteraction.EventIDIn(q.eventID...))
		}
	}
	// Type
	if q.kind != nil {
		if len(q.kind) == 1 {
			where = append(where, discordinteraction.TypeEQ(q.kind[0]))
		} else {
			where = append(where, discordinteraction.TypeIn(q.kind...))
		}
	}
	// Type
	if q.createdAfter != nil {
		where = append(where, discordinteraction.CreatedAtGT(*q.createdAfter))
	}
	return where
}

type InteractionQuery func(*interactionQuery)

func WithLimit(limit int) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.limit = limit
	}
}
func WithID(id ...string) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.id = append(iq.id, id...)
	}
}
func WithUserID(id ...string) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.userID = append(iq.userID, id...)
	}
}
func WithGuildID(id ...string) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.guildID = append(iq.guildID, id...)
	}
}
func WithChannelID(id ...string) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.channelID = append(iq.channelID, id...)
	}
}
func WithMessageID(id ...string) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.messageID = append(iq.messageID, id...)
	}
}
func WithEventID(id ...string) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.eventID = append(iq.eventID, id...)
	}
}
func WithType(types ...models.DiscordInteractionType) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.kind = append(iq.kind, types...)
	}
}
func WithSentAfter(after time.Time) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.createdAfter = &after
	}
}

func (c client) FindDiscordInteractions(ctx context.Context, opts ...InteractionQuery) ([]models.DiscordInteraction, error) {
	var query interactionQuery
	for _, apply := range opts {
		apply(&query)
	}

	records, err := c.db.DiscordInteraction.Query().Where(query.build()...).Limit(query.limit).All(ctx)
	if err != nil {
		return nil, err
	}

	var interactions []models.DiscordInteraction
	for _, r := range records {
		interactions = append(interactions, toDiscordInteraction(r))
	}
	return interactions, nil
}
