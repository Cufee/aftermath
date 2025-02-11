package replay

import (
	"fmt"
	"image"
	"math"

	"github.com/cufee/aftermath/internal/render/assets"
	common "github.com/cufee/aftermath/internal/render/v1"
	fetch "github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
	"github.com/cufee/aftermath/internal/stats/frame"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
)

func newTitleBlock(replay replay.Cards, width float64) common.Block {
	style := defaultCardStyle(width, 80)
	style.JustifyContent = common.JustifyContentSpaceBetween
	style.Direction = common.DirectionHorizontal
	style.AlignItems = common.AlignItemsCenter
	style.PaddingX = playerCardPadding + hpBarWidth/2

	allyIcon, enemyIcon := outcomeIcons(replay.Header.Outcome)

	return common.NewBlocksContent(style,
		allyIcon,
		common.NewTextContent(common.Style{
			Font:      common.FontLarge(),
			FontColor: common.TextPrimary,
		}, fmt.Sprintf("%s - %s", replay.Header.MapName, replay.Header.GameMode)),
		enemyIcon)
}

func newPlayerCard(style common.Style, sizes map[prepare.Tag]float64, card replay.Card, player fetch.Player, ally, protagonist bool) common.Block {
	hpBarValue := float64(player.HPLeft) / float64((player.Performance.DamageReceived + player.HPLeft))
	if hpBarValue > 0 {
		hpBarValue = math.Max(hpBarValue, 0.2)
	}

	var hpBar common.Block
	if ally {
		hpBar = newProgressBar(int(hpBarHeight), int(hpBarValue*100), progressDirectionVertical, hpBarColorAllies, hpBarBgColorAllies)
	} else {
		hpBar = newProgressBar(int(hpBarHeight), int(hpBarValue*100), progressDirectionVertical, hpBarColorEnemies, hpBarBgColorEnemies)
	}

	vehicleColor := common.TextPrimary
	if player.HPLeft <= 0 {
		vehicleColor = common.TextSecondary
	}

	leftBlock := common.NewBlocksContent(common.Style{
		Direction:  common.DirectionHorizontal,
		AlignItems: common.AlignItemsCenter,
		Gap:        defaultCardStyle(0, 0).Gap,
		Height:     80,
		// Debug:      true,
	},
		hpBar,
		common.NewBlocksContent(common.Style{Direction: common.DirectionVertical},
			common.NewTextContent(common.Style{Font: common.FontLarge(), FontColor: vehicleColor}, card.Title),
			playerNameBlock(player, protagonist),
		))

	var rightBlocks []common.Block
	for _, block := range card.Blocks {
		rightBlocks = append(rightBlocks, statsBlockToBlock(block, sizes[block.Tag]))
		if block.Tag == prepare.TagWN8 {
			rightBlocks = append(rightBlocks, playerWN8Icon(block.Value()))
		}
		if block.Tag == prepare.TagRankedRating {
			rightBlocks = append(rightBlocks, playerRatingIcon(block.Value()))
		}
	}
	rightBlock := common.NewBlocksContent(statsRowStyle(), rightBlocks...)

	style.PaddingX = (80 - hpBarHeight) / 2
	style.Direction = common.DirectionHorizontal
	style.AlignItems = common.AlignItemsCenter
	style.JustifyContent = common.JustifyContentSpaceBetween
	// style.Debug = true

	return common.NewBlocksContent(style, leftBlock, rightBlock)
}

func playerNameBlock(player fetch.Player, protagonist bool) common.Block {
	tagColor := common.TextSecondary
	nameColor := common.TextSecondary
	if protagonist {
		nameColor = protagonistColor
	} else if player.HPLeft <= 0 {
		tagColor = common.TextAlt
		nameColor = common.TextAlt
	}

	var nameBlocks []common.Block
	nameBlocks = append(nameBlocks, common.NewTextContent(common.Style{
		Font:      common.FontLarge(),
		FontColor: nameColor,
		// Debug:     true,
	}, player.Nickname))
	if player.ClanTag != "" {
		nameBlocks = append(nameBlocks, common.NewTextContent(common.Style{
			FontColor: tagColor,
			Font:      common.FontMedium(),
			// Debug:     true,
		}, fmt.Sprintf("[%s]", player.ClanTag)))
	}

	return common.NewBlocksContent(common.Style{Direction: common.DirectionHorizontal, Gap: 5, AlignItems: common.AlignItemsCenter}, nameBlocks...)
}

var outcomeIconCache image.Image

func outcomeIcons(outcome fetch.Outcome) (common.Block, common.Block) {
	if outcomeIconCache == nil {
		flagIcon, _ := assets.GetLoadedImage("flag")
		outcomeIconCache = imaging.Fit(flagIcon, int(outcomeIconSize), int(outcomeIconSize), imaging.Linear)
	}

	allyIconColor := outcomeIconBgColorGreen
	enemyIconColor := outcomeIconBgColorRed
	if outcome == fetch.OutcomeVictory {
		allyIconColor = outcomeIconColorGreen
	}
	if outcome == fetch.OutcomeDefeat {
		enemyIconColor = outcomeIconColorRed
	}
	if outcome == fetch.OutcomeDraw {
		allyIconColor = outcomeIconColorYellow
		enemyIconColor = outcomeIconColorYellow
	}

	return common.NewImageContent(common.Style{BackgroundColor: allyIconColor}, outcomeIconCache), common.NewImageContent(common.Style{BackgroundColor: enemyIconColor}, outcomeIconCache)
}

func playerWN8Icon(value frame.Value) common.Block {
	colors := common.GetWN8Colors(value.Float())
	if frame.InvalidValue.Equals(value) {
		colors.Background = common.TextAlt
	}
	icon := common.AftermathLogo(colors.Background, common.SmallLogoOptions())
	return common.NewImageContent(common.Style{Width: playerWN8IconSize, Height: playerWN8IconSize}, icon)

}

func playerRatingIcon(value frame.Value) common.Block {
	icon, ok := common.GetRatingIcon(value, playerRatingIconSize)
	if !ok {
		return common.NewEmptyContent(1, 1)
	}
	return icon

}

func playerWinrateIndicator(value frame.Value) common.Block {
	color := common.GetWinrateColor(value)
	if value.Float() == frame.InvalidValue.Float() {
		color = common.TextAlt
	}
	ctx := gg.NewContext(int(playerWinrateIndicatorSize), int(playerWinrateIndicatorSize))
	r := playerWinrateIndicatorSize / 2
	ctx.DrawCircle(r, r, r)
	ctx.SetColor(color)
	ctx.Fill()
	return common.NewImageContent(common.Style{}, ctx.Image())

}
