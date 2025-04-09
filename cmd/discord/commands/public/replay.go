package public

import (
	"bytes"
	"context"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/permissions"
	stats "github.com/cufee/aftermath/internal/stats/client/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
	"github.com/pkg/errors"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("replay").
			Middleware(middleware.RequirePermissions(permissions.UseTextCommands, permissions.UseImageCommands)).
			Options(
				builder.NewOption("file", discordgo.ApplicationCommandOptionAttachment).
					Params(builder.SetNameKey("command_option_replay_file_name"), builder.SetDescKey("command_option_replay_file_description")),

				builder.NewOption("link", discordgo.ApplicationCommandOptionString).
					Params(builder.SetNameKey("command_option_replay_link_name"), builder.SetDescKey("command_option_replay_link_description")),
			).
			Handler(func(ctx common.Context) error {
				if len(ctx.Options()) == 0 {
					return ctx.Reply().Send("command_replay_help_message")
				}

				link, linkOK := ctx.Options().Value("link").(string)
				file, fileOK := ctx.Options().Value("file").(string)
				if (!linkOK && !fileOK) || (link == "" && file == "") {
					return ctx.Reply().IsError(common.UserError).Send("replay_errors_missing_attachment")
				}
				var hintMessage string
				if linkOK && fileOK {
					hintMessage = "replay_error_too_many_options_non_blocking"
				}

				replayURL := link
				if data, ok := ctx.CommandData(); ok && data.Resolved != nil {
					if attachment, ok := data.Resolved.Attachments[file]; ok {
						replayURL = attachment.URL
					}
				}

				parsed, err := url.Parse(replayURL)
				if err != nil {
					return ctx.Reply().IsError(common.UserError).Send("replay_errors_invalid_attachment")
				}
				if !strings.Contains(parsed.Path, ".wotbreplay") {
					return ctx.Reply().IsError(common.UserError).Send("replay_errors_invalid_attachment")
				}

				image, _, err := ctx.Core().Stats(ctx.Locale()).ReplayImage(context.Background(), parsed.String(), stats.WithWN8())
				if err != nil {
					if errors.Is(err, replay.ErrInvalidReplayFile) {
						return ctx.Reply().IsError(common.UserError).Send("replay_errors_invalid_attachment")
					}
					return ctx.Err(err, common.ApplicationError)
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return ctx.Err(err, common.ApplicationError)
				}
				return ctx.Reply().WithAds().Hint(hintMessage).File(buf.Bytes(), "replay_command_by_aftermath.png").Send()
			}),
	)

}
