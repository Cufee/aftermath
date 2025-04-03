package session

import (
	"slices"

	"github.com/cufee/aftermath/cmd/frontend/assets"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"golang.org/x/text/language"
)

type cardBuilder struct {
	session  fetch.AccountStatsOverPeriod
	career   fetch.AccountStatsOverPeriod
	glossary map[string]models.Vehicle
}

func (b *cardBuilder) vehicle(id string) models.Vehicle {
	vehicle := b.glossary[id]
	vehicle.ID = id
	return vehicle
}

func NewCards(session, career fetch.AccountStatsOverPeriod, glossary map[string]models.Vehicle, opts ...common.Option) (Cards, error) {
	options := common.DefaultOptions()
	for _, apply := range opts {
		apply(&options)
	}
	if glossary == nil {
		glossary = make(map[string]models.Vehicle)
	}

	builder := cardBuilder{session, career, glossary}

	var cards Cards

	{ // Rating battles overview
		var ratingColumns = ratingOverviewBlocks
		if options.RatingColumns != nil {
			ratingColumns = options.RatingColumns
		}

		// Rating battles overview
		card, err := builder.makeOverviewCard(
			ratingColumns,
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

		blocks := unratedOverviewBlocks
		if options.UnratedColumns != nil {
			blocks = options.UnratedColumns
		}

		card, err := builder.makeOverviewCard(
			blocks,
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

	{ // Rating battles vehicles
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

			card, err := builder.makeVehicleCard(
				data.VehicleID,
				[]common.Tag{common.TagWN8},
				common.CardTypeRatingVehicle,
				data,
				career.RatingBattles.Vehicles[data.VehicleID],
				options.Printer(),
				options.Locale(),
			)
			if err != nil {
				return Cards{}, err
			}
			cards.Rating.Vehicles = append(cards.Rating.Vehicles, card)
		}
	}

	{ // Regular battles vehicles
		var unratedVehicles []frame.VehicleStatsFrame
		for _, vehicle := range session.RegularBattles.Vehicles {
			unratedVehicles = append(unratedVehicles, vehicle)
		}
		slices.SortFunc(unratedVehicles, func(a, b frame.VehicleStatsFrame) int {
			return int(b.LastBattleTime.Unix() - a.LastBattleTime.Unix())
		})

		blocks := vehicleBlocks
		if options.VehicleTags != nil {
			blocks = options.VehicleTags
		}

		for _, data := range unratedVehicles {
			if len(cards.Unrated.Vehicles) >= 10 {
				break
			}
			if data.Battles < 1 {
				continue
			}

			card, err := builder.makeVehicleCard(
				data.VehicleID,
				blocks,
				common.CardTypeVehicle,
				data,
				career.RegularBattles.Vehicles[data.VehicleID],
				options.Printer(),
				options.Locale(),
			)
			if err != nil {
				return Cards{}, err
			}
			cards.Unrated.Vehicles = append(cards.Unrated.Vehicles, card)
		}
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
		card, err := builder.makeHighlightCard(
			data.Vehicle.VehicleID,
			data.Highlight,
			data.Vehicle,
			frame.VehicleStatsFrame{},
			options.Printer(),
			options.Locale(),
		)
		if err != nil {
			return Cards{}, err
		}
		cards.Unrated.Highlights = append(cards.Unrated.Highlights, card)
	}

	return cards, nil
}

func (b *cardBuilder) makeVehicleCard(vehicleID string, presets []common.Tag, cardType common.CardType, session, career frame.VehicleStatsFrame, printer func(string) string, locale language.Tag) (VehicleCard, error) {
	var sFrame, cFrame frame.StatsFrame
	if session.StatsFrame != nil {
		sFrame = *session.StatsFrame
	}
	if career.StatsFrame != nil {
		cFrame = *career.StatsFrame
	}

	var blocks []common.StatsBlock[BlockData, string]
	for _, preset := range presets {
		block, err := b.presetToBlock(preset, sFrame, cFrame)
		if err != nil {
			return VehicleCard{}, err
		}
		block.Localize(printer)
		block.Meta = specialIconPath(preset, sFrame, true)
		blocks = append(blocks, block)
	}

	return VehicleCard{
		Meta:   logic.IntToRoman(b.vehicle(vehicleID).Tier),
		Title:  b.vehicle(vehicleID).Name(locale),
		Type:   cardType,
		Blocks: blocks,
	}, nil
}

func (b *cardBuilder) makeHighlightCard(vehicleID string, highlight common.Highlight, session, career frame.VehicleStatsFrame, printer func(string) string, locale language.Tag) (VehicleCard, error) {
	var sFrame, cFrame frame.StatsFrame
	if session.StatsFrame != nil {
		sFrame = *session.StatsFrame
	}
	if career.StatsFrame != nil {
		cFrame = *career.StatsFrame
	}

	var blocks []common.StatsBlock[BlockData, string]
	for _, preset := range highlight.Blocks {
		block, err := b.presetToBlock(preset, sFrame, cFrame)
		if err != nil {
			return VehicleCard{}, err
		}
		block.Localize(printer)
		blocks = append(blocks, block)
	}

	return VehicleCard{
		Meta:   printer(highlight.Label),
		Title:  logic.IntToRoman(b.vehicle(vehicleID).Tier) + " " + b.vehicle(vehicleID).Name(locale),
		Type:   common.CardTypeHighlight,
		Blocks: blocks,
	}, nil
}

func (b *cardBuilder) makeOverviewCard(columns []common.TagColumn[string], session, career frame.StatsFrame, label string, printer func(string) string, replace func(common.Tag) common.Tag) (OverviewCard, error) {
	var blocks []OverviewColumn
	for _, columnBlocks := range columns {
		var column []common.StatsBlock[BlockData, string]
		for _, p := range columnBlocks.Tags {
			preset := p
			if replace != nil {
				preset = replace(p)
			}
			block, err := b.presetToBlock(preset, session, career)
			if err != nil {
				return OverviewCard{}, err
			}
			block.Localize(printer)
			block.Meta = specialIconPath(preset, session, false)
			column = append(column, block)
		}
		blocks = append(blocks, OverviewColumn{
			Flavor: blockFlavor(columnBlocks.Meta),
			Blocks: column,
		})
	}
	return OverviewCard{
		Type:   common.CardTypeOverview,
		Title:  printer(label),
		Blocks: blocks,
	}, nil
}

func specialIconPath(preset common.Tag, frame frame.StatsFrame, small bool) string {
	if preset == common.TagWN8 && small {
		return assets.WN8IconPathSmall(frame.WN8().Float())
	}
	if preset == common.TagWN8 {
		return assets.WN8IconPath(frame.WN8().Float())
	}
	if preset == common.TagRankedRating {
		return assets.RatingIconPath(frame.Rating().Float())
	}
	return ""
}
