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
	var maxWidthOverviewBlock = make(map[string]float64)
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

	var statsCards []*facepaint.Block

	statsCards = append(statsCards, newPlayerNameCard(stats.Account))

	if card := newUnratedOverviewCard(cards.Overview, maxWidthOverviewBlock); card != nil {
		statsCards = append(statsCards, card)
	}
	if card := newRatingOverviewCard(cards.Rating, maxWidthOverviewBlock); card != nil {
		statsCards = append(statsCards, card)
	}
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
