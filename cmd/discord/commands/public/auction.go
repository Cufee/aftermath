package public

import (
	"errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/seasonal/auction"
	"github.com/cufee/aftermath/internal/utils"
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

				cards, data, err := auction.CurrentAuctionImages(ctx.Ctx(), realm)
				if err != nil {
					if errors.Is(err, auction.ErrRealmNotCached) {
						return ctx.Reply().Send("error_command_auction_still_refreshing")
					}
					return ctx.Err(err)
				}

				batches := utils.Batch(cards, 10)
				for i, group := range batches {
					reply := ctx.Reply()
					for i, img := range group {
						reply = reply.File(img, fmt.Sprintf("%02d_auction_by_amth.one.png", i))
					}

					var err error
					if i == 0 {
						err = reply.Format("command_auction_result_fmt", data.LastUpdate.Unix()).Send()
					} else {
						_, err = reply.Followup()
					}
					if err != nil {
						return ctx.Err(err)
					}
				}

				return nil
			}),
	)
}
