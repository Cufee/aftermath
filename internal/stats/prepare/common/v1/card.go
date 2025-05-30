package common

import (
	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

type CardType string

const (
	CardTypeVehicle        CardType = "vehicle"
	CardTypeRatingVehicle  CardType = "ratingVehicle"
	CardTypeOverview       CardType = "overview"
	CardTypeHighlight      CardType = "highlight"
	CardTypeTierPercentage CardType = "tierPercentage"
)

type StatsCard[B, M any] struct {
	Type   CardType `json:"type"`
	Title  string   `json:"title"`
	Blocks []B      `json:"blocks"`
	Meta   M        `json:"meta,omitempty"`
}

type StatsBlock[D, M any] struct {
	Data  D           `json:"data"`
	Tag   Tag         `json:"tag"`
	Label string      `json:"label"`
	V     frame.Value `json:"value"`
	Meta  M           `json:"meta,omitempty"`
}

func (b *StatsBlock[D, M]) Value() frame.Value {
	if b.V == nil {
		return frame.InvalidValue
	}
	return b.V
}

func (b *StatsBlock[D, M]) SetValue(value frame.Value) {
	b.V = value
}

func NewBlock[D, M any](tag Tag, data D) StatsBlock[D, M] {
	return StatsBlock[D, M]{Tag: tag, Label: "label_" + tag.String(), Data: data}
}

func (block *StatsBlock[D, M]) Localize(printer func(string) string) {
	block.Label = printer(block.Label)
}

func (block *StatsBlock[D, M]) FillValue(stats frame.StatsFrame, args ...any) error {
	value, err := PresetValue(block.Tag, stats, args...)
	if err != nil {
		return err
	}
	block.V = value
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
		return stats.Rating(args...), nil
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
	case TagLastBattleTime:
		if len(args) != 1 {
			return frame.InvalidValue, errors.New("invalid args length for last_battle_time")
		}
		vehicle, ok := args[0].(frame.VehicleStatsFrame)
		if !ok {
			return frame.InvalidValue, errors.New("invalid args for last_battle_time, second arg should be vehicle")
		}
		return frame.ValueInt(vehicle.LastBattleTime.Unix()), nil
	case TagAvgTier:
		if len(args) != 2 {
			return frame.InvalidValue, errors.New("invalid args length for avg_tier")
		}
		vehicles, ok := args[0].(map[string]frame.VehicleStatsFrame)
		if !ok {
			return frame.InvalidValue, errors.New("invalid args for avg_tier, first arg should be vehicles")
		}
		glossary, ok := args[1].(map[string]models.Vehicle)
		if !ok {
			return frame.InvalidValue, errors.New("invalid args for avg_tier, second arg should be glossary")
		}
		return avgTierValue(vehicles, glossary), nil

	// Some tags cannot be parsed here and should be implemented by the package
	// TagDamageBlocked - replay
	// TagDamageAssisted - replay
	// TagDamageAssistedCombined - replay
	default:
		return frame.InvalidValue, errors.New("invalid preset " + preset.String())
	}
}

func avgTierValue(vehicles map[string]frame.VehicleStatsFrame, glossary map[string]models.Vehicle) frame.Value {
	var weightedTotal, battlesTotal float32
	for _, vehicle := range vehicles {
		if data, ok := glossary[vehicle.VehicleID]; ok && data.Tier > 0 {
			battlesTotal += vehicle.Battles.Float()
			weightedTotal += vehicle.Battles.Float() * float32(data.Tier)
		}
	}
	if battlesTotal == 0 {
		return frame.InvalidValue
	}
	return frame.ValueFloatDecimal(weightedTotal / battlesTotal)
}
