package replay

import (
	"fmt"

	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

func generateCards(replay fetch.Replay, cards replay.Cards, opts common.Options) (common.Segments, error) {
	var segments common.Segments

	var alliesBlocks, enemiesBlocks []common.Block

	var playerNameWidth float64
	statsSizes := make(map[prepare.Tag]float64)
	for _, card := range append(cards.Allies, cards.Enemies...) {
		// Measure player name and tag or vehicle name
		name := card.Meta.Player.Nickname
		if card.Meta.Player.ClanTag != "" {
			name += fmt.Sprintf(" [%s]", card.Meta.Player.ClanTag)
		}
		nameSize := common.MeasureString(name, common.FontLarge())
		tankSize := common.MeasureString(card.Title, common.FontLarge())
		size := nameSize
		if tankSize.TotalWidth > nameSize.TotalWidth {
			size = tankSize
		}
		if size.TotalWidth > playerNameWidth {
			playerNameWidth = size.TotalWidth
		}

		// Measure stats value and label
		for _, block := range card.Blocks {
			valueSize := common.MeasureString(block.Value.String(), common.FontLarge())
			labelSize := common.MeasureString(block.Label, common.FontSmall())
			w := valueSize.TotalWidth
			if labelSize.TotalWidth > valueSize.TotalWidth {
				w = labelSize.TotalWidth
			}
			if w > statsSizes[block.Tag] {
				statsSizes[block.Tag] = w
			}
		}
	}

	var totalStatsWidth float64
	for _, width := range statsSizes {
		totalStatsWidth += width
	}

	playerStatsCardStyle := defaultCardStyle(playerNameWidth+(float64(len(statsSizes)*10))+totalStatsWidth, 0)
	totalCardsWidth := (playerStatsCardStyle.Width * 2) - 30

	// Allies
	for _, card := range cards.Allies {
		alliesBlocks = append(alliesBlocks, newPlayerCard(playerStatsCardStyle, statsSizes, card, card.Meta.Player, true, card.Meta.Player.ID == replay.Protagonist.ID))
	}
	// Enemies
	for _, card := range cards.Enemies {
		enemiesBlocks = append(enemiesBlocks, newPlayerCard(playerStatsCardStyle, statsSizes, card, card.Meta.Player, false, false))
	}

	// Title Card
	titleBlock := newTitleBlock(cards, totalCardsWidth)

	// Teams
	var teamsBlocks []common.Block
	teamsBlocks = append(teamsBlocks, common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, Gap: 10}, alliesBlocks...))
	teamsBlocks = append(teamsBlocks, common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, Gap: 10}, enemiesBlocks...))
	playersBlock := common.NewBlocksContent(common.Style{Direction: common.DirectionHorizontal, Gap: 10}, teamsBlocks...)
	teamsBlock := common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, Gap: 10}, playersBlock)

	segments.AddContent(common.NewBlocksContent(frameStyle, titleBlock, teamsBlock))
	return segments, nil
}
