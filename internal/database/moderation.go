package database

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/moderationrequest"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/userrestriction"
	"github.com/cufee/aftermath/internal/database/models"
)

func toModerationRequest(r *db.ModerationRequest) models.ModerationRequest {
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

	return toModerationRequest(record), nil
}

func (c *client) GetModerationRequest(ctx context.Context, id string) (models.ModerationRequest, error) {
	record, err := c.db.ModerationRequest.Get(ctx, id)
	if err != nil {
		return models.ModerationRequest{}, err
	}
	return toModerationRequest(record), nil
}

func (c *client) FindUserModerationRequests(ctx context.Context, userID string, referenceIDs []string, status []models.ModerationStatus, since time.Time) ([]models.ModerationRequest, error) {
	var where []predicate.ModerationRequest
	where = append(where, moderationrequest.RequestorID(userID), moderationrequest.UpdatedAtGT(since))
	if referenceIDs != nil {
		where = append(where, moderationrequest.ReferenceIDIn(referenceIDs...))
	}
	if status != nil {
		where = append(where, moderationrequest.ActionStatusIn(status...))
	}

	records, err := c.db.ModerationRequest.Query().Where(where...).Order(moderationrequest.ByCreatedAt(sql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}

	var requests []models.ModerationRequest
	for _, r := range records {
		requests = append(requests, toModerationRequest(r))
	}

	return requests, nil
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

	return toModerationRequest(record), nil
}

func (c *client) GetUserRestriction(ctx context.Context, id string) (models.UserRestriction, error) {
	record, err := c.db.UserRestriction.Get(ctx, id)
	if err != nil {
		return models.UserRestriction{}, err
	}

	return toUserRestriction(record), nil
}

func (c *client) GetUserRestrictions(ctx context.Context, userID string) ([]models.UserRestriction, error) {
	records, err := c.db.UserRestriction.Query().Where(userrestriction.UserID(userID)).All(ctx)
	if err != nil {
		return nil, err
	}

	var restrictions []models.UserRestriction
	for _, r := range records {
		restrictions = append(restrictions, toUserRestriction(r))
	}

	return restrictions, nil
}

func (c *client) CreateUserRestriction(ctx context.Context, data models.UserRestriction) (models.UserRestriction, error) {
	user, err := c.db.User.Get(ctx, data.UserID)
	if err != nil {
		return models.UserRestriction{}, err
	}

	record, err := c.db.UserRestriction.Create().
		SetModeratorComment(data.ModeratorComment).
		SetRestriction(data.Restriction.String()).
		SetPublicReason(data.PublicReason).
		SetExpiresAt(data.ExpiresAt).
		SetEvents(data.Events).
		SetType(data.Type).
		SetUser(user).
		Save(ctx)
	if err != nil {
		return models.UserRestriction{}, err
	}

	return toUserRestriction(record), nil
}

func (c *client) UpdateUserRestriction(ctx context.Context, data models.UserRestriction) (models.UserRestriction, error) {
	record, err := c.db.UserRestriction.UpdateOneID(data.ID).
		SetModeratorComment(data.ModeratorComment).
		SetRestriction(data.Restriction.String()).
		SetPublicReason(data.PublicReason).
		SetExpiresAt(data.ExpiresAt).
		SetEvents(data.Events).
		SetType(data.Type).
		Save(ctx)
	if err != nil {
		return models.UserRestriction{}, err
	}
	return toUserRestriction(record), nil
}
