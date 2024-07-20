package commands

import (
	"bytes"
	"context"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	stats "github.com/cufee/aftermath/internal/stats/client/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
	"github.com/pkg/errors"
)

func init() {
	LoadedPublic.add(
		builder.NewCommand("replay").
			Options(
				builder.NewOption("file", discordgo.ApplicationCommandOptionAttachment).
					Params(builder.SetNameKey("command_option_replay_file_name"), builder.SetDescKey("command_option_replay_file_description")),

				builder.NewOption("link", discordgo.ApplicationCommandOptionString).
					Params(builder.SetNameKey("command_option_replay_link_name"), builder.SetDescKey("command_option_replay_link_description")),
			).
			Handler(func(ctx *common.Context) error {
				var replayURL string
				link, linkOK := ctx.Options().Value("link").(string)
				file, fileOK := ctx.Options().Value("file").(string)
				if (!linkOK && !fileOK) || (link == "" && file == "") {
					return ctx.Reply().Send("replay_errors_missing_attachment")
				}
				if link != "" {
					replayURL = link
				} else {
					replayURL = ""
				}
				if !strings.Contains(link, ".wotbreplay") {
					return ctx.Reply().Send("replay_errors_invalid_attachment")
				}
				parsed, err := url.Parse(replayURL)
				if err != nil {
					return ctx.Reply().Send("replay_errors_invalid_attachment")
				}

				var backgroundURL = "static://bg-default"
				image, _, err := ctx.Core.Stats(ctx.Locale).ReplayImage(context.Background(), parsed.String(), stats.WithBackgroundURL(backgroundURL), stats.WithWN8())
				if err != nil {
					if errors.Is(err, replay.ErrInvalidReplayFile) {
						return ctx.Reply().Send("replay_errors_invalid_attachment")
					}
					return ctx.Err(err)
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return ctx.Err(err)
				}
				return ctx.Reply().File(buf.Bytes(), "replay_command_by_aftermath.png").Send()
			}),
	)

}
