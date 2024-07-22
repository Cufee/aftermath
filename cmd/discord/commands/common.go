package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
)

func saveInteractionData(ctx common.Context, command string, opts models.DiscordInteractionOptions) (discordgo.MessageComponent, error) {
	interaction := models.DiscordInteraction{
		Command:     command,
		Locale:      ctx.Locale(),
		UserID:      ctx.User().ID,
		ReferenceID: ctx.InteractionID(),
		Type:        models.InteractionTypeStats,
		Options:     opts,
	}
	err := ctx.Core().Database().CreateDiscordInteraction(ctx.Ctx(), interaction)
	if err != nil {
		return nil, err
	}

	return newStatsRefreshButton(interaction), nil
}
