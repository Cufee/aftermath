package replay

import (
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

func generateCards(replay fetch.Replay, cards replay.Cards) (common.Segments, error) {
	var segments common.Segments

	var alliesBlocks, enemiesBlocks []common.Block

	cardStyle := defaultCardStyle(0, 0)

	var playerNameWidth float64
	statsSizes := make(map[prepare.Tag]float64)
	for _, card := range append(cards.Allies, cards.Enemies...) {
		// Measure player name and tag or vehicle name
		nameSize := common.MeasureString(card.Meta.Player.Nickname, common.FontLarge())
		clanSize := common.MeasureString(card.Meta.Player.ClanTag, common.FontMedium())
		tankSize := common.MeasureString(card.Title, common.FontLarge())
		playerNameWidth = max(playerNameWidth, max(nameSize.TotalWidth+clanSize.TotalWidth+cardStyle.Gap, tankSize.TotalWidth))

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

	statsStyle := statsRowStyle()
	var totalStatsWidth float64 = statsStyle.Gap*float64(len(statsSizes)-1) + statsStyle.PaddingX*2
	for _, width := range statsSizes {
		totalStatsWidth += width
	}

	playerStatsCardStyle := defaultCardStyle(playerNameWidth+totalStatsWidth+hpBarWidth+cardStyle.Gap*2+cardStyle.PaddingX*2, 0)
	totalCardsWidth := (playerStatsCardStyle.Width * 2) + frameStyle.Gap

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
	playersBlock := common.NewBlocksContent(teamsRowStyle(), teamsBlocks...)
	teamsBlock := common.NewBlocksContent(teamsRowStyle(), playersBlock)

	segments.AddContent(common.NewBlocksContent(frameStyle, titleBlock, teamsBlock))
	return segments, nil
}