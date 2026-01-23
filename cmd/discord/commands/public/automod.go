package public

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/log"
)

var AutomodDeleteMessageTimeout = time.Second * 30

func AutomodHandler(helpImage []byte) common.EventHandler[discordgo.MessageCreate] {
	return common.EventHandler[discordgo.MessageCreate]{
		Match: func(db database.Client, s *discordgo.Session, event *discordgo.MessageCreate) bool {
			// only apply to the primary guild
			if event.GuildID != constants.DiscordPrimaryGuildID {
				return false
			}
			// ignore bot messages
			if event.Author == nil || event.Author.Bot {
				return false
			}
			// check if the user is verified
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()

			user, err := db.GetUserByID(ctx, event.Author.ID)
			if err != nil {
				// user doesn't exist in our database, they are not verified
				return true
			}
			return !user.AutomodVerified
		},
		Handle: func(ctx common.Context, event *discordgo.MessageCreate) error {
			// delete the user's message
			err := ctx.Rest().DeleteMessage(ctx.Ctx(), event.ChannelID, event.ID)
			if err != nil {
				log.Err(err).Str("channelID", event.ChannelID).Str("messageID", event.ID).Msg("automod: failed to delete message from unverified user")
				// continue to send explainer even if deletion fails
			}

			// send an explainer message with the help image
			content := fmt.Sprintf(ctx.Localize("automod_unverified_user_message_deleted_fmt"), event.Author.ID)
			msg, err := ctx.Rest().CreateMessage(ctx.Ctx(), event.ChannelID, discordgo.MessageSend{
				Content: content,
			}, []rest.File{{Data: helpImage, Name: "how_to_use_aftermath.png"}})
			if err != nil {
				log.Err(err).Str("channelID", event.ChannelID).Str("userID", event.Author.ID).Msg("automod: failed to send explainer message")
				return nil
			}

			// delete the explainer after a timeout
			go func() {
				time.Sleep(AutomodDeleteMessageTimeout)

				delCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()

				err := ctx.Rest().DeleteMessage(delCtx, event.ChannelID, msg.ID)
				if err != nil {
					log.Err(err).Str("channelID", event.ChannelID).Str("messageID", msg.ID).Msg("automod: failed to delete explainer message")
				}
			}()

			return nil
		},
	}
}
