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

func newRatingOverviewCardStyle(theme common.Theme) overviewCardStyle {
	return overviewCardStyle{
		styleBlock: newStyleRatingOverviewBlock(theme),
		card:       newUnratedOverviewCardStyle(theme).card,
	}
}

func newStyleRatingOverviewBlock(theme common.Theme) func(prepare.StatsBlock[period.BlockData, string]) blockStyle {
	unrated := newStyleUnratedOverviewBlock(theme)
	return func(block prepare.StatsBlock[period.BlockData, string]) blockStyle {
		stl := unrated(block)
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
}

func newUnratedOverviewCardStyle(theme common.Theme) overviewCardStyle {
	return overviewCardStyle{
		styleBlock: newStyleUnratedOverviewBlock(theme),
		card: common.ApplyTheme(style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionHorizontal,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentSpaceBetween,

			GrowHorizontal: true,
			Gap:            10,

			PaddingLeft:   common.CardPaddingX,
			PaddingRight:  common.CardPaddingX,
			PaddingTop:    common.CardPaddingY / 2,
			PaddingBottom: common.CardPaddingY / 2,
		}, theme.Card),
	}
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
			PaddingTop:    common.CardPaddingY / 2,
			PaddingBottom: common.CardPaddingY / 2,
		}
	default:
		return style.Style{
			Debug: debugOverviewCards,

			Direction:      style.DirectionVertical,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentCenter,
			GrowVertical:   false,
			Gap:            10,

			PaddingTop:    common.CardPaddingY / 2,
			PaddingBottom: common.CardPaddingY / 2,
		}
	}
}

func newStyleUnratedOverviewBlock(theme common.Theme) func(prepare.StatsBlock[period.BlockData, string]) blockStyle {
	return func(block prepare.StatsBlock[period.BlockData, string]) blockStyle {
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
					Gap:            5,
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
		case period.BlockFlavorSecondary:
			return blockStyle{
				wrapper: style.Style{},
				valueContainer: style.Style{
					Debug: debugOverviewCards,

					Direction:      style.DirectionVertical,
					AlignItems:     style.AlignItemsCenter,
					JustifyContent: style.JustifyContentCenter,
				},
				value: common.ApplyTheme(style.Style{
					Debug: debugOverviewCards,
					Font:  common.FontMedium(),
				}, theme.TextSecondary()),
				label: common.ApplyTheme(style.Style{
					Font:       common.FontSmall(),
					PaddingTop: -3,
				}, theme.TextAlt()),
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
				value: common.ApplyTheme(style.Style{
					Debug: debugOverviewCards,
					Font:  common.FontLarge(),
				}, theme.TextPrimary()),
				label: common.ApplyTheme(style.Style{
					Font:       common.FontSmall(),
					PaddingTop: -5,
				}, theme.TextAlt()),
			}
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
