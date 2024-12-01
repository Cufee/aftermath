package common

import (
	"strings"

	"github.com/fogleman/gg"
)

type stringSize struct {
	TotalWidth  float64
	TotalHeight float64
	LineOffset  float64
	LineHeight  float64
}

func MeasureString(text string, font Font) stringSize {
	if !font.Valid() {
		return stringSize{}
	}
	if text == "" {
		return stringSize{}
	}

	face, close := font.Face()
	defer close()

	measureCtx := gg.NewContext(1, 1)
	measureCtx.SetFontFace(face)

	var result stringSize
	// Account for font descender height
	result.LineOffset = float64(face.Metrics().Descent>>6) * 2

	for _, line := range strings.Split(text, "\n") {
		w, h := measureCtx.MeasureString(line)
		h += result.LineOffset
		w += 1

		if w > result.TotalWidth {
			result.TotalWidth = w
		}
		if h > result.LineHeight {
			result.LineHeight = h
		}

		result.TotalHeight += h
	}

	return result
}
