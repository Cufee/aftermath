package commands

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"net/url"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/encoding"
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
			Handler(func(ctx common.Context) error {
				helpButton := discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						common.ButtonJoinPrimaryGuild(ctx.Localize("buttons_have_a_question_question")),
					}}

				pending, err := ctx.Core().Database().FindUserModerationRequests(ctx.Ctx(), ctx.User().ID, []string{"background-upload"}, []models.ModerationStatus{models.ModerationStatusSubmitted})
				if err != nil {
					return ctx.Err(err)
				}
				if len(pending) > 0 {
					return ctx.Reply().Component(helpButton).Send("fancy_errors_pending_request_exists")
				}

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

				// save the image
				encoded, err := encoding.EncodeGob(img)
				if err != nil {
					return ctx.Err(err)
				}
				content, err := ctx.Core().Database().CreateUserContent(ctx.Ctx(), models.UserContent{
					Type:        models.UserContentTypeInModeration,
					Value:       string(encoded),
					UserID:      ctx.User().ID,
					ReferenceID: ctx.User().ID,
				})
				if err != nil {
					return ctx.Err(err)
				}

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
							CustomID: fmt.Sprintf("moderation_image_submit#%s", content.ID),
						},
						common.ButtonJoinPrimaryGuild(ctx.Localize("buttons_have_a_question_question")),
					}}

				// send a preview
				var buf bytes.Buffer
				imaging.Encode(&buf, imaging.Blur(withBg.Image(), render.DefaultBackgroundBlur), imaging.JPEG, imaging.JPEGQuality(80))
				err = ctx.Reply().File(buf.Bytes(), "preview.jpeg").Component(buttons).Hint(hintMessage).Send("fancy_preview_msg")
				if err != nil {
					return ctx.Err(err)
				}

				go func(ctx common.Context) {
					time.Sleep(time.Minute * 1)
					c, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()
					err := ctx.DeleteResponse(c)
					if err != nil {
						log.Err(err).Msg("failed to delete an interaction response")
					}
				}(ctx)

				return nil
			}),
	)

	LoadedPublic.add(
		builder.NewCommand("moderation_image_submit_button").
			Middleware(middleware.RequirePermissions(permissions.CreatePersonalContent, permissions.UseTextCommands)).
			ComponentType(func(customID string) bool {
				return strings.HasPrefix(customID, "moderation_image_submit#")
			}).
			Handler(func(ctx common.Context) error {
				data, ok := ctx.ComponentData()
				if !ok {
					return ctx.Error("failed to get component data on interaction command")
				}
				contentID := strings.ReplaceAll(data.CustomID, "moderation_image_submit#", "")
				if contentID == "" {
					return ctx.Error("failed to get content id from custom id")
				}

				content, err := ctx.Core().Database().GetUserContent(ctx.Ctx(), contentID)
				if err != nil {
					if database.IsNotFound(err) {
						return ctx.Reply().Send("fancy_errors_preview_expired")
					}
					return ctx.Err(err)
				}

				var img image.NRGBA
				err = encoding.DecodeGob([]byte(content.Value), &img)
				if err != nil {
					return ctx.Err(err)
				}

				request := models.ModerationRequest{
					RequestorID:    ctx.User().ID,
					ReferenceID:    "background-upload",
					RequestContext: "image submitted from /fancy",
					ActionStatus:   models.ModerationStatusSubmitted,
					Data:           map[string]any{"content": content.ID},
				}

				request, err = ctx.Core().Database().CreateModerationRequest(ctx.Ctx(), request)
				if err != nil {
					return ctx.Err(err)
				}

				var buf bytes.Buffer
				err = imaging.Encode(&buf, &img, imaging.PNG)
				if err != nil {
					return ctx.Err(err)
				}

				_, err = ctx.CreateMessage(ctx.Ctx(), constants.DiscordContentModerationChannelID, discordgo.MessageSend{
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{Style: discordgo.SuccessButton, Label: "Approved", CustomID: "moderation_image_approve_button#" + request.ID},
								discordgo.Button{Style: discordgo.PrimaryButton, Label: "Decline", CustomID: "moderation_image_decline_button#" + request.ID},
							},
						},
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{Style: discordgo.DangerButton, Label: "Feature Ban", CustomID: "moderation_image_feature_ban_button#" + request.ID},
							},
						},
					},
					Content: fmt.Sprintf(
						`
## New Moderation Request
### *image submitted from /fancy*
**id**: %s%s%s
**User:** <@%s>
Please review the attached image and use the buttons below to take an action.
- We should not allow any kind of NSFW or suggestive content - this includes an intent. If there is something ever remotely sus, decline it.
- Some users will upload NSFW content or attempt to scribble over the image to bypass the filer. In such cases, we should issue a permanent feature ban.

*A feature ban will disable all content uploading features for the user.*
					`, request.ID, "`", ctx.User().ID, "`",
					)},
					[]rest.File{{Data: buf.Bytes(), Name: "user_background.png"}},
				)
				if err != nil {
					return ctx.Err(err)
				}

				return ctx.Reply().Hint(request.ID).Send("fancy_submitted_msg")
			}),
	)
}
