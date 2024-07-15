package replay

import (
	"fmt"
	"slices"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1/replay"

	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

func NewCards(replay fetch.Replay, glossary map[string]models.Vehicle, opts ...common.Option) (Cards, error) {
	options := common.DefaultOptions
	for _, apply := range opts {
		apply(&options)
	}
	if glossary == nil {
		glossary = make(map[string]models.Vehicle)
	}

	var cards Cards
	cards.Header.Result = options.Printer()("label_" + string(replay.Outcome))
	cards.Header.GameMode = options.Printer()(replay.GameMode.String())
	cards.Header.MapName = "label_map_name_unknown"
	if n := replay.Map.LocalizedNames[options.Locale()]; n != "" {
		cards.Header.MapName = n
	}
	cards.Header.MapName = options.Printer()(cards.Header.MapName)

	sortTeams(replay.Teams)
	// Allies
	for _, player := range replay.Teams.Allies {
		vehicle := glossary[player.VehicleID]
		vehicle.ID = player.VehicleID
		name := fmt.Sprintf("%s %s", common.IntToRoman(vehicle.Tier), vehicle.Name(options.Locale()))
		card, err := playerToCard(player, name, nil)
		if err != nil {
			return cards, err
		}
		cards.Allies = append(cards.Allies, card)
	}
	// Enemies
	for _, player := range replay.Teams.Enemies {
		vehicle := glossary[player.VehicleID]
		vehicle.ID = player.VehicleID
		name := fmt.Sprintf("%s %s", common.IntToRoman(vehicle.Tier), vehicle.Name(options.Locale()))
		card, err := playerToCard(player, name, nil)
		if err != nil {
			return cards, err
		}
		cards.Enemies = append(cards.Enemies, card)
	}

	return cards, nil
}

func playerToCard(player replay.Player, vehicleName string, blocks []common.Tag) (Card, error) {
	card := Card{
		Title: vehicleName,
		Type:  common.CardTypeVehicle,
		Meta:  CardMeta{player, blocks},
	}
	for _, preset := range blocks {
		block, err := presetToBlock(player, preset)
		if err != nil {
			return card, err
		}
		card.Blocks = append(card.Blocks, block)
	}
	return card, nil
}

func presetToBlock(player replay.Player, preset common.Tag) (common.StatsBlock[BlockData], error) {
	block := common.StatsBlock[BlockData](common.NewBlock(preset, BlockData{}))

	switch preset {
	case TagDamageBlocked:
		block.Value = player.Performance.DamageBlocked
		return block, nil
	case TagDamageAssisted:
		block.Value = player.Performance.DamageAssisted
		return block, nil
	case TagDamageAssistedCombined:
		block.Value = player.Performance.DamageAssisted + player.Performance.DamageBlocked
		return block, nil
	}

	err := block.FillValue(player.Performance.StatsFrame)
	return block, err
}

func sortTeams(teams replay.Teams) {
	sortPlayers(teams.Allies)
	sortPlayers(teams.Enemies)
}

func sortPlayers(players []replay.Player) {
	slices.SortFunc(players, func(j, i replay.Player) int {
		return int((i.Performance.DamageDealt.Float() + i.Performance.DamageAssisted.Float() + i.Performance.DamageBlocked.Float()) - (j.Performance.DamageDealt.Float() + j.Performance.DamageAssisted.Float() + j.Performance.DamageBlocked.Float()))
	})
}
