package period

import (
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

func presetToBlock(preset common.Tag, stats frame.StatsFrame, vehicles map[string]frame.VehicleStatsFrame, glossary map[string]models.Vehicle) (common.StatsBlock[BlockData], error) {
	block := common.StatsBlock[BlockData](common.NewBlock(preset, BlockData{}))

	switch preset {
	case common.TagWN8:
		block.Data.Flavor = BlockFlavorSpecial

	case common.TagBattles:
		block.Data.Flavor = BlockFlavorSecondary

	case common.TagSurvivalRatio:
		block.Data.Flavor = BlockFlavorSecondary

	case common.TagAvgTier:
		block.Data.Flavor = BlockFlavorSecondary
		err := block.FillValue(frame.StatsFrame{}, vehicles, glossary)
		return block, err

	case common.TagSurvivalPercent:
		block.Data.Flavor = BlockFlavorSecondary

	case common.TagAccuracy:
		block.Data.Flavor = BlockFlavorSecondary

	case common.TagDamageRatio:
		block.Data.Flavor = BlockFlavorSecondary

	default:
		block.Data.Flavor = BlockFlavorDefault
	}

	err := block.FillValue(stats)
	return block, err
}
