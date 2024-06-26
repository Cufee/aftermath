package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
)

func init() {
	Loaded.add(
		builder.NewCommand("link").
			Options(
				builder.NewOption("nickname", discordgo.ApplicationCommandOptionString).
					Required().
					Min(5).
					Max(30).
					Params(
						builder.SetNameKey("common_option_stats_nickname_name"),
						builder.SetDescKey("common_option_stats_nickname_description"),
					),
				builder.NewOption("server", discordgo.ApplicationCommandOptionString).
					Required().
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
				options := getDefaultStatsOptions(ctx)
				message, valid := options.Validate(ctx)
				if !valid {
					return ctx.Reply().Send(message)
				}

				account, err := ctx.Core.Fetch().Search(ctx.Context, options.Nickname, options.Server)
				if err != nil {
					if err.Error() == "no results found" {
						return ctx.Reply().Format("stats_error_nickname_not_fount_fmt", options.Nickname, strings.ToUpper(options.Server)).Send()
					}
					return ctx.Err(err)
				}

				currentConnection, exists := ctx.User.Connection(models.ConnectionTypeWargaming)
				if !exists {
					currentConnection.UserID = ctx.User.ID
					currentConnection.Metadata = make(map[string]any)
					currentConnection.Type = models.ConnectionTypeWargaming
				}

				currentConnection.Metadata["verified"] = false
				currentConnection.ReferenceID = fmt.Sprint(account.ID)

				_, err = ctx.Core.Database().UpsertConnection(ctx.Context, currentConnection)
				if err != nil {
					return ctx.Err(err)
				}

				return ctx.Reply().Format("command_link_linked_successfully_fmt", account.Nickname, strings.ToUpper(options.Server)).Send()
			}),
	)
}
