package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/middleware"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
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
							builder.NewChoice("failed", string(models.TaskStatusFailed)),
							builder.NewChoice("complete", string(models.TaskStatusComplete)),
							builder.NewChoice("scheduled", string(models.TaskStatusScheduled)),
							builder.NewChoice("in-progress", string(models.TaskStatusInProgress)),
						).Required(),
						builder.NewOption("hours", discordgo.ApplicationCommandOptionNumber).Required(),
					),
					builder.NewOption("details", discordgo.ApplicationCommandOptionSubCommand).Options(
						builder.NewOption("id", discordgo.ApplicationCommandOptionString).Required(),
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
					hours, _ := opts.Value("hours").(float64)
					status, _ := opts.Value("status").(string)
					if hours < 1 {
						hours = 1
					}

					tasks, err := ctx.Core.Database().GetRecentTasks(ctx.Context, time.Now().Add(time.Hour*time.Duration(hours)*-1), models.TaskStatus(status))
					if err != nil {
						return ctx.Reply("Database#GetRecentTasks: " + err.Error())
					}
					if len(tasks) < 1 {
						return ctx.Reply("No recent tasks with status " + status)
					}

					ids := []string{fmt.Sprintf("total: %d", len(tasks))}
					for _, t := range tasks {
						ids = append(ids, t.ID)
					}
					return ctx.File(bytes.NewBufferString(strings.Join(ids, "\n")), "tasks.txt")

				case "tasks_details":
					id, _ := opts.Value("id").(string)
					if id == "" {
						return ctx.Reply("id cannot be blank")
					}

					tasks, err := ctx.Core.Database().GetTasks(ctx.Context, id)
					if err != nil {
						return ctx.Reply("Database#GetTasks: " + err.Error())
					}
					if len(tasks) < 1 {
						return ctx.Reply("No recent task found")
					}

					data, err := json.MarshalIndent(tasks, "", "  ")
					if err != nil {
						return ctx.Reply("json.Marshal: " + err.Error())
					}
					return ctx.File(bytes.NewReader(data), "tasks.txt")

				default:
					return ctx.Reply("invalid subcommand, thought this should never happen")
				}
			}),
	)
}
