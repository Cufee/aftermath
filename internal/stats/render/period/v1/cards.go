package period

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	v1 "github.com/cufee/aftermath/internal/render/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"

	"github.com/cufee/aftermath/internal/log"
)

func generateCards(stats fetch.AccountStatsOverPeriod, cards period.Cards, subs []models.UserSubscription, opts common.Options) (v1.Segments, error) {
	if len(cards.Overview.Blocks) == 0 && len(cards.Highlights) == 0 {
		log.Error().Msg("player cards slice is 0 length, this should not happen")
		return v1.Segments{}, errors.New("no cards provided")
	}

	var segments v1.Segments

	// Calculate minimal card width to fit all the content
	var cardWidth float64
	overviewColumnWidth := float64(common.DefaultLogoOptions().Width())
	{
		{
			titleStyle := v1.DefaultPlayerTitleStyle(stats.Account.Nickname, titleCardStyle(cardWidth))
			clanSize := v1.MeasureString(stats.Account.ClanTag, titleStyle.ClanTag.Font)
			nameSize := v1.MeasureString(stats.Account.Nickname, titleStyle.Nickname.Font)
			cardWidth = max(cardWidth, titleStyle.TotalPaddingAndGaps()+nameSize.TotalWidth+clanSize.TotalWidth*2)
		}
		{
			rowStyle := getOverviewStyle(cardWidth)
			for _, column := range append(cards.Overview.Blocks, cards.Rating.Blocks...) {
				for _, block := range column.Blocks {
					valueStyle, labelStyle := rowStyle.block(block)

					label := block.Label
					if block.Tag == prepare.TagWN8 {
						label = common.GetWN8TierName(block.Value().Float())
					}
					if block.Tag == prepare.TagRankedRating {
						label = common.GetRatingTierName(block.Value().Float())
					}
					labelSize := v1.MeasureString(label, labelStyle.Font)
					valueSize := v1.MeasureString(block.Value().String(), valueStyle.Font)

					overviewColumnWidth = max(overviewColumnWidth, max(labelSize.TotalWidth+overviewSpecialRatingPillStyle().PaddingX*2, valueSize.TotalWidth))
				}
			}

			cardStyle := overviewCardBlocksStyle(cardWidth)
			paddingAndGaps := (cardStyle.PaddingX+rowStyle.container.PaddingX+rowStyle.blockContainer.PaddingX)*2 + float64(len(cards.Overview.Blocks)-1)*(cardStyle.Gap+rowStyle.container.Gap+rowStyle.blockContainer.Gap)

			overviewCardContentWidth := overviewColumnWidth * float64(len(cards.Overview.Blocks))
			cardWidth = max(cardWidth, overviewCardContentWidth+paddingAndGaps)
		}

		{
			highlightStyle := highlightCardStyle(defaultCardStyle(0))
			var highlightBlocksMaxCount, highlightTitleMaxWidth, highlightBlockMaxSize float64
			for _, highlight := range cards.Highlights {
				// Title and tank name
				metaSize := v1.MeasureString(highlight.Meta, highlightStyle.cardTitle.Font)
				titleSize := v1.MeasureString(highlight.Title, highlightStyle.tankName.Font)
				highlightTitleMaxWidth = max(highlightTitleMaxWidth, metaSize.TotalWidth, titleSize.TotalWidth)

				// Blocks
				highlightBlocksMaxCount = max(highlightBlocksMaxCount, float64(len(highlight.Blocks)))
				for _, block := range highlight.Blocks {
					labelSize := v1.MeasureString(block.Label, highlightStyle.blockLabel.Font)
					valueSize := v1.MeasureString(block.Value().String(), highlightStyle.blockValue.Font)
					highlightBlockMaxSize = max(highlightBlockMaxSize, valueSize.TotalWidth, labelSize.TotalWidth)
				}
			}

			highlightCardWidthMax := (highlightStyle.container.PaddingX * 2) + (highlightStyle.container.Gap * highlightBlocksMaxCount) + (highlightBlockMaxSize * highlightBlocksMaxCount) + highlightTitleMaxWidth
			cardWidth = max(cardWidth, highlightCardWidthMax)
		}
	}

	// We first render a footer in order to calculate the minimum required width
	{
		var footer []string
		if opts.VehicleID != "" {
			footer = append(footer, cards.Overview.Title)
		}

		sessionTo := stats.PeriodEnd.Format("Jan 2, 2006")
		sessionFrom := stats.PeriodStart.Format("Jan 2, 2006")
		if sessionFrom == sessionTo {
			footer = append(footer, sessionTo)
		} else {
			footer = append(footer, sessionFrom+" - "+sessionTo)
		}
		footerBlock := v1.NewFooterCard(strings.Join(footer, " • "))
		footerImage, err := footerBlock.Render()
		if err != nil {
			return segments, err
		}

		cardWidth = max(cardWidth, float64(footerImage.Bounds().Dx()))
		segments.AddFooter(v1.NewImageContent(v1.Style{}, footerImage))
	}

	// Header card
	if headerCard, headerCardExists := v1.NewHeaderCard(cardWidth, subs, opts.PromoText); headerCardExists {
		segments.AddHeader(headerCard)
	}

	// Player Title card
	segments.AddContent(v1.NewPlayerTitleCard(v1.DefaultPlayerTitleStyle(stats.Account.Nickname, titleCardStyle(cardWidth)), stats.Account.Nickname, stats.Account.ClanTag, subs))

	// Overview Card
	if len(cards.Overview.Blocks) > 0 {
		var overviewStatsBlocks []v1.Block
		for _, column := range cards.Overview.Blocks {
			columnBlock, err := statsBlocksToColumnBlock(getOverviewStyle(overviewColumnWidth), column.Blocks)
			if err != nil {
				return segments, err
			}
			overviewStatsBlocks = append(overviewStatsBlocks, columnBlock)
		}
		var overviewCardBlocks []v1.Block
		overviewCardBlocks = append(overviewCardBlocks, v1.NewBlocksContent(overviewCardBlocksStyle(cardWidth), overviewStatsBlocks...))
		segments.AddContent(v1.NewBlocksContent(overviewCardStyle(), overviewCardBlocks...))
	}

	// Rating Card -- only when player has current season rating
	if cards.Rating.Meta {
		var ratingStatsBlocks []v1.Block
		for _, column := range cards.Rating.Blocks {
			columnBlock, err := statsBlocksToColumnBlock(getOverviewStyle(overviewColumnWidth), column.Blocks)
			if err != nil {
				return segments, err
			}
			ratingStatsBlocks = append(ratingStatsBlocks, columnBlock)
		}
		var ratingCardBlocks []v1.Block
		ratingCardBlocks = append(ratingCardBlocks, v1.NewBlocksContent(overviewCardBlocksStyle(cardWidth), ratingStatsBlocks...))
		segments.AddContent(v1.NewBlocksContent(overviewCardStyle(), ratingCardBlocks...))
	}

	// Highlights
	for i, card := range cards.Highlights {
		if i > 0 && cards.Rating.Meta {
			break // only show 1 highlight when rating battles card is visible
		}
		segments.AddContent(newHighlightCard(highlightCardStyle(defaultCardStyle(cardWidth)), card))
	}

	return segments, nil
}

func newHighlightCard(style highlightStyle, card period.VehicleCard) v1.Block {
	titleBlock :=
		v1.NewBlocksContent(v1.Style{
			Direction: v1.DirectionVertical,
		},
			v1.NewTextContent(style.cardTitle, card.Meta),
			v1.NewTextContent(style.tankName, card.Title),
		)

	var contentRow []v1.Block
	for _, block := range card.Blocks {
		contentRow = append(contentRow, v1.NewBlocksContent(v1.Style{Direction: v1.DirectionVertical, AlignItems: v1.AlignItemsCenter},
			v1.NewTextContent(style.blockValue, block.Value().String()),
			v1.NewTextContent(style.blockLabel, block.Label),
		))
	}

	return v1.NewBlocksContent(style.container, titleBlock, v1.NewBlocksContent(v1.Style{
		Gap: style.container.Gap,
	}, contentRow...))
}
