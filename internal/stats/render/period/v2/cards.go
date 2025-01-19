package period

import (
	"errors"
	"strconv"

	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/facepaint/style"
	"github.com/nao1215/imaging"

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
	for _, column := range cards.Overview.Blocks {
		for _, block := range column.Blocks {
			switch block.Tag {
			case prepare.TagWN8:
				block.Label = common.GetWN8TierName(block.Value().Float())
				maxWidthOverviewBlock = max(maxWidthOverviewBlock, iconSizeWN8)
			}
			maxWidthOverviewBlock = max(maxWidthOverviewBlock, facepaint.MeasureString(block.Label, styledUnratedOverviewCard.styleBlock(block).label.Font).TotalWidth)
			maxWidthOverviewBlock = max(maxWidthOverviewBlock, facepaint.MeasureString(block.Value().String(), styledUnratedOverviewCard.styleBlock(block).value.Font).TotalWidth)
		}
	}
	for _, column := range cards.Rating.Blocks {
		for _, block := range column.Blocks {
			switch block.Tag {
			case prepare.TagRankedRating:
				block.Label = common.GetRatingTierName(block.Value().Float())
				maxWidthOverviewBlock = max(maxWidthOverviewBlock, iconSizeRating)
			}
			maxWidthOverviewBlock = max(maxWidthOverviewBlock, facepaint.MeasureString(block.Label, styledRatingOverviewCard.styleBlock(block).label.Font).TotalWidth)
			maxWidthOverviewBlock = max(maxWidthOverviewBlock, facepaint.MeasureString(block.Value().String(), styledRatingOverviewCard.styleBlock(block).value.Font).TotalWidth)
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

	var statsCards []*facepaint.Block

	statsCards = append(statsCards, newPlayerNameCard(stats.Account))

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

	if card := newUnratedOverviewCard(cards.Overview, maxWidthOverviewBlock); card != nil {
		statsCards = append(statsCards, card)
	}
	if card := newRatingOverviewCard(cards.Rating, maxWidthOverviewBlock); card != nil {
		statsCards = append(statsCards, card)
	}

	// // Highlights
	// for i, card := range cards.Highlights {
	// 	if i > 0 && cards.Rating.Meta {
	// 		break // only show 1 highlight when rating battles card is visible
	// 	}
	// 	segments.AddContent(newHighlightCard(highlightCardStyle(defaultCardStyle(cardWidth)), card))
	// }

	if len(statsCards) == 0 {
		return nil, errors.New("no cards to render")
	}

	cardsFrame := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsFrame)), statsCards...)

	// add background branding
	if opts.Background != nil && !opts.BackgroundIsCustom {
		cardsFrameSize := cardsFrame.Dimensions()
		seed, _ := strconv.Atoi(stats.Account.ID)
		opts.Background = imaging.Resize(opts.Background, cardsFrameSize.Width, cardsFrameSize.Height, imaging.Lanczos)
		opts.Background = addBackgroundBranding(opts.Background, stats.RegularBattles.Vehicles, seed)
	}
	// add background
	if opts.Background != nil {
		cardsFrame = facepaint.NewBlocksContent(style.NewStyle(),
			facepaint.MustNewImageContent(styledCardsBackground, opts.Background), cardsFrame,
		)
	}

	var frameCards []*facepaint.Block
	frameCards = append(frameCards, cardsFrame)

	return facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledFinalFrame)), frameCards...), nil

}
