package session

import (
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
	"golang.org/x/text/language"
)

func NewCards(session, career fetch.AccountStatsOverPeriod, glossary map[string]database.Vehicle, opts ...common.Option) (Cards, error) {
	options := common.DefaultOptions
	for _, apply := range opts {
		apply(&options)
	}
	if glossary == nil {
		glossary = make(map[string]database.Vehicle)
	}

	var cards Cards

	// Rating battles overview
	if session.RatingBattles.Battles > 0 {
		card, err := makeOverviewCard(
			ratingOverviewBlocks,
			session.RatingBattles.StatsFrame,
			career.RatingBattles.StatsFrame,
			"label_overview_rating",
			options.Printer(),
			func(t common.Tag) common.Tag {
				if t == common.TagWN8 {
					return common.TagRankedRating
				}
				return t
			},
		)
		if err != nil {
			return Cards{}, err
		}
		cards.Rating.Overview = card
	}
	// Regular battles overview
	if session.RegularBattles.Battles > 0 {
		card, err := makeOverviewCard(
			unratedOverviewBlocks,
			session.RegularBattles.StatsFrame,
			career.RegularBattles.StatsFrame,
			"label_overview_unrated",
			options.Printer(),
			nil,
		)
		if err != nil {
			return Cards{}, err
		}
		cards.Unrated.Overview = card
	}

	// Rating battles vehicles
	for id, data := range session.RatingBattles.Vehicles {
		glossary := glossary[id]
		glossary.ID = id

		card, err := makeVehicleCard(
			[]common.Tag{common.TagWN8},
			common.CardTypeRatingVehicle,
			data,
			career.RatingBattles.Vehicles[id],
			options.Printer(),
			options.Locale(),
			glossary,
		)
		if err != nil {
			return Cards{}, err
		}
		cards.Rating.Vehicles = append(cards.Rating.Vehicles, card)
	}
	// Regular battles vehicles
	for id, data := range session.RegularBattles.Vehicles {
		glossary := glossary[id]
		glossary.ID = id

		card, err := makeVehicleCard(
			vehicleBlocks,
			common.CardTypeVehicle,
			data,
			career.RegularBattles.Vehicles[id],
			options.Printer(),
			options.Locale(),
			glossary,
		)
		if err != nil {
			return Cards{}, err
		}
		cards.Unrated.Vehicles = append(cards.Unrated.Vehicles, card)
	}

	return cards, nil
}

func makeVehicleCard(presets []common.Tag, cardType common.CardType, session, career frame.VehicleStatsFrame, printer func(string) string, locale language.Tag, glossary database.Vehicle) (VehicleCard, error) {
	var blocks []common.StatsBlock[BlockData]
	for _, preset := range presets {
		block, err := presetToBlock(preset, *session.StatsFrame, *career.StatsFrame)
		if err != nil {
			return VehicleCard{}, err
		}
		block.Localize(printer)
		blocks = append(blocks, block)
	}

	return VehicleCard{
		Meta:   common.IntToRoman(glossary.Tier),
		Title:  glossary.Name(locale),
		Type:   cardType,
		Blocks: blocks,
	}, nil
}

func makeOverviewCard(columns [][]common.Tag, session, career frame.StatsFrame, label string, printer func(string) string, replace func(common.Tag) common.Tag) (OverviewCard, error) {
	var blocks [][]common.StatsBlock[BlockData]
	for _, presets := range columns {
		var column []common.StatsBlock[BlockData]
		for _, p := range presets {
			preset := p
			if replace != nil {
				preset = replace(p)
			}
			block, err := presetToBlock(preset, session, career)
			if err != nil {
				return OverviewCard{}, err
			}
			block.Localize(printer)
			column = append(column, block)
		}
		blocks = append(blocks, column)
	}
	return OverviewCard{
		Type:   common.CardTypeOverview,
		Title:  printer(label),
		Blocks: blocks,
	}, nil
}
