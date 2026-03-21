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
	iconWrapper    style.Style
	valueContainer style.Style
	value          style.Style
	label          style.Style
}

type overviewCardStyle struct {
	card       style.Style
	styleBlock func(block prepare.StatsBlock[session.BlockData, string]) blockStyle
}

func newOverviewCardStyle(theme common.Theme) overviewCardStyle {
	return overviewCardStyle{
		styleBlock: newStyleOverviewBlock(theme),
		card: common.ApplyTheme(style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionHorizontal,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentSpaceBetween,

			PaddingLeft:   common.CardPaddingX,
			PaddingRight:  common.CardPaddingX,
			PaddingTop:    common.CardPaddingY / 2,
			PaddingBottom: common.CardPaddingY / 2,

			GrowHorizontal: true,
			Gap:            15,
		}, theme.Card),
	}
}

func (overviewCardStyle) column(column session.OverviewColumn) style.Style {
	switch column.Flavor {
	case session.BlockFlavorRating, session.BlockFlavorWN8:
		return style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionVertical,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentCenter,
			GrowHorizontal: true,

			BorderRadiusTopLeft:     common.BorderRadiusSM,
			BorderRadiusTopRight:    common.BorderRadiusSM,
			BorderRadiusBottomLeft:  common.BorderRadiusSM,
			BorderRadiusBottomRight: common.BorderRadiusSM,

			PaddingLeft:   10,
			PaddingRight:  10,
			PaddingTop:    common.CardPaddingY / 2,
			PaddingBottom: common.CardPaddingY / 2,

			Gap: 15,
		}
	default:
		return style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionVertical,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentCenter,
			Gap:            15,

			PaddingTop:    common.CardPaddingY / 2,
			PaddingBottom: common.CardPaddingY / 2,
		}
	}
}

func newStyleOverviewBlock(theme common.Theme) func(prepare.StatsBlock[session.BlockData, string]) blockStyle {
	return func(block prepare.StatsBlock[session.BlockData, string]) blockStyle {
		defaultStyle := blockStyle{
			wrapper: style.Style{},
			valueContainer: style.Style{
				Debug: debugOverviewCards,

				Gap: 6,

				Direction:      style.DirectionVertical,
				AlignItems:     style.AlignItemsCenter,
				JustifyContent: style.JustifyContentCenter,
			},
			value: common.ApplyTheme(style.Style{
				Debug: debugOverviewCards,
				Font:  common.FontLarge(),
			}, theme.TextPrimary()),
			label: common.ApplyTheme(style.Style{
				Font:       common.FontSmall(),
				PaddingTop: -5,
			}, theme.TextAlt()),
		}

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
				value: common.ApplyTheme(style.Style{
					Debug:      debugOverviewCards,
					PaddingTop: -6,
					Font:       common.FontXL(),
				}, theme.TextPrimary()),
				label: common.ApplyTheme(style.Style{
					Font:       common.FontSmall(),
					PaddingTop: -6,
				}, theme.TextAlt()),
			}
		default:
			return defaultStyle
		}
	}
}

var styledOverviewSpecialBlockWrapper = style.Style{
	Debug: debugOverviewCards,

	Direction:      style.DirectionVertical,
	AlignItems:     style.AlignItemsCenter,
	JustifyContent: style.JustifyContentCenter,
	Gap:            10,
}
