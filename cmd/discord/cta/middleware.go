package cta

import (
	"context"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/guregu/null/v6"
)

var DefaultCooldownDM = time.Hour * 24 * 7
var DefaultCooldownServer = time.Hour * 24 * 3

var followUpResponseDelay = time.Second * 15

func Middleware(cooldownServer, cooldownDM time.Duration) middleware.MiddlewareFunc {
	return func(ctx common.Context, next func(common.Context) error) func(common.Context) error {
		return func(ctx common.Context) error {
			err := next(ctx)
			if err != nil {
				return err
			}
			if !ctx.ShowAds() || ctx.HasError() {
				return nil
			}

			guildID := ctx.GuildID()
			channelID := ctx.ChannelID()

			cooldown := cooldownDM
			if guildID.Valid {
				cooldown = cooldownServer
			}

			cctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			// check if this channel had a recent ad
			lastRun, err := ctx.Core().Database().GetChannelLastAdRun(cctx, channelID)
			if err != nil && !database.IsNotFound(err) {
				log.Err(err).Msg("failed to get channel last ad run")
				return nil
			}
			if time.Since(lastRun) < cooldown {
				return nil
			}

			// check if a user should see an ad
			var tags []string
			if guildID.Valid {
				tags = append(tags, "guild")
			}
			if _, isCommand := ctx.CommandData(); isCommand {
				tags = append(tags, "command_"+ctx.ID())
			}
			if strings.HasPrefix(ctx.ID(), "refresh_stats_button") {
				tags = append(tags, "command_stats", "command_session")
			}
			embed, buttons, ok := defaultCTACollection.NewEmbed(ctx.Locale(), tags)
			if !ok {
				return nil
			}

			// make sure we do not server more ads later
			ctx.WithAds(false)

			reply := ctx.Reply().Embed(&embed).Component(buttons...)
			data := models.DiscordAdRun{
				CampaignID: "cta",
				ContentID:  ctx.ID(),
				GuildID:    guildID,
				ChannelID:  channelID,
				Locale:     ctx.Locale(),
				Tags:       tags,
			}

			data, err = ctx.Core().Database().CreateDiscordAdRun(cctx, data)
			if err != nil {
				log.Err(err).Msg("failed to save an ad run record")
				return nil
			}

			go func(core core.Client, reply common.Reply, data models.DiscordAdRun) {
				// allow the user time to interact with the application and view the command output
				time.Sleep(followUpResponseDelay)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
				defer cancel()

				msg, err := reply.Followup(ctx)
				if err != nil {
					log.Err(err).Msg("failed to send an interaction ad followup")
					return
				}
				data.MessageID = null.StringFrom(msg.ID)

				_, err = core.Database().UpdateDiscordAdRun(ctx, data)
				if err != nil {
					log.Err(err).Msg("failed to save an ad run record")
				}
			}(ctx.Core(), reply, data)

			return nil
		}
	}
}
