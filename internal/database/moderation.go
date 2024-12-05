package database

import (
	"context"
	"time"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/sqlite"
)

func (c *client) CreateModerationRequest(ctx context.Context, request models.ModerationRequest) (models.ModerationRequest, error) {
	model := request.Model()
	stmt := t.ModerationRequest.
		INSERT(t.ModerationRequest.AllColumns).
		MODEL(model).
		RETURNING(t.ModerationRequest.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.ModerationRequest{}, err
	}

	return models.ToModerationRequest(&model), nil
}

func (c *client) GetModerationRequest(ctx context.Context, id string) (models.ModerationRequest, error) {
	stmt := t.ModerationRequest.
		SELECT(t.ModerationRequest.AllColumns).
		WHERE(t.ModerationRequest.ID.EQ(s.String(id)))

	var result m.ModerationRequest
	err := c.query(ctx, stmt, &result)
	if err != nil {
		return models.ModerationRequest{}, err
	}

	return models.ToModerationRequest(&result), nil
}

func (c *client) FindUserModerationRequests(ctx context.Context, userID string, referenceIDs []string, status []models.ModerationStatus, since time.Time) ([]models.ModerationRequest, error) {
	where := []s.BoolExpression{
		t.ModerationRequest.RequestorID.EQ(s.String(userID)),
		t.ModerationRequest.UpdatedAt.GT(s.DATETIME(since)),
	}

	if referenceIDs != nil {
		where = append(where, t.ModerationRequest.ReferenceID.IN(stringsToExp(referenceIDs)...))
	}
	if status != nil {
		var s []string
		for _, st := range status {
			s = append(s, string(st))
		}
		where = append(where, t.ModerationRequest.ActionStatus.IN(stringsToExp(s)...))
	}

	var records []m.ModerationRequest
	stmt := t.ModerationRequest.
		SELECT(t.ModerationRequest.AllColumns).
		WHERE(s.AND(where...)).
		ORDER_BY(t.ModerationRequest.CreatedAt.DESC())

	err := c.query(ctx, stmt, &records)
	if err != nil {
		return nil, err
	}

	var requests []models.ModerationRequest
	for _, r := range records {
		requests = append(requests, models.ToModerationRequest(&r))
	}
	return requests, nil
}

func (c *client) UpdateModerationRequest(ctx context.Context, request models.ModerationRequest) (models.ModerationRequest, error) {
	model := request.Model()
	stmt := t.ModerationRequest.
		UPDATE(
			t.ModerationRequest.UpdatedAt,
			t.ModerationRequest.ReferenceID,
			t.ModerationRequest.RequestorID,
			t.ModerationRequest.Context,
			t.ModerationRequest.ActionReason,
			t.ModerationRequest.ActionStatus,
			t.ModerationRequest.ModeratorComment,
			t.ModerationRequest.ModeratorID,
			t.ModerationRequest.Data,
		).
		MODEL(model).
		WHERE(t.ModerationRequest.ID.EQ(s.String(request.ID))).
		RETURNING(t.ModerationRequest.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.ModerationRequest{}, err
	}

	return models.ToModerationRequest(&model), nil
}

func (c *client) GetUserRestriction(ctx context.Context, id string) (models.UserRestriction, error) {
	var record m.UserRestriction
	err := c.query(ctx, t.UserRestriction.SELECT(t.UserRestriction.AllColumns).WHERE(t.UserRestriction.ID.EQ(s.String(id))), &record)
	if err != nil {
		return models.UserRestriction{}, err
	}

	return models.ToUserRestriction(&record), nil
}

func (c *client) GetUserRestrictions(ctx context.Context, userID string) ([]models.UserRestriction, error) {
	var record []m.UserRestriction
	err := c.query(ctx, t.UserRestriction.SELECT(t.UserRestriction.AllColumns).WHERE(t.UserRestriction.UserID.EQ(s.String(userID))), &record)
	if err != nil {
		return nil, err
	}

	var restrictions []models.UserRestriction
	for _, r := range record {
		restrictions = append(restrictions, models.ToUserRestriction(&r))
	}

	return restrictions, nil
}

func (c *client) CreateUserRestriction(ctx context.Context, data models.UserRestriction) (models.UserRestriction, error) {
	model := data.Model()
	stmt := t.UserRestriction.INSERT(t.UserRestriction.AllColumns).MODEL(model).RETURNING(t.UserRestriction.AllColumns)
	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.UserRestriction{}, err
	}
	return models.ToUserRestriction(&model), nil
}

func (c *client) UpdateUserRestriction(ctx context.Context, data models.UserRestriction) (models.UserRestriction, error) {
	model := data.Model()
	stmt := t.UserRestriction.
		UPDATE(
			t.UserRestriction.UpdatedAt,
			t.UserRestriction.PublicReason,
			t.UserRestriction.ModeratorComment,
			t.UserRestriction.ExpiresAt,
			t.UserRestriction.Events,
		).
		WHERE(t.UserRestriction.ID.EQ(s.String(data.ID))).
		MODEL(model).
		RETURNING(t.UserRestriction.AllColumns)
	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.UserRestriction{}, err
	}
	return models.ToUserRestriction(&model), nil
}
