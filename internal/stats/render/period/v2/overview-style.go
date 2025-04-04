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
	iconSizeRating = 64.0
)

type blockStyle struct {
	wrapper        style.Style
	valueContainer style.Style
	value          style.Style
	label          style.Style
}

type overviewCardStyle struct {
	card       style.Style
	styleBlock func(block prepare.StatsBlock[period.BlockData, string]) blockStyle
}

// rating

var styledRatingOverviewCard = overviewCardStyle{
	styleBlock: styleRatingOverviewBlock,
	card:       styledUnratedOverviewCard.card,
}

func styleRatingOverviewBlock(block prepare.StatsBlock[period.BlockData, string]) blockStyle {
	stl := styleUnratedOverviewBlock(block)
	if block.Data.Flavor != period.BlockFlavorSpecial {
		return stl
	}
	stl.wrapper = style.Style{
		Debug: debugOverviewCards,

		Direction:      style.DirectionVertical,
		AlignItems:     style.AlignItemsCenter,
		JustifyContent: style.JustifyContentCenter,
		Gap:            10,
	}
	return stl
}

// unrated

var styledUnratedOverviewCard = overviewCardStyle{
	styleBlock: styleUnratedOverviewBlock,
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
		Gap:            10,

		PaddingLeft:   cardPaddingX,
		PaddingRight:  cardPaddingX,
		PaddingTop:    cardPaddingY / 2,
		PaddingBottom: cardPaddingY / 2,
	},
}

func (overviewCardStyle) column(column period.OverviewColumn) style.Style {
	switch column.Flavor {
	case period.BlockFlavorSpecial:
		return style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionVertical,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentCenter,
			GrowVertical:   true,
			GrowHorizontal: true,
			Gap:            15,

			PaddingLeft:   10,
			PaddingRight:  10,
			PaddingTop:    cardPaddingY / 2,
			PaddingBottom: cardPaddingY / 2,
		}
	default:
		return style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionVertical,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentCenter,
			GrowVertical:   false,
			Gap:            10,

			PaddingTop:    cardPaddingY / 2,
			PaddingBottom: cardPaddingY / 2,
		}
	}
}

func styleUnratedOverviewBlock(block prepare.StatsBlock[period.BlockData, string]) blockStyle {
	switch block.Data.Flavor {
	case period.BlockFlavorSpecial:
		return blockStyle{
			wrapper: style.Style{
				Debug: debugOverviewCards,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentSpaceAround,
				GrowVertical:   true,
				Gap:            5,
			},
			valueContainer: style.Style{
				Debug: debugOverviewCards,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentEnd,
				// GrowVertical:   true,
				Gap: 5,
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

// wrapped around special block text and icon
var styledOverviewSpecialBlockWrapper = style.Style{
	Debug: debugOverviewCards,

	Direction:      style.DirectionVertical,
	AlignItems:     style.AlignItemsCenter,
	JustifyContent: style.JustifyContentCenter,
	Gap:            10,
}
