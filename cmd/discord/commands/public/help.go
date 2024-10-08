package public

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

func sessionResetTimes(localize func(string) string) string {
	return fmt.Sprintf(localize("commands_help_refresh_times_fmt"),
		naCron.Next(now()).Unix(), naCron.Next(now()).Unix(),
		euCron.Next(now()).Unix(), euCron.Next(now()).Unix(),
		asiaCron.Next(now()).Unix(), asiaCron.Next(now()).Unix())
}

func now() time.Time {
	return time.Now().UTC()
}

func sendHelpResponse(ctx common.Context) error {
	return reply(ctx).Send()
}

func reply(ctx common.Context) common.Reply {
	return ctx.Reply().Format("commands_help_message_fmt", sessionResetTimes(ctx.Localize)).
		Component(
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					common.ButtonInviteAftermath(ctx.Localize("buttons_add_aftermath_to_your_server")),
					common.ButtonJoinPrimaryGuild(ctx.Localize("buttons_join_primary_guild")),
				}},
		)
}

func Help() builder.Builder {
	return builder.NewCommand("help").
		Ephemeral().
		Middleware(middleware.RequirePermissions(permissions.UseTextCommands)).
		Handler(func(ctx common.Context) error {
			return sendHelpResponse(ctx)
		})

}
