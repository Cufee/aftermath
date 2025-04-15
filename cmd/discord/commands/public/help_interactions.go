package public

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/log"
	"golang.org/x/text/language"
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
