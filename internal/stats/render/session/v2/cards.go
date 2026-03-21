package session

import (
	"strconv"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
	"github.com/nao1215/imaging"
)

func generateCards(sessionData, careerData fetch.AccountStatsOverPeriod, cards session.Cards, _ []models.UserSubscription, opts common.Options) (*facepaint.Block, error) {
	theme := opts.Theme
	hlStyle := common.NewHighlightCardStyle(theme)
	styledOverviewCard := newOverviewCardStyle(theme)
	vStyle := newVehicleCardStyle(theme)
	legendPillText := newVehicleLegendPillText(theme)

	var (
		renderUnratedVehiclesCount    = 8
		shouldRenderUnratedOverview   = sessionData.RegularBattles.Battles > 0 || sessionData.RatingBattles.Battles < 1
		shouldRenderUnratedHighlights = sessionData.RegularBattles.Battles > 0 && len(cards.Unrated.Vehicles) > len(cards.Unrated.Highlights)
		shouldRenderRatingOverview    = sessionData.RatingBattles.Battles > 0
	)

	var maxWidthOverviewColumn = make(map[bool]float64)
	for _, column := range append(cards.Unrated.Overview.Blocks, cards.Rating.Overview.Blocks...) {
		for _, block := range column.Blocks {
			key := column.Flavor == session.BlockFlavorDefault
			blockStyle := styledOverviewCard.styleBlock(block)
			switch block.Tag {
			case prepare.TagWN8:
				block.Label = common.GetWN8TierName(block.Value().Float())
				maxWidthOverviewColumn[key] = max(maxWidthOverviewColumn[key], iconSizeWN8)
			case prepare.TagRankedRating:
				block.Label = common.GetRatingTierName(block.Value().Float())
				maxWidthOverviewColumn[key] = max(maxWidthOverviewColumn[key], iconSizeRating)
			}
			maxWidthOverviewColumn[key] = max(maxWidthOverviewColumn[key],
				facepaint.MeasureStringWidth(block.Label, blockStyle.label.Font),
				facepaint.MeasureStringWidth(block.Value().String(), blockStyle.value.Font),
			)
		}
	}

	var vehicleBlockWidth = make(map[prepare.Tag]float64)
	for _, card := range cards.Unrated.Vehicles {
		for _, block := range card.Blocks {
			vehicleBlockWidth[block.Tag] = max(vehicleBlockWidth[block.Tag],
				facepaint.MeasureBlockWidth(block.Label, *legendPillText),
				facepaint.MeasureStringWidth(block.Value().String(), vStyle.value().Font),
			)
		}
	}

	var overviewCards = []*facepaint.Block{common.NewPlayerNameBlock(careerData.Account, theme)}
	if shouldRenderUnratedOverview {
		if card := newUnratedOverviewCard(styledOverviewCard, cards.Unrated.Overview, maxWidthOverviewColumn); card != nil {
			overviewCards = append(overviewCards, card)
		}
	}
	if shouldRenderRatingOverview {
		if card := newRatingOverviewCard(styledOverviewCard, cards.Rating, maxWidthOverviewColumn); card != nil {
			overviewCards = append(overviewCards, card)
		}
	}
	if shouldRenderUnratedHighlights {
		var highlightBlockWidth = make(map[prepare.Tag]float64)
		for _, highlight := range cards.Unrated.Highlights {
			for _, block := range highlight.Blocks {
				highlightBlockWidth[block.Tag] = max(highlightBlockWidth[block.Tag],
					facepaint.MeasureStringWidth(block.Label, hlStyle.BlockLabel().Font),
					facepaint.MeasureStringWidth(block.Value().String(), hlStyle.BlockValue().Font),
				)
			}
		}

		for _, card := range cards.Unrated.Highlights {
			overviewCards = append(overviewCards, newHighlightCard(hlStyle, card, highlightBlockWidth))
		}
	}

	var vehicleCards []*facepaint.Block
	for i, card := range cards.Unrated.Vehicles {
		if i == renderUnratedVehiclesCount {
			break
		}
		vehicleCards = append(vehicleCards, newVehicleCard(vStyle, card, vehicleBlockWidth))
	}

	var sectionBlocks []*facepaint.Block
	sectionBlocks = append(sectionBlocks, facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsSection)), overviewCards...))
	if len(vehicleCards) > 0 {
		vehicleCards = append(vehicleCards, newVehicleLegendCard(vStyle, legendPillText, cards.Unrated.Vehicles[0], vehicleBlockWidth))
		sectionBlocks = append(sectionBlocks, facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsSection)), vehicleCards...))
	}
	statsCardsBlock := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsSectionsWrapper)), sectionBlocks...)

	cardsFrame := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledStatsFrame)), statsCardsBlock)

	if opts.Background != nil {
		bgDims := cardsFrame.Dimensions()
		seed, _ := strconv.Atoi(careerData.Account.ID)
		opts.Background = imaging.Fill(opts.Background, bgDims.Width, bgDims.Height, imaging.Center, imaging.Lanczos)
		if !opts.BackgroundIsCustom {
			opts.Background = common.AddWN8BackgroundBranding(opts.Background, sessionData.RegularBattles.Vehicles, seed)
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
	frameCards = append(frameCards, common.NewFooterBlock(sessionData, opts))

	frameStyle := common.FinalFrameStyle(theme)
	return facepaint.NewBlocksContent(style.NewStyle(style.Parent(frameStyle)), frameCards...), nil
}
