package public

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"golang.org/x/text/language"

	stats "github.com/cufee/aftermath/internal/stats/client/v1"

	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"
)

func MentionHandler(errorImage []byte) func(s *discordgo.Session, e *discordgo.MessageCreate) {
	return func(s *discordgo.Session, e *discordgo.MessageCreate) {
		for _, mention := range e.Mentions {
			if mention.ID == s.State.User.ID {
				// Use the user locale selection by default with fallback to English
				locale := language.English
				if mention.Locale != "" {
					locale = common.LocaleToLanguageTag(discordgo.Locale(mention.Locale))
				}

				printer, err := localization.NewPrinter("discord", locale)
				if err != nil {
					log.Err(err).Msg("failed to get a localization printer for context")
					printer = func(s string) string { return s }
				}

				channel, err := s.UserChannelCreate(e.Author.ID)
				if err != nil {
					log.Warn().Str("userId", e.Author.ID).Err(err).Msg("failed to create a DM channel for a user")
					data := discordgo.MessageSend{Files: []*discordgo.File{{Name: "how-to-use-commands.png", Reader: bytes.NewReader(errorImage)}}, Content: fmt.Sprintf(printer("errors_help_missing_dm_permissions_fmt"), e.Author.Mention())}
					_, _ = s.ChannelMessageSendComplex(e.ChannelID, &data)
					return
				}

				_, err = s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{Content: fmt.Sprintf(printer("commands_help_message_fmt"), sessionResetTimes(printer), backgroundResetTime()), Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							common.ButtonInviteAftermath(printer("buttons_add_aftermath_to_your_server")),
							common.ButtonJoinPrimaryGuild(printer("buttons_join_primary_guild")),
						}}}})
				if err != nil {
					log.Warn().Str("userId", e.Author.ID).Err(err).Msg("failed to DM a user")
					data := discordgo.MessageSend{Files: []*discordgo.File{{Name: "how-to-use-commands.png", Reader: bytes.NewReader(errorImage)}}, Content: fmt.Sprintf(printer("errors_help_missing_dm_permissions_fmt"), e.Author.Mention())}
					_, _ = s.ChannelMessageSendComplex(e.ChannelID, &data)
				}
				return
			}
		}
	}
}

func newStatsRefreshButton(data models.DiscordInteraction) discordgo.MessageComponent {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{discordgo.Button{
			Style: discordgo.SecondaryButton,
			Emoji: &discordgo.ComponentEmoji{
				ID:   "1255647885723435048",
				Name: "aftermath_refresh",
			},
			CustomID: fmt.Sprintf("refresh_stats_from_button_%s", data.ID),
		}},
	}
}

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("refresh_stats_from_button").
			Middleware(middleware.RequirePermissions(permissions.UseImageCommands, permissions.UseTextCommands)).
			ComponentType(func(customID string) bool {
				return strings.HasPrefix(customID, "refresh_stats_from_button_")
			}).
			Handler(func(ctx common.Context) error {
				data, ok := ctx.ComponentData()
				if !ok {
					return ctx.Error("failed to get component data on interaction command")
				}
				interactionID := strings.ReplaceAll(data.CustomID, "refresh_stats_from_button_", "")
				if interactionID == "" {
					return ctx.Error("failed to get interaction id from custom id")
				}

				interaction, err := ctx.Core().Database().GetDiscordInteraction(ctx.Ctx(), interactionID)
				if err != nil {
					return ctx.Reply().Send("stats_refresh_interaction_error_expired")
				}

				ioptions, err := statsOptions{}.fromInteraction(interaction)
				if err != nil {
					log.Warn().Err(err).Msg("failed to decode stats options for a refresh button")
					return ctx.Reply().Send("stats_refresh_interaction_error_expired")
				}

				var opts = []stats.RequestOption{stats.WithWN8(), stats.WithVehicleID(ioptions.TankID)}

				if ioptions.BackgroundID != "" {
					background, _ := ctx.Core().Database().GetUserContent(ctx.Ctx(), ioptions.BackgroundID)
					if img, err := logic.UserContentToImage(background); err == nil {
						opts = append(opts, stats.WithBackground(img))
					}
				}

				var image stats.Image
				var meta stats.Metadata
				switch interaction.EventID {
				case "stats":
					img, mt, err := ctx.Core().Stats(ctx.Locale()).PeriodImage(context.Background(), ioptions.AccountID, ioptions.PeriodStart, opts...)
					if err != nil {
						return ctx.Err(err)
					}
					image = img
					meta = mt

				case "session":
					img, mt, err := ctx.Core().Stats(ctx.Locale()).SessionImage(context.Background(), ioptions.AccountID, ioptions.PeriodStart, opts...)
					if err != nil {
						if errors.Is(err, fetch.ErrSessionNotFound) || errors.Is(err, stats.ErrAccountNotTracked) {
							return ctx.Reply().Send("stats_refresh_interaction_error_expired")
						}
						return ctx.Err(err)
					}
					image = img
					meta = mt

				default:
					log.Error().Str("customId", data.CustomID).Str("command", interaction.EventID).Msg("received an unexpected component interaction callback")
					return ctx.Reply().Send("stats_refresh_interaction_error_expired")
				}

				button, saveErr := ioptions.refreshButton(ctx)
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
