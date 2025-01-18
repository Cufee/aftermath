package period

import (
	"github.com/cufee/aftermath/internal/render/common"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func newOverviewCard(data period.OverviewCard, columnWidth float64) *facepaint.Block {
	if len(data.Blocks) == 0 {
		return nil
	}

	var columns []*facepaint.Block
	for _, column := range data.Blocks {
		columns = append(columns, newOverviewColumn(column, columnWidth))
	}
	// card
	return facepaint.NewBlocksContent(styledOverviewCard.card.Options(), columns...)
}

func newOverviewColumn(data period.OverviewColumn, columnWidth float64) *facepaint.Block {
	var columnBlocks []*facepaint.Block
	for _, block := range data.Blocks {
		switch block.Tag {
		default:
			columnBlocks = append(columnBlocks, newOverviewBlockWithIcon(block, nil))
		case prepare.TagWN8:
			columnBlocks = append(columnBlocks, newOverviewWN8Block(block))
		case prepare.TagRankedRating:
			columnBlocks = append(columnBlocks, newOverviewRatingBlock(block))
		}
	}
	// column
	return facepaint.NewBlocksContent(style.NewStyle(
		style.Parent(styledOverviewCard.column),
		style.SetWidth(columnWidth),
	), columnBlocks...)
}

func newOverviewBlockWithIcon(block prepare.StatsBlock[period.BlockData, string], icon *facepaint.Block) *facepaint.Block {
	blockStyle := styledOverviewCard.styleBlock(block)
	if icon == nil {
		// block
		return facepaint.NewBlocksContent(blockStyle.valueContainer.Options(),
			// value
			facepaint.MustNewTextContent(blockStyle.value.Options(), block.V.String()),
			// label
			facepaint.MustNewTextContent(blockStyle.label.Options(), block.Label),
		)
	}
	// wrapper
	return facepaint.NewBlocksContent(blockStyle.wrapper.Options(),
		icon,
		// block
		facepaint.NewBlocksContent(blockStyle.valueContainer.Options(),
			// value
			facepaint.MustNewTextContent(blockStyle.value.Options(), block.V.String()),
			// label
			facepaint.MustNewTextContent(blockStyle.label.Options(), block.Label),
		))
}

func newOverviewWN8Block(block prepare.StatsBlock[period.BlockData, string]) *facepaint.Block {
	ratingColors := common.GetWN8Colors(block.Value().Float())
	if block.Value().Float() <= 0 {
		ratingColors.Background = common.TextAlt
	}
	icon, _ := facepaint.NewImageContent(
		style.NewStyle(style.SetWidth(iconSizeWN8)),
		common.AftermathLogo(ratingColors.Background, common.DefaultLogoOptions()),
	)
	block.Label = common.GetWN8TierName(block.Value().Float())
	return newOverviewBlockWithIcon(block, icon)
}

func newOverviewRatingBlock(block prepare.StatsBlock[period.BlockData, string]) *facepaint.Block {
	icon, _ := common.GetRatingIcon(block.V, iconSizeRating)
	block.Label = common.GetRatingTierName(block.Value().Float())
	return newOverviewBlockWithIcon(block, icon)
}

// func newHighlightCard(style highlightStyle, card period.VehicleCard) common.Block {
// 	titleBlock :=
// 		common.NewBlocksContent(common.Style{
// 			Direction: common.DirectionVertical,
// 		},
// 			common.NewTextContent(style.cardTitle, card.Meta),
// 			common.NewTextContent(style.tankName, card.Title),
// 		)

// 	var contentRow []common.Block
// 	for _, block := range card.Blocks {
// 		contentRow = append(contentRow, common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter},
// 			common.NewTextContent(style.blockValue, block.Value().String()),
// 			common.NewTextContent(style.blockLabel, block.Label),
// 		))
// 	}

// 	return common.NewBlocksContent(style.container, titleBlock, common.NewBlocksContent(common.Style{
// 		Gap: style.container.Gap,
// 	}, contentRow...))
// }
