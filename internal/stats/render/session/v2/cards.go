package session

import (
	"strconv"

	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/facepaint/style"
	"github.com/nao1215/imaging"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/facepaint"
)

func generateCards(sessionData, careerData fetch.AccountStatsOverPeriod, cards session.Cards, _ []models.UserSubscription, opts common.Options) (*facepaint.Block, error) {
	var (
		renderUnratedVehiclesCount = 3 // minimum number of vehicle cards
		// primary cards
		// when there are some unrated battles or no battles at all
		shouldRenderUnratedOverview = sessionData.RegularBattles.Battles > 0 || sessionData.RatingBattles.Battles < 1
		// when there are 3 vehicle cards and no rating overview cards or there are 6 vehicle cards and some rating battles
		shouldRenderUnratedHighlights = (sessionData.RegularBattles.Battles > 0 && sessionData.RatingBattles.Battles < 1 && len(cards.Unrated.Vehicles) > renderUnratedVehiclesCount) ||
			(sessionData.RegularBattles.Battles > 0 && len(cards.Unrated.Vehicles) > 3)
		shouldRenderRatingOverview = sessionData.RatingBattles.Battles > 0
		// secondary cards
		shouldRenderUnratedVehicles = sessionData.RegularBattles.Battles > 0 && len(cards.Unrated.Vehicles) > 0
	)

	// try to make the columns height roughly similar to primary column
	if shouldRenderUnratedHighlights {
		renderUnratedVehiclesCount += len(cards.Unrated.Highlights)
	}
	if shouldRenderRatingOverview {
		renderUnratedVehiclesCount += 1
	}

	// calculate max overview block width to make all blocks the same size
	var maxWidthOverviewColumn = make(map[bool]float64)
	for _, column := range cards.Unrated.Overview.Blocks {
		for _, block := range column.Blocks {
			switch block.Tag {
			case prepare.TagWN8:
				block.Label = common.GetWN8TierName(block.Value().Float())
				maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault] = max(maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault], iconSizeWN8)
			}
			maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault] = max(maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault], facepaint.MeasureString(block.Label, styledOverviewCard.styleBlock(block).label.Font).TotalWidth)
			maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault] = max(maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault], facepaint.MeasureString(block.Value().String(), styledOverviewCard.styleBlock(block).value.Font).TotalWidth)
		}
	}
	for _, column := range cards.Rating.Overview.Blocks {
		for _, block := range column.Blocks {
			switch block.Tag {
			case prepare.TagRankedRating:
				block.Label = common.GetRatingTierName(block.Value().Float())
				maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault] = max(maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault], iconSizeRating)
			}
			maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault] = max(maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault], facepaint.MeasureString(block.Label, styledOverviewCard.styleBlock(block).label.Font).TotalWidth)
			maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault] = max(maxWidthOverviewColumn[column.Flavor == session.BlockFlavorDefault], facepaint.MeasureString(block.Value().String(), styledOverviewCard.styleBlock(block).value.Font).TotalWidth)
		}
	}

	// calculate per block type width of highlight stats to make things even
	var highlightBlockWidth = make(map[prepare.Tag]float64)
	for _, highlight := range cards.Unrated.Highlights {
		for _, block := range highlight.Blocks {
			label := facepaint.MeasureString(block.Label, styledHighlightCard.blockLabel().Font).TotalWidth
			value := facepaint.MeasureString(block.Value().String(), styledHighlightCard.blockValue().Font).TotalWidth
			highlightBlockWidth[block.Tag] = max(highlightBlockWidth[block.Tag], label, value)
		}
	}

	// calculate per block type width of vehicle stats to make things even
	var vehicleBlockWidth = make(map[prepare.Tag]float64)
	for _, card := range cards.Unrated.Vehicles {
		for _, block := range card.Blocks {
			labelStyle := styledVehicleLegendPillText()
			label := facepaint.MeasureString(block.Label, labelStyle.Font).TotalWidth + labelStyle.PaddingLeft + labelStyle.PaddingRight
			value := facepaint.MeasureString(block.Value().String(), styledVehicleCard.value().Font).TotalWidth
			vehicleBlockWidth[block.Tag] = max(vehicleBlockWidth[block.Tag], label, value)
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
		for _, card := range cards.Unrated.Highlights {
			overviewCards = append(overviewCards, newHighlightCard(card, highlightBlockWidth))
		}
	}

	// vehicles
	var vehicleCards []*facepaint.Block
	if shouldRenderUnratedVehicles {
		for i, card := range cards.Unrated.Vehicles {
			if i == renderUnratedVehiclesCount {
				break
			}
			vehicleCards = append(vehicleCards, newVehicleCard(card, vehicleBlockWidth))
		}
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
			opts.Background = addBackgroundBranding(opts.Background, sessionData.RegularBattles.Vehicles, seed)
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
