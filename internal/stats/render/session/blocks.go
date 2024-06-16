package session

import (
	"image/color"

	prepare "github.com/cufee/aftermath/internal/stats/prepare/common"
	"github.com/cufee/aftermath/internal/stats/prepare/session"
	"github.com/cufee/aftermath/internal/stats/render/common"
)

func makeSpecialRatingColumn(block prepare.StatsBlock[session.BlockData], width float64) common.Block {
	switch block.Tag {
	case prepare.TagWN8:
		ratingColors := common.GetWN8Colors(block.Value.Float())
		if block.Value.Float() <= 0 {
			ratingColors.Content = common.TextAlt
			ratingColors.Background = common.TextAlt
		}

		var column []common.Block
		iconTop := common.AftermathLogo(ratingColors.Background, common.DefaultLogoOptions())
		column = append(column, common.NewImageContent(common.Style{Width: specialRatingIconSize, Height: specialRatingIconSize}, iconTop))

		pillColor := ratingColors.Background
		if block.Value.Float() < 0 {
			pillColor = color.Transparent
		}
		column = append(column, common.NewBlocksContent(overviewColumnStyle(width),
			common.NewTextContent(vehicleBlockStyle.session, block.Data.Session.String()),
			common.NewBlocksContent(
				overviewSpecialRatingPillStyle(pillColor),
				common.NewTextContent(overviewSpecialRatingLabelStyle(ratingColors.Content), common.GetWN8TierName(block.Value.Float())),
			),
		))
		return common.NewBlocksContent(specialRatingColumnStyle, column...)

	case prepare.TagRankedRating:
		var column []common.Block
		icon, ok := getRatingIcon(block.Value)
		if ok {
			column = append(column, icon)
		}
		column = append(column, common.NewBlocksContent(overviewColumnStyle(width),
			common.NewTextContent(vehicleBlockStyle.session, block.Data.Session.String()),
		))
		return common.NewBlocksContent(specialRatingColumnStyle, column...)

	default:
		return common.NewBlocksContent(statsBlockStyle(width),
			common.NewTextContent(vehicleBlockStyle.session, block.Data.Session.String()),
			// common.NewTextContent(vehicleBlockStyle.career, block.Data.Career.String()),
			common.NewTextContent(vehicleBlockStyle.label, block.Label),
		)
	}
}
