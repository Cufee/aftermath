package commands

import (
	"bytes"
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/internal/stats/render"
	"github.com/cufee/aftermath/internal/stats/render/assets"
)

func init() {
	Loaded.add(
		builder.NewCommand("stats").
			Options(
				builder.NewOption("days", discordgo.ApplicationCommandOptionInteger).
					Min(1).
					Max(90).
					Params(
						builder.SetNameKey("common_option_stats_days_name"),
						builder.SetDescKey("common_option_stats_days_description"),
					),
				builder.NewOption("user", discordgo.ApplicationCommandOptionUser).
					Params(
						builder.SetNameKey("common_option_stats_user_name"),
						builder.SetDescKey("common_option_stats_user_description"),
					),
				builder.NewOption("nickname", discordgo.ApplicationCommandOptionString).
					Min(5).
					Max(30).
					Params(
						builder.SetNameKey("common_option_stats_nickname_name"),
						builder.SetDescKey("common_option_stats_nickname_description"),
					),
				builder.NewOption("server", discordgo.ApplicationCommandOptionString).
					Params(
						builder.SetNameKey("common_option_stats_realm_name"),
						builder.SetDescKey("common_option_stats_realm_description"),
					).
					Choices(
						builder.NewChoice("realm_na", "na").Params(builder.SetNameKey("common_label_realm_na")),
						builder.NewChoice("realm_eu", "eu").Params(builder.SetNameKey("common_label_realm_eu")),
						builder.NewChoice("realm_as", "as").Params(builder.SetNameKey("common_label_realm_as")),
					),
			).
			Handler(func(ctx *common.Context) error {
				var periodStart time.Time
				if days, _ := common.GetOption[float64](ctx, "days"); days > 0 {
					periodStart = time.Now().Add(time.Hour * 24 * time.Duration(days) * -1)
				}

				background, _ := assets.GetLoadedImage("bg-default")
				image, _, err := ctx.Core.Render(ctx.Locale).Period(context.Background(), "1013072123", periodStart, render.WithBackground(background))
				if err != nil {
					return ctx.Err(err, "some_error_message_key")
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return ctx.Err(err, "some_error_message_key")
				}

				return ctx.File(&buf, "image_by_aftermath.png")
			}),
	)
}
