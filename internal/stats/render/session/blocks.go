package session

import (
	"image/color"

	"github.com/cufee/aftermath/internal/stats/frame"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common"
	"github.com/cufee/aftermath/internal/stats/prepare/session"
	"github.com/cufee/aftermath/internal/stats/render/assets"
	"github.com/cufee/aftermath/internal/stats/render/common"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
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

func blankBlock(width, height float64) common.Block {
	return common.NewImageContent(common.Style{Width: width, Height: height}, gg.NewContext(int(width), int(height)).Image())
}

func vehicleWN8Icon(wn8 frame.Value) common.Block {
	ratingColors := common.GetWN8Colors(wn8.Float())
	if wn8.Float() <= 0 {
		ratingColors.Content = common.TextAlt
		ratingColors.Background = common.TextAlt
	}
	iconTop := common.AftermathLogo(ratingColors.Background, common.SmallLogoOptions())
	return common.NewImageContent(common.Style{Width: vehicleWN8IconSize, Height: vehicleWN8IconSize}, iconTop)
}

func vehicleComparisonIcon(session, career frame.Value) common.Block {
	ctx := gg.NewContext(int(vehicleComparisonIconSize), int(vehicleComparisonIconSize))
	switch {
	case session.Float() < 0, career.Float() < 0:
		fallthrough
	default:
		return common.NewImageContent(common.Style{}, ctx.Image())

	case session.Float() > career.Float()*1.05:
		icon, _ := assets.GetLoadedImage("triangle-up-solid")
		ctx.DrawImage(imaging.Fill(icon, int(vehicleComparisonIconSize), int(vehicleComparisonIconSize), imaging.Center, imaging.Linear), 0, 0)
		return common.NewImageContent(common.Style{BackgroundColor: color.RGBA{3, 201, 169, 255}}, ctx.Image())

	case session.Float() < career.Float()*0.95:
		icon, _ := assets.GetLoadedImage("triangle-down-solid")
		ctx.DrawImage(imaging.Fill(icon, int(vehicleComparisonIconSize), int(vehicleComparisonIconSize), imaging.Center, imaging.Linear), 0, 0)
		return common.NewImageContent(common.Style{BackgroundColor: color.RGBA{231, 130, 141, 255}}, ctx.Image())
	}
}

func blockWithVehicleIcon(block common.Block, session, career frame.Value) common.Block {
	return common.NewBlocksContent(common.Style{}, block, vehicleComparisonIcon(session, career))
}

func blockShouldHaveCompareIcon(block prepare.StatsBlock[session.BlockData]) bool {
	return block.Tag != prepare.TagBattles
}
