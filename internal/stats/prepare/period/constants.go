package period

import (
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

const TagAvgTier common.Tag = "avg_tier"

var selectedBlocks = [][]common.Tag{{common.TagDamageRatio, common.TagAvgDamage, common.TagAccuracy}, {common.TagWN8, common.TagBattles}, {TagAvgTier, common.TagWinrate, common.TagSurvivalPercent}}
var selectedHighlights = []highlight{HighlightBattles, HighlightWN8, HighlightAvgDamage}

type Cards struct {
	Overview   OverviewCard  `json:"overview"`
	Highlights []VehicleCard `json:"highlights"`
}

type OverviewCard common.StatsCard[[]common.StatsBlock[blockData], string]
type VehicleCard common.StatsCard[common.StatsBlock[blockData], string]

type blockData struct {
	Flavor blockFlavor `json:"flavor"`
}

type blockFlavor string

const (
	blockFlavorDefault   blockFlavor = "default"
	blockFlavorSpecial   blockFlavor = "special"
	blockFlavorSecondary blockFlavor = "secondary"
)
