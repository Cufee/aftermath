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
		// Debug: true,

		Direction:  style.DirectionHorizontal,
		AlignItems: style.AlignItemsCenter,

		BackgroundColor: common.DefaultCardColor,

		BorderRadiusTopLeft:     common.BorderRadiusLG,
		BorderRadiusTopRight:    common.BorderRadiusLG,
		BorderRadiusBottomLeft:  common.BorderRadiusLG,
		BorderRadiusBottomRight: common.BorderRadiusLG,

		GrowHorizontal: true,

		PaddingLeft:   20,
		PaddingRight:  20,
		PaddingTop:    15,
		PaddingBottom: 15,
	},
	titleWrapper: style.Style{
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
