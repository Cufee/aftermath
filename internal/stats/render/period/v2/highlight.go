package period

import (
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func newHighlightCard(data period.VehicleCard, blockSizes map[prepare.Tag]float64) *facepaint.Block {
	leftSide := facepaint.NewBlocksContent(styledHighlightCard.titleWrapper.Options(),
		facepaint.MustNewTextContent(styledHighlightCard.titleLabel().Options(), data.Meta),
		facepaint.MustNewTextContent(styledHighlightCard.titleVehicle().Options(), data.Title),
	)

	var rightSide []*facepaint.Block
	for _, block := range data.Blocks {
		rightSide = append(rightSide, facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledHighlightCard.stats), style.SetWidth(blockSizes[block.Tag])),
			facepaint.MustNewTextContent(styledHighlightCard.blockValue().Options(), block.V.String()),
			facepaint.MustNewTextContent(styledHighlightCard.blockLabel().Options(), block.Label),
		))
	}

	return facepaint.NewBlocksContent(styledHighlightCard.card.Options(),
		append([]*facepaint.Block{leftSide}, rightSide...)...,
	)

}
