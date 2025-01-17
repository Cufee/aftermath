package render

import (
	"image/color"
	"time"
)

func getDebugColor() color.Color {
	ns := time.Now().Nanosecond()
	return color.NRGBA{uint8(ns%120) + 120, uint8(ns % 200), uint8(ns % 200), 255}
}
