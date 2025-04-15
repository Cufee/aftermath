package public

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/emoji"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/glossary"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
	"github.com/cufee/aftermath/internal/utils"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/language"

	stats "github.com/cufee/aftermath/internal/stats/client/common"

	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"
)

func MentionHandler(errorImage []byte) common.EventHandler[discordgo.MessageCreate] {
	return common.EventHandler[discordgo.MessageCreate]{
		Match: func(c database.Client, s *discordgo.Session, event *discordgo.MessageCreate) bool {
			return len(event.Mentions) > 0 && // has mentions
				event.Type != discordgo.MessageTypeReply && // not a reply
				event.Author.ID != constants.DiscordBotUserID && // not from self
				len(event.Content) < len(constants.DiscordBotUserID)+4 // only includes the mention text
		},
		Handle: func(ctx common.Context, event *discordgo.MessageCreate) error {
			for _, mention := range event.Mentions {
				if mention.ID == constants.DiscordBotUserID {
					// Use the user locale selection by default with fallback to English
					locale := language.English
					if mention.Locale != "" {
						locale = common.LocaleToLanguageTag(discordgo.Locale(mention.Locale))
					}

					printer, err := localization.NewPrinterWithFallback("discord", locale, language.English)
					if err != nil {
						log.Err(err).Msg("failed to get a localization printer for context")
						printer = func(s string) string { return s }
					}

					channel, err := ctx.Rest().CreateDMChannel(ctx.Ctx(), ctx.User().ID)
					if err != nil {
						log.Warn().Str("userId", ctx.User().ID).Err(err).Msg("failed to create a DM channel for a user")
						data := discordgo.MessageSend{Content: fmt.Sprintf(printer("errors_help_missing_dm_permissions_fmt"), "<@"+ctx.User().ID+">"), Flags: discordgo.MessageFlagsEphemeral}
						_, err = ctx.Rest().CreateMessage(ctx.Ctx(), event.ChannelID, data, []rest.File{{Data: errorImage, Name: "how_to_use_aftermath.png"}})
						if err != nil {
							log.Err(err).Msg("failed to send a channel message")
						}
						return nil
					}

					_, err = ctx.Rest().CreateMessage(
						ctx.Ctx(),
						channel.ID,
						discordgo.MessageSend{Content: fmt.Sprintf(printer("commands_help_message_fmt"), sessionResetTimes(printer)), Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									common.ButtonInviteAftermath(printer("buttons_add_aftermath_to_your_server")),
									common.ButtonJoinPrimaryGuild(printer("buttons_join_primary_guild")),
								},
							}}},
						nil)

					if err != nil {
						log.Warn().Str("userId", ctx.User().ID).Err(err).Msg("failed to DM a user")
						data := discordgo.MessageSend{Content: fmt.Sprintf(printer("errors_help_missing_dm_permissions_fmt"), "<@"+ctx.User().ID+">"), Flags: discordgo.MessageFlagsEphemeral}
						_, err = ctx.Rest().CreateMessage(ctx.Ctx(), event.ChannelID, data, []rest.File{{Data: errorImage, Name: "how_to_use_aftermath.png"}})
						if err != nil {
							log.Err(err).Msg("failed to send a channel message")
						}
					}
					return nil
				}
			}
			return nil
		},
	}
}

var replayReactionEmoji = emoji.AftermathLogoDefault()
var replayReactionEmojiString string = fmt.Sprintf("%s:%s", replayReactionEmoji.Name, replayReactionEmoji.ID)

var ReplayFileHandler = common.EventHandler[discordgo.MessageCreate]{
	Match: func(db database.Client, s *discordgo.Session, event *discordgo.MessageCreate) bool {
		for _, file := range event.Attachments {
			if !strings.Contains(file.Filename, ".wotbreplay") {
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			prev, _ := db.FindDiscordInteractions(ctx, database.WithChannelID(event.ChannelID), database.WithType(models.InteractionTypeGatewayEvent))
			for _, record := range prev {
				messageID, _ := record.Meta["eventMessageId"].(string)
				if messageID == event.ID {
					return false // We already worked on this event
				}
			}
			return true
		}
		return false
	},
	Handle: func(ctx common.Context, event *discordgo.MessageCreate) error {
		err := ctx.Rest().CreateMessageReaction(ctx.Ctx(), event.ChannelID, event.ID, replayReactionEmojiString)
		if err != nil && !errors.Is(err, rest.ErrMissingPermissions) && !errors.Is(err, rest.ErrMissingUserUnreachable) {
			return err
		}
		return nil
	},
}

type imageWithIndex struct {
	stats.Image
	index int
}

var ReplayInteractionHandler = common.EventHandler[discordgo.MessageReactionAdd]{
	Match: func(db database.Client, s *discordgo.Session, event *discordgo.MessageReactionAdd) bool {
		if event.Emoji.ID != replayReactionEmoji.ID || event.UserID == constants.DiscordBotUserID {
			return false
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		prev, _ := db.FindDiscordInteractions(ctx, database.WithChannelID(event.ChannelID), database.WithType(models.InteractionTypeGatewayEvent))
		for _, record := range prev {
			messageID, _ := record.Meta["eventMessageId"].(string)
			if messageID == event.MessageID {
				return false // We already worked on this event
			}
		}
		return true
	},
	Handle: func(ctx common.Context, event *discordgo.MessageReactionAdd) error {
		message, err := ctx.Rest().GetMessage(ctx.Ctx(), event.ChannelID, event.MessageID)
		if err != nil {
			if errors.Is(err, rest.ErrMissingPermissions) {
				return ctx.Reply().IsError(common.UserError).Send("replay_errors_missing_permissions_vague")
			}
			log.Err(err).Msg("failed to get channel message")
			return nil // this is a silent error for the user
		}

		var hasReaction bool
		for _, reaction := range message.Reactions {
			if reaction.Me {
				hasReaction = true
				break
			}
		}
		if !hasReaction {
			return nil
		}

		err = ctx.Rest().DeleteOwnMessageReaction(ctx.Ctx(), event.ChannelID, event.MessageID, replayReactionEmojiString)
		if err != nil && !errors.Is(err, rest.ErrMissingPermissions) {
			log.Err(err).Msg("failed to delete a message reaction")
			return nil // this is a silent error for the user
		}

		var attachments []string
		attachmentHashes := make(map[string]struct{})
		for _, attachment := range message.Attachments {
			name := strings.ToLower(attachment.Filename)
			if !strings.Contains(name, ".wotbreplay") {
				continue
			}

			hash := fmt.Sprintf("%s-%d", name, attachment.Size)
			if _, ok := attachmentHashes[hash]; ok {
				continue
			}
			attachmentHashes[hash] = struct{}{}

			attachments = append(attachments, attachment.URL)
		}

		images := make(chan imageWithIndex, len(attachments))
		var group errgroup.Group

		// only process up to 5 files
		for i, url := range utils.Batch(attachments, 5)[0] {
			group.Go(func() error {
				c, cancel := context.WithTimeout(context.Background(), time.Second*15)
				defer cancel()

				image, _, err := ctx.Core().Stats(ctx.Locale()).ReplayImage(c, url, stats.WithWN8())
				if err != nil {
					return err
				}
				images <- imageWithIndex{image, i}
				return nil
			})
		}
		err = group.Wait()
		close(images)

		if errors.Is(err, replay.ErrInvalidReplayFile) || len(images) == 0 {
			return ctx.Reply().IsError(common.UserError).Send("replay_errors_all_attachments_invalid")
		}

		reply := ctx.Reply()
		if len(attachments) > 5 {
			reply = reply.Hint("replay_errors_too_many_files")
		}
		if err != nil {
			reply = reply.Hint("replay_errors_some_attachments_invalid")
		}

		var files = make([]rest.File, len(images))
		var filesMx sync.Mutex

		for img := range images {
			group.Go(func() error {
				var buf bytes.Buffer
				file := rest.File{Name: fmt.Sprintf("replay_command_by_aftermath_%d.png", img.index)}
				err := img.PNG(&buf)
				if err != nil {
					return err
				}
				file.Data = buf.Bytes()

				filesMx.Lock()
				files[img.index] = file
				defer filesMx.Unlock()
				return nil
			})
		}
		err = group.Wait()
		if err != nil {
			return ctx.Reply().IsError(common.UserError).Send("replay_errors_all_attachments_invalid")
		}

		for _, file := range files {
			if file.Data == nil {
				continue
			}
			reply = reply.File(file.Data, file.Name)
		}
		return reply.WithAds().Reference(event.MessageID, event.ChannelID, event.GuildID).Format("-# <@%s> used a replay reaction", ctx.User().ID).Send()
	},
}

func newStatsRefreshButton(data models.DiscordInteraction) discordgo.MessageComponent {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{discordgo.Button{
			Style: discordgo.SecondaryButton,
			Emoji: &discordgo.ComponentEmoji{
				ID:   "1255647885723435048",
				Name: "aftermath_refresh",
			},
			CustomID: fmt.Sprintf("refresh_stats_from_button#%s", data.ID),
		}},
	}
}

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
					img, mt, err := ctx.Core().Stats(ctx.Locale()).PeriodImage(context.Background(), ioptions.AccountID, ioptions.PeriodStart, opts...)
					if errors.Is(err, blitzstars.ErrServiceUnavailable) {
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
					img, mt, err := ctx.Core().Stats(ctx.Locale()).SessionImage(context.Background(), ioptions.AccountID, ioptions.PeriodStart, opts...)
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

				button, saveErr := ioptions.refreshButton(ctx, interaction.EventID)
				if saveErr != nil {
					// nil button will not cause an error and will be ignored
					log.Err(err).Str("interactionId", ctx.ID()).Str("command", "session").Msg("failed to save discord interaction")
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
