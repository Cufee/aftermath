package public

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/encoding"
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

	raw, ok := data.Meta["stats_options"].([]byte)
	if !ok {
		return statsOptions{}, errors.New("invalid stats options data type")
	}

	return o, encoding.DecodeGob(raw, &o)
}

func (o statsOptions) refreshButton(ctx common.Context) (discordgo.MessageComponent, error) {
	encoded, err := encoding.EncodeGob(o)
	if err != nil {
		return nil, err
	}

	interaction := models.DiscordInteraction{
		Result:    "generated-refresh-button",
		EventID:   ctx.ID(),
		Locale:    ctx.Locale(),
		UserID:    ctx.User().ID,
		GuildID:   ctx.RawInteraction().GuildID,
		ChannelID: ctx.RawInteraction().ChannelID,
		MessageID: "not-available",
		Type:      models.InteractionTypeComponent,
		Meta:      map[string]any{"stats_options": encoded},
	}
	if ctx.RawInteraction().Message != nil {
		interaction.MessageID = ctx.RawInteraction().Message.ID
	}

	err = ctx.Core().Database().CreateDiscordInteraction(ctx.Ctx(), interaction)
	if err != nil {
		return nil, err
	}

	return newStatsRefreshButton(interaction), nil
}
