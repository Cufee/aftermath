package common

import (
	"image"
	"strings"

	"github.com/cufee/aftermath/internal/render/assets"
)

type Options struct {
	VehicleID          string
	PromoText          []string
	Background         image.Image
	BackgroundIsCustom bool
	Printer            func(string) string
}

func DefaultOptions() Options {
	return Options{Printer: func(s string) string { return s }}
}

type Option func(*Options)

func WithPromoText(text ...string) Option {
	return func(o *Options) {
		o.PromoText = text
	}
}
func WithVehicleID(vid string) Option {
	return func(o *Options) {
		o.VehicleID = vid
	}
}

func WithPrinter(printer func(string) string) Option {
	return func(o *Options) {
		o.Printer = printer
	}
}

func WithBackgroundURL(bgURL string, isCustom bool) Option {
	if strings.HasPrefix(bgURL, "static://") {
		img, ok := assets.GetLoadedImage(strings.ReplaceAll(bgURL, "static://", ""))
		if ok {
			return func(o *Options) {
				o.Background = img
				o.BackgroundIsCustom = isCustom
			}
		}
	}
	return func(o *Options) {}
}

func WithBackground(image image.Image, isCustom bool) Option {
	return func(o *Options) {
		o.BackgroundIsCustom = isCustom
		o.Background = image
	}
}
