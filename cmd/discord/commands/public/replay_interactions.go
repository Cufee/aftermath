package public

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/emoji"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	stats "github.com/cufee/aftermath/internal/stats/client/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
	"github.com/cufee/aftermath/internal/utils"
	"golang.org/x/sync/errgroup"
)

var replayReactionEmoji = emoji.AftermathLogoDefault()
var replayReactionEmojiString string = fmt.Sprintf("%s:%s", replayReactionEmoji.Name, replayReactionEmoji.ID)
var replayReactionMissingPermissionsEmoji string = fmt.Sprintf("missing_permissions:%s", constants.MustGetEnv("EMOJI_MISSING_PERMISSIONS_ID"))

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
	Handle: commonErrorsMiddleware(func(ctx common.Context, event *discordgo.MessageReactionAdd) error {
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
	}),
}

func commonErrorsMiddleware(next func(ctx common.Context, event *discordgo.MessageReactionAdd) error) func(ctx common.Context, event *discordgo.MessageReactionAdd) error {
	return func(ctx common.Context, event *discordgo.MessageReactionAdd) error {
		err := next(ctx, event)

		// missing permissions to message in the channel
		if errors.Is(err, rest.ErrMissingPermissions) {
			return ctx.Rest().CreateMessageReaction(ctx.Ctx(), event.ChannelID, event.MessageID, replayReactionMissingPermissionsEmoji)
		}

		return err
	}
}
