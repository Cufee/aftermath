package public

import (
	"encoding/json"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
)

type statsOptions struct {
	commands.StatsOptions
	BackgroundID string
	ReferenceID  string
	AccountID    string
}

func (o statsOptions) fromInteraction(data models.DiscordInteraction) (statsOptions, error) {
	if data.Type != models.InteractionTypeComponent || data.Meta == nil || data.Meta["stats_options"] == nil {
		return statsOptions{}, errors.New("interactions contains no stats options")
	}

	raw, ok := data.Meta["stats_options"].(string)
	if !ok {
		return statsOptions{}, errors.New("invalid stats options data type")
	}

	return o, json.Unmarshal([]byte(raw), &o)
}

func (o statsOptions) refreshButton(ctx common.Context, id string) (discordgo.MessageComponent, error) {
	encoded, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	interaction := models.DiscordInteraction{
		Result:    "generated-refresh-button",
		EventID:   id,
		Locale:    ctx.Locale(),
		UserID:    ctx.User().ID,
		GuildID:   ctx.RawInteraction().GuildID,
		ChannelID: ctx.RawInteraction().ChannelID,
		MessageID: "not-available",
		Type:      models.InteractionTypeComponent,
		Meta:      map[string]any{"stats_options": string(encoded)},
	}
	if ctx.RawInteraction().Message != nil {
		interaction.MessageID = ctx.RawInteraction().Message.ID
	}

	interaction, err = ctx.Core().Database().CreateDiscordInteraction(ctx.Ctx(), interaction)
	if err != nil {
		return nil, err
	}

	return newStatsRefreshButton(interaction), nil
}
