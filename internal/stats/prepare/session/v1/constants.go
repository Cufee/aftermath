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
	{[]common.Tag{common.TagWN8}, BlockFlavorWN8},
	{[]common.Tag{common.TagAvgDamage, common.TagDamageRatio}, BlockFlavorDefault},
}

var ratingOverviewBlocks = []overviewColumnBlocks{
	{[]common.Tag{common.TagBattles, common.TagWinrate}, BlockFlavorDefault},
	{[]common.Tag{common.TagRankedRating}, BlockFlavorRating},
	{[]common.Tag{common.TagAvgDamage, common.TagDamageRatio}, BlockFlavorDefault},
}

var vehicleBlocks = []common.Tag{common.TagBattles, common.TagWinrate, common.TagAvgDamage, common.TagWN8}
var highlights = []common.Highlight{common.HighlightWN8, common.HighlightAvgDamage, common.HighlightBattles}

type Cards struct {
	Unrated UnratedCards `json:"unrated"`
	Rating  RatingCards  `json:"rating"`
}

type UnratedCards struct {
	Overview   OverviewCard  `json:"overview"`
	Vehicles   []VehicleCard `json:"vehicles"`
	Highlights []VehicleCard `json:"highlights"`
}

type RatingCards struct {
	Overview OverviewCard  `json:"overview"`
	Vehicles []VehicleCard `json:"vehicles"`
}

type OverviewColumn struct {
	Blocks []common.StatsBlock[BlockData] `json:"blocks"`
	Flavor blockFlavor                    `json:"flavor"`
}

type OverviewCard common.StatsCard[OverviewColumn, string]
type VehicleCard common.StatsCard[common.StatsBlock[BlockData], string]

type BlockData struct {
	Session frame.Value `json:"session"`
	Career  frame.Value `json:"career"`
}

type blockFlavor string

const (
	BlockFlavorDefault blockFlavor = "default"
	BlockFlavorRating  blockFlavor = "rating"
	BlockFlavorWN8     blockFlavor = "wn8"
)
