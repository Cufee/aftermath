package common

import (
	"image"
	"sync"

	"github.com/cufee/aftermath/internal/retry"
	"github.com/pkg/errors"
)

type Segments struct {
	header  []Block
	content []Block
	footer  []Block
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

func (s *Segments) Render(opts ...Option) (image.Image, error) {
	if len(s.content) < 1 {
		return nil, errors.New("segments.content cannot be empty")
	}

	options := DefaultOptions()
	for _, apply := range opts {
		apply(&options)
	}

	var wg sync.WaitGroup
	var headerBlock *Block
	var contentBlock retry.DataWithErr[Block]
	var footerBlock *Block

	wg.Add(3)
	go func() {
		defer wg.Done()
		mainSegment := NewBlocksContent(
			Style{
				Direction:  DirectionVertical,
				AlignItems: AlignItemsCenter,
				PaddingX:   20,
				PaddingY:   20,
				Gap:        10,
			}, s.content...)
		mainSegmentImg, err := mainSegment.Render()
		if err != nil {
			contentBlock = retry.DataWithErr[Block]{Err: err}
			return
		}
		mainSegmentBlock := NewImageContent(Style{}, AddBackground(mainSegmentImg, options.Background, Style{Blur: 10, BorderRadius: 42.5}))
		contentBlock = retry.DataWithErr[Block]{Data: mainSegmentBlock}
	}()
	go func() {
		defer wg.Done()
		if len(s.header) > 0 {
			header := NewBlocksContent(
				Style{
					Direction:  DirectionVertical,
					AlignItems: AlignItemsCenter,
					Gap:        10,
				}, s.header...)
			headerBlock = &header
		}
	}()
	go func() {
		defer wg.Done()
		if len(s.footer) > 0 {
			footer := NewBlocksContent(
				Style{
					Direction:  DirectionVertical,
					AlignItems: AlignItemsCenter,
					Gap:        10,
				}, s.footer...)
			footerBlock = &footer
		}
	}()
	wg.Wait()

	var frameBlocks []Block
	if headerBlock != nil {
		frameBlocks = append(frameBlocks, *headerBlock)
	}
	if contentBlock.Err != nil {
		return nil, contentBlock.Err
	}
	frameBlocks = append(frameBlocks, contentBlock.Data)
	if footerBlock != nil {
		frameBlocks = append(frameBlocks, *footerBlock)
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
