package session

import (
	"fmt"
	"strings"

	"github.com/cufee/aftermath/internal/database/models"
	common "github.com/cufee/aftermath/internal/render/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
)

func cardsToSegments(session, _ fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts common.Options) (common.Segments, error) {
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

	var segments common.Segments
	var primaryColumn []common.Block
	var secondaryColumn []common.Block

	// Calculate minimal card width to fit all the content
	overviewColumnSizes := make(map[string]float64)
	primaryCardBlockSizes := make(map[string]float64)
	secondaryCardBlockSizes := make(map[string]float64)
	highlightCardBlockSizes := make(map[string]float64)
	var primaryCardWidth float64 = minPrimaryCardWidth
	var secondaryCardWidth, totalFrameWidth float64

	{
		titleStyle := common.DefaultPlayerTitleStyle(session.Account.Nickname, playerNameCardStyle(0))
		clanSize := common.MeasureString(session.Account.ClanTag, titleStyle.ClanTag.Font)
		nameSize := common.MeasureString(session.Account.Nickname, titleStyle.Nickname.Font)
		primaryCardWidth = max(primaryCardWidth, titleStyle.TotalPaddingAndGaps()+nameSize.TotalWidth+clanSize.TotalWidth*2)
	}
	{
		for _, text := range opts.PromoText {
			size := common.MeasureString(text, promoTextStyle().Font)
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
			labelSize := common.MeasureString(card.Meta, highlightCardTitleTextStyle().Font)
			titleSize := common.MeasureString(card.Title, highlightVehicleNameTextStyle().Font)

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

			titleSize := common.MeasureString(card.Title, vehicleCardTitleTextStyle().Font)
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
		} else {
			footer = append(footer, common.RealmLabel(session.Realm))
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
			segments.AddFooter(common.NewFooterCard(strings.Join(footer, " • ")))
		}
	}

	frameWidth := secondaryCardWidth + primaryCardWidth
	if secondaryCardWidth > 0 && primaryCardWidth > 0 {
		frameWidth += frameStyle().Gap
	}
	totalFrameWidth = max(totalFrameWidth, frameWidth)

	// header card
	if headerCard, headerCardExists := common.NewHeaderCard(totalFrameWidth, subs, opts.PromoText); headerCardExists {
		segments.AddHeader(headerCard)
	}

	// player title
	primaryColumn = append(primaryColumn,
		common.NewPlayerTitleCard(common.DefaultPlayerTitleStyle(session.Account.Nickname, playerNameCardStyle(primaryCardWidth)), session.Account.Nickname, session.Account.ClanTag, subs),
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

	columns := []common.Block{common.NewBlocksContent(overviewColumnStyle(primaryCardWidth), primaryColumn...)}
	if len(secondaryColumn) > 0 {
		columns = append(columns, common.NewBlocksContent(overviewColumnStyle(secondaryCardWidth), secondaryColumn...))
	}
	segments.AddContent(common.NewBlocksContent(frameStyle(), columns...))

	return segments, nil
}

func vehicleBlocksWidth(blocks []prepare.StatsBlock[session.BlockData, string], sessionStyle, careerStyle, labelStyle, containerStyle common.Style) (map[string]float64, float64) {
	presetBlockWidth := make(map[string]float64)
	var contentWidth float64
	var maxBlockWidth float64
	for _, block := range blocks {
		var width float64
		{
			size := common.MeasureString(block.Data.Session().String(), sessionStyle.Font)
			width = max(width, size.TotalWidth+sessionStyle.PaddingX*2)
		}
		{
			size := common.MeasureString(block.Data.Career().String(), careerStyle.Font)
			width = max(width, size.TotalWidth+careerStyle.PaddingX*2)
		}
		{
			size := common.MeasureString(block.Label, labelStyle.Font)
			width = max(width, size.TotalWidth+labelStyle.PaddingX*2+vehicleLegendLabelContainer.PaddingX*2)
		}
		maxBlockWidth = max(maxBlockWidth, width)
		presetBlockWidth[block.Tag.String()] = max(presetBlockWidth[block.Tag.String()], width)
	}

	if containerStyle.Direction == common.DirectionHorizontal {
		for _, w := range presetBlockWidth {
			contentWidth += w
		}
	} else {
		contentWidth = maxBlockWidth
	}

	return presetBlockWidth, contentWidth
}

func highlightedVehicleBlocksWidth(blocks []prepare.StatsBlock[session.BlockData, string], sessionStyle, _, labelStyle, containerStyle common.Style) (map[string]float64, float64) {
	presetBlockWidth := make(map[string]float64)
	var contentWidth float64
	var maxBlockWidth float64
	for _, block := range blocks {
		var width float64
		{
			size := common.MeasureString(block.Data.Session().String(), sessionStyle.Font)
			width = max(width, size.TotalWidth+sessionStyle.PaddingX*2)
		}
		{
			size := common.MeasureString(block.Label, labelStyle.Font)
			width = max(width, size.TotalWidth+labelStyle.PaddingX*2)
		}
		maxBlockWidth = max(maxBlockWidth, width)
		presetBlockWidth[block.Tag.String()] = max(presetBlockWidth[block.Tag.String()], width)
	}

	if containerStyle.Direction == common.DirectionHorizontal {
		for _, w := range presetBlockWidth {
			contentWidth += w
		}
	} else {
		contentWidth = maxBlockWidth
	}

	return presetBlockWidth, contentWidth
}

func overviewColumnBlocksWidth(blocks []prepare.StatsBlock[session.BlockData, string], sessionStyle, careerStyle, labelStyle, containerStyle common.Style) (map[string]float64, float64) {
	presetBlockWidth, contentWidth := vehicleBlocksWidth(blocks, sessionStyle, careerStyle, labelStyle, containerStyle)
	for _, block := range blocks {
		// adjust width if this column includes a special icon
		if block.Tag == prepare.TagWN8 {
			tierNameSize := common.MeasureString(common.GetWN8TierName(block.Value().Float()), overviewSpecialRatingLabelStyle(nil).Font)
			tierNameWithPadding := tierNameSize.TotalWidth + overviewSpecialRatingPillStyle(nil).PaddingX*2
			presetBlockWidth[block.Tag.String()] = max(presetBlockWidth[block.Tag.String()], specialRatingIconSize, tierNameWithPadding)
			contentWidth = max(contentWidth, tierNameWithPadding)
		}
		if block.Tag == prepare.TagRankedRating {
			valueSize := common.MeasureString(block.Value().String(), overviewSpecialRatingLabelStyle(nil).Font)
			presetBlockWidth[block.Tag.String()] = max(presetBlockWidth[block.Tag.String()], specialRatingIconSize, valueSize.TotalWidth)
			contentWidth = max(contentWidth, valueSize.TotalWidth)
		}
	}
	return presetBlockWidth, contentWidth
}

func makeVehicleCard(vehicle session.VehicleCard, blockSizes map[string]float64, cardWidth float64) common.Block {
	var vehicleWN8 frame.Value = frame.InvalidValue
	var content []common.Block
	for _, block := range vehicle.Blocks {
		style := vehicleBlockStyle()
		var blockContent []common.Block
		if blockShouldHaveCompareIcon(block) {
			blockContent = append(blockContent, blockWithVehicleIcon(common.NewTextContent(style.session, block.Data.Session().String()), block.Data.Session(), block.Data.Career()))
		} else {
			blockContent = append(blockContent, common.NewTextContent(style.session, block.Data.Session().String()))
		}

		containerStyle := statsBlockStyle(blockSizes[block.Tag.String()])
		content = append(content,
			common.NewBlocksContent(containerStyle, blockContent...),
		)

		if block.Tag == prepare.TagWN8 {
			vehicleWN8 = block.Value()
		}
	}

	titleStyle := vehicleCardTitleContainerStyle(cardWidth)
	titleStyle.Width -= vehicleCardStyle(0).PaddingX * 2
	return common.NewBlocksContent(vehicleCardStyle(cardWidth),
		common.NewBlocksContent(titleStyle,
			common.NewTextContent(vehicleCardTitleTextStyle(), fmt.Sprintf("%s %s", vehicle.Meta, vehicle.Title)), // name and tier
			vehicleWN8Icon(vehicleWN8),
		),
		common.NewBlocksContent(vehicleBlocksRowStyle(0), content...),
	)
}

func makeVehicleHighlightCard(vehicle session.VehicleCard, blockSizes map[string]float64, cardWidth float64) common.Block {
	var content []common.Block
	style := vehicleBlockStyle()
	for _, block := range vehicle.Blocks {
		content = append(content,
			common.NewBlocksContent(statsBlockStyle(blockSizes[block.Tag.String()]),
				common.NewTextContent(style.session, block.Data.Session().String()),
				common.NewTextContent(style.label, block.Label),
			),
		)
	}

	return common.NewBlocksContent(highlightedVehicleCardStyle(cardWidth),
		common.NewBlocksContent(common.Style{Direction: common.DirectionVertical},
			common.NewTextContent(highlightCardTitleTextStyle(), vehicle.Meta),
			common.NewTextContent(highlightVehicleNameTextStyle(), vehicle.Title),
		),
		common.NewBlocksContent(highlightedVehicleBlockRowStyle(0), content...),
	)
}

func makeVehicleLegendCard(reference session.VehicleCard, blockSizes map[string]float64, cardWidth float64) common.Block {
	var content []common.Block
	style := vehicleBlockStyle()
	for _, block := range reference.Blocks {
		label := common.NewBlocksContent(vehicleLegendLabelContainer, common.NewTextContent(style.label, block.Label))
		if blockShouldHaveCompareIcon(block) {
			label = blockWithVehicleIcon(common.NewBlocksContent(vehicleLegendLabelContainer, common.NewTextContent(style.label, block.Label)), frame.InvalidValue, frame.InvalidValue)
		}
		containerStyle := statsBlockStyle(blockSizes[block.Tag.String()])
		content = append(content,
			common.NewBlocksContent(containerStyle, label),
		)
	}
	containerStyle := vehicleCardStyle(cardWidth)
	containerStyle.BackgroundColor = nil
	containerStyle.PaddingY = 0
	containerStyle.PaddingX = 0
	return common.NewBlocksContent(containerStyle, common.NewBlocksContent(vehicleBlocksRowStyle(0), content...))
}

func makeOverviewCard(card session.OverviewCard, columnSizes map[string]float64, style common.Style) common.Block {
	// made all columns the same width for things to be centered
	var content []common.Block // add a blank block to balance the offset added from icons
	blockStyle := vehicleBlockStyle()
	for _, column := range card.Blocks {
		var columnContent []common.Block
		for _, block := range column.Blocks {
			var col common.Block
			blockWidth := columnSizes[string(column.Flavor)] // fit the block to column width to make things look even
			if block.Tag == prepare.TagWN8 || block.Tag == prepare.TagRankedRating {
				col = makeSpecialRatingColumn(block, blockWidth)
			} else if blockShouldHaveCompareIcon(block) {
				col = common.NewBlocksContent(statsBlockStyle(blockWidth),
					blockWithDoubleVehicleIcon(common.NewTextContent(blockStyle.session, block.Data.Session().String()), block.Data.Session(), block.Data.Career()),
					common.NewTextContent(blockStyle.label, block.Label),
				)
			} else {
				col = common.NewBlocksContent(statsBlockStyle(blockWidth),
					common.NewTextContent(blockStyle.session, block.Data.Session().String()),
					common.NewTextContent(blockStyle.label, block.Label),
				)
			}
			columnContent = append(columnContent, col)
		}
		content = append(content, common.NewBlocksContent(overviewColumnStyle(0), columnContent...))
	}
	return common.NewBlocksContent(style, content...)
}
