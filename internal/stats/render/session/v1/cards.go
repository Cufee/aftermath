package session

import (
	"fmt"
	"strings"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	v1 "github.com/cufee/aftermath/internal/render/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
)

func cardsToSegments(session, _ fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts common.Options) (v1.Segments, error) {
	var (
		renderUnratedVehiclesCount = 3 // minimum number of vehicle cards
		// primary cards
		// when there are some unrated battles or no battles at all
		shouldRenderUnratedOverview = session.RegularBattles.Battles > 0 || session.RatingBattles.Battles < 1
		// when there are 3 vehicle cards and no rating overview cards or there are 6 vehicle cards and some rating battles
		shouldRenderUnratedHighlights = (session.RegularBattles.Battles > 0 && session.RatingBattles.Battles < 1 && len(cards.Unrated.Vehicles) > renderUnratedVehiclesCount) ||
			(session.RegularBattles.Battles > 0 && len(cards.Unrated.Vehicles) > 3)
		shouldRenderRatingOverview = session.RatingBattles.Battles > 0 && opts.VehicleID == ""
		// secondary cards
		shouldRenderUnratedVehicles = session.RegularBattles.Battles > 0 && len(cards.Unrated.Vehicles) > 0
	)

	// try to make the columns height roughly similar to primary column
	if shouldRenderUnratedHighlights {
		renderUnratedVehiclesCount += len(cards.Unrated.Highlights)
	}
	if shouldRenderRatingOverview {
		renderUnratedVehiclesCount += 1
	}

	var segments v1.Segments
	var primaryColumn []v1.Block
	var secondaryColumn []v1.Block

	// Calculate minimal card width to fit all the content
	overviewColumnSizes := make(map[string]float64)
	primaryCardBlockSizes := make(map[string]float64)
	secondaryCardBlockSizes := make(map[string]float64)
	highlightCardBlockSizes := make(map[string]float64)
	var primaryCardWidth float64 = minPrimaryCardWidth
	var secondaryCardWidth, totalFrameWidth float64

	{
		titleStyle := v1.DefaultPlayerTitleStyle(session.Account.Nickname, playerNameCardStyle(0))
		clanSize := v1.MeasureString(session.Account.ClanTag, titleStyle.ClanTag.Font)
		nameSize := v1.MeasureString(session.Account.Nickname, titleStyle.Nickname.Font)
		primaryCardWidth = max(primaryCardWidth, titleStyle.TotalPaddingAndGaps()+nameSize.TotalWidth+clanSize.TotalWidth*2)
	}
	{
		for _, text := range opts.PromoText {
			size := v1.MeasureString(text, promoTextStyle().Font)
			totalFrameWidth = max(size.TotalWidth, totalFrameWidth)
		}
	}

	if shouldRenderUnratedOverview {
		for _, column := range cards.Unrated.Overview.Blocks {
			styleWithIconOffset := overviewStatsBlockStyle()
			styleWithIconOffset.session.PaddingX += vehicleComparisonIconSize

			presetBlockWidth, contentWidth := overviewColumnBlocksWidth(column.Blocks, styleWithIconOffset.session, styleWithIconOffset.career, styleWithIconOffset.label, overviewColumnStyle(0))
			overviewColumnSizes[string(column.Flavor)] = max(overviewColumnSizes[string(column.Flavor)], contentWidth)
			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = max(primaryCardBlockSizes[key], width)
			}
		}
	}
	if shouldRenderRatingOverview {
		for _, column := range cards.Rating.Overview.Blocks {
			styleWithIconOffset := overviewStatsBlockStyle()
			styleWithIconOffset.session.PaddingX += vehicleComparisonIconSize

			presetBlockWidth, contentWidth := overviewColumnBlocksWidth(column.Blocks, styleWithIconOffset.session, styleWithIconOffset.career, styleWithIconOffset.label, overviewColumnStyle(0))
			overviewColumnSizes[string(column.Flavor)] = max(overviewColumnSizes[string(column.Flavor)], contentWidth)
			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = max(primaryCardBlockSizes[key], width)
			}
		}
	}
	// we now have column width for both unrated and rating overviews
	if shouldRenderUnratedOverview {
		var totalContentWidth float64 = overviewCardStyle(0).Gap*float64(len(cards.Unrated.Overview.Blocks)-1) + overviewCardStyle(0).PaddingX*2
		for _, column := range cards.Unrated.Overview.Blocks {
			totalContentWidth += overviewColumnSizes[string(column.Flavor)]
		}
		primaryCardWidth = max(primaryCardWidth, totalContentWidth)
	}
	if shouldRenderRatingOverview {
		var totalContentWidth float64 = overviewCardStyle(0).Gap*float64(len(cards.Rating.Overview.Blocks)-1) + overviewCardStyle(0).PaddingX*2
		for _, column := range cards.Rating.Overview.Blocks {
			totalContentWidth += overviewColumnSizes[string(column.Flavor)]
		}
		primaryCardWidth = max(primaryCardWidth, totalContentWidth)
	}

	// highlighted vehicles go on the primary block
	if shouldRenderUnratedHighlights {
		for _, card := range cards.Unrated.Highlights {
			labelSize := v1.MeasureString(card.Meta, highlightCardTitleTextStyle().Font)
			titleSize := v1.MeasureString(card.Title, highlightVehicleNameTextStyle().Font)

			style := vehicleBlockStyle()
			presetBlockWidth, contentWidth := highlightedVehicleBlocksWidth(card.Blocks, style.session, style.career, style.label, highlightedVehicleBlockRowStyle(0))
			// add the gap and card padding, the gap here accounts for title/label being inline with content
			contentWidth += highlightedVehicleBlockRowStyle(0).Gap*float64(len(card.Blocks)-1) + highlightedVehicleCardStyle(0).Gap + highlightedVehicleCardStyle(0).PaddingX*2 + max(titleSize.TotalWidth, labelSize.TotalWidth)
			primaryCardWidth = max(primaryCardWidth, contentWidth)

			for key, width := range presetBlockWidth {
				highlightCardBlockSizes[key] = max(highlightCardBlockSizes[key], width)
			}
		}
	}
	if shouldRenderUnratedVehicles { // unrated vehicles go on the secondary block
		for _, card := range cards.Unrated.Vehicles {
			styleWithIconOffset := vehicleBlockStyle()
			// icon is only on one side, so we divide by 2
			styleWithIconOffset.label.PaddingX += vehicleComparisonIconSize / 2
			styleWithIconOffset.session.PaddingX += vehicleComparisonIconSize / 2

			presetBlockWidth, contentWidth := vehicleBlocksWidth(card.Blocks, styleWithIconOffset.session, styleWithIconOffset.career, styleWithIconOffset.label, vehicleBlocksRowStyle(0))
			contentWidth += vehicleBlocksRowStyle(0).Gap*float64(len(card.Blocks)-1) + vehicleCardStyle(0).PaddingX*2

			titleSize := v1.MeasureString(card.Title, vehicleCardTitleTextStyle().Font)
			secondaryCardWidth = max(secondaryCardWidth, contentWidth, titleSize.TotalWidth+vehicleCardStyle(0).PaddingX*2)

			for key, width := range presetBlockWidth {
				secondaryCardBlockSizes[key] = max(secondaryCardBlockSizes[key], width)
			}
		}
	}

	{
		var footer []string
		if opts.VehicleID != "" {
			footer = append(footer, cards.Unrated.Overview.Title)
		}
		if session.LastBattleTime.Unix() > 1 {
			sessionTo := session.PeriodEnd.Format("Jan 2")
			sessionFrom := session.PeriodStart.Format("Jan 2")
			if sessionFrom == sessionTo {
				footer = append(footer, sessionTo)
			} else {
				footer = append(footer, sessionFrom+" - "+sessionTo)
			}
		}

		if len(footer) > 0 {
			segments.AddFooter(v1.NewFooterCard(strings.Join(footer, " • ")))
		}
	}

	frameWidth := secondaryCardWidth + primaryCardWidth
	if secondaryCardWidth > 0 && primaryCardWidth > 0 {
		frameWidth += frameStyle().Gap
	}
	totalFrameWidth = max(totalFrameWidth, frameWidth)

	// header card
	if headerCard, headerCardExists := v1.NewHeaderCard(totalFrameWidth, subs, opts.PromoText); headerCardExists {
		segments.AddHeader(headerCard)
	}

	// player title
	primaryColumn = append(primaryColumn,
		v1.NewPlayerTitleCard(v1.DefaultPlayerTitleStyle(session.Account.Nickname, playerNameCardStyle(primaryCardWidth)), session.Account.Nickname, session.Account.ClanTag, subs),
	)

	// overview cards
	if shouldRenderUnratedOverview {
		primaryColumn = append(primaryColumn, makeOverviewCard(cards.Unrated.Overview, overviewColumnSizes, overviewCardStyle(primaryCardWidth)))
	}
	if shouldRenderRatingOverview {
		primaryColumn = append(primaryColumn, makeOverviewCard(cards.Rating.Overview, overviewColumnSizes, overviewRatingCardStyle(primaryCardWidth)))
	}

	// highlights
	if shouldRenderUnratedHighlights {
		for _, vehicle := range cards.Unrated.Highlights {
			primaryColumn = append(primaryColumn, makeVehicleHighlightCard(vehicle, highlightCardBlockSizes, primaryCardWidth))
		}
	}
	// unrated vehicles
	if shouldRenderUnratedVehicles {
		for i, vehicle := range cards.Unrated.Vehicles {
			if i >= renderUnratedVehiclesCount {
				break
			}
			secondaryColumn = append(secondaryColumn, makeVehicleCard(vehicle, secondaryCardBlockSizes, secondaryCardWidth))
		}
		if len(cards.Unrated.Vehicles) > 0 {
			secondaryColumn = append(secondaryColumn, makeVehicleLegendCard(cards.Unrated.Vehicles[0], secondaryCardBlockSizes, secondaryCardWidth))
		}
	}

	columns := []v1.Block{v1.NewBlocksContent(overviewColumnStyle(primaryCardWidth), primaryColumn...)}
	if len(secondaryColumn) > 0 {
		columns = append(columns, v1.NewBlocksContent(overviewColumnStyle(secondaryCardWidth), secondaryColumn...))
	}
	segments.AddContent(v1.NewBlocksContent(frameStyle(), columns...))

	return segments, nil
}

func vehicleBlocksWidth(blocks []prepare.StatsBlock[session.BlockData, string], sessionStyle, careerStyle, labelStyle, containerStyle v1.Style) (map[string]float64, float64) {
	presetBlockWidth := make(map[string]float64)
	var contentWidth float64
	var maxBlockWidth float64
	for _, block := range blocks {
		var width float64
		{
			size := v1.MeasureString(block.Data.Session().String(), sessionStyle.Font)
			width = max(width, size.TotalWidth+sessionStyle.PaddingX*2)
		}
		{
			size := v1.MeasureString(block.Data.Career().String(), careerStyle.Font)
			width = max(width, size.TotalWidth+careerStyle.PaddingX*2)
		}
		{
			size := v1.MeasureString(block.Label, labelStyle.Font)
			width = max(width, size.TotalWidth+labelStyle.PaddingX*2+vehicleLegendLabelContainer.PaddingX*2)
		}
		maxBlockWidth = max(maxBlockWidth, width)
		presetBlockWidth[block.Tag.String()] = max(presetBlockWidth[block.Tag.String()], width)
	}

	if containerStyle.Direction == v1.DirectionHorizontal {
		for _, w := range presetBlockWidth {
			contentWidth += w
		}
	} else {
		contentWidth = maxBlockWidth
	}

	return presetBlockWidth, contentWidth
}

func highlightedVehicleBlocksWidth(blocks []prepare.StatsBlock[session.BlockData, string], sessionStyle, _, labelStyle, containerStyle v1.Style) (map[string]float64, float64) {
	presetBlockWidth := make(map[string]float64)
	var contentWidth float64
	var maxBlockWidth float64
	for _, block := range blocks {
		var width float64
		{
			size := v1.MeasureString(block.Data.Session().String(), sessionStyle.Font)
			width = max(width, size.TotalWidth+sessionStyle.PaddingX*2)
		}
		{
			size := v1.MeasureString(block.Label, labelStyle.Font)
			width = max(width, size.TotalWidth+labelStyle.PaddingX*2)
		}
		maxBlockWidth = max(maxBlockWidth, width)
		presetBlockWidth[block.Tag.String()] = max(presetBlockWidth[block.Tag.String()], width)
	}

	if containerStyle.Direction == v1.DirectionHorizontal {
		for _, w := range presetBlockWidth {
			contentWidth += w
		}
	} else {
		contentWidth = maxBlockWidth
	}

	return presetBlockWidth, contentWidth
}

func overviewColumnBlocksWidth(blocks []prepare.StatsBlock[session.BlockData, string], sessionStyle, careerStyle, labelStyle, containerStyle v1.Style) (map[string]float64, float64) {
	presetBlockWidth, contentWidth := vehicleBlocksWidth(blocks, sessionStyle, careerStyle, labelStyle, containerStyle)
	for _, block := range blocks {
		// adjust width if this column includes a special icon
		if block.Tag == prepare.TagWN8 {
			tierNameSize := v1.MeasureString(v1.GetWN8TierName(block.Value().Float()), overviewSpecialRatingLabelStyle().Font)
			tierNameWithPadding := tierNameSize.TotalWidth + overviewSpecialRatingPillStyle().PaddingX*2
			presetBlockWidth[block.Tag.String()] = max(presetBlockWidth[block.Tag.String()], specialWN8IconSize, tierNameWithPadding)
			contentWidth = max(contentWidth, tierNameWithPadding)
		}
		if block.Tag == prepare.TagRankedRating {
			valueSize := v1.MeasureString(v1.GetRatingTierName(block.Value().Float()), overviewSpecialRatingLabelStyle().Font)
			tierNameWithPadding := valueSize.TotalWidth + overviewSpecialRatingPillStyle().PaddingX*2
			presetBlockWidth[block.Tag.String()] = max(presetBlockWidth[block.Tag.String()], specialRatingIconSize, tierNameWithPadding)
			contentWidth = max(contentWidth, tierNameWithPadding)
		}
	}
	return presetBlockWidth, contentWidth
}

func makeVehicleCard(vehicle session.VehicleCard, blockSizes map[string]float64, cardWidth float64) v1.Block {
	var vehicleWN8 frame.Value = frame.InvalidValue
	var content []v1.Block
	for _, block := range vehicle.Blocks {
		style := vehicleBlockStyle()
		var blockContent []v1.Block
		if blockShouldHaveCompareIcon(block) {
			blockContent = append(blockContent, blockWithVehicleIcon(v1.NewTextContent(style.session, block.Data.Session().String()), block.Data.Session(), block.Data.Career()))
		} else {
			blockContent = append(blockContent, v1.NewTextContent(style.session, block.Data.Session().String()))
		}

		containerStyle := statsBlockStyle(blockSizes[block.Tag.String()])
		content = append(content,
			v1.NewBlocksContent(containerStyle, blockContent...),
		)

		if block.Tag == prepare.TagWN8 {
			vehicleWN8 = block.Value()
		}
	}

	titleStyle := vehicleCardTitleContainerStyle(cardWidth)
	titleStyle.Width -= vehicleCardStyle(0).PaddingX * 2
	return v1.NewBlocksContent(vehicleCardStyle(cardWidth),
		v1.NewBlocksContent(titleStyle,
			v1.NewTextContent(vehicleCardTitleTextStyle(), fmt.Sprintf("%s %s", vehicle.Meta, vehicle.Title)), // name and tier
			vehicleWN8Icon(vehicleWN8),
		),
		v1.NewBlocksContent(vehicleBlocksRowStyle(0), content...),
	)
}

func makeVehicleHighlightCard(vehicle session.VehicleCard, blockSizes map[string]float64, cardWidth float64) v1.Block {
	var content []v1.Block
	style := vehicleBlockStyle()
	for _, block := range vehicle.Blocks {
		content = append(content,
			v1.NewBlocksContent(statsBlockStyle(blockSizes[block.Tag.String()]),
				v1.NewTextContent(style.session, block.Data.Session().String()),
				v1.NewTextContent(style.label, block.Label),
			),
		)
	}

	return v1.NewBlocksContent(highlightedVehicleCardStyle(cardWidth),
		v1.NewBlocksContent(v1.Style{Direction: v1.DirectionVertical},
			v1.NewTextContent(highlightCardTitleTextStyle(), vehicle.Meta),
			v1.NewTextContent(highlightVehicleNameTextStyle(), vehicle.Title),
		),
		v1.NewBlocksContent(highlightedVehicleBlockRowStyle(0), content...),
	)
}

func makeVehicleLegendCard(reference session.VehicleCard, blockSizes map[string]float64, cardWidth float64) v1.Block {
	var content []v1.Block
	style := vehicleBlockStyle()
	for _, block := range reference.Blocks {
		label := v1.NewBlocksContent(vehicleLegendLabelContainer, v1.NewTextContent(style.label, block.Label))
		if blockShouldHaveCompareIcon(block) {
			label = blockWithVehicleIcon(v1.NewBlocksContent(vehicleLegendLabelContainer, v1.NewTextContent(style.label, block.Label)), frame.InvalidValue, frame.InvalidValue)
		}
		containerStyle := statsBlockStyle(blockSizes[block.Tag.String()])
		content = append(content,
			v1.NewBlocksContent(containerStyle, label),
		)
	}
	containerStyle := vehicleCardStyle(cardWidth)
	containerStyle.BackgroundColor = nil
	containerStyle.PaddingY = 0
	containerStyle.PaddingX = 0
	return v1.NewBlocksContent(containerStyle, v1.NewBlocksContent(vehicleBlocksRowStyle(0), content...))
}

func makeOverviewCard(card session.OverviewCard, columnSizes map[string]float64, style v1.Style) v1.Block {
	// made all columns the same width for things to be centered
	var content []v1.Block // add a blank block to balance the offset added from icons
	blockStyle := vehicleBlockStyle()
	for _, column := range card.Blocks {
		var columnContent []v1.Block
		for _, block := range column.Blocks {
			var col v1.Block
			blockWidth := columnSizes[string(column.Flavor)] // fit the block to column width to make things look even
			if block.Tag == prepare.TagWN8 || block.Tag == prepare.TagRankedRating {
				col = makeSpecialRatingColumn(block, blockWidth)
			} else if blockShouldHaveCompareIcon(block) {
				col = v1.NewBlocksContent(statsBlockStyle(blockWidth),
					blockWithDoubleVehicleIcon(v1.NewTextContent(blockStyle.session, block.Data.Session().String()), block.Data.Session(), block.Data.Career()),
					v1.NewTextContent(blockStyle.label, block.Label),
				)
			} else {
				col = v1.NewBlocksContent(statsBlockStyle(blockWidth),
					v1.NewTextContent(blockStyle.session, block.Data.Session().String()),
					v1.NewTextContent(blockStyle.label, block.Label),
				)
			}
			columnContent = append(columnContent, col)
		}
		content = append(content, v1.NewBlocksContent(overviewColumnStyle(0), columnContent...))
	}
	return v1.NewBlocksContent(style, content...)
}
