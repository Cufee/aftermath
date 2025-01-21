package sakura

import (
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/internal/stats/render/session/special/sakura/assets"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func newOverviewCard(data session.OverviewCard, columnWidth map[bool]float64) *facepaint.Block {
	if len(data.Blocks) == 0 {
		return nil
	}

	var columns []*facepaint.Block
	for i, column := range data.Blocks {
		columns = append(columns, newOverviewColumn(i, column, columnWidth[column.Flavor == session.BlockFlavorDefault]))
	}
	// card
	return facepaint.NewBlocksContent(overviewStyle.card(), columns...)
}

func newOverviewColumn(index int, data session.OverviewColumn, columnWidth float64) *facepaint.Block {
	var columnBlocks []*facepaint.Block
	for _, block := range data.Blocks {
		switch block.Tag {
		default:
			columnBlocks = append(columnBlocks, newOverviewBlockWithIcon(index, block, nil))
		case prepare.TagWN8:
			// color := common.GetWN8Colors(block.Value().Float())
			img, _ := assets.Image("cherry")
			icon := facepaint.MustNewImageContent(style.NewStyle(
				style.Parent(style.Style{
					// Color: color.Background,
				}),
				style.SetWidth(overviewSpecialIconSize),
			), img)
			columnBlocks = append(columnBlocks, newOverviewBlockWithIcon(index, block, icon))
		case prepare.TagRankedRating:
			columnBlocks = append(columnBlocks, newOverviewBlockWithIcon(index, block, nil))
		}
	}
	// column
	return facepaint.NewBlocksContent(overviewStyle.column(index, string(data.Flavor)).Chain(style.SetMinWidth(columnWidth)), columnBlocks...)
}

func newOverviewBlockWithIcon(colIndex int, block prepare.StatsBlock[session.BlockData, string], icon *facepaint.Block) *facepaint.Block {
	if icon == nil {
		return newOverviewBlock(colIndex, block)
	}
	// wrapper
	return facepaint.NewBlocksContent(overviewStyle.blockContainer(colIndex, block.Tag),
		// icon
		icon,
		// block
		newOverviewBlock(colIndex, block),
	)
}

func newOverviewBlock(colIndex int, block prepare.StatsBlock[session.BlockData, string]) *facepaint.Block {
	return facepaint.NewBlocksContent(overviewStyle.blockContainer(colIndex, block.Tag),
		// value
		facepaint.MustNewTextContent(overviewStyle.blockValue(block.Tag), block.V.String()),
		// label
		facepaint.MustNewTextContent(overviewStyle.blockLabel(block.Tag), block.Label),
	)
}
