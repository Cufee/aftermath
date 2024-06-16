package session

import (
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

var unratedOverviewBlocks = [][]common.Tag{{common.TagBattles, common.TagWinrate}, {common.TagWN8}, {common.TagAvgDamage, common.TagDamageRatio}}
var ratingOverviewBlocks = [][]common.Tag{{common.TagBattles, common.TagWinrate}, {common.TagRankedRating}, {common.TagAvgDamage, common.TagDamageRatio}}
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

type OverviewCard common.StatsCard[[]common.StatsBlock[BlockData], string]
type VehicleCard common.StatsCard[common.StatsBlock[BlockData], string]

type BlockData struct {
	Session frame.Value
	Career  frame.Value
}
