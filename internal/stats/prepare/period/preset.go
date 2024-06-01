package period

import (
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

func presetToBlock(preset common.Tag, stats frame.StatsFrame) common.StatsBlock[BlockData] {
	block := common.StatsBlock[BlockData](common.NewBlock(preset, BlockData{}))
	block.FillValue(stats)

	switch preset {
	case common.TagWN8:
		block.Data.Flavor = BlockFlavorSpecial

	case common.TagBattles:
		block.Data.Flavor = BlockFlavorSpecial

	case common.TagSurvivalRatio:
		block.Data.Flavor = BlockFlavorSecondary

	case common.TagSurvivalPercent:
		block.Data.Flavor = BlockFlavorSecondary

	case common.TagAccuracy:
		block.Data.Flavor = BlockFlavorSecondary

	case common.TagDamageRatio:
		block.Data.Flavor = BlockFlavorSecondary

	default:
		block.Data.Flavor = BlockFlavorDefault
	}

	return block
}
