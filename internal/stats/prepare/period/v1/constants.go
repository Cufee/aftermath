package period

import (
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

const TagAvgTier common.Tag = "avg_tier"

var overviewBlocks = []common.TagColumn[string]{
	{Tags: []common.Tag{common.TagBattles, common.TagWinrate, common.TagAccuracy}, Meta: string(BlockFlavorDefault)},
	{Tags: []common.Tag{common.TagWN8}, Meta: string(BlockFlavorSpecial)},
	{Tags: []common.Tag{TagAvgTier, common.TagAvgDamage, common.TagDamageRatio}, Meta: string(BlockFlavorDefault)},
}

var overviewBlocksRating = []common.TagColumn[string]{
	{Tags: []common.Tag{common.TagWinrate, common.TagBattles}, Meta: string(BlockFlavorDefault)},
	{Tags: []common.Tag{common.TagRankedRating}, Meta: string(BlockFlavorSpecial)},
	{Tags: []common.Tag{common.TagAvgDamage, common.TagDamageRatio}, Meta: string(BlockFlavorDefault)},
}

var highlights = []common.Highlight{common.HighlightBattles, common.HighlightWN8, common.HighlightAvgDamage}

type Cards struct {
	Rating     RatingOverviewCard `json:"rating"`
	Overview   OverviewCard       `json:"overview"`
	Highlights []VehicleCard      `json:"highlights"`
}

type OverviewColumn struct {
	Blocks []common.StatsBlock[BlockData, string]
	Flavor blockFlavor
}

type OverviewCard common.StatsCard[OverviewColumn, string]
type RatingOverviewCard common.StatsCard[OverviewColumn, bool]
type VehicleCard common.StatsCard[common.StatsBlock[BlockData, string], string]

type BlockData struct {
	Flavor blockFlavor `json:"flavor"`
}

type blockFlavor string

const (
	BlockFlavorDefault   blockFlavor = "default"
	BlockFlavorSpecial   blockFlavor = "special"
	BlockFlavorSecondary blockFlavor = "secondary"
)
