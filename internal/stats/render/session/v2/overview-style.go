package session

import (
	"github.com/cufee/aftermath/internal/render/common"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
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
	card       style.Style
	column     style.Style
	styleBlock func(block prepare.StatsBlock[session.BlockData, string]) blockStyle
}

// unrated

var styledOverviewCard = overviewCardStyle{
	styleBlock: styleOverviewBlock,
	card: style.Style{
		Debug: debugOverviewCards,

		Direction:      style.DirectionHorizontal,
		AlignItems:     style.AlignItemsCenter,
		JustifyContent: style.JustifyContentSpaceBetween,

		BackgroundColor: common.DefaultCardColor,
		BlurBackground:  cardBackgroundBlur,

		BorderRadiusTopLeft:     common.BorderRadiusLG,
		BorderRadiusTopRight:    common.BorderRadiusLG,
		BorderRadiusBottomLeft:  common.BorderRadiusLG,
		BorderRadiusBottomRight: common.BorderRadiusLG,

		GrowHorizontal: true,
		Gap:            15,

		PaddingLeft:   cardPaddingX,
		PaddingRight:  cardPaddingX,
		PaddingTop:    cardPaddingY,
		PaddingBottom: cardPaddingY,
	},
	column: style.Style{
		Debug: debugOverviewCards,

		Direction:      style.DirectionVertical,
		AlignItems:     style.AlignItemsCenter,
		JustifyContent: style.JustifyContentSpaceAround,
		GrowVertical:   true,
		Gap:            10,
	},
}

func styleOverviewBlock(block prepare.StatsBlock[session.BlockData, string]) blockStyle {
	switch block.Tag {
	case prepare.TagWN8, prepare.TagRankedRating:
		return blockStyle{
			wrapper: style.Style{
				Debug: debugOverviewCards,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentCenter,
				Gap:            15,
			},
			valueContainer: style.Style{
				Debug: debugOverviewCards,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentEnd,
				Gap:            0,
			},
			value: style.Style{
				Debug: debugOverviewCards,

				PaddingTop: -6,
				Color:      common.TextPrimary,
				Font:       common.FontXL(),
			},
			label: style.Style{
				Color:      common.TextAlt,
				Font:       common.FontSmall(),
				PaddingTop: -6,
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

// wrapped around special block text and icon
var styledOverviewSpecialBlockWrapper = style.Style{
	Debug: debugOverviewCards,

	Direction:      style.DirectionVertical,
	AlignItems:     style.AlignItemsCenter,
	JustifyContent: style.JustifyContentCenter,
	Gap:            10,
}
