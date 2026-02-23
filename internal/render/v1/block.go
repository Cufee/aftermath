package render

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/pkg/errors"

	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
)

type blockContentType int

func (t blockContentType) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "%d", t), nil
}
func (t blockContentType) String() string {
	return fmt.Sprintf("%d", t)
}

const (
	BlockContentTypeText blockContentType = iota
	BlockContentTypeImage
	BlockContentTypeBlocks
	BlockContentTypeEmpty
)

type measured struct {
	height float64
	width  float64
}

type BlockContent interface {
	Render(Style) (image.Image, error)
	Type() blockContentType

	measure(Style) (*measured, error)
}

type Block struct {
	ContentType blockContentType
	content     BlockContent
	Style       Style
}

func (block *Block) Render() (image.Image, error) {
	return block.content.Render(block.Style)
}

func NewBlock(content BlockContent, style Style) Block {
	return Block{
		ContentType: content.Type(),
		content:     content,
		Style:       style,
	}
}

type contentText struct {
	dimensions *measured
	value      string
}

func NewTextContent(style Style, value string) Block {
	return NewBlock(&contentText{
		value: value,
	}, style)
}

func (content *contentText) measure(style Style) (*measured, error) {
	if content.dimensions != nil {
		return content.dimensions, nil
	}
	if !style.Font.Valid() {
		return nil, errors.New("font not valid")
	}

	dimensions := &measured{width: style.Width, height: style.Height}
	if dimensions.height > 0 && dimensions.width > 0 {
		content.dimensions = dimensions
		return dimensions, nil
	}

	size := MeasureString(content.value, style.Font)
	if dimensions.height == 0 {
		dimensions.height = size.TotalHeight
	}
	if dimensions.width == 0 {
		dimensions.width = size.TotalWidth
	}

	content.dimensions = dimensions
	return dimensions, nil
}

func (content *contentText) Render(style Style) (image.Image, error) {
	if !style.Font.Valid() {
		return nil, errors.New("font not valid")
	}

	size := MeasureString(content.value, style.Font)

	width := size.TotalWidth + (style.PaddingX * 2)
	height := size.TotalHeight + (style.PaddingY * 2)
	if style.Width > 0 {
		width = style.Width
	}
	if style.Height > 0 {
		height = style.Height
	}

	ctx := gg.NewContext(int(width), int(height))

	// Render text
	face, close := style.Font.Face()
	defer close()

	ctx.SetFontFace(face)
	ctx.SetColor(style.FontColor)

	var lastX, lastY float64 = style.PaddingX, style.PaddingY + 1

	switch style.JustifyContent {
	case JustifyContentEnd:
		lastX += width - size.TotalWidth
	case JustifyContentCenter:
		lastX += (width - size.TotalWidth) / 2
	}
	switch style.AlignItems {
	case AlignItemsEnd:
		lastY += height - size.TotalHeight
	case AlignItemsCenter:
		lastY += (height - size.TotalHeight) / 2
	}

	for str := range strings.SplitSeq(content.value, "\n") {
		lastY += size.LineHeight
		ctx.DrawString(str, lastX, lastY-size.LineOffset)
	}

	if style.Debug {
		ctx.SetColor(getDebugColor())
		ctx.DrawRectangle(0, 0, float64(ctx.Width()), float64(ctx.Height()))
		ctx.Stroke()
	}

	return ctx.Image(), nil
}

func (content contentText) Type() blockContentType {
	return BlockContentTypeText
}

type contentBlocks struct {
	dimensions *measured
	blocks     []Block
}

func NewBlocksContent(style Style, blocks ...Block) Block {
	return NewBlock(&contentBlocks{
		blocks: blocks,
	}, style)
}

func (content *contentBlocks) measure(style Style) (*measured, error) {
	if content.dimensions != nil {
		return content.dimensions, nil
	}

	dimensions := &measured{width: style.Width, height: style.Height}
	if (len(content.blocks) < 1) || (dimensions.height > 0 && dimensions.width > 0) {
		content.dimensions = dimensions
		return dimensions, nil
	}

	var measuredWidth, measuredHeight float64

	// padding / gaps
	if style.Direction == DirectionHorizontal {
		measuredWidth += (style.Gap * float64(len(content.blocks)-1)) + (style.PaddingX * 2)
		measuredHeight += style.PaddingY * 2
	} else {
		measuredWidth += style.PaddingX * 2
		measuredHeight += (style.Gap * float64(len(content.blocks)-1)) + (style.PaddingY * 2)
	}

	// pure content size
	for _, block := range content.blocks {
		size, err := block.content.measure(block.Style)
		if err != nil {
			return nil, err
		}
		if style.Direction == DirectionHorizontal {
			measuredHeight = max(measuredHeight, size.height)
			measuredWidth += size.width
		} else {
			measuredHeight += size.height
			measuredWidth = max(measuredWidth, size.width)
		}
	}

	// make sure we respect explicit values
	if dimensions.height == 0 {
		dimensions.height = measuredHeight
	}
	if dimensions.width == 0 {
		dimensions.width = measuredWidth
	}

	content.dimensions = dimensions
	return dimensions, nil
}

func (content *contentBlocks) Render(style Style) (image.Image, error) {
	if len(content.blocks) < 1 {
		return nil, errors.New("block content cannot be empty")
	}
	containerSize, err := content.measure(style)
	if err != nil {
		return nil, err
	}

	var blocksWithGrowX, blocksWithGrowY float64
	for _, block := range content.blocks {
		if block.Style.GrowX {
			blocksWithGrowX++
		}
		if block.Style.GrowY {
			blocksWithGrowY++
		}
	}

	var perBlockGrowSpaceX, perBlockGrowSpaceY float64
	if style.Width > containerSize.width {
		perBlockGrowSpaceX = (style.Width - containerSize.width) / blocksWithGrowX
	}
	if style.Height > containerSize.height {
		perBlockGrowSpaceY = (style.Height - containerSize.height) / blocksWithGrowY
	}

	var images []image.Image
	for _, block := range content.blocks {
		size, err := block.content.measure(block.Style)
		if err != nil {
			return nil, errors.Wrap(err, "failed to measure block")
		}

		if style.Direction == DirectionHorizontal {
			if block.Style.GrowY {
				block.Style.Height = containerSize.height
			}
			if block.Style.GrowX {
				block.Style.Width = size.width + perBlockGrowSpaceX
			}
		} else {
			if block.Style.GrowX {
				block.Style.Width = containerSize.width
			}
			if block.Style.GrowY {
				block.Style.Height = size.height + perBlockGrowSpaceY
			}
		}

		img, err := block.Render()
		if err != nil {
			return nil, err
		}
		if img == nil {
			return nil, errors.New("image is nil for content type " + content.Type().String())
		}
		images = append(images, img)
	}
	return renderImages(images, style)
}

// func (content *contentBlocks) renderAsync(style Style) (image.Image, error) {
// 	images := make([]image.Image, len(content.blocks))
// 	var g errgroup.Group

// 	for i, block := range content.blocks {
// 		i, block := i, block // Create new variables for closure
// 		g.Go(func() error {
// 			img, err := block.Render()
// 			if err != nil {
// 				return err
// 			}
// 			if img == nil {
// 				return fmt.Errorf("image is nil for content type %s", content.Type())
// 			}
// 			images[i] = img
// 			return nil
// 		})
// 	}

// 	if err := g.Wait(); err != nil {
// 		return nil, err
// 	}

// 	return renderImages(images, style)
// }

func (content contentBlocks) Type() blockContentType {
	return BlockContentTypeBlocks
}

type contentImage struct {
	dimensions *measured
	image      image.Image
}

func NewImageContent(style Style, image image.Image) Block {
	return NewBlock(&contentImage{
		image: image,
	}, style)
}

func (content *contentImage) measure(style Style) (*measured, error) {
	if content.dimensions != nil {
		return content.dimensions, nil
	}

	dimensions := &measured{width: style.Width, height: style.Height}
	if dimensions.height > 0 && dimensions.width > 0 {
		content.dimensions = dimensions
		return dimensions, nil
	}

	if dimensions.height == 0 {
		dimensions.height = float64(content.image.Bounds().Dy())
	}
	if dimensions.width == 0 {
		dimensions.width = float64(content.image.Bounds().Dx())
	}

	content.dimensions = dimensions
	return dimensions, nil
}

func (content *contentImage) Render(style Style) (image.Image, error) {
	if content.image == nil {
		return nil, errors.New("image content cannot be nil")
	}

	if style.Width == 0 {
		style.Width = float64(content.image.Bounds().Dx())
	}
	if style.Height == 0 {
		style.Height = float64(content.image.Bounds().Dy())
	}

	// Type cast to image.Image for gg
	newHeight := int(style.Height)
	newWidth := int(style.Width)
	if newHeight > newWidth {
		newWidth = 0
	} else {
		newHeight = 0
	}

	var image image.Image = imaging.Resize(content.image, newWidth, newHeight, imaging.CatmullRom)
	if style.BackgroundColor != nil {
		mask := gg.NewContextForImage(image)
		ctx := gg.NewContext(image.Bounds().Dx(), image.Bounds().Dy())
		ctx.SetMask(mask.AsMask())
		ctx.SetColor(style.BackgroundColor)
		ctx.DrawRectangle(0, 0, float64(ctx.Width()), float64(ctx.Height()))
		ctx.Fill()
		image = ctx.Image()
	}

	return image, nil
}

func NewEmptyContent(width, height float64) Block {
	return NewBlock(&contentEmpty{dimensions: &measured{width: width, height: height}}, Style{Width: width, Height: height})
}

func (content *contentImage) Type() blockContentType {
	return BlockContentTypeBlocks
}

type contentEmpty struct {
	dimensions *measured
}

func (content contentEmpty) Type() blockContentType {
	return BlockContentTypeEmpty
}
func (content *contentEmpty) measure(style Style) (*measured, error) {
	if content.dimensions != nil {
		return content.dimensions, nil
	}

	dimensions := &measured{width: style.Width, height: style.Height}
	content.dimensions = dimensions
	return dimensions, nil
}

func (content *contentEmpty) Render(style Style) (image.Image, error) {
	return imaging.New(int(style.Width), int(style.Height), color.Transparent), nil
}
