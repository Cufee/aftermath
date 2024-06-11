package period

import (
	"fmt"
	"math"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
)

func NewCards(stats fetch.AccountStatsOverPeriod, glossary map[string]database.Vehicle, opts ...Option) (Cards, error) {
	options := defaultOptions
	for _, apply := range opts {
		apply(&options)
	}
	if glossary == nil {
		glossary = make(map[string]database.Vehicle)
	}

	var cards Cards

	// Overview Card
	for _, column := range selectedBlocks {
		var columnBlocks []common.StatsBlock[BlockData]
		for _, preset := range column {
			block := presetToBlock(preset, stats.RegularBattles.StatsFrame)
			if preset == TagAvgTier {
				block = avgTierBlock(stats.RegularBattles.Vehicles, glossary)
			}

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
	var minimumBattles float64 = 5
	periodDays := stats.PeriodEnd.Sub(stats.PeriodStart).Hours() / 24
	withFallback := func(battles float64) float64 {
		return math.Min(battles, float64(stats.RegularBattles.Battles.Float())/float64(len(selectedHighlights)))
	}
	if periodDays > 90 {
		minimumBattles = withFallback(100)
	} else if periodDays > 60 {
		minimumBattles = withFallback(75)
	} else if periodDays > 30 {
		minimumBattles = withFallback(50)
	} else if periodDays > 14 {
		minimumBattles = withFallback(25)
	} else if periodDays > 7 {
		minimumBattles = withFallback(10)
	}

	highlightedVehicles := getHighlightedVehicles(selectedHighlights, stats.RegularBattles.Vehicles, int(minimumBattles))
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
			Title:  fmt.Sprintf("%s %s", common.IntToRoman(glossary.Tier), glossary.Name(options.locale)),
			Type:   common.CardTypeVehicle,
			Blocks: vehicleBlocks,
			Meta:   options.localePrinter(data.highlight.label),
		})
	}

	return cards, nil
}
