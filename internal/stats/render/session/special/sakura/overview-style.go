package sakura

import (
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"

	"github.com/cufee/facepaint/style"
)

var (
	overviewSpecialIconSize = 40.0
)

var overviewStyle overviewStyleType

type overviewStyleType struct{}

func (overviewStyleType) card() style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		Direction:      style.DirectionHorizontal,
		AlignItems:     style.AlignItemsCenter,
		JustifyContent: style.JustifyContentCenter,
		// Debug:          true,

		BackgroundColor: cardBackgroundColor,
		BlurBackground:  cardBackgroundBlur,

		GrowHorizontal: true,
	}),
		style.SetPaddingX(30),
		style.SetPaddingY(20),
		style.SetBorderRadius(25),
	)
}

func (overviewStyleType) column(index int, flavor string) style.StyleOptions {
	align := style.AlignItemsCenter
	if index == 0 {
		align = style.AlignItemsStart
	}
	if index > 1 {
		align = style.AlignItemsEnd
	}

	return style.NewStyle(style.Parent(style.Style{
		// Debug:     true,
		Direction: style.DirectionVertical,

		AlignItems:     align,
		GrowHorizontal: true,
		Gap:            10,
	}))
}

func (overviewStyleType) blockContainer(index int, tag prepare.Tag) style.StyleOptions {
	align := style.AlignItemsCenter
	if index == 0 {
		align = style.AlignItemsStart
	}
	if index > 1 {
		align = style.AlignItemsEnd
	}
	return style.NewStyle(style.Parent(style.Style{
		Direction:      style.DirectionVertical,
		AlignItems:     align,
		GrowHorizontal: true,
		// Debug:          true,
	}))
}

func (overviewStyleType) blockValue(tag prepare.Tag) style.StyleOptions {
	switch tag {
	default:
		return style.NewStyle(style.Parent(style.Style{
			Color: textPrimary,
			Font:  FontLarge(),
		}))

	case prepare.TagWN8, prepare.TagRankedRating:
		return style.NewStyle(style.Parent(style.Style{
			Color: textPrimary,
			Font:  FontXL(),
		}))
	}
}

func (overviewStyleType) blockLabel(tag prepare.Tag) style.StyleOptions {
	switch tag {
	default:
		return style.NewStyle(style.Parent(style.Style{
			Color: textSecondary,
			Font:  FontSmall(),
		}))
	case prepare.TagWN8, prepare.TagRankedRating:
		return style.NewStyle(style.Parent(style.Style{
			Color: textSecondary,
			Font:  FontMedium(),
		}))
	}
}
