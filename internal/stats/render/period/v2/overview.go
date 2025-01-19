package period

import (
	"github.com/cufee/aftermath/internal/render/common"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func newRatingOverviewCard(data period.RatingOverviewCard, columnWidth map[string]float64) *facepaint.Block {
	if len(data.Blocks) == 0 {
		return nil
	}

	var columns []*facepaint.Block
	for _, column := range data.Blocks {
		columns = append(columns, newOverviewColumn(styledRatingOverviewCard, column, columnWidth[string(column.Flavor)]))
	}
	// card
	return facepaint.NewBlocksContent(styledRatingOverviewCard.card.Options(), columns...)
}

func newUnratedOverviewCard(data period.OverviewCard, columnWidth map[string]float64) *facepaint.Block {
	if len(data.Blocks) == 0 {
		return nil
	}

	var columns []*facepaint.Block
	for _, column := range data.Blocks {
		columns = append(columns, newOverviewColumn(styledUnratedOverviewCard, column, columnWidth[string(column.Flavor)]))
	}
	// card
	return facepaint.NewBlocksContent(styledUnratedOverviewCard.card.Options(), columns...)
}

func newOverviewColumn(stl overviewCardStyle, data period.OverviewColumn, columnWidth float64) *facepaint.Block {
	var columnBlocks []*facepaint.Block
	for _, block := range data.Blocks {
		switch block.Tag {
		default:
			columnBlocks = append(columnBlocks, newOverviewBlockWithIcon(stl.styleBlock(block), block, nil))
		case prepare.TagWN8:
			columnBlocks = append(columnBlocks, newOverviewWN8Block(stl.styleBlock(block), block))
		case prepare.TagRankedRating:
			columnBlocks = append(columnBlocks, newOverviewRatingBlock(stl.styleBlock(block), block))
		}
	}
	// column
	return facepaint.NewBlocksContent(style.NewStyle(
		style.Parent(stl.column),
		style.SetWidth(columnWidth),
	), columnBlocks...)
}

func newOverviewBlockWithIcon(blockStyle blockStyle, block prepare.StatsBlock[period.BlockData, string], icon *facepaint.Block) *facepaint.Block {
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

func newOverviewWN8Block(blockStyle blockStyle, block prepare.StatsBlock[period.BlockData, string]) *facepaint.Block {
	ratingColors := common.GetWN8Colors(block.Value().Float())
	if block.Value().Float() <= 0 {
		ratingColors.Background = common.TextAlt
	}
	icon, _ := facepaint.NewImageContent(
		style.NewStyle(style.SetWidth(iconSizeWN8)),
		common.AftermathLogo(ratingColors.Background, common.DefaultLogoOptions()),
	)
	block.Label = common.GetWN8TierName(block.Value().Float())
	return newOverviewBlockWithIcon(blockStyle, block, icon)
}

func newOverviewRatingBlock(blockStyle blockStyle, block prepare.StatsBlock[period.BlockData, string]) *facepaint.Block {
	icon, _ := common.GetRatingIcon(block.V, iconSizeRating)
	block.Label = common.GetRatingTierName(block.Value().Float())
	return newOverviewBlockWithIcon(blockStyle, block, icon)
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
