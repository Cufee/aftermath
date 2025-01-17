package render

import (
	"errors"
	"fmt"

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

			gaps := computed.Gap * float64(len(content.value)-1)
			dimensions.width += ceil(gaps)
			dimensions.paddingAndGapsX += gaps

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

			gaps := computed.Gap * float64(len(content.value)-1)
			dimensions.height += ceil(gaps)
			dimensions.paddingAndGapsY += gaps

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

	var originX, originY float64 = pos.X + computed.PaddingLeft, pos.Y + computed.PaddingTop
	if computed.Position == style.PositionAbsolute {
		originX += computed.MarginLeft
		originY += computed.MarginTop
	}

	if computed.BackgroundColor != nil {
		ctx.SetColor(computed.BackgroundColor)
		ctx.DrawRectangle(originX, originY, float64(dimensions.width), float64(dimensions.height))
		ctx.Fill()
	}
	if computed.BackgroundImage != nil {
		background := imaging.Fill(computed.BackgroundImage, dimensions.width, dimensions.height, imaging.Center, imaging.Lanczos)
		ctx.DrawImage(background, ceil(originX), ceil(originY))
	}

	if computed.Debug {
		ctx.SetColor(getDebugColor())
		ctx.DrawRectangle(pos.X, pos.Y, float64(dimensions.width), float64(dimensions.height))
		ctx.Stroke()
	}

	applyGrowth(computed, dimensions, content.value...)
	return renderBlocksContent(ctx, computed, dimensions, pos, content.value...)
}

func renderBlocksContent(ctx *gg.Context, containerStyle style.Style, container contentDimensions, pos Position, blocks ...*Block) error {
	if len(blocks) < 1 {
		return errors.New("no blocks to render")
	}

	var originX, originY = pos.X + containerStyle.PaddingLeft, pos.Y + containerStyle.PaddingTop

	var lastX, lastY float64 = originX, originY
	var justifyOffsetX, justifyOffsetY float64

	var freeSpaceX, freeSpaceY = float64(container.width) - container.paddingAndGapsX, float64(container.height) - container.paddingAndGapsY

	// Set correct gaps and offsets based on justify content
	switch containerStyle.JustifyContent {
	case style.JustifyContentCenter:
		lastX += freeSpaceX / 2
		lastY += freeSpaceY / 2
	case style.JustifyContentEnd:
		lastX += freeSpaceX
		lastY += freeSpaceY
	case style.JustifyContentSpaceBetween:
		if len(blocks) > 0 {
			justifyOffsetX = float64(freeSpaceX / float64(len(blocks)-1))
			justifyOffsetY = float64(freeSpaceY / float64(len(blocks)-1))
		}
	case style.JustifyContentSpaceAround:
		spacingX := float64(freeSpaceX / float64(len(blocks)+1))
		spacingY := float64(freeSpaceY / float64(len(blocks)+1))
		justifyOffsetX = spacingX
		justifyOffsetY = spacingY
		lastX += spacingX
		lastY += spacingY
	default: // JustifyContentStart
	}

	for i, block := range blocks {
		blockSize := block.content.dimensions()
		posX, posY := lastX, lastY

		switch containerStyle.Direction {
		case style.DirectionVertical:
			if i > 0 {
				posY += justifyOffsetY + containerStyle.Gap
			}
			lastY = posY + float64(blockSize.height)

			switch containerStyle.AlignItems {
			case style.AlignItemsCenter:
				posX = float64(container.width-blockSize.width) / 2
			case style.AlignItemsEnd:
				posX = float64(container.width-blockSize.width) - containerStyle.PaddingRight
			default: // AlignItemsStart
				posX = containerStyle.PaddingLeft
			}
		default: // DirectionHorizontal
			if i > 0 {
				posX += justifyOffsetX + containerStyle.Gap
			}
			lastX = posX + float64(blockSize.width)

			switch containerStyle.AlignItems {
			case style.AlignItemsCenter:
				posY = float64(container.height-blockSize.height) / 2
			case style.AlignItemsEnd:
				posY = float64(container.height-blockSize.height) - containerStyle.PaddingBottom
			default: // AlignItemsStart
				posY = containerStyle.PaddingTop
			}

		}

		err := block.content.Render(ctx, Position{posX, posY})
		if err != nil {
			return err
		}
	}

	return nil
}

func applyGrowth(containerStyle style.Style, container contentDimensions, blocks ...*Block) {
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

	// calculate empty space blocks can use to grow
	var growSpaceX, growSpaceY = 0, 0
	if growBlocksX > 0 {
		switch containerStyle.Direction {
		case style.DirectionHorizontal:
			growSpaceX = container.width - ceil(container.paddingAndGapsX) - blockWidthTotal
		case style.DirectionVertical:
			growSpaceX = container.width - ceil(container.paddingAndGapsX) - blockWidthMax
		}
	}
	if growBlocksY > 0 {
		switch containerStyle.Direction {
		case style.DirectionHorizontal:
			growSpaceY = container.height - ceil(container.paddingAndGapsY) - blockHeightMax
		case style.DirectionVertical:
			growSpaceY = container.height - ceil(container.paddingAndGapsY) - blockWidthTotal
		}
	}

	fmt.Printf("grow x %v container %v %v blocks %v \n", growSpaceX, container.width, container.paddingAndGapsX, blockWidthTotal)

	// apply growth to blocks
	if growBlocksX > 0 || growBlocksY > 0 {
		var blockGrowX, blockGrowY = max(0, growSpaceX) / max(1, growBlocksX), max(0, growSpaceY) / max(1, growBlocksY)
		for _, block := range blocks {
			blockStyle := block.Style()
			blockComputed := blockStyle.Computed()

			// update the block width
			if blockComputed.GrowHorizontal {
				blockStyle.Add(style.SetWidth(blockComputed.Width + float64(blockGrowX)))
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
