package common

import (
	"github.com/cufee/aftermath/internal/stats/frame"
)

type Highlight struct {
	CompareWith      Tag
	Blocks           []Tag
	Label            string
	ignoreMinBattles bool
}

var (
	HighlightRecentBattle = Highlight{TagLastBattleTime, []Tag{TagBattles, TagAvgDamage, TagWN8}, "label_highlight_recent_battle", true}
	HighlightAvgDamage    = Highlight{TagAvgDamage, []Tag{TagBattles, TagAvgDamage, TagWN8}, "label_highlight_avg_damage", false}
	HighlightBattles      = Highlight{TagBattles, []Tag{TagBattles, TagAvgDamage, TagWN8}, "label_highlight_battles", false}
	HighlightWN8          = Highlight{TagWN8, []Tag{TagBattles, TagAvgDamage, TagWN8}, "label_highlight_wn8", false}
)

type highlightedVehicle struct {
	Highlight Highlight
	Vehicle   frame.VehicleStatsFrame
	Value     frame.Value
}

func GetHighlightedVehicles(highlights []Highlight, vehicles map[string]frame.VehicleStatsFrame, minBattles int) ([]highlightedVehicle, error) {
	leadersMap := make(map[string]highlightedVehicle)
	for _, vehicle := range vehicles {
		for _, highlight := range highlights {
			if int(vehicle.Battles.Float()) < minBattles && !highlight.ignoreMinBattles {
				continue
			}

			currentLeader, leaderExists := leadersMap[highlight.Label]

			value, err := PresetValue(highlight.CompareWith, *vehicle.StatsFrame, vehicle)
			if err != nil {
				return nil, err
			}

			if !frame.InvalidValue.Equals(value) && (!leaderExists || value.Float() > currentLeader.Value.Float()) {
				currentLeader.Highlight = highlight
				currentLeader.Vehicle = vehicle
				currentLeader.Value = value
				leadersMap[highlight.Label] = currentLeader
			}
		}
	}

	nominateVehicles := make(map[string]struct{})
	var highlightedVehicles []highlightedVehicle
	for _, highlight := range highlights {
		leader, leaderExists := leadersMap[highlight.Label]
		if !leaderExists {
			continue
		}
		if _, nominated := nominateVehicles[leader.Vehicle.VehicleID]; nominated {
			continue
		}
		highlightedVehicles = append(highlightedVehicles, leader)
		nominateVehicles[leader.Vehicle.VehicleID] = struct{}{}
	}
	return highlightedVehicles, nil
}
