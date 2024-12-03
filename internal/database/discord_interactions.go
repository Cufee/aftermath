package database

import (
	"context"

	"time"

	"github.com/cufee/aftermath/internal/database/models"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	s "github.com/go-jet/jet/v2/sqlite"
)

func (c client) CreateDiscordInteraction(ctx context.Context, data models.DiscordInteraction) (models.DiscordInteraction, error) {
	model := models.FromDiscordInteraction(&data)
	stmt := t.DiscordInteraction.
		INSERT(t.DiscordInteraction.AllColumns).
		MODEL(model).
		RETURNING(t.DiscordInteraction.AllColumns)

	var result m.DiscordInteraction
	err := c.query(ctx, stmt, &result)
	if err != nil {
		return models.DiscordInteraction{}, err
	}
	return models.ToDiscordInteraction(&result), nil

}

func (c client) GetDiscordInteraction(ctx context.Context, id string) (models.DiscordInteraction, error) {
	stmt := t.DiscordInteraction.
		SELECT(t.DiscordInteraction.AllColumns).
		WHERE(t.DiscordInteraction.ID.EQ(s.String(id)))

	var result m.DiscordInteraction
	err := c.query(ctx, stmt, &result)
	if err != nil {
		return models.DiscordInteraction{}, err
	}
	return models.ToDiscordInteraction(&result), nil
}

type interactionQuery struct {
	id           []string
	snowflake    []string
	userID       []string
	guildID      []string
	channelID    []string
	messageID    []string
	eventID      []string
	kind         []string
	createdAfter *time.Time
	limit        int
}

func (q *interactionQuery) build() s.SelectStatement {
	stmt := t.DiscordInteraction.SELECT(t.DiscordInteraction.AllColumns)

	if q.limit == 0 {
		q.limit = 10
	}

	var where []s.BoolExpression
	// ID
	if q.id != nil {
		if len(q.id) == 1 {
			where = append(where, t.DiscordInteraction.ID.EQ(s.String(q.id[0])))
		} else {
			where = append(where, t.DiscordInteraction.ID.IN(toStringSlice(q.id...)...))
		}
	}
	// Snowflake
	if q.snowflake != nil {
		if len(q.snowflake) == 1 {
			where = append(where, t.DiscordInteraction.Snowflake.EQ(s.String(q.snowflake[0])))
		} else {
			where = append(where, t.DiscordInteraction.Snowflake.IN(toStringSlice(q.snowflake...)...))
		}
	}
	// UserID
	if q.userID != nil {
		if len(q.userID) == 1 {
			where = append(where, t.DiscordInteraction.UserID.EQ(s.String(q.userID[0])))
		} else {
			where = append(where, t.DiscordInteraction.UserID.IN(toStringSlice(q.userID...)...))
		}
	}
	// GuildID
	if q.guildID != nil {
		if len(q.guildID) == 1 {
			where = append(where, t.DiscordInteraction.GuildID.EQ(s.String(q.guildID[0])))
		} else {
			where = append(where, t.DiscordInteraction.GuildID.IN(toStringSlice(q.guildID...)...))
		}
	}
	// ChannelID
	if q.channelID != nil {
		if len(q.channelID) == 1 {
			where = append(where, t.DiscordInteraction.ChannelID.EQ(s.String(q.channelID[0])))
		} else {
			where = append(where, t.DiscordInteraction.ChannelID.IN(toStringSlice(q.channelID...)...))
		}
	}
	// MessageID
	if q.messageID != nil {
		if len(q.messageID) == 1 {
			where = append(where, t.DiscordInteraction.MessageID.EQ(s.String(q.messageID[0])))
		} else {
			where = append(where, t.DiscordInteraction.MessageID.IN(toStringSlice(q.messageID...)...))
		}
	}
	// EventID
	if q.eventID != nil {
		if len(q.eventID) == 1 {
			where = append(where, t.DiscordInteraction.EventID.EQ(s.String(q.eventID[0])))
		} else {
			where = append(where, t.DiscordInteraction.EventID.IN(toStringSlice(q.eventID...)...))
		}
	}
	// Type
	if q.kind != nil {
		if len(q.kind) == 1 {
			where = append(where, t.DiscordInteraction.Type.EQ(s.String(q.kind[0])))
		} else {
			where = append(where, t.DiscordInteraction.Type.IN(toStringSlice(q.kind...)...))
		}
	}
	// Type
	if q.createdAfter != nil {
		where = append(where, t.DiscordInteraction.CreatedAt.GT(s.DATETIME(q.createdAfter)))
	}

	return stmt.WHERE(s.AND(where...))
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
		for _, t := range types {
			iq.kind = append(iq.kind, string(t))
		}

	}
}
func WithSentAfter(after time.Time) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.createdAfter = &after
	}
}
func WithSnowflake(id string) InteractionQuery {
	return func(iq *interactionQuery) {
		iq.snowflake = append(iq.snowflake, id)
	}
}

func (c client) FindDiscordInteractions(ctx context.Context, opts ...InteractionQuery) ([]models.DiscordInteraction, error) {
	var query interactionQuery
	for _, apply := range opts {
		apply(&query)
	}

	var record []m.DiscordInteraction
	c.query(ctx, query.build(), &record)

	var interactions []models.DiscordInteraction
	for _, r := range record {
		interactions = append(interactions, models.ToDiscordInteraction(&r))
	}
	return interactions, nil
}
