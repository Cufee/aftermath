package period

import (
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

type highlight struct {
	compareWith common.Tag
	blocks      []common.Tag
	label       string
}

var (
	HighlightAvgDamage = highlight{common.TagAvgDamage, []common.Tag{common.TagBattles, common.TagAvgDamage, common.TagWN8}, "label_highlight_avg_damage"}
	HighlightBattles   = highlight{common.TagBattles, []common.Tag{common.TagBattles, common.TagAvgDamage, common.TagWN8}, "label_highlight_battles"}
	HighlightWN8       = highlight{common.TagWN8, []common.Tag{common.TagBattles, common.TagAvgDamage, common.TagWN8}, "label_highlight_wn8"}
)

type highlightedVehicle struct {
	highlight highlight
	vehicle   frame.VehicleStatsFrame
	value     frame.Value
}

func getHighlightedVehicles(highlights []highlight, vehicles map[string]frame.VehicleStatsFrame, minBattles int) []highlightedVehicle {
	leadersMap := make(map[string]highlightedVehicle)
	for _, vehicle := range vehicles {
		if int(vehicle.Battles.Float()) < minBattles {
			continue
		}

		for _, highlight := range highlights {
			currentLeader, leaderExists := leadersMap[highlight.label]
			block := presetToBlock(highlight.compareWith, *vehicle.StatsFrame)

			if !leaderExists || block.Value.Float() > currentLeader.value.Float() {
				currentLeader.highlight = highlight
				currentLeader.value = block.Value
				currentLeader.vehicle = vehicle
				leadersMap[highlight.label] = currentLeader
			}
		}
	}

	nominateVehicles := make(map[string]struct{})
	var highlightedVehicles []highlightedVehicle
	for _, highlight := range highlights {
		leader, leaderExists := leadersMap[highlight.label]
		if !leaderExists {
			continue
		}
		if _, nominated := nominateVehicles[leader.vehicle.VehicleID]; nominated {
			continue
		}
		highlightedVehicles = append(highlightedVehicles, leader)
		nominateVehicles[leader.vehicle.VehicleID] = struct{}{}
	}
	return highlightedVehicles
}
