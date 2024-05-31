package period

import (
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common"
	"github.com/cufee/aftermath/internal/stats/prepare/period"
	"github.com/cufee/aftermath/internal/stats/render/common"
)

type overviewStyle struct {
	container      common.Style
	blockContainer common.Style
}

type highlightStyle struct {
	container  common.Style
	cardTitle  common.Style
	tankName   common.Style
	blockLabel common.Style
	blockValue common.Style
}

func (s *overviewStyle) block(block prepare.StatsBlock[period.BlockData]) (common.Style, common.Style) {
	switch block.Data.Flavor {
	case period.BlockFlavorSpecial:
		return common.Style{FontColor: common.TextPrimary, Font: &common.FontXL}, common.Style{FontColor: common.TextAlt, Font: &common.FontSmall}
	case period.BlockFlavorSecondary:
		return common.Style{FontColor: common.TextSecondary, Font: &common.FontMedium}, common.Style{FontColor: common.TextAlt, Font: &common.FontSmall}
	default:
		return common.Style{FontColor: common.TextPrimary, Font: &common.FontLarge}, common.Style{FontColor: common.TextAlt, Font: &common.FontSmall}
	}
}

func getOverviewStyle(columnWidth float64) overviewStyle {
	return overviewStyle{common.Style{
		Direction:      common.DirectionVertical,
		AlignItems:     common.AlignItemsCenter,
		JustifyContent: common.JustifyContentCenter,
		PaddingX:       15,
		PaddingY:       0,
		Gap:            10,
		Width:          columnWidth,
	}, common.Style{
		Direction:  common.DirectionVertical,
		AlignItems: common.AlignItemsCenter,
		// Debug:      true,
	}}
}

func defaultCardStyle(width float64) common.Style {
	style := common.Style{
		JustifyContent:  common.JustifyContentCenter,
		AlignItems:      common.AlignItemsCenter,
		Direction:       common.DirectionVertical,
		BackgroundColor: common.DefaultCardColor,
		BorderRadius:    20,
		PaddingY:        10,
		PaddingX:        20,
		Gap:             20,
		Width:           width,
		// Debug:           true,
	}
	return style
}

func titleCardStyle(width float64) common.Style {
	style := defaultCardStyle(width)
	style.PaddingX = style.PaddingY
	return style
}

func overviewCardStyle(width float64) common.Style {
	style := defaultCardStyle(width)
	style.AlignItems = common.AlignItemsEnd
	style.Direction = common.DirectionHorizontal
	style.JustifyContent = common.JustifyContentSpaceAround
	style.PaddingY = 25
	style.PaddingX = 0
	style.Gap = 0
	// style.Debug = true
	return style
}

func highlightCardStyle(containerStyle common.Style) highlightStyle {
	container := containerStyle
	container.Gap = 5
	container.PaddingX = 20
	container.PaddingY = 15
	container.Direction = common.DirectionHorizontal
	container.JustifyContent = common.JustifyContentSpaceBetween

	return highlightStyle{
		container:  container,
		cardTitle:  common.Style{Font: &common.FontSmall, FontColor: common.TextSecondary},
		tankName:   common.Style{Font: &common.FontMedium, FontColor: common.TextPrimary},
		blockValue: common.Style{Font: &common.FontMedium, FontColor: common.TextPrimary},
		blockLabel: common.Style{Font: &common.FontSmall, FontColor: common.TextAlt},
	}
}
