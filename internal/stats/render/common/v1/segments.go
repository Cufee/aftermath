package common

import (
	"image"
	"image/color"
	"sync"

	"github.com/cufee/aftermath/internal/retry"
	"github.com/pkg/errors"
)

type Segments struct {
	header  []Block
	content []Block
	footer  []Block

	rendered struct {
		header  image.Image
		content image.Image
		footer  image.Image
	}
}

func (s *Segments) AddHeader(blocks ...Block) {
	s.header = append(s.header, blocks...)
}

func (s *Segments) AddContent(blocks ...Block) {
	s.content = append(s.content, blocks...)
}

func (s *Segments) AddFooter(blocks ...Block) {
	s.footer = append(s.footer, blocks...)
}

func (s *Segments) ContentMask(opts ...Option) (*image.Alpha, error) {
	options := DefaultOptions()
	for _, apply := range opts {
		apply(&options)
	}

	if s.rendered.content == nil {
		content, err := s.renderContent(options)
		if err != nil {
			return nil, err
		}
		s.rendered.content = content
	}

	bounds := s.rendered.content.Bounds()
	mask := image.NewAlpha(bounds)

	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			if _, _, _, a := s.rendered.content.At(x, y).RGBA(); a > 0 {
				// Check neighboring pixels for alpha
				if _, _, _, a := s.rendered.content.At(x-1, y).RGBA(); a == 0 { // left
					mask.SetAlpha(x, y, color.Alpha{A: 0})
					continue
				}
				if _, _, _, a := s.rendered.content.At(x+1, y).RGBA(); a == 0 { // right
					mask.SetAlpha(x, y, color.Alpha{A: 0})
					continue
				}
				if _, _, _, a := s.rendered.content.At(x, y+1).RGBA(); a == 0 { // top
					mask.SetAlpha(x, y, color.Alpha{A: 0})
					continue
				}
				if _, _, _, a := s.rendered.content.At(x, y-1).RGBA(); a == 0 { // bottom
					mask.SetAlpha(x, y, color.Alpha{A: 0})
					continue
				}
				mask.SetAlpha(x, y, color.Alpha{A: 255})
			}
		}
	}

	return mask, nil
}

func (s *Segments) renderHeader(_ Options) (image.Image, error) {
	header := NewBlocksContent(
		Style{
			Direction:  DirectionVertical,
			AlignItems: AlignItemsCenter,
			Gap:        10,
		}, s.header...)
	return header.Render()
}

func (s *Segments) renderFooter(_ Options) (image.Image, error) {
	footer := NewBlocksContent(
		Style{
			Direction:  DirectionVertical,
			AlignItems: AlignItemsCenter,
			Gap:        10,
		}, s.footer...)
	return footer.Render()
}

func (s *Segments) renderContent(_ Options) (image.Image, error) {
	mainSegment := NewBlocksContent(
		Style{
			Direction:  DirectionVertical,
			AlignItems: AlignItemsCenter,
			PaddingX:   20,
			PaddingY:   20,
			Gap:        10,
		}, s.content...)
	return mainSegment.Render()
}

func (s *Segments) Render(opts ...Option) (image.Image, error) {
	if len(s.content) < 1 {
		return nil, errors.New("segments.content cannot be empty")
	}

	options := DefaultOptions()
	for _, apply := range opts {
		apply(&options)
	}

	var wg sync.WaitGroup
	var headerBlock retry.DataWithErr[*Block]
	var footerBlock retry.DataWithErr[*Block]
	var contentBlock retry.DataWithErr[Block]

	wg.Add(3)
	go func(block *retry.DataWithErr[Block]) {
		defer wg.Done()

		if s.rendered.content == nil {
			content, err := s.renderContent(options)
			if err != nil {
				*block = retry.DataWithErr[Block]{Err: err}
				return
			}
			s.rendered.content = content
		}

		bgMask, err := s.ContentMask(opts...)
		if err != nil {
			block.Err = err
			return
		}

		block.Data = NewImageContent(Style{}, AddGlassMaskedBackground(s.rendered.content, options.Background, bgMask, Style{BorderRadius: 42.5}))
	}(&contentBlock)

	go func(block *retry.DataWithErr[*Block]) {
		defer wg.Done()
		if len(s.header) > 0 {
			if s.rendered.header == nil {
				header, err := s.renderHeader(options)
				if err != nil {
					block.Err = err
					return
				}
				s.rendered.header = header
			}

			b := NewImageContent(Style{AlignItems: AlignItemsCenter, JustifyContent: JustifyContentCenter}, s.rendered.header)
			block.Data = &b
		}
	}(&headerBlock)
	go func(block *retry.DataWithErr[*Block]) {
		defer wg.Done()
		if len(s.footer) > 0 {
			if s.rendered.footer == nil {
				footer, err := s.renderFooter(options)
				if err != nil {
					block.Err = err
					return
				}
				s.rendered.footer = footer
			}

			b := NewImageContent(Style{AlignItems: AlignItemsCenter, JustifyContent: JustifyContentCenter}, s.rendered.footer)
			block.Data = &b
		}
	}(&footerBlock)
	wg.Wait()

	var frameBlocks []Block
	if headerBlock.Err != nil {
		return nil, headerBlock.Err
	}
	if headerBlock.Data != nil {
		frameBlocks = append(frameBlocks, *headerBlock.Data)
	}

	if contentBlock.Err != nil {
		return nil, contentBlock.Err
	}
	frameBlocks = append(frameBlocks, contentBlock.Data)

	if footerBlock.Err != nil {
		return nil, footerBlock.Err
	}
	if footerBlock.Data != nil {
		frameBlocks = append(frameBlocks, *footerBlock.Data)
	}

	frame := NewBlocksContent(
		Style{
			Direction:  DirectionVertical,
			AlignItems: AlignItemsCenter,
			Gap:        5,
		},
		frameBlocks...,
	)
	return frame.Render()
}
