package session

import (
	"github.com/cufee/aftermath/internal/render/common"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func newRatingOverviewCard(data session.RatingCards, columnWidth map[string]float64) *facepaint.Block {
	if len(data.Overview.Blocks) == 0 {
		return nil
	}

	var columns []*facepaint.Block
	for _, column := range data.Overview.Blocks {
		columns = append(columns, newOverviewColumn(styledOverviewCard, column, columnWidth[string(column.Flavor)]))
	}
	// card
	return facepaint.NewBlocksContent(styledOverviewCard.card.Options(), columns...)
}

func newUnratedOverviewCard(data session.OverviewCard, columnWidth map[string]float64) *facepaint.Block {
	if len(data.Blocks) == 0 {
		return nil
	}

	var columns []*facepaint.Block
	for _, column := range data.Blocks {
		columns = append(columns, newOverviewColumn(styledOverviewCard, column, columnWidth[string(column.Flavor)]))
	}
	// card
	return facepaint.NewBlocksContent(styledOverviewCard.card.Options(), columns...)
}

func newOverviewColumn(stl overviewCardStyle, data session.OverviewColumn, columnWidth float64) *facepaint.Block {
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

func newOverviewBlockWithIcon(blockStyle blockStyle, block prepare.StatsBlock[session.BlockData, string], icon *facepaint.Block) *facepaint.Block {
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
		// icon
		icon,
		// block
		facepaint.NewBlocksContent(blockStyle.valueContainer.Options(),
			// value
			facepaint.MustNewTextContent(blockStyle.value.Options(), block.V.String()),
			// label
			facepaint.MustNewTextContent(blockStyle.label.Options(), block.Label),
		))
}

func newOverviewWN8Block(blockStyle blockStyle, block prepare.StatsBlock[session.BlockData, string]) *facepaint.Block {
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

func newOverviewRatingBlock(blockStyle blockStyle, block prepare.StatsBlock[session.BlockData, string]) *facepaint.Block {
	icon, _ := common.GetRatingIcon(block.V, iconSizeRating)
	block.Label = common.GetRatingTierName(block.Value().Float())
	return newOverviewBlockWithIcon(blockStyle, block, icon)
}
