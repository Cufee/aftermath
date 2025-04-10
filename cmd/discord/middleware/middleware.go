package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/cta"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/permissions"
)

var adCooldownPeriod = time.Second * 6

// var adCooldownPeriod = time.Hour * 6

type MiddlewareFunc func(common.Context, func(common.Context) error) func(common.Context) error

func RequirePermissions(required ...permissions.Permissions) MiddlewareFunc {
	return func(ctx common.Context, next func(common.Context) error) func(common.Context) error {
		if !ctx.User().HasPermission(required...) {
			return func(ctx common.Context) error {
				return ctx.Reply().Send("common_error_command_missing_permissions")
			}
		}
		return next
	}
}

func ExpireAfter(expiration time.Time) MiddlewareFunc {
	return func(ctx common.Context, next func(common.Context) error) func(common.Context) error {
		if time.Now().After(expiration) {
			return func(ctx common.Context) error {
				return ctx.Reply().Format("common_error_command_expired_fmt", expiration.Unix()).Send()
			}
		}
		return next
	}
}

func ServeAds() MiddlewareFunc {
	return func(ctx common.Context, next func(common.Context) error) func(common.Context) error {
		return func(ctx common.Context) error {
			err := next(ctx)
			if err != nil {
				return err
			}

			guildID := ctx.GuildID()
			channelID := ctx.ChannelID()

			cctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			// check if this channel had a recent ad
			lastRun, err := ctx.Core().Database().GetChannelLastAdRun(cctx, channelID)
			if err != nil && !database.IsNotFound(err) {
				log.Err(err).Msg("failed to get channel last ad run")
				return nil
			}
			if time.Since(lastRun) < adCooldownPeriod {
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
			embed, buttons, ok := cta.DefaultCTACollection.NewEmbed(ctx.Locale(), tags)
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

			go func(core core.Client, reply common.Reply, data models.DiscordAdRun) {
				// allow the user time to interact with the application and view the command output
				time.Sleep(time.Minute * 1)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
				defer cancel()

				// check last ad run time again
				lastRun, err := core.Database().GetChannelLastAdRun(ctx, data.ChannelID)
				if err != nil && !database.IsNotFound(err) {
					log.Err(err).Msg("failed to get channel last ad run")
					return
				}
				if time.Since(lastRun) < adCooldownPeriod {
					return
				}

				msg, err := reply.Followup(ctx)
				if err != nil {
					log.Err(err).Msg("failed to send an interaction ad followup")
					return
				}
				data.MessageID = msg.ID

				_, err = core.Database().CreateDiscordAdRun(ctx, data)
				if err != nil {
					log.Err(err).Msg("failed to save an ad run record")
				}
			}(ctx.Core(), reply, data)

			return nil
		}
	}
}
