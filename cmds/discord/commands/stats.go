package commands

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/internal/stats/render/assets"
	"github.com/disgoorg/disgo/discord"
)

func init() {
	Loaded.add(
		cmd().
			Name("stats").
			Option("days", discord.ApplicationCommandOptionInt{
				MaxValue: common.Pointer(90),
				MinValue: common.Pointer(1),
				Required: false,
			},
				SetNameKey("common_option_days_name"),
				SetDescKey("common_option_days_description"),
			).
			Option("user", discord.ApplicationCommandOptionUser{
				Required: false,
			},

				SetNameKey("common_option_user_name"),
				SetDescKey("common_option_user_description"),
			).
			Option("nickname", discord.ApplicationCommandOptionString{
				Required: false,
			},
				SetNameKey("common_option_nickname_name"),
				SetDescKey("common_option_nickname_description")).
			Option("server", discord.ApplicationCommandOptionString{
				Required: false,
				Choices: []discord.ApplicationCommandOptionChoiceString{
					{
						Name:  "North America",
						Value: "na",
					},
					{
						Name:  "Europe",
						Value: "eu",
					},
					{
						Name:  "Asia",
						Value: "as",
					},
				},
			},

				SetNameKey("common_option_server_name"),
				SetDescKey("common_option_server_description"),
			).
			Handler(func(ctx *ctx) error {
				var periodStart time.Time
				if days := time.Duration(ctx.Options().Int("days")); days > 0 {
					periodStart = time.Now().Add(time.Hour * 24 * days * -1)
				}

				fmt.Printf("%#v", ctx.Options().All())

				image, _, err := ctx.Core.Render(ctx.Locale).Period(context.Background(), "579553160", periodStart)
				if err != nil {
					return err
				}

				background, _ := assets.GetLoadedImage("bg-light")
				err = image.AddBackground(background)
				if err != nil {
					return err
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return err
				}

				return ctx.File(&buf, "image_by_aftermath.png")
			}),
	)
}
