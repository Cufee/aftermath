package period

import "github.com/cufee/aftermath/internal/stats/render/common"

type options struct {
	cardStyle common.Style
	promoText []string
}

type Option func(*options)

func WithStyle(style common.Style) Option {
	return func(o *options) {
		o.cardStyle = style
	}
}

func WithPromoText(text ...string) Option {
	return func(o *options) {
		o.promoText = text
	}
}
