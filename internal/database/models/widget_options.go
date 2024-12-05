package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
)

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

	Title        string    `json:"title"`
	UserID       string    `json:"userId"`
	AccountID    string    `json:"accountId"`
	SessionFrom  time.Time `json:"sessionFrom"`
	SessionRefID string    `json:"sessionReferenceId"`

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

func ToWidgetOptions(record *model.WidgetSettings) WidgetOptions {
	o := WidgetOptions{
		ID:        record.ID,
		CreatedAt: StringToTime(record.CreatedAt),
		UpdatedAt: StringToTime(record.UpdatedAt),

		UserID:    record.UserID,
		AccountID: record.ReferenceID,
	}
	if record.Title != nil {
		o.Title = *record.Title
	}
	if record.SessionFrom != nil {
		o.SessionFrom = StringToTime(*record.SessionFrom)
	}
	if record.SessionReferenceID != nil {
		o.SessionRefID = *record.SessionReferenceID
	}
	json.Unmarshal(record.Styles, &o.Style)
	json.Unmarshal(record.Metadata, &o.Meta)

	if o.Meta == nil {
		o.Meta = make(map[string]any, 0)
	}
	return o
}

func (record *WidgetOptions) Model() model.WidgetSettings {
	s := model.WidgetSettings{
		ID:          utils.StringOr(record.ID, cuid.New()),
		CreatedAt:   TimeToString(time.Now()),
		UpdatedAt:   TimeToString(time.Now()),
		ReferenceID: record.AccountID,
		UserID:      record.UserID,
	}
	if record.Title != "" {
		s.Title = &record.Title
	}
	if !record.SessionFrom.IsZero() {
		f := TimeToString(record.SessionFrom)
		s.SessionFrom = &f
	}
	if record.SessionRefID != "" {
		s.SessionReferenceID = &record.SessionRefID
	}
	s.Metadata, _ = json.Marshal(record.Meta)
	s.Styles, _ = json.Marshal(record.Style)
	return s
}
