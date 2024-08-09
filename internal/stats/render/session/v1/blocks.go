package session

import (
	"image/color"

	"github.com/cufee/aftermath/internal/stats/frame"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/internal/stats/render/assets"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

func makeSpecialRatingColumn(block prepare.StatsBlock[session.BlockData], width float64) common.Block {
	blockStyle := vehicleBlockStyle()
	switch block.Tag {
	case prepare.TagWN8:
		ratingColors := common.GetWN8Colors(block.Value().Float())
		if block.Value().Float() <= 0 {
			ratingColors.Content = common.TextAlt
			ratingColors.Background = common.TextAlt
		}

		var column []common.Block
		iconTop := common.AftermathLogo(ratingColors.Background, common.DefaultLogoOptions())
		column = append(column, common.NewImageContent(common.Style{Width: specialRatingIconSize, Height: specialRatingIconSize}, iconTop))

		pillColor := ratingColors.Background
		if block.Value().Float() < 0 {
			pillColor = color.Transparent
		}
		column = append(column, common.NewBlocksContent(overviewColumnStyle(width),
			common.NewTextContent(blockStyle.session, block.Data.Session().String()),
			common.NewBlocksContent(
				overviewSpecialRatingPillStyle(pillColor),
				common.NewTextContent(overviewSpecialRatingLabelStyle(ratingColors.Content), common.GetWN8TierName(block.Value().Float())),
			),
		))
		return common.NewBlocksContent(specialRatingColumnStyle(), column...)

	case prepare.TagRankedRating:
		var column []common.Block
		icon, ok := common.GetRatingIcon(block.Value(), specialRatingIconSize)
		if ok {
			column = append(column, icon)
		}
		column = append(column, common.NewBlocksContent(overviewColumnStyle(width),
			blockWithDoubleVehicleIcon(common.NewTextContent(blockStyle.session, block.Data.Session().String()), block.Data.Session(), block.Data.Career()),
		))
		return common.NewBlocksContent(specialRatingColumnStyle(), column...)

	default:
		return common.NewBlocksContent(statsBlockStyle(width),
			common.NewTextContent(blockStyle.session, block.Data.Session().String()),
			common.NewTextContent(blockStyle.label, block.Label),
		)
	}
}

func blankBlock(width, height float64) common.Block {
	return common.NewImageContent(common.Style{Width: width, Height: height}, gg.NewContext(int(width), int(height)).Image())
}

func vehicleWN8Icon(wn8 frame.Value) common.Block {
	ratingColors := common.GetWN8Colors(wn8.Float())
	if wn8.Float() <= 0 {
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

	case session.Float() > career.Float():
		icon, _ := assets.GetLoadedImage("triangle-up-solid")
		ctx.DrawImage(imaging.Fill(icon, int(vehicleComparisonIconSize), int(vehicleComparisonIconSize), imaging.Center, imaging.Linear), 0, 0)
		return common.NewImageContent(common.Style{BackgroundColor: color.NRGBA{3, 201, 169, 255}}, ctx.Image())

	case session.Float() < career.Float():
		icon, _ := assets.GetLoadedImage("triangle-down-solid")
		ctx.DrawImage(imaging.Fill(icon, int(vehicleComparisonIconSize), int(vehicleComparisonIconSize), imaging.Center, imaging.Linear), 0, 0)
		return common.NewImageContent(common.Style{BackgroundColor: color.NRGBA{231, 130, 141, 255}}, ctx.Image())

	case session.Float() == career.Float():
		ctx.DrawRectangle(vehicleComparisonIconSize*0.2, (vehicleComparisonIconSize-2)/2, vehicleComparisonIconSize*0.8, 2)
		ctx.SetColor(common.TextAlt)
		ctx.Fill()
		return common.NewImageContent(common.Style{}, ctx.Image())
	}
}

func blockWithVehicleIcon(block common.Block, session, career frame.Value) common.Block {
	return common.NewBlocksContent(common.Style{}, block, vehicleComparisonIcon(session, career))
}

func blockWithDoubleVehicleIcon(block common.Block, session, career frame.Value) common.Block {
	return common.NewBlocksContent(common.Style{}, blankBlock(vehicleComparisonIconSize, 1), block, vehicleComparisonIcon(session, career))
}

func blockShouldHaveCompareIcon(block prepare.StatsBlock[session.BlockData]) bool {
	return block.Tag != prepare.TagBattles
}
