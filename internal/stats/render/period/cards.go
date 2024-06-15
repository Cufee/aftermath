package period

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common"
	"github.com/cufee/aftermath/internal/stats/prepare/period"
	"github.com/cufee/aftermath/internal/stats/render"
	"github.com/cufee/aftermath/internal/stats/render/common"

	"github.com/rs/zerolog/log"
)

func generateCards(stats fetch.AccountStatsOverPeriod, cards period.Cards, subs []database.UserSubscription, opts render.Options) (render.Segments, error) {
	if len(cards.Overview.Blocks) == 0 && len(cards.Highlights) == 0 {
		log.Error().Msg("player cards slice is 0 length, this should not happen")
		return render.Segments{}, errors.New("no cards provided")
	}

	var segments render.Segments

	// Calculate minimal card width to fit all the content
	var cardWidth float64
	overviewColumnWidth := float64(common.DefaultLogoOptions().Width())
	{
		{
			titleStyle := common.DefaultPlayerTitleStyle(titleCardStyle(cardWidth))
			clanSize := common.MeasureString(stats.Account.ClanTag, *titleStyle.ClanTag.Font)
			nameSize := common.MeasureString(stats.Account.Nickname, *titleStyle.Nickname.Font)
			cardWidth = common.Max(cardWidth, titleStyle.TotalPaddingAndGaps()+nameSize.TotalWidth+clanSize.TotalWidth*2)
		}
		{
			rowStyle := getOverviewStyle(cardWidth)
			for _, column := range cards.Overview.Blocks {
				for _, block := range column {
					valueStyle, labelStyle := rowStyle.block(block)

					label := block.Label
					if block.Tag == prepare.TagWN8 {
						label = common.GetWN8TierName(block.Value.Float())
					}
					labelSize := common.MeasureString(label, *labelStyle.Font)
					valueSize := common.MeasureString(block.Value.String(), *valueStyle.Font)

					overviewColumnWidth = common.Max(overviewColumnWidth, common.Max(labelSize.TotalWidth, valueSize.TotalWidth)+(rowStyle.container.PaddingX*2))
				}
			}

			cardStyle := overviewCardStyle(cardWidth)
			paddingAndGaps := (cardStyle.PaddingX+rowStyle.container.PaddingX+rowStyle.blockContainer.PaddingX)*2 + float64(len(cards.Overview.Blocks)-1)*(cardStyle.Gap+rowStyle.container.Gap+rowStyle.blockContainer.Gap)

			overviewCardContentWidth := overviewColumnWidth * float64(len(cards.Overview.Blocks))
			cardWidth = common.Max(cardWidth, overviewCardContentWidth+paddingAndGaps)
		}

		{
			highlightStyle := highlightCardStyle(defaultCardStyle(0))
			var highlightBlocksMaxCount, highlightTitleMaxWidth, highlightBlockMaxSize float64
			for _, highlight := range cards.Highlights {
				// Title and tank name
				metaSize := common.MeasureString(highlight.Meta, *highlightStyle.cardTitle.Font)
				titleSize := common.MeasureString(highlight.Title, *highlightStyle.tankName.Font)
				highlightTitleMaxWidth = common.Max(highlightTitleMaxWidth, metaSize.TotalWidth, titleSize.TotalWidth)

				// Blocks
				highlightBlocksMaxCount = common.Max(highlightBlocksMaxCount, float64(len(highlight.Blocks)))
				for _, block := range highlight.Blocks {
					labelSize := common.MeasureString(block.Label, *highlightStyle.blockLabel.Font)
					valueSize := common.MeasureString(block.Value.String(), *highlightStyle.blockValue.Font)
					highlightBlockMaxSize = common.Max(highlightBlockMaxSize, valueSize.TotalWidth, labelSize.TotalWidth)
				}
			}

			highlightCardWidthMax := (highlightStyle.container.PaddingX * 2) + (highlightStyle.container.Gap * highlightBlocksMaxCount) + (highlightBlockMaxSize * highlightBlocksMaxCount) + highlightTitleMaxWidth
			cardWidth = common.Max(cardWidth, highlightCardWidthMax)
		}
	}

	// We first render a footer in order to calculate the minimum required width
	{
		var footer []string
		switch strings.ToLower(stats.Realm) {
		case "na":
			footer = append(footer, "North America")
		case "eu":
			footer = append(footer, "Europe")
		case "as":
			footer = append(footer, "Asia")
		}

		sessionTo := stats.PeriodEnd.Format("Jan 2, 2006")
		sessionFrom := stats.PeriodStart.Format("Jan 2, 2006")
		if sessionFrom == sessionTo {
			footer = append(footer, sessionTo)
		} else {
			footer = append(footer, sessionFrom+" - "+sessionTo)
		}
		footerBlock := common.NewFooterCard(strings.Join(footer, " • "))
		footerImage, err := footerBlock.Render()
		if err != nil {
			return segments, err
		}

		cardWidth = max(cardWidth, float64(footerImage.Bounds().Dx()))
		segments.AddFooter(common.NewImageContent(common.Style{Width: cardWidth, Height: float64(footerImage.Bounds().Dy())}, footerImage))
	}

	// Header card
	if headerCard, headerCardExists := common.NewHeaderCard(cardWidth, subs, opts.PromoText); headerCardExists {
		segments.AddHeader(headerCard)
	}

	// Player Title card
	segments.AddContent(common.NewPlayerTitleCard(common.DefaultPlayerTitleStyle(titleCardStyle(cardWidth)), stats.Account.Nickname, stats.Account.ClanTag, subs))

	// Overview Card
	{
		var overviewCardBlocks []common.Block
		for _, column := range cards.Overview.Blocks {
			columnBlock, err := statsBlocksToColumnBlock(getOverviewStyle(overviewColumnWidth), column)
			if err != nil {
				return segments, err
			}
			overviewCardBlocks = append(overviewCardBlocks, columnBlock)
		}
		segments.AddContent(common.NewBlocksContent(overviewCardStyle(cardWidth), overviewCardBlocks...))
	}

	// Highlights
	for _, card := range cards.Highlights {
		segments.AddContent(newHighlightCard(highlightCardStyle(defaultCardStyle(cardWidth)), card))
	}

	return segments, nil
}

func newHighlightCard(style highlightStyle, card period.VehicleCard) common.Block {
	titleBlock :=
		common.NewBlocksContent(common.Style{
			Direction: common.DirectionVertical,
		},
			common.NewTextContent(style.cardTitle, card.Meta),
			common.NewTextContent(style.tankName, card.Title),
		)

	var contentRow []common.Block
	for _, block := range card.Blocks {
		contentRow = append(contentRow, common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter},
			common.NewTextContent(style.blockValue, block.Value.String()),
			common.NewTextContent(style.blockLabel, block.Label),
		))
	}

	return common.NewBlocksContent(style.container, titleBlock, common.NewBlocksContent(common.Style{
		Gap: style.container.Gap,
	}, contentRow...))
}
