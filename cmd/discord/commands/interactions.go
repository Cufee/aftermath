package commands

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"

	stats "github.com/cufee/aftermath/internal/stats/client/v1"

	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"
)

func newStatsRefreshButton(data models.DiscordInteraction) discordgo.MessageComponent {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{discordgo.Button{
			Style: discordgo.SecondaryButton,
			Emoji: &discordgo.ComponentEmoji{
				ID:   "1255647885723435048",
				Name: "aftermath_refresh",
			},
			CustomID: fmt.Sprintf("refresh_stats_from_button_%s", data.ReferenceID),
		}},
	}
}

func init() {
	LoadedPublic.add(
		builder.NewCommand("refresh_stats_from_button").
			ComponentType(func(customID string) bool {
				return strings.HasPrefix(customID, "refresh_stats_from_button_")
			}).
			Handler(func(ctx *common.Context) error {
				data, ok := ctx.ComponentData()
				if !ok {
					return ctx.Error("failed to get component data on interaction command")
				}
				interactionID := strings.ReplaceAll(data.CustomID, "refresh_stats_from_button_", "")
				if interactionID == "" {
					return ctx.Error("failed to get interaction id from custom id")
				}

				interaction, err := ctx.Core.Database().GetDiscordInteraction(ctx.Context, interactionID)
				if err != nil {
					return ctx.Reply().Send("stats_refresh_interaction_error_expired")
				}

				var image stats.Image
				var meta stats.Metadata
				switch interaction.Command {
				case "stats":
					img, mt, err := ctx.Core.Stats(ctx.Locale).PeriodImage(context.Background(), interaction.Options.AccountID, interaction.Options.PeriodStart, stats.WithBackgroundURL(interaction.Options.BackgroundImageURL), stats.WithWN8())
					if err != nil {
						return ctx.Err(err)
					}
					image = img
					meta = mt

				case "session":
					img, mt, err := ctx.Core.Stats(ctx.Locale).SessionImage(context.Background(), interaction.Options.AccountID, interaction.Options.PeriodStart, stats.WithBackgroundURL(interaction.Options.BackgroundImageURL), stats.WithWN8())
					if err != nil {
						if errors.Is(err, fetch.ErrSessionNotFound) || errors.Is(err, stats.ErrAccountNotTracked) {
							return ctx.Reply().Send("stats_refresh_interaction_error_expired")
						}
						return ctx.Err(err)
					}
					image = img
					meta = mt

				default:
					log.Error().Str("customId", data.CustomID).Str("command", interaction.Command).Msg("received an unexpected component interaction callback")
					return ctx.Reply().Send("stats_refresh_interaction_error_expired")
				}

				button, saveErr := saveInteractionData(ctx, interaction.Command, interaction.Options)
				if saveErr != nil {
					// nil button will not cause an error and will be ignored
					log.Err(err).Str("interactionId", ctx.ID()).Str("command", "session").Msg("failed to save discord interaction")
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return ctx.Err(err)
				}

				var timings []string
				if ctx.User.HasPermission(permissions.UseDebugFeatures) {
					timings = append(timings, "```")
					for name, duration := range meta.Timings {
						timings = append(timings, fmt.Sprintf("%s: %v", name, duration.Milliseconds()))
					}
					timings = append(timings, "```")
				}

				return ctx.Reply().File(buf.Bytes(), interaction.Command+"_command_by_aftermath.png").Component(button).Text(timings...).Send()
			}),
	)
}
