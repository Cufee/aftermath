package session

import (
	"image/color"

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

var (
	iconBackgroundColorOverview = color.NRGBA{40, 40, 40, 50}
)

type blockStyle struct {
	wrapper        style.Style
	iconWrapper    style.Style
	valueContainer style.Style
	value          style.Style
	label          style.Style
}

type overviewCardStyle struct {
	card       style.Style
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

		PaddingLeft:   cardPaddingX,
		PaddingRight:  cardPaddingX,
		PaddingTop:    cardPaddingY / 2,
		PaddingBottom: cardPaddingY / 2,

		GrowHorizontal: true,
		Gap:            15,
	},
}

func (overviewCardStyle) column(column session.OverviewColumn) style.Style {
	switch column.Flavor {
	case session.BlockFlavorRating, session.BlockFlavorWN8:
		return style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionVertical,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentCenter,
			// GrowVertical:   true,
			GrowHorizontal: true,

			BlurBackground: cardBackgroundBlur,

			BorderRadiusTopLeft:     common.BorderRadiusSM,
			BorderRadiusTopRight:    common.BorderRadiusSM,
			BorderRadiusBottomLeft:  common.BorderRadiusSM,
			BorderRadiusBottomRight: common.BorderRadiusSM,

			BackgroundColor: iconBackgroundColorOverview,

			PaddingLeft:   20,
			PaddingRight:  20,
			PaddingTop:    cardPaddingY / 2,
			PaddingBottom: cardPaddingY / 2,

			Gap: 15,
		}
	default:
		return style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionVertical,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentCenter,
			Gap:            15,

			PaddingTop:    cardPaddingY / 2,
			PaddingBottom: cardPaddingY / 2,
		}
	}
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
				Gap:            10,

				GrowVertical: true,
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
