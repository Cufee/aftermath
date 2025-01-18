package period

import (
	"errors"

	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/facepaint/style"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint"
)

func generateCards(stats fetch.AccountStatsOverPeriod, cards period.Cards, subs []models.UserSubscription, opts common.Options) (*facepaint.Block, error) {
	if len(cards.Overview.Blocks) == 0 && len(cards.Highlights) == 0 {
		log.Error().Msg("player cards slice is 0 length, this should not happen")
		return nil, errors.New("no cards provided")
	}

	// calculate max overview block width to make all blocks the same size
	var maxWidthOverviewBlock float64
	for _, column := range append(cards.Overview.Blocks, cards.Rating.Blocks...) {
		for _, block := range column.Blocks {
			switch block.Tag {
			case prepare.TagWN8:
				block.Label = common.GetWN8TierName(block.Value().Float())
				maxWidthOverviewBlock = max(maxWidthOverviewBlock, iconSizeWN8)

			case prepare.TagRankedRating:
				block.Label = common.GetRatingTierName(block.Value().Float())
				maxWidthOverviewBlock = max(maxWidthOverviewBlock, iconSizeRating)
			}
			maxWidthOverviewBlock = max(maxWidthOverviewBlock, facepaint.MeasureString(block.Label, styledOverviewCard.styleBlock(block).label.Font).TotalWidth)
			maxWidthOverviewBlock = max(maxWidthOverviewBlock, facepaint.MeasureString(block.Value().String(), styledOverviewCard.styleBlock(block).value.Font).TotalWidth)
		}
	}

	// 	{
	// 		highlightStyle := highlightCardStyle(defaultCardStyle(0))
	// 		var highlightBlocksMaxCount, highlightTitleMaxWidth, highlightBlockMaxSize float64
	// 		for _, highlight := range cards.Highlights {
	// 			// Title and tank name
	// 			metaSize := common.MeasureString(highlight.Meta, highlightStyle.cardTitle.Font)
	// 			titleSize := common.MeasureString(highlight.Title, highlightStyle.tankName.Font)
	// 			highlightTitleMaxWidth = max(highlightTitleMaxWidth, metaSize.TotalWidth, titleSize.TotalWidth)

	// 			// Blocks
	// 			highlightBlocksMaxCount = max(highlightBlocksMaxCount, float64(len(highlight.Blocks)))
	// 			for _, block := range highlight.Blocks {
	// 				labelSize := common.MeasureString(block.Label, highlightStyle.blockLabel.Font)
	// 				valueSize := common.MeasureString(block.Value().String(), highlightStyle.blockValue.Font)
	// 				highlightBlockMaxSize = max(highlightBlockMaxSize, valueSize.TotalWidth, labelSize.TotalWidth)
	// 			}
	// 		}

	// 		highlightCardWidthMax := (highlightStyle.container.PaddingX * 2) + (highlightStyle.container.Gap * highlightBlocksMaxCount) + (highlightBlockMaxSize * highlightBlocksMaxCount) + highlightTitleMaxWidth
	// 		cardWidth = max(cardWidth, highlightCardWidthMax)
	// 	}
	// }

	var finalCards []*facepaint.Block

	// // We first render a footer in order to calculate the minimum required width
	// {
	// 	var footer []string
	// 	if opts.VehicleID != "" {
	// 		footer = append(footer, cards.Overview.Title)
	// 	}

	// 	sessionTo := stats.PeriodEnd.Format("Jan 2, 2006")
	// 	sessionFrom := stats.PeriodStart.Format("Jan 2, 2006")
	// 	if sessionFrom == sessionTo {
	// 		footer = append(footer, sessionTo)
	// 	} else {
	// 		footer = append(footer, sessionFrom+" - "+sessionTo)
	// 	}
	// 	footerBlock := common.NewFooterCard(strings.Join(footer, " • "))
	// 	footerImage, err := footerBlock.Render()
	// 	if err != nil {
	// 		return segments, err
	// 	}

	// 	cardWidth = max(cardWidth, float64(footerImage.Bounds().Dx()))
	// 	segments.AddFooter(common.NewImageContent(common.Style{}, footerImage))
	// }

	// // Header card
	// if headerCard, headerCardExists := common.NewHeaderCard(cardWidth, subs, opts.PromoText); headerCardExists {
	// 	segments.AddHeader(headerCard)
	// }

	// // Player Title card
	// segments.AddContent(common.NewPlayerTitleCard(common.DefaultPlayerTitleStyle(stats.Account.Nickname, titleCardStyle(cardWidth)), stats.Account.Nickname, stats.Account.ClanTag, subs))

	// unrated overview card
	if card := newOverviewCard(cards.Overview, maxWidthOverviewBlock); card != nil {
		finalCards = append(finalCards, card)
	}

	// // Rating Card -- only when player has current season rating
	// if cards.Rating.Meta {
	// 	var ratingStatsBlocks []common.Block
	// 	for _, column := range cards.Rating.Blocks {
	// 		columnBlock, err := statsBlocksToColumnBlock(getOverviewStyle(overviewColumnWidth), column.Blocks)
	// 		if err != nil {
	// 			return segments, err
	// 		}
	// 		ratingStatsBlocks = append(ratingStatsBlocks, columnBlock)
	// 	}
	// 	var ratingCardBlocks []common.Block
	// 	ratingCardBlocks = append(ratingCardBlocks, common.NewBlocksContent(overviewCardBlocksStyle(cardWidth), ratingStatsBlocks...))
	// 	segments.AddContent(common.NewBlocksContent(overviewCardStyle(), ratingCardBlocks...))
	// }

	// // Highlights
	// for i, card := range cards.Highlights {
	// 	if i > 0 && cards.Rating.Meta {
	// 		break // only show 1 highlight when rating battles card is visible
	// 	}
	// 	segments.AddContent(newHighlightCard(highlightCardStyle(defaultCardStyle(cardWidth)), card))
	// }

	if len(finalCards) == 0 {
		return nil, errors.New("no cards to render")
	}

	return facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsFrame)), finalCards...), nil

}
