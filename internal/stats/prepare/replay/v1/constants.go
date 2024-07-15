package replay

import (
	"github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

const (
	TagDamageBlocked          common.Tag = "blocked"
	TagDamageAssisted         common.Tag = "assisted"
	TagDamageAssistedCombined common.Tag = "assisted_combined"
)

type Cards struct {
	Header  HeaderCard `json:"header"`
	Allies  []Card     `json:"allies"`
	Enemies []Card     `json:"enemies"`
}

type Card common.StatsCard[common.StatsBlock[BlockData], CardMeta]

type HeaderCard struct {
	Result   string `json:"result"`
	MapName  string `json:"map"`
	GameMode string `json:"gameMode"`
}

type CardMeta struct {
	Player replay.Player `json:"player"`
	Tags   []common.Tag  `json:"tags"`
}

type BlockData struct{}
