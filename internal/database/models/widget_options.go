package models

import "time"

type WidgetOptions struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID    string
	AccountID string

	Style WidgetStyling
}

type WidgetStyling struct {
	Flavor          string
	UnratedOverview WidgetCardStyle
	RatingOverview  WidgetCardStyle
	Vehicles        WidgetCardStyle
}

type WidgetCardStyle struct {
	ShowTitle  bool
	ShowCareer bool
	ShowLabel  bool
	Blocks     []string
}
