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

func generateCards(stats fetch.AccountStatsOverPeriod, cards period.Cards, _ []models.UserSubscription, opts common.Options) (*facepaint.Block, error) {
	if len(cards.Overview.Blocks) == 0 && len(cards.Highlights) == 0 {
		log.Error().Msg("player cards slice is 0 length, this should not happen")
		return nil, errors.New("no cards provided")
	}

	var (
		// renderUnratedVehiclesCount = 3 // minimum number of vehicle cards
		shouldRenderUnratedOverview = stats.RegularBattles.Battles > 0 || stats.RatingBattles.Battles < 1
		shouldRenderRatingOverview  = cards.Rating.Meta && stats.RatingBattles.Battles > 0 && opts.VehicleIDs == nil
		highlightCardsCount         = 3
	)
	if shouldRenderRatingOverview {
		highlightCardsCount = 1
	}
	if len(opts.VehicleIDs) == 1 {
		highlightCardsCount = 0
	}

	// calculate max overview block width to make all blocks the same size
	var maxWidthOverviewBlock = make(map[string]float64)

	if shouldRenderUnratedOverview {
		for _, column := range cards.Overview.Blocks {
			for _, block := range column.Blocks {
				switch block.Tag {
				case prepare.TagWN8:
					block.Label = common.GetWN8TierName(block.Value().Float())
					maxWidthOverviewBlock[string(block.Data.Flavor)] = max(maxWidthOverviewBlock[string(block.Data.Flavor)], iconSizeWN8)
				}
				maxWidthOverviewBlock[string(block.Data.Flavor)] = max(maxWidthOverviewBlock[string(block.Data.Flavor)], facepaint.MeasureString(block.Label, styledUnratedOverviewCard.styleBlock(block).label.Font).TotalWidth)
				maxWidthOverviewBlock[string(block.Data.Flavor)] = max(maxWidthOverviewBlock[string(block.Data.Flavor)], facepaint.MeasureString(block.Value().String(), styledUnratedOverviewCard.styleBlock(block).value.Font).TotalWidth)
			}
		}
	}

	if shouldRenderRatingOverview {
		for _, column := range cards.Rating.Blocks {
			for _, block := range column.Blocks {
				switch block.Tag {
				case prepare.TagRankedRating:
					block.Label = common.GetRatingTierName(block.Value().Float())
					maxWidthOverviewBlock[string(block.Data.Flavor)] = max(maxWidthOverviewBlock[string(block.Data.Flavor)], iconSizeRating)
				}
				maxWidthOverviewBlock[string(block.Data.Flavor)] = max(maxWidthOverviewBlock[string(block.Data.Flavor)], facepaint.MeasureString(block.Label, styledRatingOverviewCard.styleBlock(block).label.Font).TotalWidth)
				maxWidthOverviewBlock[string(block.Data.Flavor)] = max(maxWidthOverviewBlock[string(block.Data.Flavor)], facepaint.MeasureString(block.Value().String(), styledRatingOverviewCard.styleBlock(block).value.Font).TotalWidth)
			}
		}
	}

	// calculate per block type width of highlight stats to make things even
	var highlightBlockWidth = make(map[prepare.Tag]float64)
	for i, highlight := range cards.Highlights {
		if i >= highlightCardsCount {
			break
		}

		for _, block := range highlight.Blocks {
			label := facepaint.MeasureString(block.Label, styledHighlightCard.blockLabel().Font).TotalWidth
			value := facepaint.MeasureString(block.Value().String(), styledHighlightCard.blockValue().Font).TotalWidth
			highlightBlockWidth[block.Tag] = max(highlightBlockWidth[block.Tag], label, value)
		}
	}
	var statsCards []*facepaint.Block

	// player name card
	statsCards = append(statsCards, newPlayerNameCard(stats.Account))

	if shouldRenderUnratedOverview {
		if card := newUnratedOverviewCard(cards.Overview, maxWidthOverviewBlock); card != nil {
			statsCards = append(statsCards, card)
		}
	}

	// rating battles
	if shouldRenderRatingOverview {
		if card := newRatingOverviewCard(cards.Rating, maxWidthOverviewBlock); card != nil {
			statsCards = append(statsCards, card)
		}
	}

	// highlights
	for i, card := range cards.Highlights {
		if i >= highlightCardsCount {
			break
		}

		statsCards = append(statsCards, newHighlightCard(card, highlightBlockWidth))
	}

	if len(statsCards) == 0 {
		return nil, errors.New("no cards to render")
	}

	footer := newFooterCard(stats, cards, opts, cards.Overview.Meta)
	cardsFrame := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsFrame)), statsCards...)

	// resize and place background
	if opts.Background != nil {
		cardsFrameSize := cardsFrame.Dimensions()
		opts.Background = imaging.Fill(opts.Background, cardsFrameSize.Width, cardsFrameSize.Height, imaging.Center, imaging.Lanczos)
		if !opts.BackgroundIsCustom {
			seed, _ := strconv.Atoi(stats.Account.ID)
			opts.Background = addBackgroundBranding(opts.Background, stats.RegularBattles.Vehicles, seed)
		}
		cardsFrame = facepaint.NewBlocksContent(style.NewStyle(),
			facepaint.MustNewImageContent(styledCardsBackground, opts.Background), cardsFrame,
		)
	}

	var frameCards []*facepaint.Block
	frameCards = append(frameCards, cardsFrame)
	frameCards = append(frameCards, footer)

	return facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledFinalFrame)), frameCards...), nil

}
