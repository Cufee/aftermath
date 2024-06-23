package common

import (
	"image"
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

func WithBackground(bg image.Image) Option {
	return func(o *Options) {
		o.Background = bg
	}
}
