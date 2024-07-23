package public

import (
	"fmt"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/permissions"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("moderation_image_action_button").
			Middleware(middleware.RequirePermissions(permissions.ContentModerator, permissions.UseTextCommands)).
			ComponentType(func(customID string) bool {
				for _, id := range []string{"moderation_image_approve_button#", "moderation_image_decline_button#", "moderation_image_feature_ban_button#"} {
					if strings.HasPrefix(customID, id) {
						return true
					}
				}
				return false
			}).
			Handler(func(ctx common.Context) error {
				data, ok := ctx.ComponentData()
				if !ok {
					return ctx.Reply().Hint("no action was performed").Send("Failed to get component data")
				}

				var messageID string
				if ctx.RawInteraction().Message != nil {
					messageID = ctx.RawInteraction().Message.ID
				}
				if messageID == "" {
					return ctx.Reply().Hint("no action was performed").Send("Failed to get message id")
				}

				// parse request id and action
				var action string
				var requestID string
				switch {
				default:
					return ctx.Reply().Hint("no action was performed").Send("Invalid button custom id")

				case strings.HasPrefix(data.CustomID, "moderation_image_approve_button#"):
					requestID = strings.TrimPrefix(data.CustomID, "moderation_image_approve_button#")
					action = "approve"

				case strings.HasPrefix(data.CustomID, "moderation_image_decline_button#"):
					requestID = strings.TrimPrefix(data.CustomID, "moderation_image_decline_button#")
					action = "decline"

				case strings.HasPrefix(data.CustomID, "moderation_image_feature_ban_button#"):
					requestID = strings.TrimPrefix(data.CustomID, "moderation_image_feature_ban_button#")
					action = "feature-ban"
				}
				if requestID == "" {
					return ctx.Reply().Hint("no action was performed").Send("Invalid request id extracted from custom id")
				}

				// get the request
				request, err := ctx.Core().Database().GetModerationRequest(ctx.Ctx(), requestID)
				if err != nil {
					return ctx.Reply().Hint("no action was performed").Send("Failed to get a request", err.Error())
				}
				if request.ActionStatus != models.ModerationStatusSubmitted {
					return ctx.Reply().Hint("no action was performed").Send("This moderation request was already reviewed", "result: "+string(request.ActionStatus), fmt.Sprintf("moderator: <@%s>", *request.ModeratorID))
				}
				var contentID string
				if id, ok := request.Data["content"].(string); ok {
					contentID = id
				}
				if contentID == "" {
					return ctx.Reply().Hint("no action was performed").Send("Failed to get content id from request")
				}

				// get user content
				content, err := ctx.Core().Database().GetUserContent(ctx.Ctx(), contentID)
				if err != nil {
					return ctx.Reply().Hint("no action was performed").Send("Failed to get content id from request")
				}
				content.Meta["moderation_request"] = request.ID

				// update the request / issue actions
				var directMessageContent string
				switch action {
				default:
					return ctx.Reply().Hint("no action was performed").Send("Invalid action")
				case "approve":
					request.ActionStatus = models.ModerationStatusApproved
					// update user content
					content.ReferenceID = request.RequestorID
					content.Type = models.UserContentTypePersonalBackground
					content.Meta["moderation_request_approved_time"] = time.Now()
					content.Meta["moderation_request_approved"] = true
					content.Meta["value_type"] = "gob_image"
					_, err := ctx.Core().Database().UpdateUserContent(ctx.Ctx(), content)
					if err != nil {
						return ctx.Reply().Hint("no action was performed").Send("Failed to update user content")
					}
					directMessageContent = "fancy_moderation_request_approved"

				case "feature-ban":
					// issue a feature ban
					restriction := models.UserRestriction{
						UserID:           request.RequestorID,
						Type:             models.RestrictionTypePartial,
						Restriction:      permissions.CreatePersonalContent.Add(permissions.UpdatePersonalContent),
						PublicReason:     "Uploading inappropriate content with /fancy command. This restriction cannot be appealed.",
						ModeratorComment: "moderator issued a feature ban using an action button",
						ExpiresAt:        time.Now().Add(time.Hour * 24 * 180), // 180 days
					}
					restriction.AddEvent(ctx.User().ID, "restriction created", "restriction issued from a /fancy image review message")
					restriction, err = ctx.Core().Database().CreateUserRestriction(ctx.Ctx(), restriction)
					if err != nil {
						return ctx.Reply().Send("Failed to create a user restriction", err.Error())
					}
					directMessageContent = fmt.Sprintf(ctx.Localize("fancy_moderation_request_declined_and_banned_fmt"), "180")
					request.ModeratorComment = "moderator chose to issue a feature ban"
					fallthrough
				case "decline":
					// delete user content
					request.ActionStatus = models.ModerationStatusDeclined
					err := ctx.Core().Database().DeleteUserContent(ctx.Ctx(), content.ID)
					if err != nil {
						return ctx.Reply().Send("Failed to delete user content")
					}
					if directMessageContent == "" {
						directMessageContent = ctx.Localize("fancy_moderation_request_declined")
					}
				}

				// send a DM
				var replyHint string
				dmChan, err := ctx.CreateDMChannel(ctx.Ctx(), request.RequestorID)
				if err != nil {
					log.Err(err).Msg("failed to create a DM channel")
					replyHint = "failed to dm the user, this does not cancel the action."
				} else {
					_, err := ctx.CreateMessage(ctx.Ctx(), dmChan.ID, ctx.Reply().Text(directMessageContent))
					if err != nil {
						log.Err(err).Msg("failed to send a DM")
						replyHint = "failed to dm the user, this does not cancel the action."
					}
				}

				// update moderation request
				id := ctx.User().ID
				request.ModeratorID = &id
				request, err = ctx.Core().Database().UpdateModerationRequest(ctx.Ctx(), request)
				if err != nil {
					return ctx.Reply().Hint("no action was performed").Send("Failed to update request", err.Error())
				}

				_, err = ctx.UpdateMessage(ctx.Ctx(), ctx.RawInteraction().ChannelID, messageID, ctx.Reply().
					Format("## Moderation Request (Actioned)\n*image submitted from /fancy*\n**User:** <@%s>\n**Action:** %s\n**Moderator:** <@%s>", request.RequestorID, string(request.ActionStatus), ctx.User().ID).
					Component(nil),
				)
				if err != nil {
					return ctx.Reply().Send("Failed to update request message", err.Error())
				}

				return ctx.Reply().Hint(replyHint).Send("request actioned")
			}),
	)

}
