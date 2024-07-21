package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/gorhill/cronexpr"
)

var euCron = cronexpr.MustParse("0 1 * * *")
var naCron = cronexpr.MustParse("0 9 * * *")
var asiaCron = cronexpr.MustParse("0 18 * * *")
var fancyCron = cronexpr.MustParse("0 0 */7 * *")

func sessionResetTimes(localize func(string) string) string {
	return fmt.Sprintf(localize("commands_help_refresh_times_fmt"),
		naCron.Next(now()).Unix(), naCron.Next(now()).Unix(),
		euCron.Next(now()).Unix(), euCron.Next(now()).Unix(),
		asiaCron.Next(now()).Unix(), asiaCron.Next(now()).Unix())
}

func backgroundResetTime() string {
	return fmt.Sprintf("<t:%d:R>", fancyCron.Next(now()).Unix())
}

func now() time.Time {
	return time.Now().UTC()
}

func sendHelpResponse(ctx *common.Context) error {
	return ctx.Reply().
		Format("commands_help_message_fmt", sessionResetTimes(ctx.Localize), backgroundResetTime()).
		Component(
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Style: discordgo.LinkButton,
						Label: ctx.Localize("buttons_add_aftermath_to_your_server"),
						Emoji: &discordgo.ComponentEmoji{Name: "aftermath", ID: "1214348603104034876"},
						URL:   "https://amth.one/invite",
					},
					discordgo.Button{
						Style: discordgo.LinkButton,
						Label: ctx.Localize("buttons_join_primary_guild"),
						Emoji: &discordgo.ComponentEmoji{Name: "aftermath_yellow", ID: "1214621621659238460"},
						URL:   "https://amth.one/join",
					},
				}},
		).
		Send()
}

func Help() builder.Builder {
	return builder.NewCommand("help").
		Middleware(middleware.RequirePermissions(permissions.UseTextCommands)).
		Ephemeral().
		Handler(func(ctx *common.Context) error {
			return sendHelpResponse(ctx)
		})

}
