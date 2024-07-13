package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database/models"
)

func init() {
	LoadedPublic.add(
		builder.NewCommand("widget").
			Ephemeral().
			Options(
				builder.NewOption("account", discordgo.ApplicationCommandOptionString).
					Params(builder.SetNameKey("command_option_widget_account_name"), builder.SetDescKey("command_option_widget_account_description")).
					Autocomplete(),
			).
			Handler(func(ctx *common.Context) error {
				var accountID string
				value, _ := ctx.Options().Value("account").(string)
				parts := strings.Split(value, "#")
				if len(parts) == 4 && parts[0] == "valid" {
					accountID = parts[1]
				} else {
					defaultAccount, hasDefaultAccount := ctx.User.Connection(models.ConnectionTypeWargaming, map[string]any{"default": true})
					if !hasDefaultAccount {
						return ctx.Reply().Format("commands_widget_message_fmt", constants.FrontendURL+"/widget/").Send()
					}
					accountID = defaultAccount.ReferenceID
				}

				return ctx.Reply().Format("commands_widget_message_fmt", constants.FrontendURL+"/widget/"+accountID).Send()
			}))
}