package commands

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/permissions"
	render "github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

func init() {
	LoadedPublic.add(
		builder.NewCommand("fancy").
			Middleware(middleware.RequirePermissions(permissions.UseTextCommands, permissions.CreatePersonalContent, permissions.RemovePersonalContent, permissions.UpdatePersonalContent)).
			Ephemeral().
			Options(
				builder.NewOption("file", discordgo.ApplicationCommandOptionAttachment).
					Params(builder.SetNameKey("command_option_fancy_file_name"), builder.SetDescKey("command_option_fancy_file_description")),

				builder.NewOption("link", discordgo.ApplicationCommandOptionString).
					Params(builder.SetNameKey("command_option_fancy_link_name"), builder.SetDescKey("command_option_fancy_link_description")),
			).
			Handler(func(ctx *common.Context) error {
				link, linkOK := ctx.Options().Value("link").(string)
				file, fileOK := ctx.Options().Value("file").(string)
				if (!linkOK && !fileOK) || (link == "" && file == "") {
					return ctx.Reply().Send("fancy_errors_missing_attachment")
				}
				var hintMessage string = "fancy_hint_image_transformation"
				if linkOK && fileOK {
					hintMessage = "fancy_error_too_many_options_non_blocking"
				}

				imageURL := link
				if data, ok := ctx.CommandData(); ok && data.Resolved != nil {
					if attachment, ok := data.Resolved.Attachments[file]; ok {
						imageURL = attachment.URL
					}
				}

				helpButton := discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						ButtonJoinPrimaryGuild(ctx.Localize("buttons_have_a_question_question")),
					}}

				parsed, err := url.Parse(imageURL)
				if err != nil {
					return ctx.Reply().Component(helpButton).Send("fancy_errors_invalid_link")
				}

				// download and resize the image
				img, err := common.SafeLoadRemoteImage(parsed)
				if err != nil {
					if errors.Is(err, common.ErrUnsupportedImageFormat) {
						return ctx.Reply().Component(helpButton).Send("fancy_errors_invalid_format")
					}
					if errors.Is(err, common.ErrInvalidImage) {
						return ctx.Reply().Component(helpButton).Send("fancy_errors_invalid_image")
					}
					return ctx.Err(err)
				}
				img = imaging.Fill(img, 300, 300, imaging.Center, imaging.Linear)
				withBg := gg.NewContext(300, 300)
				withBg.SetColor(render.DiscordBackgroundColor)
				withBg.Clear()
				withBg.DrawImage(img, 0, 0)

				buttons := discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style: discordgo.SecondaryButton,
							Label: ctx.Localize("buttons_submit"),
							Emoji: &discordgo.ComponentEmoji{
								ID:       "1264733320713867365",
								Name:     "check",
								Animated: true,
							},
							CustomID: fmt.Sprintf("moderation_image_submit#id:%s", "id"),
						},
						ButtonJoinPrimaryGuild(ctx.Localize("buttons_have_a_question_question")),
					}}

				// send a preview
				var buf bytes.Buffer
				imaging.Encode(&buf, imaging.Blur(withBg.Image(), render.DefaultBackgroundBlur), imaging.JPEG, imaging.JPEGQuality(80))
				err = ctx.Reply().File(buf.Bytes(), "preview.jpeg").Component(buttons).Hint(hintMessage).Send("fancy_preview_msg")
				if err != nil {
					return ctx.Err(err)
				}

				go func(rest *rest.Client, appID, token string) {
					time.Sleep(time.Minute * 1)
					c, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()
					err := rest.DeleteInteractionResponse(c, appID, token)
					if err != nil {
						log.Err(err).Msg("failed to delete an interaction response")
					}
				}(ctx.Rest(), ctx.RawInteraction().AppID, ctx.RawInteraction().Token)

				return nil
			}),
	)

}
