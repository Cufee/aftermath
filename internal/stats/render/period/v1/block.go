package period

import (
	common "github.com/cufee/aftermath/internal/render/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
)

func statsBlocksToColumnBlock(style overviewStyle, statsBlocks []prepare.StatsBlock[period.BlockData, string]) (common.Block, error) {
	var content []common.Block

	for _, statsBlock := range statsBlocks {
		if statsBlock.Data.Flavor == period.BlockFlavorSpecial {
			content = append(content, uniqueStatsBlock(style, statsBlock))
		} else {
			content = append(content, defaultStatsBlock(style, statsBlock))
		}
	}
	return common.NewBlocksContent(style.container, content...), nil
}

func uniqueStatsBlock(style overviewStyle, stats prepare.StatsBlock[period.BlockData, string]) common.Block {
	switch stats.Tag {
	case prepare.TagWN8:
		return uniqueBlockWN8(style, stats)
	default:
		return defaultStatsBlock(style, stats)
	}
}

func defaultStatsBlock(style overviewStyle, stats prepare.StatsBlock[period.BlockData, string]) common.Block {
	valueStyle, labelStyle := style.block(stats)

	blocks := []common.Block{common.NewTextContent(valueStyle, stats.Value().String())}
	blocks = append(blocks, common.NewTextContent(labelStyle, stats.Label))

	return common.NewBlocksContent(style.blockContainer, blocks...)
}

func uniqueBlockWN8(style overviewStyle, stats prepare.StatsBlock[period.BlockData, string]) common.Block {
	var blocks []common.Block

	valueStyle, labelStyle := style.block(stats)
	valueBlock := common.NewTextContent(valueStyle, stats.Value().String())

	ratingColors := common.GetWN8Colors(stats.Value().Float())
	if stats.Value().Float() <= 0 {
		ratingColors.Content = common.TextAlt
		ratingColors.Background = common.TextAlt
	}

	iconTop := common.AftermathLogo(ratingColors.Background, common.DefaultLogoOptions())
	iconBlockTop := common.NewImageContent(common.Style{Width: float64(iconTop.Bounds().Dx()), Height: float64(iconTop.Bounds().Dy())}, iconTop)

	style.blockContainer.Gap = 5
	blocks = append(blocks, common.NewBlocksContent(style.blockContainer, iconBlockTop, valueBlock))

	if stats.Value().Float() >= 0 {
		labelStyle.FontColor = ratingColors.Content
		blocks = append(blocks, common.NewBlocksContent(overviewSpecialRatingPillStyle(ratingColors.Background), common.NewTextContent(labelStyle, common.GetWN8TierName(stats.Value().Float()))))
	}

	return common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter, Gap: 0}, blocks...)
}
