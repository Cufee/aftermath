package public

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/seasonal/auction"
)

var auctionCommandExpiration = time.Date(2025, 1, 11, 0, 0, 0, 0, time.UTC)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("auction").
			Middleware(middleware.ExpireAfter(auctionCommandExpiration)).
			Params(builder.SetNameKey("command_auction_name"), builder.SetDescKey("command_auction_description")).
			Options(
				builder.NewOption("server", discordgo.ApplicationCommandOptionString).
					Required().
					Params(
						builder.SetNameKey("common_option_stats_realm_name"),
						builder.SetDescKey("common_option_stats_realm_description"),
					).
					Choices(
						builder.NewChoice("realm_na", "NA").Params(builder.SetNameKey("common_label_realm_na")),
						builder.NewChoice("realm_eu", "EU").Params(builder.SetNameKey("common_label_realm_eu")),
						builder.NewChoice("realm_as", "AS").Params(builder.SetNameKey("common_label_realm_as")),
					),
			).
			Handler(func(ctx common.Context) error {
				server, _ := ctx.Options().Value("server").(string)
				realm, _ := ctx.Core().Wargaming().ParseRealm(server)

				data, err := auction.CurrentAuction(ctx.Ctx(), realm)
				if err != nil {
					return ctx.Err(err)
				}

				var message string
				for _, vehicle := range data.Vehicles {
					message += fmt.Sprintf("%s %d <:wg_gold:1317571455097110629>\n", vehicle.Name(ctx.Locale()), vehicle.Price.Current.Value)
				}

				return ctx.Reply().Send(message)
			}),
	)
}
