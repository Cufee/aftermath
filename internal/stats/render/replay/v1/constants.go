package replay

import (
	"image/color"

	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

var (
	frameStyle = common.Style{Direction: common.DirectionVertical, PaddingX: 30, PaddingY: 30, Gap: 10}

	hpBarColorAllies  = color.NRGBA{R: 120, G: 255, B: 120, A: 255}
	hpBarColorEnemies = color.NRGBA{R: 255, G: 120, B: 120, A: 255}

	protagonistColor = color.NRGBA{255, 223, 0, 255}
)

func defaultCardStyle(width, height float64) common.Style {
	return common.Style{
		Direction:       common.DirectionVertical,
		Width:           width + 40,
		Height:          height,
		BackgroundColor: common.DefaultCardColor,
		PaddingX:        10,
		BorderRadius:    15,
	}
}
