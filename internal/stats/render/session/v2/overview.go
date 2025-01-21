package session

import (
	"image/color"

	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/frame"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
)

func newRatingOverviewCard(data session.RatingCards, columnWidth map[bool]float64) *facepaint.Block {
	if len(data.Overview.Blocks) == 0 {
		return nil
	}

	var columns []*facepaint.Block
	for _, column := range data.Overview.Blocks {
		columns = append(columns, newOverviewColumn(styledOverviewCard, column, columnWidth[column.Flavor == session.BlockFlavorDefault]))
	}
	// card
	return facepaint.NewBlocksContent(styledOverviewCard.card.Options(), columns...)
}

func newUnratedOverviewCard(data session.OverviewCard, columnWidth map[bool]float64) *facepaint.Block {
	if len(data.Blocks) == 0 {
		return nil
	}

	var columns []*facepaint.Block
	for _, column := range data.Blocks {
		columns = append(columns, newOverviewColumn(styledOverviewCard, column, columnWidth[column.Flavor == session.BlockFlavorDefault]))
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
		style.Parent(stl.column(data)),
		style.SetMinWidth(columnWidth),
	), columnBlocks...)
}

func newOverviewBlockWithIcon(blockStyle blockStyle, block prepare.StatsBlock[session.BlockData, string], icon *facepaint.Block) *facepaint.Block {
	if icon == nil {
		// block
		return facepaint.NewBlocksContent(blockStyle.wrapper.Options(), newOverviewBlock(blockStyle, block))
	}
	// wrapper
	return facepaint.NewBlocksContent(blockStyle.wrapper.Options(),
		// icon
		facepaint.NewBlocksContent(blockStyle.iconWrapper.Options(), icon),
		// block
		newOverviewBlock(blockStyle, block),
	)
}

func newOverviewBlock(blockStyle blockStyle, block prepare.StatsBlock[session.BlockData, string]) *facepaint.Block {
	switch block.Tag {
	case prepare.TagBattles, prepare.TagWN8, prepare.TagRankedRating:
		return facepaint.NewBlocksContent(blockStyle.wrapper.Options(),
			facepaint.NewBlocksContent(blockStyle.valueContainer.Options(),
				// value
				facepaint.MustNewTextContent(blockStyle.value.Options(), block.V.String()),
				// label
				facepaint.MustNewTextContent(blockStyle.label.Options(), block.Label),
			),
		)

	default:
		var indicatorColor color.Color = color.Transparent
		if block.Data.S.Float() != frame.InvalidValue.Float() && block.Data.C.Float() != frame.InvalidValue.Float() {
			if block.Data.S.Float() > block.Data.C.Float() {
				indicatorColor = color.NRGBA{163, 235, 177, 255}
			}
			if block.Data.S.Float() < block.Data.C.Float() {
				indicatorColor = color.NRGBA{219, 154, 156, 255}
			}
		}

		indicator := facepaint.NewEmptyContent(style.NewStyle(style.Parent(style.Style{
			Position:                style.PositionAbsolute,
			BackgroundColor:         indicatorColor,
			MinWidth:                14,
			Height:                  3,
			BorderRadiusTopLeft:     1.5,
			BorderRadiusTopRight:    1.5,
			BorderRadiusBottomLeft:  1.5,
			BorderRadiusBottomRight: 1.5,
			Bottom:                  4,
		})))

		return facepaint.NewBlocksContent(blockStyle.wrapper.Options(),
			facepaint.NewBlocksContent(blockStyle.valueContainer.Options(),
				// value
				facepaint.NewBlocksContent(style.NewStyle(), indicator, facepaint.MustNewTextContent(blockStyle.value.Options(), block.V.String())),
				// label
				facepaint.MustNewTextContent(blockStyle.label.Options(), block.Label),
			),
		)
	}
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
