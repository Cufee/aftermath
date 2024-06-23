package common

import (
	"github.com/cufee/aftermath/internal/stats/frame"
)

type Highlight struct {
	CompareWith Tag
	Blocks      []Tag
	Label       string
}

var (
	HighlightAvgDamage = Highlight{TagAvgDamage, []Tag{TagBattles, TagAvgDamage, TagWN8}, "label_highlight_avg_damage"}
	HighlightBattles   = Highlight{TagBattles, []Tag{TagBattles, TagAvgDamage, TagWN8}, "label_highlight_battles"}
	HighlightWN8       = Highlight{TagWN8, []Tag{TagBattles, TagAvgDamage, TagWN8}, "label_highlight_wn8"}
)

type highlightedVehicle struct {
	Highlight Highlight
	Vehicle   frame.VehicleStatsFrame
	Value     frame.Value
}

func GetHighlightedVehicles(highlights []Highlight, vehicles map[string]frame.VehicleStatsFrame, minBattles int) ([]highlightedVehicle, error) {
	leadersMap := make(map[string]highlightedVehicle)
	for _, vehicle := range vehicles {
		if int(vehicle.Battles.Float()) < minBattles {
			continue
		}

		for _, highlight := range highlights {
			currentLeader, leaderExists := leadersMap[highlight.Label]
			value, err := PresetValue(highlight.CompareWith, *vehicle.StatsFrame)
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
