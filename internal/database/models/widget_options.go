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
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Title       string    `json:"title"`
	UserID      string    `json:"userId"`
	AccountID   string    `json:"accountId"`
	SessionFrom time.Time `json:"sessionFrom"`

	Style WidgetStyling `json:"style"`

	Meta map[string]any `json:"meta"`
}

type WidgetStyling struct {
	Flavor          WidgetFlavor           `json:"flavor"`
	UnratedOverview WidgetCardStyle        `json:"unrated"`
	RatingOverview  WidgetCardStyle        `json:"rating"`
	Vehicles        WidgetVehicleCardStyle `json:"vehicles"`
}

type WidgetCardStyle struct {
	Visible    bool     `json:"visible"`
	ShowTitle  bool     `json:"showTitle"`
	ShowCareer bool     `json:"showCareer"`
	ShowLabel  bool     `json:"showLabel"`
	Blocks     []string `json:"blocks"`
}

type WidgetVehicleCardStyle struct {
	WidgetCardStyle
	Limit int `json:"limit"`
}
