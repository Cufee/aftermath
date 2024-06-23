package period

import (
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

const TagAvgTier common.Tag = "avg_tier"

var overviewBlocks = [][]common.Tag{{common.TagDamageRatio, common.TagAvgDamage, common.TagAccuracy}, {common.TagWN8, common.TagBattles}, {TagAvgTier, common.TagWinrate, common.TagSurvivalPercent}}
var highlights = []common.Highlight{common.HighlightBattles, common.HighlightWN8, common.HighlightAvgDamage}

type Cards struct {
	Overview   OverviewCard  `json:"overview"`
	Highlights []VehicleCard `json:"highlights"`
}

type OverviewCard common.StatsCard[[]common.StatsBlock[BlockData], string]
type VehicleCard common.StatsCard[common.StatsBlock[BlockData], string]

type BlockData struct {
	Flavor blockFlavor `json:"flavor"`
}

type blockFlavor string

const (
	BlockFlavorDefault   blockFlavor = "default"
	BlockFlavorSpecial   blockFlavor = "special"
	BlockFlavorSecondary blockFlavor = "secondary"
)
