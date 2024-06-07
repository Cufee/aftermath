package render

import (
	"errors"
	"image"

	"github.com/cufee/aftermath/internal/stats/render/common"
)

type Segments struct {
	header  []common.Block
	content []common.Block
	footer  []common.Block
}

func (s *Segments) AddHeader(blocks ...common.Block) {
	s.header = append(s.header, blocks...)
}

func (s *Segments) AddContent(blocks ...common.Block) {
	s.content = append(s.content, blocks...)
}

func (s *Segments) AddFooter(blocks ...common.Block) {
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

	var frameParts []common.Block

	mainSegment := common.NewBlocksContent(
		common.Style{
			Direction:  common.DirectionVertical,
			AlignItems: common.AlignItemsCenter,
			PaddingX:   20,
			PaddingY:   20,
			Gap:        10,
		}, s.content...)
	mainSegmentImg, err := mainSegment.Render()
	if err != nil {
		return nil, err
	}
	mainSegmentImg = common.AddBackground(mainSegmentImg, options.Background, common.Style{Blur: 10, BorderRadius: 42.5})

	frameParts = append(frameParts, common.NewImageContent(common.Style{}, mainSegmentImg))
	if len(s.header) > 0 {
		frameParts = append(frameParts, common.NewBlocksContent(
			common.Style{
				Direction:  common.DirectionVertical,
				AlignItems: common.AlignItemsCenter,
				Gap:        10,
			}, s.header...))
	}
	if len(s.footer) > 0 {
		frameParts = append(frameParts, common.NewBlocksContent(
			common.Style{
				Direction:  common.DirectionVertical,
				AlignItems: common.AlignItemsCenter,
				Gap:        10,
			}, s.footer...))
	}

	frame := common.NewBlocksContent(
		common.Style{
			Direction:  common.DirectionVertical,
			AlignItems: common.AlignItemsCenter,
			Gap:        10,
		},
		frameParts...,
	)
	return frame.Render()
}
