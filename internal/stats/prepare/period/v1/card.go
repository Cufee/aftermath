package period

import (
	"fmt"
	"math"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

func NewCards(stats fetch.AccountStatsOverPeriod, glossary map[string]models.Vehicle, opts ...common.Option) (Cards, error) {
	options := common.DefaultOptions()
	for _, apply := range opts {
		apply(&options)
	}
	if glossary == nil {
		glossary = make(map[string]models.Vehicle)
	}

	var cards Cards
	cards.Rating.Meta = stats.RatingBattles.Rating() != frame.InvalidValue
	for _, column := range overviewBlocksRating {
		var columnBlocks []common.StatsBlock[BlockData, string]
		for _, preset := range column.Tags {
			var block common.StatsBlock[BlockData, string]
			b, err := presetToBlock(preset, stats.RatingBattles.StatsFrame, stats.RatingBattles.Vehicles, glossary)
			if err != nil {
				return Cards{}, err
			}
			block = b

			block.Localize(options.Printer())
			columnBlocks = append(columnBlocks, block)
		}
		cards.Rating.Type = common.CardTypeOverview
		cards.Rating.Blocks = append(cards.Rating.Blocks, OverviewColumn{columnBlocks, blockFlavor(column.Meta)})
	}

	for _, column := range overviewBlocks {
		var columnBlocks []common.StatsBlock[BlockData, string]
		for _, preset := range column.Tags {
			var block common.StatsBlock[BlockData, string]
			b, err := presetToBlock(preset, stats.RegularBattles.StatsFrame, stats.RegularBattles.Vehicles, glossary)
			if err != nil {
				return Cards{}, err
			}
			block = b

			block.Localize(options.Printer())
			columnBlocks = append(columnBlocks, block)
		}

		cards.Overview.Type = common.CardTypeOverview
		cards.Overview.Blocks = append(cards.Overview.Blocks, OverviewColumn{columnBlocks, blockFlavor(column.Meta)})
	}

	if len(stats.RegularBattles.Vehicles) < 1 || len(highlights) < 1 {
		return cards, nil
	}

	// Vehicle Highlights
	var minimumBattles float64 = 5
	periodDays := stats.PeriodEnd.Sub(stats.PeriodStart).Hours() / 24
	withFallback := func(battles float64) float64 {
		return math.Min(battles, float64(stats.RegularBattles.Battles.Float())/float64(len(highlights)))
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

	highlightedVehicles, err := common.GetHighlightedVehicles(highlights, stats.RegularBattles.Vehicles, int(minimumBattles))
	if err != nil {
		return Cards{}, err
	}

	for _, data := range highlightedVehicles {
		var vehicleBlocks []common.StatsBlock[BlockData, string]

		for _, preset := range data.Highlight.Blocks {
			block, err := presetToBlock(preset, *data.Vehicle.StatsFrame, nil, nil)
			if err != nil {
				return Cards{}, err
			}
			block.Localize(options.Printer())
			vehicleBlocks = append(vehicleBlocks, block)
		}

		glossary := glossary[data.Vehicle.VehicleID]
		glossary.ID = data.Vehicle.VehicleID

		cards.Highlights = append(cards.Highlights, VehicleCard{
			Title:  fmt.Sprintf("%s %s", logic.IntToRoman(glossary.Tier), glossary.Name(options.Locale())),
			Meta:   options.Printer()(data.Highlight.Label),
			Type:   common.CardTypeVehicle,
			Blocks: vehicleBlocks,
		})
	}

	return cards, nil
}
