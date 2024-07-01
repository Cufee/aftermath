package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
)

func init() {
	LoadedPublic.add(
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
				var wgConnections []models.UserConnection
				for _, conn := range ctx.User.Connections {
					if conn.Type != models.ConnectionTypeWargaming {
						continue
					}
					wgConnections = append(wgConnections, conn)
				}
				if len(wgConnections) >= 3 {
					return ctx.Reply().Send("link_error_too_many_connections")
				}

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

				var found bool
				for _, conn := range wgConnections {
					if conn.ReferenceID == fmt.Sprint(account.ID) {
						conn.Metadata["verified"] = false
						found = true
					}
					conn.Metadata["default"] = conn.ReferenceID == fmt.Sprint(account.ID)

					_, err = ctx.Core.Database().UpsertConnection(ctx.Context, conn)
					if err != nil {
						return ctx.Err(err)
					}
				}
				if !found {
					meta := make(map[string]any)
					meta["verified"] = false
					meta["default"] = true
					_, err = ctx.Core.Database().UpsertConnection(ctx.Context, models.UserConnection{
						Type:        models.ConnectionTypeWargaming,
						ReferenceID: fmt.Sprint(account.ID),
						UserID:      ctx.User.ID,
						Metadata:    meta,
					})
					if err != nil {
						return ctx.Err(err)
					}
				}

				go func(id string) {
					c, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()
					_, _ = ctx.Core.Fetch().Account(c, id) // Make sure the account is cached
				}(fmt.Sprint(account.ID))

				return ctx.Reply().Format("command_link_linked_successfully_fmt", account.Nickname, strings.ToUpper(options.Server)).Send()
			}),
	)
}
