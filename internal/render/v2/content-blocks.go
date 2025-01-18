package render

import (
	"errors"

	"github.com/cufee/aftermath/internal/render/v2/style"
	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
)

var _ BlockContent = &contentBlocks{}

func NewBlocksContent(style style.StyleOptions, value ...*Block) *Block {
	return NewBlock(&contentBlocks{
		value: value,
		style: style,
	})
}

type contentBlocks struct {
	style style.StyleOptions
	value []*Block
}

func (content *contentBlocks) setStyle(style style.StyleOptions) {
	content.style = style
}

func (content *contentBlocks) dimensions() contentDimensions {
	if len(content.value) == 0 {
		return contentDimensions{}
	}

	computed := content.style.Computed()
	dimensions := contentDimensions{
		width:           ceil(computed.Width),
		height:          ceil(computed.Height),
		paddingAndGapsX: computed.PaddingLeft + computed.PaddingRight,
		paddingAndGapsY: computed.PaddingTop + computed.PaddingBottom,
	}

	switch computed.Direction {
	case style.DirectionHorizontal:
		dimensions.paddingAndGapsX += computed.Gap * float64(len(content.value)-1)
	case style.DirectionVertical:
		dimensions.paddingAndGapsY += computed.Gap * float64(len(content.value)-1)
	}

	if dimensions.width > 0 && dimensions.height > 0 {
		return dimensions
	}

	// add content dimensions of each block to the total
	var blockWidthTotal, blockWidthMax, blockHeightTotal, blockHeightMax int
	for _, block := range content.value {
		blockDimensions := block.content.dimensions()

		blockWidthTotal += blockDimensions.width
		blockWidthMax = max(blockWidthMax, blockDimensions.width)

		blockHeightTotal += blockDimensions.height
		blockHeightMax = max(blockHeightMax, blockDimensions.height)
	}

	// calculate final block width if it was not set already
	if dimensions.width == 0 {
		dimensions.width = ceil(computed.PaddingLeft) + ceil(computed.PaddingRight)

		switch computed.Direction {
		case style.DirectionHorizontal:
			dimensions.width += blockWidthTotal

		case style.DirectionVertical:
			dimensions.width += blockWidthMax
		}
	}
	// calculate final block height if it was not set already
	if dimensions.height == 0 {
		dimensions.height = ceil(computed.PaddingTop + computed.PaddingBottom)

		switch computed.Direction {
		case style.DirectionHorizontal:
			dimensions.height += blockHeightMax
		case style.DirectionVertical:
			dimensions.height += blockHeightTotal
		}
	}

	return dimensions
}

func (content *contentBlocks) Type() blockContentType {
	return BlockContentTypeBlocks
}

func (content *contentBlocks) Style() style.StyleOptions {
	return content.style
}

func (content *contentBlocks) Render(ctx *gg.Context, pos Position) error {
	computed := content.style.Computed()
	dimensions := content.dimensions()

	if computed.Blur > 0 {
		blur := computed.Blur
		computed.Blur = 0
		// render the content onto a new image, blur it, render onto parent
		child := gg.NewContext(dimensions.width, dimensions.height)
		err := content.Render(child, Position{0, 0})
		if err != nil {
			return err
		}
		img := imaging.Blur(ctx.Image(), blur)
		ctx.DrawImage(img, ceil(pos.X), ceil(pos.Y))
		return nil
	}

	// if computed.Position == style.PositionAbsolute {
	// pos.X += computed.MarginLeft
	// pos.Y += computed.MarginTop
	// }

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

	applyBlocksGrowth(computed, dimensions, content.value...)

	var originX, originY = pos.X + computed.PaddingLeft, pos.Y + computed.PaddingTop
	return renderBlocksContent(ctx, computed, dimensions, Position{X: originX, Y: originY}, content.value...)
}

func renderBlocksContent(ctx *gg.Context, containerStyle style.Style, container contentDimensions, pos Position, blocks ...*Block) error {
	if len(blocks) < 1 {
		return errors.New("no blocks to render")
	}

	var lastX, lastY float64 = pos.X, pos.Y
	for i, block := range blocks {
		blockSize := block.content.dimensions()
		posX, posY := lastX, lastY

		switch containerStyle.Direction {
		case style.DirectionVertical:
			if i > 0 {
				posY += containerStyle.Gap
			}

			// align content vertically
			switch containerStyle.JustifyContent {
			case style.JustifyContentCenter:
				posY += float64(container.height-blockSize.height) / 2
			case style.JustifyContentEnd:
				posY += float64(container.height - blockSize.height)
			case style.JustifyContentSpaceAround:
				posY += float64((container.height - blockSize.height) / (len(blocks) + 1))
			case style.JustifyContentSpaceBetween:
				if len(blocks) > 1 {
					posY += float64((container.height - blockSize.height) / (len(blocks) - 1))
				}
			}

			// align content horizontally
			posX = pos.X
			switch containerStyle.AlignItems {
			case style.AlignItemsCenter:
				posX += float64(container.width-blockSize.width) / 2
			case style.AlignItemsEnd:
				posX += float64(blockSize.width)
			}
		default: // DirectionHorizontal
			if i > 0 {
				posX += containerStyle.Gap
			}

			// align content horizontally
			switch containerStyle.JustifyContent {
			case style.JustifyContentCenter:
				posX += float64(container.width-blockSize.width) / 2
			case style.JustifyContentEnd:
				posX += float64(container.width - blockSize.width)
			case style.JustifyContentSpaceAround:
				posX += float64((container.width - blockSize.width) / (len(blocks) + 1))
			case style.JustifyContentSpaceBetween:
				if len(blocks) > 1 {
					posX += float64((container.width - blockSize.width) / (len(blocks) - 1))
				}
			}

			// align content vertically
			posY = pos.Y
			switch containerStyle.AlignItems {
			case style.AlignItemsCenter:
				posY += (float64(container.height-blockSize.height) / 2)
			case style.AlignItemsEnd:
				posY += float64(blockSize.height)
			}

		}

		err := block.content.Render(ctx, Position{posX, posY})
		if err != nil {
			return err
		}

		// save the position we rendered at
		switch containerStyle.Direction {
		case style.DirectionVertical:
			lastY = posY + float64(blockSize.height)
		default:
			lastX = posX + float64(blockSize.width)
		}
	}

	return nil
}

func applyBlocksGrowth(containerStyle style.Style, container contentDimensions, blocks ...*Block) {
	// calculate content dimensions before growth
	var blockWidthTotal, blockWidthMax, blockHeightTotal, blockHeightMax int
	var growBlocksX, growBlocksY = 0, 0
	for _, block := range blocks {
		blockDimensions := block.content.dimensions()

		blockWidthTotal += blockDimensions.width
		blockWidthMax = max(blockWidthMax, blockDimensions.width)

		blockHeightTotal += blockDimensions.height
		blockHeightMax = max(blockHeightMax, blockDimensions.height)

		style := block.Style().Computed()
		if style.GrowHorizontal {
			growBlocksX++
		}
		if style.GrowVertical {
			growBlocksY++
		}
	}

	// apply growth to blocks
	if growBlocksX > 0 || growBlocksY > 0 {
		blockGrowX := max(0, container.width-ceil(container.paddingAndGapsX)-blockWidthTotal) / max(1, growBlocksX)
		blockGrowY := max(0, container.height-ceil(container.paddingAndGapsY)-blockHeightTotal) / max(1, growBlocksY)

		for _, block := range blocks {
			blockStyle := block.Style()
			blockComputed := blockStyle.Computed()

			if !blockComputed.GrowHorizontal && !blockComputed.GrowVertical {
				continue
			}

			switch containerStyle.Direction {
			case style.DirectionHorizontal:
				// update the block width
				if blockComputed.GrowHorizontal {
					blockStyle.Add(style.SetWidth(blockComputed.Width + float64(blockGrowX)))
					block.content.setStyle(blockStyle)
				}
				// update the block height
				if blockComputed.GrowVertical {
					blockStyle.Add(style.SetHeight(float64(blockHeightMax)))
					block.content.setStyle(blockStyle)
				}
			case style.DirectionVertical:
				// update the block width
				if blockComputed.GrowHorizontal {
					blockStyle.Add(style.SetWidth(float64(blockWidthMax)))
					block.content.setStyle(blockStyle)
				}
				// update the block height
				if blockComputed.GrowVertical {
					blockStyle.Add(style.SetHeight(blockComputed.Height + float64(blockGrowY)))
					block.content.setStyle(blockStyle)
				}
			}

		}
	}
}
