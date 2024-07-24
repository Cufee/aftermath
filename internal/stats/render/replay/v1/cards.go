package replay

import (
	"fmt"
	"strings"

	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	rp "github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

func generateCards(replay fetch.Replay, cards rp.Cards, printer func(string) string) (common.Segments, error) {
	var segments common.Segments

	var alliesBlocks, enemiesBlocks []common.Block

	statsStyle := statsRowStyle()
	cardStyle := defaultCardStyle(0, 0)

	var footerHintTag string

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
			w := max(valueSize.TotalWidth, labelSize.TotalWidth)
			if block.Tag == prepare.TagWN8 {
				statsSizes["wn8_icon"] = playerWN8IconSize + statsStyle.Gap/2
			}
			statsSizes[block.Tag] = max(statsSizes[block.Tag], w)

			if block.Tag == rp.TagDamageAssistedCombined {
				footerHintTag = fmt.Sprintf("%s + %s", printer("label_"+rp.TagDamageAssisted.String()), printer("label_"+rp.TagDamageBlocked.String()))
			}
		}
	}

	var headerCardWidth float64
	{
		headerSize := common.MeasureString(fmt.Sprintf("%s - %s", cards.Header.MapName, cards.Header.GameMode), common.FontLarge())
		headerCardWidth = headerSize.TotalHeight + 2*outcomeIconSize + playerCardPadding*2 + cardStyle.Gap*2
	}

	var footerWidth float64
	// We first render a footer in order to calculate the minimum required width
	{
		var footer []string

		switch strings.ToLower(replay.Realm) {
		case "na":
			footer = append(footer, "North America")
		case "eu":
			footer = append(footer, "Europe")
		case "as":
			footer = append(footer, "Asia")
		}
		footer = append(footer, replay.BattleTime.Format("Jan 2, 2006"))
		if footerHintTag != "" {
			footer = append(footer, footerHintTag)
		}

		footerBlock := common.NewFooterCard(strings.Join(footer, " • "))
		footerImage, err := footerBlock.Render()
		if err != nil {
			return segments, err
		}

		footerWidth = float64(footerImage.Bounds().Dx())
		segments.AddFooter(common.NewImageContent(common.Style{}, footerImage))
	}

	var totalStatsWidth float64 = statsStyle.Gap*float64(len(statsSizes)-1) + statsStyle.PaddingX*2
	for _, width := range statsSizes {
		totalStatsWidth += width
	}

	playerStatsCardStyle := defaultCardStyle(playerNameWidth+totalStatsWidth+hpBarWidth+cardStyle.Gap*2+cardStyle.PaddingX*2, 0)
	var totalCardsWidth = max(footerWidth, headerCardWidth, playerStatsCardStyle.Width)
	if len(cards.Allies) > 0 && len(cards.Enemies) > 0 {
		totalCardsWidth = max(totalCardsWidth, (playerStatsCardStyle.Width*2)+frameStyle.Gap)
	}

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
	if len(cards.Allies) > 0 {
		teamsBlocks = append(teamsBlocks, common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, Gap: 10}, alliesBlocks...))
	}
	if len(cards.Enemies) > 0 {
		teamsBlocks = append(teamsBlocks, common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, Gap: 10}, enemiesBlocks...))
	}
	playersBlock := common.NewBlocksContent(teamsRowStyle(), teamsBlocks...)
	teamsBlock := common.NewBlocksContent(teamsRowStyle(), playersBlock)

	segments.AddContent(common.NewBlocksContent(frameStyle, titleBlock, teamsBlock))
	return segments, nil
}
