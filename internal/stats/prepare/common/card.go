package common

import (
	"errors"

	"github.com/cufee/aftermath/internal/stats/fetch"
)

type cardType string

const (
	CardTypeVehicle        cardType = "vehicle"
	CardTypeRatingVehicle  cardType = "ratingVehicle"
	CardTypeOverview       cardType = "overview"
	CardTypeHighlight      cardType = "overview"
	CardTypeTierPercentage cardType = "tierPercentage"
)

type StatsCard[B, M any] struct {
	Type   cardType `json:"type"`
	Title  string   `json:"title"`
	Blocks []B      `json:"blocks"`
	Meta   M        `json:"meta,omitempty"`
}

type StatsBlock[D any] struct {
	Data  D           `json:"data"`
	Tag   Tag         `json:"tag"`
	Label string      `json:"label"`
	Value fetch.Value `json:"value"`
}

func NewBlock[D any](tag Tag, data D) StatsBlock[D] {
	return StatsBlock[D]{Tag: tag, Label: "label_" + tag.String(), Data: data}
}

func (block *StatsBlock[D]) Localize(printer func(string) string) {
	block.Label = printer(block.Label)
}

func (block *StatsBlock[D]) FillValue(stats fetch.StatsFrame) error {
	switch block.Tag {
	case TagWN8:
		block.Value = stats.WN8()
	case TagFrags:
		block.Value = stats.Frags
	case TagBattles:
		block.Value = stats.Battles
	case TagWinrate:
		block.Value = stats.WinRate()
	case TagAccuracy:
		block.Value = stats.Accuracy()
	case TagRankedRating:
		block.Value = stats.Rating
	case TagAvgDamage:
		block.Value = stats.AvgDamage()
	case TagDamageRatio:
		block.Value = stats.DamageRatio()
	case TagSurvivalRatio:
		block.Value = stats.SurvivalRatio()
	case TagSurvivalPercent:
		block.Value = stats.Survival()
	case TagDamageDealt:
		block.Value = stats.DamageDealt
	case TagDamageTaken:
		block.Value = stats.DamageReceived

	// Some tags cannot be parsed here and should be implemented by the package
	// TagAvgTier - period
	// TagDamageBlocked - replay
	// TagDamageAssisted - replay
	// TagDamageAssistedCombined - replay
	default:
		return errors.New("invalid preset")
	}

	return nil

}
