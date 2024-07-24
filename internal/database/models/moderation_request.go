package models

import "time"

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
