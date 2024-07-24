package session

import (
	"fmt"
	"slices"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"golang.org/x/text/language"
)

func NewCards(session, career fetch.AccountStatsOverPeriod, glossary map[string]models.Vehicle, opts ...common.Option) (Cards, error) {
	options := common.DefaultOptions
	for _, apply := range opts {
		apply(&options)
	}
	if glossary == nil {
		glossary = make(map[string]models.Vehicle)
	}

	var cards Cards

	// Rating battles overview
	if options.VehicleID == "" {
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
	{ // Regular battles overview
		sessionFrame := session.RegularBattles.StatsFrame
		careerFrame := career.RegularBattles.StatsFrame
		overviewLabel := "label_overview_unrated"
		if options.VehicleID != "" {
			sessionFrame = frame.StatsFrame{}
			careerFrame = frame.StatsFrame{}
			s, ok := session.RegularBattles.Vehicles[options.VehicleID]
			if ok {
				sessionFrame = *s.StatsFrame
			}
			c, ok := career.RegularBattles.Vehicles[options.VehicleID]
			if ok {
				careerFrame = *c.StatsFrame
			}
			glossary := glossary[options.VehicleID]
			glossary.ID = options.VehicleID
			overviewLabel = fmt.Sprintf("%s %s", common.IntToRoman(glossary.Tier), glossary.Name(options.Locale()))
		}

		card, err := makeOverviewCard(
			unratedOverviewBlocks,
			sessionFrame,
			careerFrame,
			overviewLabel,
			options.Printer(),
			nil,
		)
		if err != nil {
			return Cards{}, err
		}
		cards.Unrated.Overview = card
	}

	if options.VehicleID != "" { // we don't vehicle cards if a specific vehicle was selected
		return cards, nil
	}

	// Rating battles vehicles
	var ratingVehicles []frame.VehicleStatsFrame
	for _, vehicle := range session.RatingBattles.Vehicles {
		ratingVehicles = append(ratingVehicles, vehicle)
	}
	slices.SortFunc(ratingVehicles, func(a, b frame.VehicleStatsFrame) int {
		return int(b.LastBattleTime.Unix() - a.LastBattleTime.Unix())
	})
	for _, data := range ratingVehicles {
		if len(cards.Rating.Vehicles) >= 10 {
			break
		}

		glossary := glossary[data.VehicleID]
		glossary.ID = data.VehicleID

		card, err := makeVehicleCard(
			[]common.Tag{common.TagWN8},
			common.CardTypeRatingVehicle,
			data,
			career.RatingBattles.Vehicles[data.VehicleID],
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
	var unratedVehicles []frame.VehicleStatsFrame
	for _, vehicle := range session.RegularBattles.Vehicles {
		unratedVehicles = append(unratedVehicles, vehicle)
	}
	slices.SortFunc(unratedVehicles, func(a, b frame.VehicleStatsFrame) int {
		return int(b.LastBattleTime.Unix() - a.LastBattleTime.Unix())
	})
	for _, data := range unratedVehicles {
		if len(cards.Unrated.Vehicles) >= 10 {
			break
		}
		if data.Battles < 1 {
			continue
		}

		glossary := glossary[data.VehicleID]
		glossary.ID = data.VehicleID

		card, err := makeVehicleCard(
			vehicleBlocks,
			common.CardTypeVehicle,
			data,
			career.RegularBattles.Vehicles[data.VehicleID],
			options.Printer(),
			options.Locale(),
			glossary,
		)
		if err != nil {
			return Cards{}, err
		}
		cards.Unrated.Vehicles = append(cards.Unrated.Vehicles, card)
	}

	// Vehicle Highlights
	var minVehicleBattles = 1
	if len(session.RegularBattles.Vehicles) > 0 {
		minVehicleBattles = int(session.RegularBattles.Battles.Float()) / len(session.RegularBattles.Vehicles)
	}

	highlightedVehicles, err := common.GetHighlightedVehicles(highlights, session.RegularBattles.Vehicles, minVehicleBattles)
	if err != nil {
		return Cards{}, err
	}

	for _, data := range highlightedVehicles {
		glossary := glossary[data.Vehicle.VehicleID]
		glossary.ID = data.Vehicle.VehicleID

		card, err := makeHighlightCard(
			data.Highlight,
			data.Vehicle,
			frame.VehicleStatsFrame{},
			options.Printer(),
			options.Locale(),
			glossary,
		)
		if err != nil {
			return Cards{}, err
		}
		cards.Unrated.Highlights = append(cards.Unrated.Highlights, card)
	}

	return cards, nil
}

func makeVehicleCard(presets []common.Tag, cardType common.CardType, session, career frame.VehicleStatsFrame, printer func(string) string, locale language.Tag, glossary models.Vehicle) (VehicleCard, error) {
	var sFrame, cFrame frame.StatsFrame
	if session.StatsFrame != nil {
		sFrame = *session.StatsFrame
	}
	if career.StatsFrame != nil {
		cFrame = *career.StatsFrame
	}

	var blocks []common.StatsBlock[BlockData]
	for _, preset := range presets {
		block, err := presetToBlock(preset, sFrame, cFrame)
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

func makeHighlightCard(highlight common.Highlight, session, career frame.VehicleStatsFrame, printer func(string) string, locale language.Tag, glossary models.Vehicle) (VehicleCard, error) {
	var sFrame, cFrame frame.StatsFrame
	if session.StatsFrame != nil {
		sFrame = *session.StatsFrame
	}
	if career.StatsFrame != nil {
		cFrame = *career.StatsFrame
	}

	var blocks []common.StatsBlock[BlockData]
	for _, preset := range highlight.Blocks {
		block, err := presetToBlock(preset, sFrame, cFrame)
		if err != nil {
			return VehicleCard{}, err
		}
		block.Localize(printer)
		blocks = append(blocks, block)
	}

	return VehicleCard{
		Meta:   printer(highlight.Label),
		Title:  common.IntToRoman(glossary.Tier) + " " + glossary.Name(locale),
		Type:   common.CardTypeHighlight,
		Blocks: blocks,
	}, nil
}

func makeOverviewCard(columns []overviewColumnBlocks, session, career frame.StatsFrame, label string, printer func(string) string, replace func(common.Tag) common.Tag) (OverviewCard, error) {
	var blocks []OverviewColumn
	for _, columnBlocks := range columns {
		var column []common.StatsBlock[BlockData]
		for _, p := range columnBlocks.blocks {
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
		blocks = append(blocks, OverviewColumn{
			Flavor: columnBlocks.flavor,
			Blocks: column,
		})
	}
	return OverviewCard{
		Type:   common.CardTypeOverview,
		Title:  printer(label),
		Blocks: blocks,
	}, nil
}
