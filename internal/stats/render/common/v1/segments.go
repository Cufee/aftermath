package common

import (
	"image"

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

	var frameParts []Block

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
		return nil, err
	}
	mainSegmentImg = AddBackground(mainSegmentImg, options.Background, Style{Blur: 10, BorderRadius: 42.5})

	frameParts = append(frameParts, NewImageContent(Style{}, mainSegmentImg))
	if len(s.header) > 0 {
		frameParts = append(frameParts, NewBlocksContent(
			Style{
				Direction:  DirectionVertical,
				AlignItems: AlignItemsCenter,
				Gap:        10,
			}, s.header...))
	}
	if len(s.footer) > 0 {
		frameParts = append(frameParts, NewBlocksContent(
			Style{
				Direction:  DirectionVertical,
				AlignItems: AlignItemsCenter,
				Gap:        10,
			}, s.footer...))
	}

	frame := NewBlocksContent(
		Style{
			Direction:  DirectionVertical,
			AlignItems: AlignItemsCenter,
			Gap:        5,
		},
		frameParts...,
	)
	return frame.Render()
}