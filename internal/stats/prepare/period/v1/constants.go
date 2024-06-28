package period

import (
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

const TagAvgTier common.Tag = "avg_tier"

type overviewColumnBlocks struct {
	blocks []common.Tag
	flavor blockFlavor
}

var overviewBlocks = []overviewColumnBlocks{
	{[]common.Tag{common.TagDamageRatio, common.TagAvgDamage, common.TagAccuracy}, BlockFlavorDefault},
	{[]common.Tag{common.TagWN8, common.TagBattles}, BlockFlavorSpecial},
	{[]common.Tag{TagAvgTier, common.TagWinrate, common.TagSurvivalPercent}, BlockFlavorDefault},
}
var highlights = []common.Highlight{common.HighlightBattles, common.HighlightWN8, common.HighlightAvgDamage}

type Cards struct {
	Overview   OverviewCard  `json:"overview"`
	Highlights []VehicleCard `json:"highlights"`
}

type OverviewColumn struct {
	Blocks []common.StatsBlock[BlockData]
	Flavor blockFlavor
}

type OverviewCard common.StatsCard[OverviewColumn, string]
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
