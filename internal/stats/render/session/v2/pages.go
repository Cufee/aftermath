package session

import (
	"fmt"
	"image"
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

const maxCardsPerPage = 7
const maxPages = 10 // Discord allows up to 10 attachments per message

func generateSessionPages(sessionData, careerData fetch.AccountStatsOverPeriod, cards session.Cards, _ []models.UserSubscription, opts common.Options) ([]*facepaint.Block, error) {
	shouldRenderUnratedOverview := sessionData.RegularBattles.Battles > 0 || sessionData.RatingBattles.Battles < 1
	shouldRenderRatingOverview := sessionData.RatingBattles.Battles > 0 && opts.VehicleIDs == nil
	showHighlights := len(cards.Unrated.Vehicles) > len(cards.Unrated.Highlights)

	// Measure shared widths for consistent block sizing
	maxWidthOverviewColumn := make(map[bool]float64)
	for _, column := range append(cards.Rating.Overview.Blocks, cards.Unrated.Overview.Blocks...) {
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

	highlightBlockWidth := make(map[prepare.Tag]float64)
	for _, highlight := range cards.Unrated.Highlights {
		for _, block := range highlight.Blocks {
			highlightBlockWidth[block.Tag] = max(highlightBlockWidth[block.Tag],
				facepaint.MeasureStringWidth(block.Label, styledHighlightCard.blockLabel().Font),
				facepaint.MeasureStringWidth(block.Value().String(), styledHighlightCard.blockValue().Font),
			)
		}
	}

	vehicleBlockWidth := make(map[prepare.Tag]float64)
	for _, card := range cards.Unrated.Vehicles {
		for _, block := range card.Blocks {
			vehicleBlockWidth[block.Tag] = max(vehicleBlockWidth[block.Tag],
				facepaint.MeasureBlockWidth(block.Label, *styledVehicleLegendPillText()),
				facepaint.MeasureStringWidth(block.Value().String(), styledVehicleCard.value().Font),
			)
		}
	}

	// Build all blocks
	var overviewBlocks []*facepaint.Block
	if shouldRenderUnratedOverview {
		if card := newUnratedOverviewCard(cards.Unrated.Overview, maxWidthOverviewColumn); card != nil {
			overviewBlocks = append(overviewBlocks, card)
		}
	}
	if shouldRenderRatingOverview {
		if card := newRatingOverviewCard(cards.Rating, maxWidthOverviewColumn); card != nil {
			overviewBlocks = append(overviewBlocks, card)
		}
	}

	var highlightBlocks []*facepaint.Block
	if showHighlights {
		for _, card := range cards.Unrated.Highlights {
			highlightBlocks = append(highlightBlocks, newHighlightCard(card, highlightBlockWidth))
		}
	}

	var vehicleBlocks []*facepaint.Block
	for _, card := range cards.Unrated.Vehicles {
		vehicleBlocks = append(vehicleBlocks, newVehicleCard(card, vehicleBlockWidth))
	}

	footer := newFooterCard(sessionData, cards, opts)

	// Page 0 always has name + overviews
	page0Cards := []*facepaint.Block{newPlayerNameCard(careerData.Account)}
	page0Cards = append(page0Cards, overviewBlocks...)

	// Decide what goes on page 0 vs continuation pages.
	//
	// No highlights → not enough content for multiple pages, backfill vehicles onto page 0.
	// 2 overview cards → perfect first page as-is; highlights + vehicles go to continuation.
	// 1 overview card → shift highlights to page 0; vehicles go to continuation.
	var restCards []*facepaint.Block
	if len(highlightBlocks) == 0 {
		// No highlights: everything on page 0
		page0Cards = append(page0Cards, vehicleBlocks...)
	} else if len(overviewBlocks) >= 2 {
		// 2 overviews fill page 0; highlights + vehicles go to continuation
		restCards = append(restCards, highlightBlocks...)
		restCards = append(restCards, vehicleBlocks...)
	} else {
		// 1 overview: highlights join page 0, vehicles go to continuation
		page0Cards = append(page0Cards, highlightBlocks...)
		restCards = append(restCards, vehicleBlocks...)
	}

	pr := opts.Printer
	if pr == nil {
		pr = func(s string) string { return s }
	}
	bgSource := opts.Background

	var pages []*facepaint.Block

	// Build page 0 (always present)
	cardsFrame := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledStatsFrame)), page0Cards...)
	cardsFrame = wrapCardsFrameWithBackground(cardsFrame, bgSource, sessionData, careerData, opts)
	pages = append(pages, facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledFinalFrame)), cardsFrame, footer))

	if len(restCards) == 0 {
		return pages, nil
	}

	// Cap continuation cards so total pages stay within the limit
	maxRestCards := (maxPages - 1) * maxCardsPerPage
	if len(restCards) > maxRestCards {
		restCards = restCards[:maxRestCards]
	}

	// Split continuation cards into balanced pages
	pageSizes := balancePages(len(restCards), maxCardsPerPage)

	off := 0
	totalRest := len(restCards)
	for _, sz := range pageSizes {
		slice := restCards[off : off+sz]
		off += sz

		remaining := totalRest - off

		cf := facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledStatsFrame)), slice...)
		cf = wrapCardsFrameWithBackground(cf, bgSource, sessionData, careerData, opts)

		var frameParts []*facepaint.Block
		frameParts = append(frameParts, cf)
		if remaining > 0 {
			frameParts = append(frameParts, newSessionNavPill(fmt.Sprintf(pr("session_page_nav_vehicles_remaining_fmt"), remaining)))
		}
		pages = append(pages, facepaint.NewBlocksContent(style.NewStyle(style.Parent(styledFinalFrame)), frameParts...))
	}

	return pages, nil
}

// balancePages splits n items into pages of at most maxPerPage,
// rebalancing the last two pages so neither is too small.
func balancePages(n, maxPerPage int) []int {
	if n <= maxPerPage {
		return []int{n}
	}

	var counts []int
	rest := n
	for rest > 0 {
		take := min(maxPerPage, rest)
		counts = append(counts, take)
		rest -= take
	}

	// Rebalance last two pages: split evenly, extra goes to earlier page
	if len(counts) >= 2 {
		i, j := len(counts)-2, len(counts)-1
		total := counts[i] + counts[j]
		counts[i] = (total + 1) / 2
		counts[j] = total / 2
	}

	return counts
}

func wrapCardsFrameWithBackground(cardsFrame *facepaint.Block, bgSource image.Image, sessionData, careerData fetch.AccountStatsOverPeriod, opts common.Options) *facepaint.Block {
	if bgSource == nil {
		return cardsFrame
	}
	sz := cardsFrame.Dimensions()
	var bg image.Image = imaging.Fill(bgSource, sz.Width, sz.Height, imaging.Center, imaging.Lanczos)
	if !opts.BackgroundIsCustom {
		seed, _ := strconv.Atoi(careerData.Account.ID)
		bg = addBackgroundBranding(bg, sessionData.RegularBattles.Vehicles, seed)
	}
	return facepaint.NewBlocksContent(style.NewStyle(),
		facepaint.MustNewImageContent(styledCardsBackground, bg), cardsFrame,
	)
}
