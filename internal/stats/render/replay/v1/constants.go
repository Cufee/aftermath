package replay

import (
	"image/color"

	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

var (
	frameStyle = common.Style{Direction: common.DirectionVertical, PaddingX: 20, PaddingY: 20, Gap: 10}

	hpBarColorAllies  = color.NRGBA{R: 120, G: 255, B: 120, A: 255}
	hpBarColorEnemies = color.NRGBA{R: 255, G: 120, B: 120, A: 255}

	protagonistColor = color.NRGBA{255, 223, 0, 255}
)

func defaultCardStyle(width, height float64) common.Style {
	return common.Style{
		Direction:       common.DirectionVertical,
		Width:           width,
		Height:          height,
		BackgroundColor: common.DefaultCardColor,
		PaddingX:        common.BorderRadiusLG,
		BorderRadius:    common.BorderRadiusLG,
		Gap:             15,
		// Debug:           true,
	}
}

func statsRowStyle() common.Style {
	return common.Style{
		Direction: common.DirectionHorizontal,
		Gap:       10,
	}
}
