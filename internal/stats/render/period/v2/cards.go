package period

import (
	"errors"
	"strconv"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
	"github.com/nao1215/imaging"
)

func generateCards(stats fetch.AccountStatsOverPeriod, cards period.Cards, _ []models.UserSubscription, opts common.Options) (*facepaint.Block, error) {
	if len(cards.Overview.Blocks) == 0 && len(cards.Highlights) == 0 {
		log.Error().Msg("player cards slice is 0 length, this should not happen")
		return nil, errors.New("no cards provided")
	}

	theme := opts.Theme
	hlStyle := common.NewHighlightCardStyle(theme)
	styledUnratedOverviewCard := newUnratedOverviewCardStyle(theme)
	styledRatingOverviewCard := newRatingOverviewCardStyle(theme)

	var (
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

	var maxWidthOverviewBlock = make(map[string]float64)

	if shouldRenderUnratedOverview {
		for _, column := range cards.Overview.Blocks {
			for _, block := range column.Blocks {
				key := string(block.Data.Flavor)
				blockStyle := styledUnratedOverviewCard.styleBlock(block)
				switch block.Tag {
				case prepare.TagWN8:
					block.Label = common.GetWN8TierName(block.Value().Float())
					maxWidthOverviewBlock[key] = max(maxWidthOverviewBlock[key], iconSizeWN8)
				}
				maxWidthOverviewBlock[key] = max(maxWidthOverviewBlock[key],
					facepaint.MeasureStringWidth(block.Label, blockStyle.label.Font),
					facepaint.MeasureStringWidth(block.Value().String(), blockStyle.value.Font),
				)
			}
		}
	}

	if shouldRenderRatingOverview {
		for _, column := range cards.Rating.Blocks {
			for _, block := range column.Blocks {
				key := string(block.Data.Flavor)
				blockStyle := styledRatingOverviewCard.styleBlock(block)
				switch block.Tag {
				case prepare.TagRankedRating:
					block.Label = common.GetRatingTierName(block.Value().Float())
					maxWidthOverviewBlock[key] = max(maxWidthOverviewBlock[key], iconSizeRating)
				}
				maxWidthOverviewBlock[key] = max(maxWidthOverviewBlock[key],
					facepaint.MeasureStringWidth(block.Label, blockStyle.label.Font),
					facepaint.MeasureStringWidth(block.Value().String(), blockStyle.value.Font),
				)
			}
		}
	}

	var highlightBlockWidth = make(map[prepare.Tag]float64)
	for i, highlight := range cards.Highlights {
		if i >= highlightCardsCount {
			break
		}

		for _, block := range highlight.Blocks {
			highlightBlockWidth[block.Tag] = max(highlightBlockWidth[block.Tag],
				facepaint.MeasureStringWidth(block.Label, hlStyle.BlockLabel().Font),
				facepaint.MeasureStringWidth(block.Value().String(), hlStyle.BlockValue().Font),
			)
		}
	}
	var statsCards []*facepaint.Block

	statsCards = append(statsCards, common.NewPlayerNameBlock(stats.Account, theme))

	if shouldRenderUnratedOverview {
		if card := newUnratedOverviewCard(styledUnratedOverviewCard, cards.Overview, maxWidthOverviewBlock); card != nil {
			statsCards = append(statsCards, card)
		}
	}

	if shouldRenderRatingOverview {
		if card := newRatingOverviewCard(styledRatingOverviewCard, cards.Rating, maxWidthOverviewBlock); card != nil {
			statsCards = append(statsCards, card)
		}
	}

	for i, card := range cards.Highlights {
		if i >= highlightCardsCount {
			break
		}

		statsCards = append(statsCards, newHighlightCard(hlStyle, card, highlightBlockWidth))
	}

	if len(statsCards) == 0 {
		return nil, errors.New("no cards to render")
	}

	footer := common.NewFooterBlock(stats, opts)
	cardsFrame := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsFrame)), statsCards...)

	if opts.Background != nil {
		bgDims := cardsFrame.Dimensions()
		seed, _ := strconv.Atoi(stats.Account.ID)
		opts.Background = imaging.Fill(opts.Background, bgDims.Width, bgDims.Height, imaging.Center, imaging.Lanczos)
		if !opts.BackgroundIsCustom {
			opts.Background = common.AddWN8BackgroundBranding(opts.Background, stats.RegularBattles.Vehicles, seed)
		}

		var layers []*facepaint.Block
		layers = append(layers, facepaint.MustNewImageContent(common.CardsBackgroundStyle, opts.Background))
		if theme.BackgroundOverlay != nil {
			if overlay := theme.BackgroundOverlay(opts.Background.Bounds(), seed); overlay != nil {
				layers = append(layers, facepaint.MustNewImageContent(common.CardsBackgroundStyle, overlay))
			}
		}
		layers = append(layers, cardsFrame)
		cardsFrame = facepaint.NewBlocksContent(style.NewStyle(), layers...)
	}

	var frameCards []*facepaint.Block
	frameCards = append(frameCards, cardsFrame)
	frameCards = append(frameCards, footer)

	frameStyle := common.FinalFrameStyle(theme)
	return facepaint.NewBlocksContent(style.NewStyle(style.Parent(frameStyle)), frameCards...), nil
}
