package replay

import (
	"fmt"
	"math"

	fetch "github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

func newTitleBlock(replay replay.Cards, width float64) common.Block {
	var titleBlocks []common.Block
	titleBlocks = append(titleBlocks, common.NewTextContent(common.Style{
		Font:      common.FontLarge(),
		FontColor: common.TextPrimary,
	}, replay.Header.Result))

	titleBlocks = append(titleBlocks, common.NewTextContent(common.Style{
		Font:      common.FontLarge(),
		FontColor: common.TextSecondary,
	}, fmt.Sprintf(" - %s", replay.Header.GameMode)))

	style := defaultCardStyle(width, 75)
	style.JustifyContent = common.JustifyContentCenter
	style.Direction = common.DirectionHorizontal
	style.AlignItems = common.AlignItemsCenter

	return common.NewBlocksContent(style, titleBlocks...)
}

func newPlayerCard(style common.Style, sizes map[prepare.Tag]float64, card replay.Card, player fetch.Player, ally, protagonist bool) common.Block {
	hpBarValue := float64(player.HPLeft) / float64((player.Performance.DamageReceived + player.HPLeft))
	if hpBarValue > 0 {
		hpBarValue = math.Max(hpBarValue, 0.2)
	}

	var hpBar common.Block
	if ally {
		hpBar = newProgressBar(60, int(hpBarValue*100), progressDirectionVertical, hpBarColorAllies)
	} else {
		hpBar = newProgressBar(60, int(hpBarValue*100), progressDirectionVertical, hpBarColorEnemies)
	}

	vehicleColor := common.TextPrimary
	if player.HPLeft <= 0 {
		vehicleColor = common.TextSecondary
	}

	leftBlock := common.NewBlocksContent(common.Style{
		Direction:  common.DirectionHorizontal,
		AlignItems: common.AlignItemsCenter,
		Gap:        10,
		Height:     80,
		// Debug:      true,
	}, hpBar, common.NewBlocksContent(common.Style{Direction: common.DirectionVertical},
		common.NewTextContent(common.Style{Font: common.FontLarge(), FontColor: vehicleColor}, card.Title),
		playerNameBlock(player, protagonist),
	))

	var rightBlocks []common.Block
	for _, block := range card.Blocks {
		rightBlocks = append(rightBlocks, statsBlockToBlock(block, sizes[block.Tag]))
	}
	rightBlock := common.NewBlocksContent(common.Style{
		JustifyContent: common.JustifyContentCenter,
		AlignItems:     common.AlignItemsCenter,
		Gap:            10,
		// Debug: true,
	}, rightBlocks...)

	style.Direction = common.DirectionHorizontal
	style.AlignItems = common.AlignItemsCenter
	style.JustifyContent = common.JustifyContentSpaceBetween
	// style.Debug = true

	return common.NewBlocksContent(style, leftBlock, rightBlock)
}

func playerNameBlock(player fetch.Player, protagonist bool) common.Block {
	nameColor := common.TextSecondary
	if protagonist {
		nameColor = protagonistColor
	}

	var nameBlocks []common.Block
	nameBlocks = append(nameBlocks, common.NewTextContent(common.Style{
		Font:      common.FontLarge(),
		FontColor: nameColor,
		// Debug:     true,
	}, player.Nickname))
	if player.ClanTag != "" {
		nameBlocks = append(nameBlocks, common.NewTextContent(common.Style{
			FontColor: common.TextSecondary,
			Font:      common.FontLarge(),
			// Debug:     true,
		}, fmt.Sprintf("[%s]", player.ClanTag)))
	}
	return common.NewBlocksContent(common.Style{Direction: common.DirectionHorizontal, Gap: 5, AlignItems: common.AlignItemsCenter}, nameBlocks...)
}
