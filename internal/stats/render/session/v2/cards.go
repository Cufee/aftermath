package session

import (
	"strconv"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	recom "github.com/cufee/aftermath/internal/stats/render/common"
	"github.com/cufee/facepaint"
	"github.com/cufee/facepaint/style"
	"github.com/nao1215/imaging"
)

func generateCards(sessionData, careerData fetch.AccountStatsOverPeriod, cards session.Cards, _ []models.UserSubscription, opts common.Options) (*facepaint.Block, error) {
	var (
		renderUnratedVehiclesCount = 10 // maximum number of vehicle cards
		// primary cards
		// when there are some unrated battles or no battles at all
		shouldRenderUnratedOverview = sessionData.RegularBattles.Battles > 0 || sessionData.RatingBattles.Battles < 1
		// more highlights than vehicles implies we will render some vehicles, at least 3 in most cases
		shouldRenderUnratedHighlights = sessionData.RegularBattles.Battles > 0 && len(cards.Unrated.Vehicles) > len(cards.Unrated.Highlights)
		// no vehicle stats are available for rating battles
		shouldRenderRatingOverview = sessionData.RatingBattles.Battles > 0
	)

	// calculate max overview block width to make all blocks the same size
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

	// calculate per block type width of vehicle stats to make things even
	var vehicleBlockWidth = make(map[prepare.Tag]float64)
	for _, card := range cards.Unrated.Vehicles {
		for _, block := range card.Blocks {
			vehicleBlockWidth[block.Tag] = max(vehicleBlockWidth[block.Tag],
				facepaint.MeasureBlockWidth(block.Label, *styledVehicleLegendPillText()),
				facepaint.MeasureStringWidth(block.Value().String(), styledVehicleCard.value().Font),
			)
		}
	}

	var overviewCards = []*facepaint.Block{newPlayerNameCard(careerData.Account)}
	// unrated overview
	if shouldRenderUnratedOverview {
		if card := newUnratedOverviewCard(cards.Unrated.Overview, maxWidthOverviewColumn); card != nil {
			overviewCards = append(overviewCards, card)
		}
	}
	// rating battles
	if shouldRenderRatingOverview {
		if card := newRatingOverviewCard(cards.Rating, maxWidthOverviewColumn); card != nil {
			overviewCards = append(overviewCards, card)
		}
	}
	// highlights
	if shouldRenderUnratedHighlights {
		// calculate per block type width of highlight stats to make things even
		var highlightBlockWidth = make(map[prepare.Tag]float64)
		for _, highlight := range cards.Unrated.Highlights {
			for _, block := range highlight.Blocks {
				highlightBlockWidth[block.Tag] = max(highlightBlockWidth[block.Tag],
					facepaint.MeasureStringWidth(block.Label, styledHighlightCard.blockLabel().Font),
					facepaint.MeasureStringWidth(block.Value().String(), styledHighlightCard.blockValue().Font),
				)
			}
		}

		for _, card := range cards.Unrated.Highlights {
			overviewCards = append(overviewCards, newHighlightCard(card, highlightBlockWidth))
		}
	}

	// vehicles
	var vehicleCards []*facepaint.Block
	for i, card := range cards.Unrated.Vehicles {
		if i == renderUnratedVehiclesCount {
			break
		}
		vehicleCards = append(vehicleCards, newVehicleCard(card, vehicleBlockWidth))
	}

	var sectionBlocks []*facepaint.Block
	sectionBlocks = append(sectionBlocks, facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsSection)), overviewCards...))
	if len(vehicleCards) > 0 {
		vehicleCards = append(vehicleCards, newVehicleLegendCard(cards.Unrated.Vehicles[0], vehicleBlockWidth))
		sectionBlocks = append(sectionBlocks, facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsSection)), vehicleCards...))
	}
	statsCardsBlock := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledCardsSectionsWrapper)), sectionBlocks...)

	cardsFrame := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledStatsFrame)), statsCardsBlock)

	// resize and place background
	if opts.Background != nil {
		cardsFrameSize := cardsFrame.Dimensions()
		opts.Background = imaging.Fill(opts.Background, cardsFrameSize.Width, cardsFrameSize.Height, imaging.Center, imaging.Lanczos)
		if !opts.BackgroundIsCustom {
			seed, _ := strconv.Atoi(careerData.Account.ID)
			opts.Background = recom.AddWN8BackgroundBranding(opts.Background, sessionData.RegularBattles.Vehicles, seed)
		}
		cardsFrame = facepaint.NewBlocksContent(style.NewStyle(),
			facepaint.MustNewImageContent(styledCardsBackground, opts.Background), cardsFrame,
		)
	}

	var frameCards []*facepaint.Block
	frameCards = append(frameCards, cardsFrame)
	frameCards = append(frameCards, newFooterCard(sessionData, cards, opts))

	return facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledFinalFrame)), frameCards...), nil
}
