package render

import (
	"image/color"

	"github.com/cufee/aftermath/internal/stats/frame"
)

func GetWinrateColor(value frame.Value) color.Color {
	if value.Float() > 70.0 {
		return color.NRGBA{165, 105, 189, 255}
	}
	if value.Float() > 60.0 {
		return color.NRGBA{93, 173, 226, 255}
	}
	if value.Float() > 50.0 {
		return color.NRGBA{88, 214, 141, 255}
	}
	if value.Float() > 45.0 {
		return color.NRGBA{244, 208, 63, 255}
	}
	return color.NRGBA{241, 148, 138, 255}
}
