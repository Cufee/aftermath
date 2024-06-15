package session

import (
	"image"
	"image/color"

	"github.com/cufee/aftermath/internal/stats/render/common"
	"github.com/fogleman/gg"
)

type blockStyle struct {
	container common.Style
	session   common.Style
	career    common.Style
	label     common.Style
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
	iconSize       = 25
	wn8Icon        image.Image
	blankIconBlock common.Block
)

var (
	promoTextStyle = common.Style{Font: &common.FontMedium, FontColor: common.TextPrimary}
)

func frameStyle() common.Style {
	return common.Style{Gap: 10, Direction: common.DirectionHorizontal}
}

func columnStyle(width float64) common.Style {
	return common.Style{Gap: 10, Direction: common.DirectionVertical, Width: width}
}

var (
	vehicleCardTitleStyle = common.Style{Font: &common.FontLarge, FontColor: common.TextSecondary, PaddingX: 5}
	vehicleBlockStyle     = blockStyle{
		common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter},
		common.Style{Font: &common.FontLarge, FontColor: common.TextPrimary},
		common.Style{Font: &common.FontMedium, FontColor: common.TextSecondary},
		common.Style{Font: &common.FontSmall, FontColor: common.TextAlt},
	}
)

func vehicleCardStyle(width float64) common.Style {
	return defaultCardStyle(width)
}

func vehicleBlocksRowStyle(width float64) common.Style {
	return common.Style{
		JustifyContent: common.JustifyContentSpaceBetween,
		Direction:      common.DirectionHorizontal,
		AlignItems:     common.AlignItemsCenter,
		Width:          width,
		Gap:            10,
	}
}

var (
	ratingVehicleCardTitleStyle = common.Style{Font: &common.FontMedium, FontColor: common.TextSecondary, PaddingX: 5}
	ratingVehicleBlockStyle     = blockStyle{
		common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter},
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
		PaddingX:        20,
		PaddingY:        20,
		BackgroundColor: common.DefaultCardColor,
		BorderRadius:    20,
		Width:           width,
		// Debug:           true,
	}
	return style
}

func playerNameCardStyle(width float64) common.Style {
	style := defaultCardStyle(width)
	style.PaddingX, style.PaddingY = 10, 10
	return style
}
