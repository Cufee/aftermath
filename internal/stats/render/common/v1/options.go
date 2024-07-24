package common

import (
	"image"
	"strings"

	"github.com/cufee/aftermath/internal/stats/render/assets"
)

type Options struct {
	VehicleID  string
	PromoText  []string
	Background image.Image
	Printer    func(string) string
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

func WithBackgroundURL(bgURL string) Option {
	if strings.HasPrefix(bgURL, "static://") {
		img, ok := assets.GetLoadedImage(strings.ReplaceAll(bgURL, "static://", ""))
		if ok {
			return func(o *Options) {
				o.Background = img
			}
		}
	}
	return func(o *Options) {}
}

func WithBackground(image image.Image) Option {
	return func(o *Options) {
		o.Background = image
	}
}
