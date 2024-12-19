package render

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
)

type frameOptions struct {
	shadowSize   float64
	background   image.Image
	borderRadius float64
	width        float64
}

type FrameOption func(o *frameOptions)

func WithFrameBackground(img image.Image) FrameOption {
	return func(o *frameOptions) {
		o.background = img
	}
}
func WithFrameWidth(width float64) FrameOption {
	return func(o *frameOptions) {
		o.width = width
	}
}
func WithBorderRadius(radius float64) FrameOption {
	return func(o *frameOptions) {
		o.borderRadius = radius
	}
}
func WithShadow(size float64) FrameOption {
	return func(o *frameOptions) {
		o.shadowSize = size
	}
}

type Frame struct {
	Width   float64
	options frameOptions

	contextWidth  float64
	contextHeight float64
	targetWidth   float64
	targetHeight  float64

	content   image.Image
	baseLayer *gg.Context
	layers    []*gg.Context
	topLayer  *gg.Context
}

func NewFrameContext(content image.Image, opts ...FrameOption) Frame {
	o := frameOptions{width: 40}
	for _, apply := range opts {
		apply(&o)
	}

	contentWidth, contentHeight := float64(content.Bounds().Dx()), float64(content.Bounds().Dy())

	frame := Frame{
		options: o,
		content: content,
		Width:   o.width,
		// contextWidth: contentWidth,
		contextWidth: contentWidth + o.width,
		// contextHeight: contentHeight,
		contextHeight: contentHeight + o.width,
		targetWidth:   contentWidth,
		targetHeight:  contentHeight,
		// shadows:      gg.NewContext(int(contentWidth), int(contentHeight)),
		// shadows: gg.NewContext(int(contentWidth+o.width), int(contentHeight+o.width)),
		// content: gg.NewContext(int(contentWidth), int(contentHeight)),
		// content: gg.NewContext(int(contentWidth+o.width), int(contentHeight+o.width)),
	}

	frame.baseLayer = gg.NewContext(int(frame.contextWidth+o.width), int(frame.contextHeight+o.width))
	// clip the target content
	frame.baseLayer.DrawRoundedRectangle(frame.Width, frame.Width, frame.targetWidth, frame.targetHeight, frame.options.borderRadius)
	frame.baseLayer.Clip()
	frame.baseLayer.InvertMask()
	// clip the border radius
	clipRoundedRect(frame.baseLayer, frame.options.borderRadius)
	// draw background
	if frame.options.background != nil {
		bg := imaging.Fill(frame.options.background, int(frame.contextWidth), int(frame.contextHeight), imaging.TopLeft, imaging.Linear)
		frame.baseLayer.DrawImage(bg, 0, 0)
	}

	if frame.options.shadowSize > 0 {
		{ // frame content shadow layer
			shadowMask := gg.NewContext(int(frame.contextWidth+10), int(frame.contextHeight+10))
			shadowMask.DrawRoundedRectangle(frame.Width+5, frame.Width+5, frame.targetWidth, frame.targetHeight, frame.options.borderRadius)
			shadowMask.Clip()
			shadowMask.InvertMask()
			shadowMask.DrawRoundedRectangle(5, 5, frame.contextWidth, frame.contextHeight, frame.options.borderRadius)
			shadowMask.Clip()

			shadowMask.DrawRectangle(0, 0, float64(shadowMask.Width()), float64(shadowMask.Height()))
			shadowMask.SetColor(color.White)
			shadowMask.Fill()

			shadowCaster := gg.NewContext(shadowMask.Width(), shadowMask.Height())
			shadowCaster.SetMask(shadowMask.AsMask())
			shadowCaster.InvertMask()
			shadowCaster.DrawRectangle(0, 0, float64(shadowCaster.Width()), float64(shadowCaster.Height()))
			shadowCaster.SetColor(color.Black)
			shadowCaster.Fill()
			shadowImg := imaging.Blur(shadowCaster.Image(), frame.options.shadowSize)

			shadows := gg.NewContext(int(frame.contextWidth), int(frame.contextHeight))
			shadows.DrawImage(shadowImg, -5, -5)
			frame.layers = append(frame.layers, shadows)
		}

		{ // container content shadow layer
			shadowMask := gg.NewContext(int(frame.contextWidth+10), int(frame.contextHeight+10))
			shadowMask.DrawRoundedRectangle(5, 5, frame.contextWidth, frame.contextHeight, frame.options.borderRadius)
			shadowMask.SetColor(color.White)
			shadowMask.Fill()

			shadowCaster := gg.NewContext(shadowMask.Width(), shadowMask.Height())
			shadowCaster.SetMask(shadowMask.AsMask())
			shadowCaster.InvertMask()
			shadowCaster.DrawRectangle(0, 0, float64(shadowCaster.Width()), float64(shadowCaster.Height()))
			shadowCaster.SetColor(color.Black)
			shadowCaster.Fill()
			shadowImg := imaging.Blur(shadowCaster.Image(), frame.options.shadowSize)

			shadows := gg.NewContext(int(frame.contextWidth), int(frame.contextHeight))
			shadows.DrawImage(shadowImg, -5, -5)
			frame.topLayer = shadows
		}
	}

	return frame
}

func (f *Frame) Ctx() *gg.Context {
	return f.baseLayer
}

func (f *Frame) Image() image.Image {
	out := gg.NewContext(int(f.contextWidth), int(f.contextHeight))
	clipRoundedRect(out, f.options.borderRadius)

	cctx := gg.NewContext(f.content.Bounds().Dx(), f.content.Bounds().Dy()) // make a content context for clipping
	clipRoundedRect(cctx, f.options.borderRadius)                           // clip border radius
	cctx.DrawImage(f.content, 0, 0)

	out.DrawImage(f.baseLayer.Image(), 0, 0) // frame content into the image
	for _, layer := range f.layers {         // draw all layers
		out.DrawImage(layer.Image(), 0, 0)
	}

	out.DrawImage(cctx.Image(), int(f.Width), int(f.Width)) // draw framed content into final image
	if f.topLayer != nil {                                  // draw top layer
		out.DrawImage(f.topLayer.Image(), 0, 0)
	}

	// out.DrawImage(cctx.Image(), int(f.Width), int(f.Width)) // draw stats content

	return out.Image()
}
