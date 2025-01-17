package render

import (
	"fmt"
	"image"

	"github.com/cufee/aftermath/internal/render/v2/style"
	"github.com/fogleman/gg"
)

func NewBlock(content BlockContent) *Block {
	return &Block{
		content: content,
	}
}

type blockContentType int

func (t blockContentType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", t)), nil
}
func (t blockContentType) String() string {
	return fmt.Sprintf("%d", t)
}

const (
	BlockContentTypeEmpty blockContentType = iota
	BlockContentTypeBlocks
	BlockContentTypeImage
	BlockContentTypeText
)

type Position struct {
	X float64
	Y float64
}

type BlockContent interface {
	Type() blockContentType

	// Renders the block onto an image
	Render(*gg.Context, Position) error

	Style() style.StyleOptions
	setStyle(style.StyleOptions)

	// returns final block image dimensions without rendering
	dimensions() contentDimensions
}

type Block struct {
	content BlockContent
}

func (b *Block) Style() style.StyleOptions {
	return b.content.Style()
}

func (b *Block) Type() blockContentType {
	return b.content.Type()
}

func (b *Block) Render() (image.Image, error) {
	dimensions := b.content.dimensions()
	ctx := gg.NewContext(dimensions.width, dimensions.height)
	err := b.RenderTo(ctx, Position{0, 0})
	if err != nil {
		return nil, err
	}
	return ctx.Image(), nil

}

func (b *Block) RenderTo(ctx *gg.Context, pos Position) error {
	return b.content.Render(ctx, pos)
}

func (b *Block) Dimensions() contentDimensions {
	return b.content.dimensions()
}

type contentDimensions struct {
	width           int
	height          int
	paddingAndGapsX float64
	paddingAndGapsY float64
}
