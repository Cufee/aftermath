package session

import (
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

type overviewColumnBlocks struct {
	blocks []common.Tag
	flavor blockFlavor
}

var unratedOverviewBlocks = []overviewColumnBlocks{
	{[]common.Tag{common.TagBattles, common.TagWinrate}, BlockFlavorDefault},
	{[]common.Tag{common.TagWN8}, BlockFlavorRating},
	{[]common.Tag{common.TagAvgDamage, common.TagDamageRatio}, BlockFlavorDefault},
}

var ratingOverviewBlocks = []overviewColumnBlocks{
	{[]common.Tag{common.TagBattles, common.TagWinrate}, BlockFlavorDefault},
	{[]common.Tag{common.TagRankedRating}, BlockFlavorRating},
	{[]common.Tag{common.TagAvgDamage, common.TagDamageRatio}, BlockFlavorDefault},
}

var vehicleBlocks = []common.Tag{common.TagBattles, common.TagWinrate, common.TagAvgDamage, common.TagWN8}
var highlights = []common.Highlight{common.HighlightBattles, common.HighlightWN8, common.HighlightAvgDamage}

type Cards struct {
	Unrated UnratedCards `json:"unrated"`
	Rating  RatingCards  `json:"rating"`
}

type UnratedCards struct {
	Overview   OverviewCard
	Vehicles   []VehicleCard
	Highlights []VehicleCard
}

type RatingCards struct {
	Overview OverviewCard
	Vehicles []VehicleCard
}

type OverviewColumn struct {
	Blocks []common.StatsBlock[BlockData]
	Flavor blockFlavor
}

type OverviewCard common.StatsCard[OverviewColumn, string]
type VehicleCard common.StatsCard[common.StatsBlock[BlockData], string]

type BlockData struct {
	Session frame.Value
	Career  frame.Value
}

type blockFlavor string

const (
	BlockFlavorDefault blockFlavor = "default"
	BlockFlavorRating  blockFlavor = "rating"
)
