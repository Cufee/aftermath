package session

import (
	"image"
	"image/color"

	"github.com/cufee/aftermath/internal/stats/render/common"
	"github.com/fogleman/gg"
)

type blockStyle struct {
	session common.Style
	career  common.Style
	label   common.Style
}

func init() {
	{
		ctx := gg.NewContext(iconSize, iconSize)
		ctx.DrawRoundedRectangle(13, 2.5, 6, 17.5, 3)
		ctx.SetColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
		ctx.Fill()
		wn8Icon = ctx.Image()
	}

	{
		ctx := gg.NewContext(iconSize, 1)
		blankIconBlock = common.NewImageContent(common.Style{Width: float64(iconSize), Height: 1}, ctx.Image())
	}
}

var (
	iconSize              = 25
	specialRatingIconSize = 60.0
	wn8Icon               image.Image
	blankIconBlock        common.Block
)

var (
	specialRatingColumnStyle = common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter, Gap: 5}
	promoTextStyle           = common.Style{Font: &common.FontMedium, FontColor: common.TextPrimary}
)

func frameStyle() common.Style {
	return common.Style{Gap: 20, Direction: common.DirectionHorizontal}
}

var (
	overviewCardTitleStyle  = common.Style{Font: &common.FontMedium, FontColor: common.TextAlt, PaddingX: 5}
	overviewStatsBlockStyle = blockStyle{
		common.Style{Font: &common.FontLarge, FontColor: common.TextPrimary},
		common.Style{Font: &common.FontMedium, FontColor: common.TextSecondary},
		common.Style{Font: &common.FontSmall, FontColor: common.TextAlt},
	}
)

func overviewSpecialRatingLabelStyle(color color.Color) common.Style {
	return common.Style{FontColor: color, Font: &common.FontSmall}
}

func overviewSpecialRatingPillStyle(color color.Color) common.Style {
	return common.Style{
		PaddingY:        2,
		PaddingX:        7.5,
		BorderRadius:    common.BorderRadiusXS,
		BackgroundColor: color,
	}
}

func overviewColumnStyle(width float64) common.Style {
	return common.Style{
		Gap:            5,
		Width:          width,
		AlignItems:     common.AlignItemsCenter,
		Direction:      common.DirectionVertical,
		JustifyContent: common.JustifyContentCenter,
	}
}

func overviewCardStyle(width float64) common.Style {
	style := defaultCardStyle(width)
	style.JustifyContent = common.JustifyContentSpaceAround
	style.Direction = common.DirectionHorizontal
	style.AlignItems = common.AlignItemsEnd
	style.PaddingY = 20
	style.PaddingX = 10
	style.Gap = 5
	// style.Debug = true
	return style
}

func overviewRatingCardStyle(width float64) common.Style {
	style := overviewCardStyle(width)
	style.AlignItems = common.AlignItemsCenter
	return style
}

func statsBlockStyle(width float64) common.Style {
	return common.Style{
		JustifyContent: common.JustifyContentCenter,
		Direction:      common.DirectionVertical,
		AlignItems:     common.AlignItemsCenter,
		Width:          width,
	}
}

var (
	vehicleCardTitleStyle = common.Style{Font: &common.FontMedium, FontColor: common.TextAlt, PaddingX: 5}
	vehicleBlockStyle     = blockStyle{
		common.Style{Font: &common.FontLarge, FontColor: common.TextPrimary},
		common.Style{Font: &common.FontMedium, FontColor: common.TextSecondary},
		common.Style{Font: &common.FontSmall, FontColor: common.TextAlt},
	}
)

func vehicleCardTitleContainerStyle(width float64) common.Style {
	return common.Style{
		JustifyContent: common.JustifyContentSpaceBetween,
		Direction:      common.DirectionHorizontal,
		PaddingX:       defaultCardStyle(0).PaddingX * 2,
		Width:          width,
		Gap:            10,
	}
}

func vehicleCardStyle(width float64) common.Style {
	style := defaultCardStyle(width)
	style.Gap = 5
	return style
}

func vehicleBlocksRowStyle(width float64) common.Style {
	return common.Style{
		JustifyContent: common.JustifyContentSpaceBetween,
		Direction:      common.DirectionHorizontal,
		AlignItems:     common.AlignItemsCenter,
		Width:          width,
	}
}

var (
	ratingVehicleCardTitleContainerStyle = common.Style{Direction: common.DirectionHorizontal, Gap: 10, JustifyContent: common.JustifyContentSpaceBetween}
	ratingVehicleCardTitleStyle          = common.Style{Font: &common.FontMedium, FontColor: common.TextSecondary, PaddingX: 5}
	ratingVehicleBlockStyle              = blockStyle{
		common.Style{Font: &common.FontLarge, FontColor: common.TextPrimary},
		common.Style{Font: &common.FontMedium, FontColor: common.TextSecondary},
		common.Style{Font: &common.FontSmall, FontColor: common.TextAlt},
	}
)

func ratingVehicleCardStyle(width float64) common.Style {
	return defaultCardStyle(width)
}

func ratingVehicleBlocksRowStyle(width float64) common.Style {
	return vehicleBlocksRowStyle(width)
}

func highlightCardStyle(width float64) common.Style {
	return defaultCardStyle(width)
}

func highlightedVehicleCardStyle(width float64) common.Style {
	return vehicleCardStyle(width)
}

func defaultCardStyle(width float64) common.Style {
	style := common.Style{
		JustifyContent:  common.JustifyContentCenter,
		AlignItems:      common.AlignItemsCenter,
		Direction:       common.DirectionVertical,
		PaddingX:        10,
		PaddingY:        15,
		BackgroundColor: common.DefaultCardColor,
		BorderRadius:    common.BorderRadiusLG,
		Width:           width,
		// Debug:           true,
	}
	return style
}

func playerNameCardStyle(width float64) common.Style {
	style := defaultCardStyle(width)
	style.PaddingX, style.PaddingY = 10, 10
	style.Gap = 10
	return style
}
