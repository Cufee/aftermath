package public

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/external/blitzkit"
	"github.com/cufee/aftermath/internal/glossary"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"

	stats "github.com/cufee/aftermath/internal/stats/client/common"

	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("refresh_stats_from_button").
			Middleware(middleware.RequirePermissions(permissions.UseImageCommands, permissions.UseTextCommands)).
			ComponentType(func(customID string) bool {
				return strings.HasPrefix(customID, "refresh_stats_from_button#") ||
					// legacy buttons
					strings.HasPrefix(customID, "refresh_stats_from_button_") ||
					strings.HasPrefix(customID, "session_refresh_")
			}).
			Handler(func(ctx common.Context) error {
				data, ok := ctx.ComponentData()
				if !ok {
					log.Error().Msg("failed to get component data on interaction command")
					return ctx.Reply().IsError(common.UserError).Send("stats_refresh_interaction_error_expired")
				}
				interactionID := strings.ReplaceAll(data.CustomID, "refresh_stats_from_button#", "")
				if interactionID == "" {
					log.Error().Str("id", data.CustomID).Msg("failed to get interaction id from custom id")
					return ctx.Reply().IsError(common.UserError).Send("stats_refresh_interaction_error_expired")
				}

				interaction, err := ctx.Core().Database().GetDiscordInteraction(ctx.Ctx(), interactionID)
				if err != nil {
					return ctx.Reply().IsError(common.UserError).Send("stats_refresh_interaction_error_expired")
				}

				ioptions, err := statsOptions{}.fromInteraction(interaction)
				if err != nil {
					log.Warn().Err(err).Msg("failed to decode stats options for a refresh button")
					return ctx.Reply().IsError(common.UserError).Send("stats_refresh_interaction_error_expired")
				}

				var opts = []stats.RequestOption{stats.WithWN8()}
				if ioptions.TankID != "" {
					opts = append(opts, stats.WithVehicleIDs(ioptions.TankID))
				}
				if ioptions.TankTier > 0 && ioptions.TankTier <= 10 {
					ids, ok := glossary.TierVehicleIDs(ioptions.TankTier)
					if ok {
						opts = append(opts, stats.WithVehicleIDs(ids...), stats.WithFooterText(ctx.Localize(fmt.Sprintf("common_label_tier_%d", ioptions.TankTier))))
					}
				}

				if ioptions.BackgroundID != "" {
					background, _ := ctx.Core().Database().GetUserContent(ctx.Ctx(), ioptions.BackgroundID)
					if img, err := logic.UserContentToImage(background); err == nil {
						opts = append(opts, stats.WithBackground(img, true))
					}
				}

				var image stats.Image
				var meta stats.Metadata
				switch interaction.EventID {
				case "career", "stats":
					img, mt, err := ctx.Core().Stats(ctx.Locale()).PeriodImage(ctx.Ctx(), ioptions.AccountID, ioptions.PeriodStart, opts...)
					if errors.Is(err, blitzkit.ErrServiceUnavailable) {
						return ctx.Reply().
							Hint(ctx.InteractionID()).
							IsError(common.ApplicationError).
							Component(discordgo.ActionsRow{Components: []discordgo.MessageComponent{common.ButtonJoinPrimaryGuild(ctx.Localize("buttons_have_a_question_question"))}}).
							Send("blitz_stars_error_service_down")
					}
					if err != nil {
						return ctx.Err(err, common.ApplicationError)
					}
					image = img
					meta = mt

				case "session":
					img, mt, err := ctx.Core().Stats(ctx.Locale()).SessionImage(ctx.Ctx(), ioptions.AccountID, ioptions.PeriodStart, opts...)
					if err != nil {
						if errors.Is(err, stats.ErrAccountNotTracked) || (errors.Is(err, fetch.ErrSessionNotFound) && ioptions.Days < 1) {
							return ctx.Reply().IsError(common.UserError).Send("session_error_account_was_not_tracked")
						}
						if errors.Is(err, fetch.ErrSessionNotFound) {
							return ctx.Reply().IsError(common.UserError).Send("session_error_no_session_for_period")
						}
						return ctx.Err(err, common.ApplicationError)
					}
					image = img
					meta = mt

				default:
					log.Error().Str("customId", data.CustomID).Str("command", interaction.EventID).Msg("received an unexpected component interaction callback")
					return ctx.Reply().IsError(common.UserError).Send("stats_refresh_interaction_error_expired")
				}

				includeSession := interaction.EventID != "session" && interaction.EventID != "replay"
				button, saveErr := ioptions.actionRow(ctx, interaction.EventID, includeSession)
				if saveErr != nil {
					// nil button will not cause an error and will be ignored
					log.Err(saveErr).Str("interactionId", ctx.ID()).Str("command", interaction.EventID).Msg("failed to save discord interaction")
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return ctx.Err(err, common.ApplicationError)
				}

				var timings []string
				if ctx.User().HasPermission(permissions.UseDebugFeatures) {
					timings = append(timings, "```")
					for name, duration := range meta.Timings {
						timings = append(timings, fmt.Sprintf("%s: %v", name, duration.Milliseconds()))
					}
					timings = append(timings, "```")
				}

				return ctx.Reply().WithAds().File(buf.Bytes(), interaction.EventID+"_command_by_aftermath.png").Component(button).Text(timings...).Send()
			}),
	)
}
