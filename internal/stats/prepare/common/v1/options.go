package common

import (
	"golang.org/x/text/language"
)

func DefaultOptions() options {
	return options{}
}

type options struct {
	localePrinter func(string) string
	locale        *language.Tag

	VehicleTags    []Tag
	RatingColumns  []TagColumn[string]
	UnratedColumns []TagColumn[string]
}

func (o options) Printer() func(s string) string {
	if o.localePrinter != nil {
		return o.localePrinter
	}
	return func(s string) string { return s }
}

func (o options) Locale() language.Tag {
	if o.locale != nil {
		return *o.locale
	}
	return language.English
}

type Option func(*options)

func WithPrinter(printer func(string) string, locale language.Tag) func(*options) {
	return func(o *options) { o.localePrinter = printer; o.locale = &locale }
}
func WithVehicleTags(tags ...Tag) func(*options) {
	return func(o *options) { o.VehicleTags = tags }
}
func WithRatingColumns(columns ...TagColumn[string]) func(*options) {
	return func(o *options) { o.RatingColumns = columns }
}
func WithUnratedColumns(columns ...TagColumn[string]) func(*options) {
	return func(o *options) { o.UnratedColumns = columns }
}
