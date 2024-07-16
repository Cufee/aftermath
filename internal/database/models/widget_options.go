package models

import "time"

type WidgetFlavor string

const (
	WidgetFlavorDefault WidgetFlavor = "default"
)

var DefaultWidgetStyle = WidgetStyling{
	Flavor: WidgetFlavorDefault,
	Vehicles: WidgetVehicleCardStyle{
		Limit: 3,
		WidgetCardStyle: WidgetCardStyle{
			Visible:    true,
			ShowTitle:  true,
			ShowLabel:  false,
			ShowCareer: false,
		},
	},
	UnratedOverview: WidgetCardStyle{
		Visible:    true,
		ShowCareer: false,
		ShowTitle:  false,
		ShowLabel:  true,
	},
	RatingOverview: WidgetCardStyle{
		Visible:    true,
		ShowCareer: false,
		ShowTitle:  false,
		ShowLabel:  true,
	},
}

type WidgetOptions struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	Title      string
	UserID     string
	AccountID  string
	SnapshotID string

	Style WidgetStyling

	Meta map[string]any
}

type WidgetStyling struct {
	Flavor          WidgetFlavor
	UnratedOverview WidgetCardStyle
	RatingOverview  WidgetCardStyle
	Vehicles        WidgetVehicleCardStyle
}

type WidgetCardStyle struct {
	Visible    bool
	ShowTitle  bool
	ShowCareer bool
	ShowLabel  bool
	Blocks     []string
}

type WidgetVehicleCardStyle struct {
	WidgetCardStyle
	Limit int
}
