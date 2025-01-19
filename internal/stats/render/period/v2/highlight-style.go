package period

import (
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/facepaint/style"
)

var styledHighlightTitle = style.Style{}
var styledHighlightStatsRow = style.Style{}

type highlightCardStyle struct {
	card         style.Style
	titleWrapper style.Style
	titleLabel   func() *style.Style
	titleVehicle func() *style.Style
	statsWrapper style.Style
	stats        style.Style
	blockValue   func() *style.Style
	blockLabel   func() *style.Style
}

var styledHighlightCard = highlightCardStyle{
	card: style.Style{
		Direction:  style.DirectionHorizontal,
		AlignItems: style.AlignItemsCenter,

		BackgroundColor: common.DefaultCardColor,
		BlurBackground:  cardBackgroundBlur,

		BorderRadiusTopLeft:     common.BorderRadiusLG,
		BorderRadiusTopRight:    common.BorderRadiusLG,
		BorderRadiusBottomLeft:  common.BorderRadiusLG,
		BorderRadiusBottomRight: common.BorderRadiusLG,

		GrowHorizontal: true,
		Gap:            20,

		PaddingLeft:   cardPaddingX / 1.5,
		PaddingRight:  cardPaddingX / 1.5,
		PaddingTop:    cardPaddingY / 1.5,
		PaddingBottom: cardPaddingY / 1.5,
	},
	titleWrapper: style.Style{
		// Debug: true,

		GrowHorizontal: true,
		Direction:      style.DirectionVertical,
	},
	titleLabel: func() *style.Style {
		return &style.Style{
			Color: common.TextSecondary,
			Font:  common.FontSmall(),
		}
	},
	titleVehicle: func() *style.Style {
		return &style.Style{
			Color: common.TextPrimary,
			Font:  common.FontMedium(),
		}
	},
	stats: style.Style{
		Direction:      style.DirectionVertical,
		AlignItems:     style.AlignItemsCenter,
		JustifyContent: style.JustifyContentCenter,
	},
	statsWrapper: style.Style{
		// Debug: true,

		Direction:  style.DirectionHorizontal,
		AlignItems: style.AlignItemsCenter,
		Gap:        10,
	},
	blockValue: func() *style.Style {
		return &style.Style{
			Color: common.TextPrimary,
			Font:  common.FontMedium(),
		}
	},
	blockLabel: func() *style.Style {
		return &style.Style{
			Color: common.TextAlt,
			Font:  common.FontSmall(),
		}
	},
}
