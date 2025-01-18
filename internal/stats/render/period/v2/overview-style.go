package period

import (
	"github.com/cufee/aftermath/internal/render/common"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint/style"
)

const (
	debugOverviewCards = false

	iconSizeWN8    = 54.0
	iconSizeRating = 60.0
)

type blockStyle struct {
	wrapper        style.Style
	valueContainer style.Style
	value          style.Style
	label          style.Style
}

type overviewCardStyle struct {
	card   style.Style
	column style.Style
}

func (s *overviewCardStyle) styleBlock(block prepare.StatsBlock[period.BlockData, string]) blockStyle {
	switch block.Data.Flavor {
	case period.BlockFlavorSpecial:
		return blockStyle{
			wrapper: style.Style{
				Debug: debugOverviewCards,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentSpaceAround,
				GrowVertical:   true,
				Gap:            10,
			},
			valueContainer: style.Style{
				Debug: debugOverviewCards,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentEnd,
				// GrowVertical:   true,
			},
			value: style.Style{
				Debug: debugOverviewCards,

				Color: common.TextPrimary,
				Font:  common.FontXL(),
			},
			label: style.Style{
				Color:      common.TextAlt,
				Font:       common.FontSmall(),
				PaddingTop: -6,
			},
		}
	case period.BlockFlavorSecondary:
		return blockStyle{
			wrapper: style.Style{},
			valueContainer: style.Style{
				Debug: debugOverviewCards,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentCenter,
			},
			value: style.Style{
				Debug: debugOverviewCards,

				Color: common.TextSecondary,
				Font:  common.FontMedium(),
			},
			label: style.Style{
				Color:      common.TextAlt,
				Font:       common.FontSmall(),
				PaddingTop: -3,
			},
		}
	default:
		return blockStyle{
			wrapper: style.Style{},
			valueContainer: style.Style{
				Debug: debugOverviewCards,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentCenter,
			},
			value: style.Style{
				Debug: debugOverviewCards,

				Color: common.TextPrimary,
				Font:  common.FontLarge(),
			},
			label: style.Style{
				Color:      common.TextAlt,
				Font:       common.FontSmall(),
				PaddingTop: -5,
			},
		}
	}
}

var styledOverviewCard = overviewCardStyle{
	card: style.Style{
		Debug: debugOverviewCards,

		Direction:               style.DirectionHorizontal,
		AlignItems:              style.AlignItemsCenter,
		JustifyContent:          style.JustifyContentSpaceAround,
		BackgroundColor:         common.DefaultCardColor,
		BorderRadiusTopLeft:     common.BorderRadiusLG,
		BorderRadiusTopRight:    common.BorderRadiusLG,
		BorderRadiusBottomLeft:  common.BorderRadiusLG,
		BorderRadiusBottomRight: common.BorderRadiusLG,
		GrowHorizontal:          true,
		Gap:                     10,
		PaddingLeft:             20,
		PaddingRight:            20,
		PaddingTop:              20,
		PaddingBottom:           20,
	},
	column: style.Style{
		Debug: debugOverviewCards,

		Direction:      style.DirectionVertical,
		AlignItems:     style.AlignItemsCenter,
		JustifyContent: style.JustifyContentCenter,
		GrowVertical:   true,
		Gap:            10,
	},
}

// wrapped around special block text and icon
var styledOverviewSpecialBlockWrapper = style.Style{
	Debug: debugOverviewCards,

	Direction:      style.DirectionVertical,
	AlignItems:     style.AlignItemsCenter,
	JustifyContent: style.JustifyContentCenter,
	Gap:            10,
}

var styledCardsFrame = style.Style{
	Debug: false,

	Direction:     style.DirectionVertical,
	Gap:           10,
	PaddingLeft:   20,
	PaddingRight:  20,
	PaddingTop:    20,
	PaddingBottom: 20,
}
