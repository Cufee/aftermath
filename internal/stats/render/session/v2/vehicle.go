package session

import (
	"github.com/cufee/aftermath/internal/render/common"
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
		statsBlocks = append(statsBlocks, facepaint.MustNewTextContent(styledVehicleCard.value(blockWidth[block.Tag]).Options(), block.Value().String()))
	}

	return facepaint.NewBlocksContent(styledVehicleCard.card,
		title,
		facepaint.NewBlocksContent(styledVehicleCard.stats, statsBlocks...),
	)

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
		icon, _ := facepaint.NewImageContent(
			style.NewStyle(style.SetWidth(vehicleIconSizeWN8)),
			common.AftermathLogo(ratingColors.Background, common.SmallLogoOptions()),
		)
		return icon
	}
	return facepaint.NewEmptyContent(style.NewStyle(style.SetWidth(vehicleIconSizeWN8)))
}
