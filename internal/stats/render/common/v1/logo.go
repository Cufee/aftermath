package common

import (
	"image"
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

type LogoSizingOptions struct {
	Lines     int
	Gap       float64
	BaseWidth float64
}

func (opts LogoSizingOptions) LineHeightAt(i int) float64 {
	h := opts.BaseWidth + (opts.BaseWidth / 2 * math.Pow(float64(i+1), 1.75))

	if i > opts.Lines/2 {
		h = opts.LineHeightAt(opts.Lines - i - 1)
	}
	return h
}
func (opts LogoSizingOptions) LineWidthAt(i int) float64 {
	return opts.BaseWidth
}
func (opts LogoSizingOptions) LineOffsetAt(i int) float64 {
	o := opts.BaseWidth * math.Pow(float64(i), 1.25)
	if i > opts.Lines/2 {
		o = opts.LineOffsetAt(opts.Lines - i - 1)
	}
	return o
}

func (opts LogoSizingOptions) Height() int {
	var maxHeight int
	for line := range opts.Lines {
		maxHeight = Max(maxHeight, int(math.RoundToEven(opts.LineHeightAt(line)+opts.LineOffsetAt(line))))
	}
	return maxHeight
}
func (opts LogoSizingOptions) Width() int {
	var totalWidth int = int(math.RoundToEven(opts.Gap * float64(opts.Lines-1)))
	for line := range opts.Lines {
		totalWidth += int(math.RoundToEven(opts.LineWidthAt(line)))
	}
	return totalWidth
}

func DefaultLogoOptions() LogoSizingOptions {
	return LogoSizingOptions{
		Lines:     7,
		BaseWidth: 6,
		Gap:       3,
	}
}

func SmallLogoOptions() LogoSizingOptions {
	return LogoSizingOptions{
		Lines:     5,
		BaseWidth: 6,
		Gap:       3,
	}
}

func LargeLogoOptions() LogoSizingOptions {
	return LogoSizingOptions{
		Lines:     7,
		BaseWidth: 60,
		Gap:       30,
	}
}

func AftermathLogo(fillColor color.Color, opts LogoSizingOptions) image.Image {
	ctx := gg.NewContext(opts.Width(), opts.Height())
	for line := range opts.Lines {
		var xPos = 0.0
		for prev := range line {
			xPos += (opts.LineWidthAt(prev) + opts.Gap)
		}
		height := opts.LineHeightAt(line)
		ctx.DrawRoundedRectangle(
			xPos, // x
			float64(opts.Height())-opts.LineOffsetAt(line)-height, // y
			opts.LineWidthAt(line),  // w
			opts.LineHeightAt(line), // h
			opts.BaseWidth/2,        // border radius
		)
		ctx.SetColor(fillColor)
		ctx.Fill()
		ctx.ClearPath()
	}

	return ctx.Image()
}
