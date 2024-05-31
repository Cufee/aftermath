package common

import (
	"errors"
)

type Tag string

func (t Tag) String() string {
	return string(t)
}

const (
	// Global
	TagWN8             Tag = "wn8"
	TagFrags           Tag = "frags"
	TagBattles         Tag = "battles"
	TagWinrate         Tag = "winrate"
	TagAccuracy        Tag = "accuracy"
	TagRankedRating    Tag = "ranked_rating"
	TagAvgDamage       Tag = "avg_damage"
	TagDamageRatio     Tag = "damage_ratio"
	TagSurvivalRatio   Tag = "survival_ratio"
	TagSurvivalPercent Tag = "survival_percent"
	TagDamageDealt     Tag = "damage_dealt"
	TagDamageTaken     Tag = "damage_taken"
)

func ParseTags(tags ...string) ([]Tag, error) {
	var parsed []Tag
	for _, tag := range tags {
		if tag == "" {
			continue
		}
		parsed = append(parsed, Tag(tag))
	}
	if len(parsed) == 0 {
		return nil, errors.New("no tags provided")
	}
	return parsed, nil
}
