package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/models"
)

func toModeratorAction(r *db.ModerationRequest) models.ModerationRequest {
	return models.ModerationRequest{
		ID:        r.ID,
		UpdateAt:  r.UpdatedAt,
		CreatedAt: r.CreatedAt,

		ReferenceID:    r.ReferenceID,
		RequestorID:    r.RequestorID,
		RequestContext: r.Context,

		ActionStatus:     r.ActionStatus,
		ActionReason:     r.ActionReason,
		ModeratorComment: r.ModeratorComment,
		ModeratorID:      r.ModeratorID,
		Data:             r.Data,
	}
}

func (c *client) CreateModerationRequest(ctx context.Context, request models.ModerationRequest) (models.ModerationRequest, error) {
	create := c.db.ModerationRequest.Create().
		SetData(request.Data).
		SetContext(request.RequestContext).
		SetReferenceID(request.ReferenceID).
		SetActionStatus(request.ActionStatus).
		SetActionReason(request.ActionReason).
		SetModeratorComment(request.ModeratorComment).
		SetRequestor(c.db.User.GetX(ctx, request.RequestorID))
	if request.ModeratorID != nil {
		mod, err := c.db.User.Get(ctx, *request.ModeratorID)
		if err != nil {
			return models.ModerationRequest{}, err
		}
		create = create.SetModerator(mod)
	}

	record, err := create.Save(ctx)
	if err != nil {
		return models.ModerationRequest{}, err
	}

	return toModeratorAction(record), nil
}

func (c *client) GetModerationRequest(ctx context.Context, id string) (models.ModerationRequest, error) {
	record, err := c.db.ModerationRequest.Get(ctx, id)
	if err != nil {
		return models.ModerationRequest{}, err
	}
	return toModeratorAction(record), nil
}

func (c *client) UpdateModerationRequest(ctx context.Context, request models.ModerationRequest) (models.ModerationRequest, error) {
	update := c.db.ModerationRequest.UpdateOneID(request.ID).
		SetData(request.Data).
		SetContext(request.RequestContext).
		SetReferenceID(request.ReferenceID).
		SetActionStatus(request.ActionStatus).
		SetActionReason(request.ActionReason).
		SetModeratorComment(request.ModeratorComment)
	if request.ModeratorID != nil {
		mod, err := c.db.User.Get(ctx, *request.ModeratorID)
		if err != nil {
			return models.ModerationRequest{}, err
		}
		update = update.SetModerator(mod)
	}

	record, err := update.Save(ctx)
	if err != nil {
		return models.ModerationRequest{}, err
	}

	return toModeratorAction(record), nil
}
