package common

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"strings"
	"sync"

	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

type blockContentType int

func (t blockContentType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", t)), nil
}
func (t blockContentType) String() string {
	return fmt.Sprintf("%d", t)
}

const (
	BlockContentTypeText blockContentType = iota
	BlockContentTypeImage
	// BlockContentTypeIcon
	BlockContentTypeBlocks
	BlockContentTypeEmpty
)

type BlockContent interface {
	Render(Style) (image.Image, error)
	Type() blockContentType
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
	value string
}

func NewTextContent(style Style, value string) Block {
	return NewBlock(contentText{
		value: value,
	}, style)
}

func (content contentText) Render(style Style) (image.Image, error) {
	if !style.Font.Valid() {
		return nil, errors.New("font not valid")
	}

	size := MeasureString(content.value, style.Font)
	ctx := gg.NewContext(int(size.TotalWidth+(style.PaddingX*2)), int(size.TotalHeight+(style.PaddingY*2)))

	// Render text
	ctx.SetFontFace(style.Font.Face())
	ctx.SetColor(style.FontColor)

	var lastX, lastY float64 = style.PaddingX, style.PaddingY + 1
	for _, str := range strings.Split(content.value, "\n") {
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
	blocks []Block
}

func NewBlocksContent(style Style, blocks ...Block) Block {
	return NewBlock(contentBlocks{
		blocks: blocks,
	}, style)
}

func (content contentBlocks) Render(style Style) (image.Image, error) {
	if len(content.blocks) < 1 {
		return nil, errors.New("block content cannot be empty")
	}

	// // avoid the overhead of mutex and goroutines if it is not required
	// if len(content.blocks) > 1 {
	// 	return content.renderAsync(style)
	// }

	var images []image.Image
	for _, block := range content.blocks {
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

func (content contentBlocks) renderAsync(style Style) (image.Image, error) {
	var mx sync.Mutex
	var wg sync.WaitGroup
	images := make([]image.Image, len(content.blocks))

	errCh := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	for i, block := range content.blocks {
		wg.Add(1)
		go func(i int, b Block) {
			defer wg.Done()

			img, err := b.Render()
			if err == nil && img == nil {
				err = errors.New("image is nil for content type " + content.Type().String())
			}
			if err != nil {
				select {
				case errCh <- err:
					cancel()
				default:
					log.Err(err).Any("type", b.content.Type()).Msg("unhandled error while rendering a block")
				}
				return
			}
			mx.Lock()
			images[i] = img
			mx.Unlock()
		}(i, block)
	}
	go func() {
		wg.Wait()
		cancel()
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		return renderImages(images, style)
	case err := <-errCh:
		return nil, err
	}
}

func (content contentBlocks) Type() blockContentType {
	return BlockContentTypeBlocks
}

type contentImage struct {
	image image.Image
}

func NewImageContent(style Style, image image.Image) Block {
	return NewBlock(contentImage{
		image: image,
	}, style)
}

func (content contentImage) Render(style Style) (image.Image, error) {
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
	var image image.Image = imaging.Fit(content.image, int(style.Width), int(style.Height), imaging.Linear)
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
	return NewBlock(contentEmpty{}, Style{Width: width, Height: height})
}

func (content contentImage) Type() blockContentType {
	return BlockContentTypeBlocks
}

type contentEmpty struct{}

func (content contentEmpty) Type() blockContentType {
	return BlockContentTypeEmpty
}

func (content contentEmpty) Render(style Style) (image.Image, error) {
	return imaging.New(int(style.Width), int(style.Height), color.Transparent), nil
}
