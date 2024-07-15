package common

import "image/color"

func getDebugColor() color.Color {
	return color.NRGBA{255, 192, 203, 255}
}

var DefaultLogoColorOptions = []color.Color{color.NRGBA{50, 50, 50, 180}, color.NRGBA{200, 200, 200, 180}}
