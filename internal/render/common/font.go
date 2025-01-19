package common

import "github.com/cufee/facepaint/style"

var defaultFont []byte
var fonts = make(map[float64]style.Font)

func Font2XL() style.Font {
	return getFont(36)
}
func FontXL() style.Font {
	return getFont(32)
}
func FontLarge() style.Font {
	return getFont(24)
}
func FontMedium() style.Font {
	return getFont(18)
}
func FontSmall() style.Font {
	return getFont(14)
}

func getFont(size float64) style.Font {
	if fonts[size] == nil {
		fonts[size], _ = style.NewFont(defaultFont, size)
	}
	return fonts[size]
}
