package commands

import (
	"encoding/json"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/middleware"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/permissions"
)

func init() {
	Loaded.add(
		builder.NewCommand("manage").
			// ExclusiveToGuilds(os.Getenv("DISCORD_PRIMARY_GUILD_ID")).
			Middleware(middleware.RequirePermissions(permissions.ContentModerator)).
			Options(
				builder.NewOption("users", discordgo.ApplicationCommandOptionSubCommandGroup).Options(
					builder.NewOption("lookup", discordgo.ApplicationCommandOptionSubCommand).Options(
						userOption,
					),
				),
				builder.NewOption("accounts", discordgo.ApplicationCommandOptionSubCommandGroup).Options(
					builder.NewOption("search", discordgo.ApplicationCommandOptionSubCommand).Options(
						nicknameAndServerOptions...,
					),
				),
				builder.NewOption("tasks", discordgo.ApplicationCommandOptionSubCommandGroup).Options(
					builder.NewOption("view", discordgo.ApplicationCommandOptionSubCommand).Options(
						builder.NewOption("status", discordgo.ApplicationCommandOptionString).Choices(
							builder.NewChoice("failed", string(database.TaskStatusFailed)),
							builder.NewChoice("complete", string(database.TaskStatusComplete)),
							builder.NewChoice("in-progress", string(database.TaskStatusInProgress)),
						).Required(),
					),
				),
				builder.NewOption("snapshots", discordgo.ApplicationCommandOptionSubCommandGroup).Options(
					builder.NewOption("view", discordgo.ApplicationCommandOptionSubCommand).Options(
						builder.NewOption("account_id", discordgo.ApplicationCommandOptionString).Required(),
					),
				),
			).
			Handler(func(ctx *common.Context) error {
				command, opts, _ := ctx.Options().Subcommand()

				switch command {
				case "users_lookup":
					userID, _ := opts.Value("user").(string)
					result, err := ctx.Core.Database().GetUserByID(ctx.Context, userID, database.WithConnections())
					if err != nil {
						return ctx.Reply("Database#GetUserByID: " + err.Error())
					}

					data, err := json.MarshalIndent(result, "", "  ")
					if err != nil {
						return ctx.Reply("MarshalIndent: " + err.Error())
					}
					return ctx.Reply("```" + string(data) + "```")

				case "accounts_search":
					nickname, _ := opts.Value("nickname").(string)
					server, _ := opts.Value("server").(string)
					result, err := ctx.Core.Fetch().Search(ctx.Context, nickname, server)
					if err != nil {
						return ctx.Reply("Fetch#Search: " + err.Error())
					}
					data, err := json.MarshalIndent(result, "", "  ")
					if err != nil {
						return ctx.Reply("MarshalIndent: " + err.Error())
					}
					return ctx.Reply("```" + string(data) + "```")

				case "snapshots_view":
					accountId, ok := opts.Value("account_id").(string)
					if !ok {
						return ctx.Reply("invalid accountId, failed to cast to string")
					}
					snapshots, err := ctx.Core.Database().GetLastAccountSnapshots(ctx.Context, accountId, 3)
					if err != nil {
						return ctx.Reply("GetLastAccountSnapshots: " + err.Error())
					}

					var data []map[string]string
					for _, snapshot := range snapshots {
						data = append(data, map[string]string{
							"id":             snapshot.ID,
							"type":           string(snapshot.Type),
							"referenceId":    snapshot.ReferenceID,
							"createdAt":      snapshot.CreatedAt.String(),
							"lastBattleTime": snapshot.LastBattleTime.String(),
							"battlesRating":  snapshot.RatingBattles.Battles.String(),
							"battlesRegular": snapshot.RegularBattles.Battles.String(),
						})
					}

					bytes, err := json.MarshalIndent(data, "", "  ")
					if err != nil {
						return ctx.Reply("json.Marshal: " + err.Error())
					}
					return ctx.Reply("```" + string(bytes) + "```")

				case "tasks_view":
					status, _ := opts.Value("status").(string)
					tasks, err := ctx.Core.Database().GetRecentTasks(ctx.Context, time.Now().Add(time.Hour*24*-1), database.TaskStatus(status))
					if err != nil {
						return ctx.Reply("Database#GetRecentTasks: " + err.Error())
					}
					if len(tasks) < 1 {
						return ctx.Reply("No recent tasks with status " + status)
					}

					bytes, err := json.MarshalIndent(tasks, "", "  ")
					if err != nil {
						return ctx.Reply("json.Marshal: " + err.Error())
					}
					if len(string(bytes)) > 1990 {
						return ctx.ReplyFmt("Too many tasks to show - %d", len(tasks))
					}
					return ctx.Reply("```" + string(bytes) + "```")

				default:
					return ctx.Reply("invalid subcommand, thought this should never happen")
				}
			}),
	)
}
