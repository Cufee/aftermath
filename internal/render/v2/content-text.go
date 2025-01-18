package render

import (
	"math"
	"strings"

	"github.com/cufee/aftermath/internal/render/v2/style"
	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
	"github.com/pkg/errors"
)

var _ BlockContent = &contentText{}

func NewTextContent(style style.StyleOptions, value string) (*Block, error) {
	computed := style.Computed()
	if !computed.Font.Valid() {
		return nil, errors.New("invalid or missing font")
	}
	if computed.Color == nil {
		return nil, errors.New("text requires a non nil color")
	}
	return NewBlock(&contentText{
		value: value,
		style: style,
	}), nil
}

func MustNewTextContent(style style.StyleOptions, value string) *Block {
	c, _ := NewTextContent(style, value)
	return c
}

type contentText struct {
	style style.StyleOptions
	value string

	dimensionsCache *contentDimensions // add cache to avoid parsing and rendering fonts repeatedly
	sizeCache       *StringSize        // add cache to avoid parsing and rendering fonts repeatedly
}

func (content *contentText) setStyle(style style.StyleOptions) {
	content.dimensionsCache = nil
	content.sizeCache = nil
	content.style = style
}

func (content *contentText) measure(font style.Font) StringSize {
	if content.sizeCache != nil {
		return *content.sizeCache
	}

	size := MeasureString(content.value, font)
	content.sizeCache = &size
	return size
}

func (content *contentText) dimensions() contentDimensions {
	if content.dimensionsCache != nil {
		return *content.dimensionsCache
	}

	computed := content.style.Computed()
	size := content.measure(computed.Font)

	var width, height = 0.0, 0.0
	if computed.Width > 0 {
		width = computed.Width
	} else {
		width = size.TotalWidth + (computed.PaddingLeft + computed.PaddingRight)
	}
	if computed.Height > 0 {
		height = computed.Height
	} else {
		height = size.TotalHeight + (computed.PaddingTop + computed.PaddingBottom)
	}

	content.dimensionsCache = &contentDimensions{width: int(math.Ceil(width)), height: int(math.Ceil(height))}
	return *content.dimensionsCache
}

func (content *contentText) Type() blockContentType {
	return BlockContentTypeText
}

func (content *contentText) Style() style.StyleOptions {
	return content.style
}

func (content *contentText) Render(ctx *gg.Context, pos Position) error {
	computed := content.style.Computed()
	size := content.measure(computed.Font)
	dimensions := content.dimensions()

	if computed.Blur > 0 {
		// reset blur
		current := content.Style()
		current.Add(style.SetBlur(0))
		content.setStyle(current)

		// render the content onto a new image, blur it, render onto parent
		child := gg.NewContext(dimensions.width, dimensions.height)
		err := content.Render(child, Position{0, 0})
		if err != nil {
			return err
		}
		img := imaging.Blur(child.Image(), computed.Blur)
		ctx.DrawImage(img, ceil(pos.X), ceil(pos.Y))
		return nil
	}

	if computed.BackgroundColor != nil {
		ctx.SetColor(computed.BackgroundColor)
		ctx.DrawRectangle(pos.X, pos.Y, float64(dimensions.width), float64(dimensions.height))
		ctx.Fill()
	}
	if computed.BackgroundImage != nil {
		background := imaging.Fill(computed.BackgroundImage, dimensions.width, dimensions.height, imaging.Center, imaging.Lanczos)
		ctx.DrawImage(background, ceil(pos.X), ceil(pos.Y))
	}

	if computed.Debug {
		ctx.SetColor(getDebugColor())
		ctx.DrawRectangle(pos.X, pos.Y, float64(dimensions.width), float64(dimensions.height))
		ctx.Stroke()
	}

	var lastX, lastY float64 = pos.X + computed.PaddingLeft, pos.Y + computed.PaddingTop + 1

	switch computed.JustifyContent {
	case style.JustifyContentEnd:
		lastX += float64(dimensions.width) - size.TotalWidth
	case style.JustifyContentCenter:
		lastX += (float64(dimensions.width) - size.TotalWidth) / 2
	}
	switch computed.AlignItems {
	case style.AlignItemsEnd:
		lastY += float64(dimensions.width) - size.TotalHeight
	case style.AlignItemsCenter:
		lastY += (float64(dimensions.width) - size.TotalHeight) / 2
	}

	// Render text
	face, close := computed.Font.Face()
	defer close()

	ctx.SetFontFace(face)
	ctx.SetColor(computed.Color)

	for _, str := range strings.Split(content.value, "\n") {
		lastY += size.LineHeight
		x, y := lastX, lastY-size.LineOffset
		ctx.DrawString(str, x, y)
	}

	return nil
}
