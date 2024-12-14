package replay

import (
	"image/color"

	common "github.com/cufee/aftermath/internal/render/v1"
	"github.com/fogleman/gg"
)

type progressDirection int

const (
	progressDirectionHorizontal progressDirection = iota
	progressDirectionVertical
)

func newProgressBar(size int, progress int, direction progressDirection, fillColor color.Color, bgColor color.Color) common.Block {
	var width, height int
	if direction == progressDirectionHorizontal {
		width = (size)
		height = int(hpBarWidth)
	} else {
		width = int(hpBarWidth)
		height = (size)
	}

	ctx := gg.NewContext((width), (height))
	ctx.SetColor(bgColor)
	ctx.DrawRoundedRectangle(0, 0, float64(width), float64(height), 5)
	ctx.Fill()

	if progress > 0 {
		ctx.SetColor(fillColor)
		if direction == progressDirectionHorizontal {
			ctx.DrawRoundedRectangle(0, 0, float64(progress)/100*float64(width), float64(height), 5)
		} else {
			ctx.DrawRoundedRectangle(0, float64(height)-float64(progress)/100*float64(height), float64(width), float64(progress)/100*float64(height), 5)
		}
		ctx.Fill()
	}

	if direction == progressDirectionHorizontal {
		return common.NewImageContent(common.Style{Width: float64(size), Height: hpBarWidth}, ctx.Image())
	}
	return common.NewImageContent(common.Style{Width: hpBarWidth, Height: float64(size)}, ctx.Image())
}
