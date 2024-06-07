package period

import (
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

func avgTierBlock(vehicles map[string]frame.VehicleStatsFrame, glossary map[string]database.Vehicle) common.StatsBlock[BlockData] {
	block := common.StatsBlock[BlockData](common.NewBlock(common.TagAvgTier, BlockData{Flavor: BlockFlavorSecondary}))

	var weightedTotal, battlesTotal float32

	for _, vehicle := range vehicles {
		if data, ok := glossary[vehicle.VehicleID]; ok && data.Tier > 0 {
			battlesTotal += vehicle.Battles.Float()
			weightedTotal += vehicle.Battles.Float() * float32(data.Tier)
		}
	}
	if battlesTotal == 0 {
		block.Value = frame.InvalidValue
		return block
	}

	block.Value = frame.ValueFloatDecimal(weightedTotal / battlesTotal)
	return block
}
