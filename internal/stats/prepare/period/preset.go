package period

import (
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

func presetToBlock(preset common.Tag, stats fetch.StatsFrame) common.StatsBlock[blockData] {
	block := common.StatsBlock[blockData](common.NewBlock(preset, blockData{}))
	block.FillValue(stats)

	switch preset {
	case common.TagWN8:
		block.Data.Flavor = blockFlavorSpecial

	case common.TagBattles:
		block.Data.Flavor = blockFlavorSpecial

	case common.TagSurvivalRatio:
		block.Data.Flavor = blockFlavorSecondary

	case common.TagSurvivalPercent:
		block.Data.Flavor = blockFlavorSecondary

	case common.TagAccuracy:
		block.Data.Flavor = blockFlavorSecondary

	case common.TagAvgDamage:
		block.Data.Flavor = blockFlavorDefault

	case common.TagDamageRatio:
		block.Data.Flavor = blockFlavorSecondary

	default:
		block.Data.Flavor = blockFlavorDefault
	}

	return block
}
