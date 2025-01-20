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

func newVehicleCard(data session.VehicleCard, blockWidth map[prepare.Tag]float64) *facepaint.Block {
	title := facepaint.NewBlocksContent(styledVehicleCard.titleWrapper,
		facepaint.MustNewTextContent(styledVehicleCard.titleText(), data.Title),
		newVehicleWN8Icon(data),
	)

	var statsBlocks []*facepaint.Block
	for _, block := range data.Blocks {
		statsBlocks = append(statsBlocks, newVehicleBlockValue(block, blockWidth))
	}

	return facepaint.NewBlocksContent(styledVehicleCard.card,
		title,
		facepaint.NewBlocksContent(styledVehicleCard.stats, statsBlocks...),
	)
}

func newVehicleBlockValue(block prepare.StatsBlock[session.BlockData, string], blockWidth map[prepare.Tag]float64) *facepaint.Block {
	switch block.Tag {
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
			MinWidth:                20,
			Height:                  3,
			BorderRadiusTopLeft:     1.5,
			BorderRadiusTopRight:    1.5,
			BorderRadiusBottomLeft:  1.5,
			BorderRadiusBottomRight: 1.5,
			Bottom:                  20,
		})))

		return facepaint.NewBlocksContent(styledVehicleCard.valueWrapper(blockWidth[block.Tag]).Options(),
			indicator,
			facepaint.MustNewTextContent(styledVehicleCard.value().Options(), block.Value().String()),
		)

	case prepare.TagBattles:
		return facepaint.NewBlocksContent(styledVehicleCard.valueWrapper(blockWidth[block.Tag]).Options(),
			facepaint.MustNewTextContent(styledVehicleCard.value().Options(), block.Value().String()),
		)
	}

}

func newVehicleLegendCard(data session.VehicleCard, blockWidth map[prepare.Tag]float64) *facepaint.Block {
	var legendBlocks []*facepaint.Block
	for _, block := range data.Blocks {
		legendBlocks = append(legendBlocks,
			facepaint.NewBlocksContent(styledVehicleLegendPill(blockWidth[block.Tag]),
				facepaint.MustNewTextContent(styledVehicleLegendPillText().Options(), block.Label),
			),
		)
	}

	return facepaint.NewBlocksContent(styledVehicleLegendPillWrapper,
		facepaint.NewBlocksContent(styledVehicleCard.stats, legendBlocks...),
	)

}

func newVehicleWN8Icon(data session.VehicleCard) *facepaint.Block {
	for _, block := range data.Blocks {
		if block.Tag != prepare.TagWN8 {
			continue
		}
		ratingColors := common.GetWN8Colors(block.Value().Float())
		if block.Value().Float() <= 0 {
			ratingColors.Background = common.TextAlt
		}
		icon := facepaint.NewBlocksContent(styledVehicleCard.titleIconWrapper,
			facepaint.MustNewImageContent(
				style.NewStyle(style.SetWidth(vehicleIconSizeWN8), style.SetWidth(vehicleIconSizeWN8)),
				common.AftermathLogo(ratingColors.Background, common.TinyLogoOptions()),
			))
		return icon
	}
	return facepaint.NewEmptyContent(style.NewStyle(style.SetWidth(vehicleIconSizeWN8)))
}
