package replay

import (
	"image/color"

	common "github.com/cufee/aftermath/internal/render/v1"
)

var (
	frameStyle = common.Style{Direction: common.DirectionVertical, PaddingX: 20, PaddingY: 20, Gap: 10}

	hpBarWidth          = 8.0
	hpBarHeight         = 45.0
	hpBarColorAllies    = color.NRGBA{R: 120, G: 255, B: 120, A: 255}
	hpBarBgColorAllies  = color.NRGBA{R: 80, G: 120, B: 80, A: 120}
	hpBarColorEnemies   = color.NRGBA{R: 255, G: 120, B: 120, A: 255}
	hpBarBgColorEnemies = color.NRGBA{R: 120, G: 80, B: 80, A: 120}

	protagonistColor = color.NRGBA{255, 223, 0, 255}

	outcomeIconSize         = 30.0
	outcomeIconColorRed     = color.NRGBA{R: 255, G: 120, B: 120, A: 255}
	outcomeIconBgColorRed   = color.NRGBA{R: 120, G: 80, B: 80, A: 120}
	outcomeIconColorGreen   = color.NRGBA{R: 120, G: 255, B: 120, A: 255}
	outcomeIconBgColorGreen = color.NRGBA{R: 80, G: 120, B: 80, A: 120}
	outcomeIconColorYellow  = color.NRGBA{255, 223, 0, 120}

	playerWN8IconSize          = 28.0
	playerRatingIconSize       = 34.0
	playerWinrateIndicatorSize = 7.5
	playerCardPadding          = (80 - hpBarHeight) / 2
)

func defaultCardStyle(width, height float64) common.Style {
	return common.Style{
		Direction:       common.DirectionVertical,
		Width:           width,
		Height:          height,
		BackgroundColor: common.DefaultCardColor,
		PaddingX:        common.BorderRadiusLG,
		BorderRadius:    common.BorderRadiusLG,
		Gap:             10,
		// Debug:           true,
	}
}

func statsRowStyle() common.Style {
	return common.Style{
		PaddingX:   common.BorderRadiusLG - playerCardPadding,
		Direction:  common.DirectionHorizontal,
		AlignItems: common.AlignItemsCenter,
		Gap:        20,
		// Debug:     true,
	}
}
func teamsRowStyle() common.Style {
	return common.Style{
		Direction: common.DirectionHorizontal,
		Gap:       10,
		// Debug:     true,
	}
}
