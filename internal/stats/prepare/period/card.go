package period

import (
	"fmt"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

func NewCards(stats fetch.AccountStatsOverPeriod, glossary map[string]database.GlossaryVehicle, opts ...Option) (Cards, error) {
	options := defaultOptions
	for _, apply := range opts {
		apply(&options)
	}
	if glossary == nil {
		glossary = make(map[string]database.GlossaryVehicle)
	}

	var cards Cards

	// Overview Card
	for _, column := range selectedBlocks {
		var columnBlocks []common.StatsBlock[BlockData]
		for _, preset := range column {
			if preset == TagAvgTier {
				// value := calculateAvgTier(input.Stats.Vehicles, input.VehicleGlossary)
				continue
			}
			block := presetToBlock(preset, stats.RegularBattles.StatsFrame)
			block.Localize(options.localePrinter)
			columnBlocks = append(columnBlocks, block)

		}

		cards.Overview.Type = common.CardTypeOverview
		cards.Overview.Blocks = append(cards.Overview.Blocks, columnBlocks)
	}

	if len(stats.RegularBattles.Vehicles) < 1 || len(selectedHighlights) < 1 {
		return cards, nil
	}

	// Vehicle Highlights
	var minimumBattles int = 5
	periodDays := stats.PeriodEnd.Sub(stats.PeriodStart).Hours() / 24
	if periodDays > 90 {
		minimumBattles = 100
	} else if periodDays > 60 {
		minimumBattles = 75
	} else if periodDays > 30 {
		minimumBattles = 50
	} else if periodDays > 14 {
		minimumBattles = 25
	} else if periodDays > 7 {
		minimumBattles = 10
	}

	highlightedVehicles := getHighlightedVehicles(selectedHighlights, stats.RegularBattles.Vehicles, minimumBattles)
	for _, data := range highlightedVehicles {
		var vehicleBlocks []common.StatsBlock[BlockData]

		for _, preset := range data.highlight.blocks {
			block := presetToBlock(preset, *data.vehicle.StatsFrame)
			block.Localize(options.localePrinter)
			vehicleBlocks = append(vehicleBlocks, block)
		}

		glossary := glossary[data.vehicle.VehicleID]
		glossary.ID = data.vehicle.VehicleID

		cards.Highlights = append(cards.Highlights, VehicleCard{
			Title:  fmt.Sprintf("%s %s", common.IntToRoman(glossary.Tier), glossary.Name(options.localePrinter)),
			Type:   common.CardTypeVehicle,
			Blocks: vehicleBlocks,
			Meta:   options.localePrinter(data.highlight.label),
		})
	}

	return cards, nil
}
