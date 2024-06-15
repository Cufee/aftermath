package common

import (
	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/stats/frame"
)

type CardType string

const (
	CardTypeVehicle        CardType = "vehicle"
	CardTypeRatingVehicle  CardType = "ratingVehicle"
	CardTypeOverview       CardType = "overview"
	CardTypeHighlight      CardType = "overview"
	CardTypeTierPercentage CardType = "tierPercentage"
)

type StatsCard[B, M any] struct {
	Type   CardType `json:"type"`
	Title  string   `json:"title"`
	Blocks []B      `json:"blocks"`
	Meta   M        `json:"meta,omitempty"`
}

type StatsBlock[D any] struct {
	Data  D           `json:"data"`
	Tag   Tag         `json:"tag"`
	Label string      `json:"label"`
	Value frame.Value `json:"value"`
}

func NewBlock[D any](tag Tag, data D) StatsBlock[D] {
	return StatsBlock[D]{Tag: tag, Label: "label_" + tag.String(), Data: data}
}

func (block *StatsBlock[D]) Localize(printer func(string) string) {
	block.Label = printer(block.Label)
}

func (block *StatsBlock[D]) FillValue(stats frame.StatsFrame, args ...any) error {
	value, err := PresetValue(block.Tag, stats)
	if err != nil {
		return err
	}
	block.Value = value
	return nil
}

func PresetValue(preset Tag, stats frame.StatsFrame, args ...any) (frame.Value, error) {
	switch preset {
	case TagWN8:
		return stats.WN8(args...), nil
	case TagFrags:
		return stats.Frags, nil
	case TagBattles:
		return stats.Battles, nil
	case TagWinrate:
		return stats.WinRate(args...), nil
	case TagAccuracy:
		return stats.Accuracy(args...), nil
	case TagRankedRating:
		return stats.Rating, nil
	case TagAvgDamage:
		return stats.AvgDamage(args...), nil
	case TagDamageRatio:
		return stats.DamageRatio(args...), nil
	case TagSurvivalRatio:
		return stats.SurvivalRatio(args...), nil
	case TagSurvivalPercent:
		return stats.Survival(args...), nil
	case TagDamageDealt:
		return stats.DamageDealt, nil
	case TagDamageTaken:
		return stats.DamageReceived, nil

	// Some tags cannot be parsed here and should be implemented by the package
	// TagAvgTier - period
	// TagDamageBlocked - replay
	// TagDamageAssisted - replay
	// TagDamageAssistedCombined - replay
	default:
		return frame.InvalidValue, errors.New("invalid preset " + preset.String())
	}
}
