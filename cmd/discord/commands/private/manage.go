package private

import (
	"fmt"
	"time"

	"github.com/cufee/aftermath/internal/json"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
)

func init() {
	commands.LoadedInternal.Add(
		builder.NewCommand("manage").
			Middleware(middleware.RequirePermissions(permissions.UseDebugFeatures)).
			Options(
				builder.NewOption("users", discordgo.ApplicationCommandOptionSubCommandGroup).Options(
					builder.NewOption("lookup", discordgo.ApplicationCommandOptionSubCommand).Options(
						commands.UserOption,
					),
				),
				builder.NewOption("accounts", discordgo.ApplicationCommandOptionSubCommandGroup).Options(
					builder.NewOption("search", discordgo.ApplicationCommandOptionSubCommand).Options(
						commands.NicknameOption,
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
			).
			Handler(func(ctx common.Context) error {
				command, opts, _ := ctx.Options().Subcommand()

				switch command {
				case "users_lookup":
					userID, _ := opts.Value("user").(string)
					result, err := ctx.Core().Database().GetUserByID(ctx.Ctx(), userID, database.WithConnections())
					if err != nil {
						return ctx.Reply().Send("Database#GetUserByID: " + err.Error())
					}

					data, err := json.MarshalIndent(result, "", "  ")
					if err != nil {
						return ctx.Reply().Send("MarshalIndent: " + err.Error())
					}
					return ctx.Reply().Send("```" + string(data) + "```")

				case "accounts_search":
					opts := commands.GetDefaultStatsOptions(ctx.Options())
					if opts.AccountID == "" {
						return ctx.Reply().Send("invalid account id")
					}

					result, err := ctx.Core().Fetch().Account(ctx.Ctx(), opts.AccountID)
					if err != nil {
						return ctx.Reply().Send("Fetch#Search: " + err.Error())
					}
					data, err := json.MarshalIndent(result, "", "  ")
					if err != nil {
						return ctx.Reply().Send("MarshalIndent: " + err.Error())
					}
					return ctx.Reply().Send("```" + string(data) + "```")

				case "tasks_view":
					if !ctx.User().HasPermission(permissions.ViewTaskLogs) {
						ctx.Reply().Send("You do not have access to this sub-command.")
					}

					hours, _ := opts.Value("hours").(float64)
					status, _ := opts.Value("status").(string)
					if hours < 1 {
						hours = 1
					}

					tasks, err := ctx.Core().Database().GetRecentTasks(ctx.Ctx(), time.Now().Add(time.Hour*time.Duration(hours)*-1), models.TaskStatus(status))
					if err != nil {
						return ctx.Reply().Send("Database#GetRecentTasks: " + err.Error())
					}
					if len(tasks) < 1 {
						return ctx.Reply().Send("No recent tasks with status " + status)
					}

					content := fmt.Sprintf("total: %d\n", len(tasks))
					var reduced []map[string]any
					for _, t := range tasks {
						reduced = append(reduced, map[string]any{
							"id":          t.ID,
							"type":        t.Type,
							"targets":     len(t.Targets),
							"referenceID": t.ReferenceID,
							"lastRun":     t.LastRun,
						})
					}
					data, err := json.MarshalIndent(reduced, "", "  ")
					if err != nil {
						return ctx.Reply().Send("json.Marshal: " + err.Error())
					}
					content += string(data)

					return ctx.Reply().File([]byte(content), "tasks.json").Send()

				case "tasks_details":
					if !ctx.User().HasPermission(permissions.ViewTaskLogs) {
						ctx.Reply().Send("You do not have access to this sub-command.")
					}

					id, _ := opts.Value("id").(string)
					if id == "" {
						return ctx.Reply().Send("id cannot be blank")
					}

					tasks, err := ctx.Core().Database().GetTasks(ctx.Ctx(), id)
					if err != nil {
						return ctx.Reply().Send("Database#GetTasks: " + err.Error())
					}
					if len(tasks) < 1 {
						return ctx.Reply().Send("No recent task found")
					}

					data, err := json.MarshalIndent(tasks, "", "  ")
					if err != nil {
						return ctx.Reply().Send("json.Marshal: " + err.Error())
					}
					return ctx.Reply().File(data, "tasks.json").Send()

				default:
					return ctx.Reply().Send("invalid subcommand, thought this should never happen")
				}
			}),
	)
}
