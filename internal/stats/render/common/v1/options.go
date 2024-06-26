package common

import (
	"image"
	"strings"

	"github.com/cufee/aftermath/internal/stats/render/assets"
)

type Options struct {
	PromoText  []string
	Background image.Image
}

func DefaultOptions() Options {
	return Options{}
}

type Option func(*Options)

func WithPromoText(text ...string) Option {
	return func(o *Options) {
		o.PromoText = text
	}
}

func WithBackground(bgURL string) Option {
	if bgURL == "" {
		bgURL = "static://bg-default"
	}

	var image image.Image
	if strings.HasPrefix(bgURL, "static://") {
		img, ok := assets.GetLoadedImage(strings.ReplaceAll(bgURL, "static://", ""))
		if ok {
			image = img
		}
	}

	if image == nil {
		return func(o *Options) {}
	}
	return func(o *Options) {
		o.Background = image
	}
}
