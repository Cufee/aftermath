package models

import (
	"time"

	"github.com/cufee/aftermath/internal/json"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
)

type ModerationStatus string

const (
	ModerationStatusSubmitted ModerationStatus = "submitted"
	ModerationStatusApproved  ModerationStatus = "approved"
	ModerationStatusDeclined  ModerationStatus = "declined"
	ModerationStatusExpired   ModerationStatus = "expired"
)

func (ModerationStatus) Values() []string {
	return []string{
		string(ModerationStatusSubmitted),
		string(ModerationStatusApproved),
		string(ModerationStatusDeclined),
		string(ModerationStatusExpired),
	}
}

type ModerationRequest struct {
	ID        string
	CreatedAt time.Time
	UpdateAt  time.Time

	ReferenceID    string
	RequestorID    string
	RequestContext string

	ActionStatus     ModerationStatus
	ActionReason     string
	ModeratorID      *string
	ModeratorComment string
	Data             map[string]any
}

func ToModerationRequest(r *model.ModerationRequest) ModerationRequest {
	req := ModerationRequest{
		ID:        r.ID,
		UpdateAt:  r.UpdatedAt,
		CreatedAt: r.CreatedAt,

		ReferenceID: r.ReferenceID,
		RequestorID: r.RequestorID,

		ActionStatus: ModerationStatus(r.ActionStatus),
		ModeratorID:  r.ModeratorID,
	}
	if r.Context != nil {
		req.RequestContext = *r.Context
	}
	if r.ActionReason != nil {
		req.ActionReason = *r.ActionReason
	}
	if r.ModeratorComment != nil {
		req.ModeratorComment = *r.ModeratorComment
	}
	json.Unmarshal(r.Data, &req.Data)

	return req
}

func (r ModerationRequest) Model() model.ModerationRequest {
	req := model.ModerationRequest{
		ID:        utils.StringOr(r.ID, cuid.New()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		ReferenceID: r.ReferenceID,
		RequestorID: r.RequestorID,

		ActionStatus: string(r.ActionStatus),
		ModeratorID:  r.ModeratorID,
	}
	if r.RequestContext != "" {
		req.Context = &r.RequestContext
	}
	if r.ActionReason != "" {
		req.ActionReason = &r.ActionReason
	}
	if r.ModeratorComment != "" {
		req.ModeratorComment = &r.ModeratorComment
	}
	req.Data, _ = json.Marshal(r.Data)

	return req
}
