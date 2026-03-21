package period

import (
	"github.com/cufee/aftermath/internal/render/common"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func newHighlightCard(hlStyle common.HighlightCardStyle, data period.VehicleCard, blockSizes map[prepare.Tag]float64) *facepaint.Block {
	leftSide := facepaint.NewBlocksContent(hlStyle.TitleWrapper.Options(),
		facepaint.MustNewTextContent(hlStyle.TitleLabel().Options(), data.Meta),
		facepaint.MustNewTextContent(hlStyle.TitleVehicle().Options(), data.Title),
	)

	var rightSide []*facepaint.Block
	for _, block := range data.Blocks {
		rightSide = append(rightSide, facepaint.NewBlocksContent(style.NewStyle(style.Parent(hlStyle.Stats), style.SetWidth(blockSizes[block.Tag])),
			facepaint.MustNewTextContent(hlStyle.BlockValue().Options(), block.V.String()),
			facepaint.MustNewTextContent(hlStyle.BlockLabel().Options(), block.Label),
		))
	}

	return facepaint.NewBlocksContent(hlStyle.Card.Options(),
		leftSide,
		facepaint.NewBlocksContent(hlStyle.StatsWrapper.Options(), rightSide...),
	)
}
