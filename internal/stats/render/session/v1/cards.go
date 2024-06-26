package session

import (
	"fmt"
	"strings"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

func cardsToSegments(session, _ fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts common.Options) (common.Segments, error) {
	var (
		renderVehiclesCount = 3 // minimum number of vehicle cards
		// primary cards
		// when there are some unrated battles or no battles at all
		shouldRenderUnratedOverview = session.RegularBattles.Battles > 0 || session.RatingBattles.Battles == 0
		// when there are 3 vehicle cards and no rating overview cards or there are 6 vehicle cards and some rating battles
		shouldRenderUnratedHighlights = (session.RegularBattles.Battles > 0 && session.RatingBattles.Battles < 1 && len(cards.Unrated.Vehicles) > renderVehiclesCount) ||
			(session.RegularBattles.Battles > 0 && len(cards.Unrated.Vehicles) > 6)
		shouldRenderRatingOverview = session.RatingBattles.Battles > 0
		shouldRenderRatingVehicles = len(cards.Unrated.Vehicles) == 0
		// secondary cards
		shouldRenderUnratedVehicles = len(cards.Unrated.Vehicles) > 0
	)

	// try to make the columns height roughly similar to primary column
	if shouldRenderUnratedHighlights && len(cards.Unrated.Highlights) > 1 {
		renderVehiclesCount += len(cards.Unrated.Highlights)
	}
	if shouldRenderRatingOverview {
		renderVehiclesCount += 2
	}
	if shouldRenderRatingVehicles {
		renderVehiclesCount += len(cards.Rating.Vehicles)
	}

	var segments common.Segments
	var primaryColumn []common.Block
	var secondaryColumn []common.Block

	// Calculate minimal card width to fit all the content
	primaryCardBlockSizes := make(map[string]float64)
	secondaryCardBlockSizes := make(map[string]float64)
	highlightCardBlockSizes := make(map[string]float64)
	var primaryCardWidth float64 = minPrimaryCardWidth
	var secondaryCardWidth, totalFrameWidth float64
	// rating and unrated battles > 0	 	   unrated battles > 0		   rating battles > 0
	// [title card		] | [vehicle]    [title card	  ] | [vehicle]    [title card	  	]
	// [overview unrated] | [vehicle] OR [overview unrated] | [vehicle] OR [overview rating ]
	// [overview rating ] | [...    ]    [highlight       ] | [...    ]    [vehicle			]

	{
		titleStyle := common.DefaultPlayerTitleStyle(playerNameCardStyle(0))
		clanSize := common.MeasureString(session.Account.ClanTag, titleStyle.ClanTag.Font)
		nameSize := common.MeasureString(session.Account.Nickname, titleStyle.Nickname.Font)
		primaryCardWidth = common.Max(primaryCardWidth, titleStyle.TotalPaddingAndGaps()+nameSize.TotalWidth+clanSize.TotalWidth*2)
	}
	{
		for _, text := range opts.PromoText {
			size := common.MeasureString(text, promoTextStyle().Font)
			totalFrameWidth = common.Max(size.TotalWidth, totalFrameWidth)
		}
	}

	if shouldRenderUnratedOverview {
		var totalContentWidth float64 = overviewCardStyle(0).Gap * float64(len(cards.Unrated.Overview.Blocks)-1)
		for _, column := range cards.Unrated.Overview.Blocks {
			styleWithIconOffset := overviewStatsBlockStyle()
			styleWithIconOffset.session.PaddingX += vehicleComparisonIconSize

			presetBlockWidth, contentWidth := overviewColumnBlocksWidth(column, styleWithIconOffset.session, styleWithIconOffset.career, styleWithIconOffset.label, overviewColumnStyle(0))
			totalContentWidth += contentWidth
			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = common.Max(primaryCardBlockSizes[key], width)
			}
		}
		primaryCardWidth = common.Max(primaryCardWidth, totalContentWidth)
	}

	if shouldRenderRatingOverview {
		var totalContentWidth float64 = overviewCardStyle(0).Gap * float64(len(cards.Rating.Overview.Blocks)-1)
		for _, column := range cards.Unrated.Overview.Blocks {
			styleWithIconOffset := overviewStatsBlockStyle()
			styleWithIconOffset.session.PaddingX += vehicleComparisonIconSize

			presetBlockWidth, contentWidth := overviewColumnBlocksWidth(column, styleWithIconOffset.session, styleWithIconOffset.career, styleWithIconOffset.label, overviewColumnStyle(0))
			totalContentWidth += contentWidth
			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = common.Max(primaryCardBlockSizes[key], width)
			}
		}
		primaryCardWidth = common.Max(primaryCardWidth, totalContentWidth)
	}
	// rating vehicle cards go on the primary block - only show if there are no unrated battles/vehicles
	if shouldRenderRatingVehicles {
		for _, card := range cards.Rating.Vehicles {
			// [title] [session]
			style := ratingVehicleBlockStyle()
			titleSize := common.MeasureString(card.Title, ratingVehicleCardTitleStyle().Font)
			presetBlockWidth, contentWidth := vehicleBlocksWidth(card.Blocks, style.session, style.career, style.label, ratingVehicleBlocksRowStyle(0))
			// add the gap and card padding, the gap here accounts for title being inline with content
			contentWidth += ratingVehicleBlocksRowStyle(0).Gap*float64(len(card.Blocks)) + ratingVehicleCardStyle(0).PaddingX*2 + titleSize.TotalWidth

			primaryCardWidth = common.Max(primaryCardWidth, contentWidth)
			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = common.Max(primaryCardBlockSizes[key], width)
			}
		}
	}
	// highlighted vehicles go on the primary block
	if shouldRenderUnratedHighlights {
		for _, card := range cards.Unrated.Highlights {
			// [card label] [session]
			// [title]  	  [label]
			titleStyle := highlightCardTitleTextStyle()
			labelSize := common.MeasureString(card.Meta, titleStyle.Font)
			titleSize := common.MeasureString(card.Title, titleStyle.Font)

			style := vehicleBlockStyle()
			presetBlockWidth, contentWidth := highlightedVehicleBlocksWidth(card.Blocks, style.session, style.career, style.label, highlightedVehicleBlockRowStyle(0))
			// add the gap and card padding, the gap here accounts for title/label being inline with content
			contentWidth += highlightedVehicleBlockRowStyle(0).Gap*float64(len(card.Blocks)) + highlightedVehicleCardStyle(0).Gap + highlightedVehicleCardStyle(0).PaddingX*2 + common.Max(titleSize.TotalWidth, labelSize.TotalWidth)
			primaryCardWidth = common.Max(primaryCardWidth, contentWidth)

			for key, width := range presetBlockWidth {
				highlightCardBlockSizes[key] = common.Max(highlightCardBlockSizes[key], width)
			}
		}
	}
	if shouldRenderUnratedVehicles { // unrated vehicles go on the secondary block
		for _, card := range cards.Unrated.Vehicles {
			// 		[ label ]
			// [session]
			// [career ]

			styleWithIconOffset := vehicleBlockStyle()
			// icon is only on one side, so we divide by 2
			styleWithIconOffset.label.PaddingX += vehicleComparisonIconSize / 2
			styleWithIconOffset.session.PaddingX += vehicleComparisonIconSize / 2

			presetBlockWidth, contentWidth := vehicleBlocksWidth(card.Blocks, styleWithIconOffset.session, styleWithIconOffset.career, styleWithIconOffset.label, vehicleBlocksRowStyle(0))
			contentWidth += vehicleBlocksRowStyle(0).Gap*float64(len(card.Blocks)-1) + vehicleCardStyle(0).PaddingX*2

			titleSize := common.MeasureString(card.Title, vehicleCardTitleTextStyle().Font)
			secondaryCardWidth = common.Max(secondaryCardWidth, contentWidth, titleSize.TotalWidth+vehicleCardStyle(0).PaddingX*2)

			for key, width := range presetBlockWidth {
				secondaryCardBlockSizes[key] = common.Max(secondaryCardBlockSizes[key], width)
			}
		}
	}

	{
		var footer []string
		switch strings.ToLower(session.Realm) {
		case "na":
			footer = append(footer, "North America")
		case "eu":
			footer = append(footer, "Europe")
		case "as":
			footer = append(footer, "Asia")
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
	totalFrameWidth = common.Max(totalFrameWidth, frameWidth)

	// header card
	if headerCard, headerCardExists := common.NewHeaderCard(totalFrameWidth, subs, opts.PromoText); headerCardExists {
		segments.AddHeader(headerCard)
	}

	// player title
	println("making player title")
	primaryColumn = append(primaryColumn,
		common.NewPlayerTitleCard(common.DefaultPlayerTitleStyle(playerNameCardStyle(primaryCardWidth)), session.Account.Nickname, session.Account.ClanTag, subs),
	)

	// overview cards
	if shouldRenderUnratedOverview {
		println("making unrated overview")
		primaryColumn = append(primaryColumn, makeOverviewCard(cards.Unrated.Overview, primaryCardBlockSizes, overviewCardStyle(primaryCardWidth)))
	}
	if shouldRenderRatingOverview {
		println("making rating overview")
		primaryColumn = append(primaryColumn, makeOverviewCard(cards.Rating.Overview, primaryCardBlockSizes, overviewRatingCardStyle(primaryCardWidth)))
	}

	// highlights
	if shouldRenderUnratedHighlights {
		println("making highlights")
		for _, vehicle := range cards.Unrated.Highlights {
			primaryColumn = append(primaryColumn, makeVehicleHighlightCard(vehicle, highlightCardBlockSizes, primaryCardWidth))
		}
	}
	// rating vehicle cards
	if shouldRenderRatingVehicles {
		//
	}

	// unrated vehicles
	if shouldRenderUnratedVehicles {
		println("making unrated vehicles")
		for i, vehicle := range cards.Unrated.Vehicles {
			if i >= renderVehiclesCount {
				break
			}
			secondaryColumn = append(secondaryColumn, makeVehicleCard(vehicle, secondaryCardBlockSizes, secondaryCardWidth))
			if i == len(cards.Unrated.Vehicles)-1 {
				secondaryColumn = append(secondaryColumn, makeVehicleLegendCard(cards.Unrated.Vehicles[0], secondaryCardBlockSizes, secondaryCardWidth))
			}
		}
	}

	columns := []common.Block{common.NewBlocksContent(overviewColumnStyle(primaryCardWidth), primaryColumn...)}
	if len(secondaryColumn) > 0 {
		columns = append(columns, common.NewBlocksContent(overviewColumnStyle(secondaryCardWidth), secondaryColumn...))
	}
	segments.AddContent(common.NewBlocksContent(frameStyle(), columns...))

	return segments, nil
}

func vehicleBlocksWidth(blocks []prepare.StatsBlock[session.BlockData], sessionStyle, careerStyle, labelStyle, containerStyle common.Style) (map[string]float64, float64) {
	presetBlockWidth := make(map[string]float64)
	var contentWidth float64
	var maxBlockWidth float64
	for _, block := range blocks {
		var width float64
		{
			size := common.MeasureString(block.Data.Session.String(), sessionStyle.Font)
			width = common.Max(width, size.TotalWidth+sessionStyle.PaddingX*2)
		}
		{
			size := common.MeasureString(block.Data.Career.String(), careerStyle.Font)
			width = common.Max(width, size.TotalWidth+careerStyle.PaddingX*2)
		}
		{
			size := common.MeasureString(block.Label, labelStyle.Font)
			width = common.Max(width, size.TotalWidth+labelStyle.PaddingX*2+vehicleLegendLabelContainer.PaddingX*2)
		}
		maxBlockWidth = common.Max(maxBlockWidth, width)
		presetBlockWidth[block.Tag.String()] = common.Max(presetBlockWidth[block.Tag.String()], width)
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

func highlightedVehicleBlocksWidth(blocks []prepare.StatsBlock[session.BlockData], sessionStyle, _, labelStyle, containerStyle common.Style) (map[string]float64, float64) {
	presetBlockWidth := make(map[string]float64)
	var contentWidth float64
	var maxBlockWidth float64
	for _, block := range blocks {
		var width float64
		{
			size := common.MeasureString(block.Data.Session.String(), sessionStyle.Font)
			width = common.Max(width, size.TotalWidth+sessionStyle.PaddingX*2)
		}
		{
			size := common.MeasureString(block.Label, labelStyle.Font)
			width = common.Max(width, size.TotalWidth+labelStyle.PaddingX*2)
		}
		maxBlockWidth = common.Max(maxBlockWidth, width)
		presetBlockWidth[block.Tag.String()] = common.Max(presetBlockWidth[block.Tag.String()], width)
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

func overviewColumnBlocksWidth(blocks []prepare.StatsBlock[session.BlockData], sessionStyle, careerStyle, labelStyle, containerStyle common.Style) (map[string]float64, float64) {
	presetBlockWidth, contentWidth := vehicleBlocksWidth(blocks, sessionStyle, careerStyle, labelStyle, containerStyle)
	for _, block := range blocks {
		// adjust width if this column includes a special icon
		if block.Tag == prepare.TagWN8 {
			tierNameSize := common.MeasureString(common.GetWN8TierName(block.Value.Float()), overviewSpecialRatingLabelStyle(nil).Font)
			tierNameWithPadding := tierNameSize.TotalWidth + overviewSpecialRatingPillStyle(nil).PaddingX*2
			presetBlockWidth[block.Tag.String()] = common.Max(presetBlockWidth[block.Tag.String()], specialRatingIconSize, tierNameWithPadding)
		}
		if block.Tag == prepare.TagRankedRating {
			presetBlockWidth[block.Tag.String()] = common.Max(presetBlockWidth[block.Tag.String()], specialRatingIconSize)
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
			blockContent = append(blockContent, blockWithVehicleIcon(common.NewTextContent(style.session, block.Data.Session.String()), block.Data.Session, block.Data.Career))
		} else {
			blockContent = append(blockContent, common.NewTextContent(style.session, block.Data.Session.String()))
		}

		containerStyle := statsBlockStyle(blockSizes[block.Tag.String()])
		content = append(content,
			common.NewBlocksContent(containerStyle, blockContent...),
		)

		if block.Tag == prepare.TagWN8 {
			vehicleWN8 = block.Value
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
				common.NewTextContent(style.session, block.Data.Session.String()),
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

func makeOverviewCard(card session.OverviewCard, blockSizes map[string]float64, style common.Style) common.Block {
	// made all columns the same width for things to be centered
	columnWidth := (style.Width - style.Gap*float64(len(card.Blocks)-1) - style.PaddingX*2) / float64(len(card.Blocks))
	var content []common.Block // add a blank block to balance the offset added from icons
	blockStyle := vehicleBlockStyle()
	for _, column := range card.Blocks {
		var columnContent []common.Block
		for _, block := range column {
			println("making block", block.Label)
			var col common.Block
			if block.Tag == prepare.TagWN8 || block.Tag == prepare.TagRankedRating {
				col = makeSpecialRatingColumn(block, blockSizes[block.Tag.String()])
			} else if blockShouldHaveCompareIcon(block) {
				col = common.NewBlocksContent(statsBlockStyle(blockSizes[block.Tag.String()]),
					blockWithDoubleVehicleIcon(common.NewTextContent(blockStyle.session, block.Data.Session.String()), block.Data.Session, block.Data.Career),
					common.NewTextContent(blockStyle.label, block.Label),
				)
			} else {
				col = common.NewBlocksContent(statsBlockStyle(blockSizes[block.Tag.String()]),
					common.NewTextContent(blockStyle.session, block.Data.Session.String()),
					common.NewTextContent(blockStyle.label, block.Label),
				)
			}
			columnContent = append(columnContent, col)
		}
		content = append(content, common.NewBlocksContent(overviewColumnStyle(columnWidth), columnContent...))
	}
	return common.NewBlocksContent(style, content...)
}
