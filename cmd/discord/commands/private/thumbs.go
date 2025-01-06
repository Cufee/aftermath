package private

import (
	"time"

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
		builder.NewCommand("thumbs").
			Middleware(middleware.RequirePermissions(permissions.GlobalAdmin)).
			Ephemeral().
			Options(
				builder.NewOption("up", discordgo.ApplicationCommandOptionSubCommand).
					Options(builder.NewOption("user_id", discordgo.ApplicationCommandOptionString)),
				builder.NewOption("clear", discordgo.ApplicationCommandOptionSubCommand).
					Options(builder.NewOption("user_id", discordgo.ApplicationCommandOptionString)),
				builder.NewOption("view", discordgo.ApplicationCommandOptionSubCommand).
					Options(builder.NewOption("user_id", discordgo.ApplicationCommandOptionString)),
			).
			Handler(func(ctx common.Context) error {
				subcommand, subOptions, _ := ctx.Options().Subcommand()
				userID, _ := subOptions.Value("user_id").(string)
				if userID == "" {
					return ctx.Error("invalid user id")
				}

				target, err := ctx.Core().Database().GetUserByID(ctx.Ctx(), userID, database.WithSubscriptions())
				if err != nil {
					return ctx.Error("user not found: " + err.Error())
				}

				subscription, ok := target.Subscription(models.SubscriptionTypeThumbsCounter)
				if !ok {
					subscription, err = ctx.Core().Database().CreateUserSubscription(ctx.Ctx(), models.UserSubscription{
						UserID:      target.ID,
						Type:        models.SubscriptionTypeThumbsCounter,
						ExpiresAt:   time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC),
						ReferenceID: target.ID,
						Meta:        map[string]any{"count": float64(0)},
					})
					if err != nil {
						return ctx.Error("failed to create a subscription: " + err.Error())
					}
				}

				switch subcommand {
				default:
					return ctx.Error("invalid subcommand: " + subcommand)
				case "up":
					count, _ := subscription.Meta["count"].(float64)
					count++
					subscription.Meta["count"] = count

				case "clear":
					subscription.Meta["count"] = float64(0)

				case "view":
					return ctx.Reply().Format("user <@%s> has %v thumbs-up", target.ID, subscription.Meta["count"]).Send()
				}

				_, err = ctx.Core().Database().UpdateUserSubscription(ctx.Ctx(), subscription.ID, subscription)
				if err != nil {
					return ctx.Error("failed to update a subscription: " + err.Error())
				}
				return ctx.Reply().Send("updated user thumbs-up subscription")
			}),
	)
}
