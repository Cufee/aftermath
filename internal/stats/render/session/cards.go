package session

import (
	"strings"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common"
	"github.com/cufee/aftermath/internal/stats/prepare/session"
	"github.com/cufee/aftermath/internal/stats/render/common"

	"github.com/cufee/aftermath/internal/stats/render"
)

func cardsToSegments(session, _ fetch.AccountStatsOverPeriod, cards session.Cards, subs []database.UserSubscription, opts render.Options) (render.Segments, error) {
	var segments render.Segments
	var primaryColumn []common.Block
	var secondaryColumn []common.Block

	// Calculate minimal card width to fit all the content
	primaryCardBlockSizes := make(map[string]float64)
	secondaryCardBlockSizes := make(map[string]float64)
	var primaryCardWidth, secondaryCardWidth, totalFrameWidth float64
	// rating and unrated battles > 0	 	   unrated battles > 0		   rating battles > 0
	// [title card		] | [vehicle]    [title card	  ] | [vehicle]    [title card	  	]
	// [overview unrated] | [vehicle] OR [overview unrated] | [vehicle] OR [overview rating ]
	// [overview rating ] | [...    ]    [highlight       ] | [...    ]    [vehicle			]

	{
		titleStyle := common.DefaultPlayerTitleStyle(playerNameCardStyle(0))
		clanSize := common.MeasureString(session.Account.ClanTag, *titleStyle.ClanTag.Font)
		nameSize := common.MeasureString(session.Account.Nickname, *titleStyle.Nickname.Font)
		primaryCardWidth = common.Max(primaryCardWidth, titleStyle.TotalPaddingAndGaps()+nameSize.TotalWidth+clanSize.TotalWidth*2)
	}
	{
		for _, text := range opts.PromoText {
			size := common.MeasureString(text, *promoTextStyle.Font)
			totalFrameWidth = common.Max(size.TotalWidth, totalFrameWidth)
		}
	}

	// unrated overview - show if there are no rating battles (so the card is not empty), or there are battles
	if session.RegularBattles.Battles > 0 || session.RatingBattles.Battles == 0 {
		for _, column := range cards.Unrated.Overview.Blocks {
			presetBlockWidth, _ := overviewColumnBlocksWidth(column, overviewStatsBlockStyle.session, overviewStatsBlockStyle.career, overviewStatsBlockStyle.label, overviewColumnStyle(0))
			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = common.Max(primaryCardBlockSizes[key], width)
			}
		}
	}
	// rating overview
	if session.RatingBattles.Battles > 0 {
		//
	}
	// rating vehicle cards go on the primary block - only show if there are no unrated battles/vehicles
	if len(cards.Unrated.Vehicles) == 0 {
		for _, card := range cards.Rating.Vehicles {
			// [title] [session]
			titleSize := common.MeasureString(card.Title, *ratingVehicleCardTitleStyle.Font)
			presetBlockWidth, contentWidth := vehicleBlocksWidth(card.Blocks, ratingVehicleBlockStyle.session, ratingVehicleBlockStyle.career, ratingVehicleBlockStyle.label, ratingVehicleBlocksRowStyle(0))
			// add the gap and card padding, the gap here accounts for title being inline with content
			contentWidth += ratingVehicleBlocksRowStyle(0).Gap*float64(len(card.Blocks)) + ratingVehicleCardStyle(0).PaddingX*2 + titleSize.TotalWidth

			primaryCardWidth = common.Max(primaryCardWidth, contentWidth)
			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = common.Max(primaryCardBlockSizes[key], width)
			}
		}
	}
	// highlighted vehicles go on the primary block
	if len(cards.Unrated.Highlights) > len(cards.Unrated.Vehicles) {
		for _, card := range cards.Unrated.Highlights {
			// [card label] [session]
			// [title]  	  [label]
			labelSize := common.MeasureString(card.Meta, *vehicleCardTitleStyle.Font)
			titleSize := common.MeasureString(card.Title, *vehicleCardTitleStyle.Font)
			presetBlockWidth, contentWidth := vehicleBlocksWidth(card.Blocks, vehicleBlockStyle.session, vehicleBlockStyle.career, vehicleBlockStyle.label, vehicleBlocksRowStyle(0))
			// add the gap and card padding, the gap here accounts for title/label being inline with content
			contentWidth += vehicleBlocksRowStyle(0).Gap*float64(len(card.Blocks)) + highlightCardStyle(0).PaddingX*2 + common.Max(titleSize.TotalWidth, labelSize.TotalWidth)
			primaryCardWidth = common.Max(primaryCardWidth, contentWidth)

			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = common.Max(primaryCardBlockSizes[key], width)
			}
		}
	}
	{ // unrated vehicles go on the secondary block
		for _, card := range cards.Unrated.Vehicles {
			// 		[ label ]
			// [session]
			// [career ]
			presetBlockWidth, contentWidth := vehicleBlocksWidth(card.Blocks, vehicleBlockStyle.session, vehicleBlockStyle.career, vehicleBlockStyle.label, vehicleBlocksRowStyle(0))
			contentWidth += vehicleBlocksRowStyle(0).Gap*float64(len(card.Blocks)-1) + vehicleCardStyle(0).PaddingX*2

			titleSize := common.MeasureString(card.Title, *vehicleCardTitleStyle.Font)
			secondaryCardWidth = common.Max(secondaryCardWidth, contentWidth, titleSize.TotalWidth+vehicleCardStyle(0).PaddingX*2)

			for key, width := range presetBlockWidth {
				secondaryCardBlockSizes[key] = common.Max(secondaryCardBlockSizes[key], width)
			}
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
	primaryColumn = append(primaryColumn,
		common.NewPlayerTitleCard(common.DefaultPlayerTitleStyle(playerNameCardStyle(primaryCardWidth)), session.Account.Nickname, session.Account.ClanTag, subs),
	)

	// overview cards
	if len(cards.Unrated.Overview.Blocks) > 0 {
		primaryColumn = append(primaryColumn, makeOverviewCard(cards.Unrated.Overview, primaryCardBlockSizes, primaryCardWidth))
	}
	if len(cards.Rating.Overview.Blocks) > 0 {
		primaryColumn = append(primaryColumn, makeOverviewCard(cards.Rating.Overview, primaryCardBlockSizes, primaryCardWidth))
	}

	// highlights
	if len(cards.Unrated.Highlights) > len(cards.Unrated.Vehicles) {
		//
	}
	// rating vehicle cards
	if len(cards.Unrated.Vehicles) == 0 {
		//
	}

	// unrated vehicles
	for i, vehicle := range cards.Unrated.Vehicles {
		if i == 0 {
			secondaryColumn = append(secondaryColumn, makeVehicleLegendCard(cards.Unrated.Vehicles[0], secondaryCardBlockSizes, secondaryCardWidth))
		}
		secondaryColumn = append(secondaryColumn, makeVehicleCard(vehicle, secondaryCardBlockSizes, secondaryCardWidth))
	}

	columns := []common.Block{common.NewBlocksContent(overviewColumnStyle(primaryCardWidth), primaryColumn...)}
	if len(secondaryColumn) > 0 {
		columns = append(columns, common.NewBlocksContent(overviewColumnStyle(secondaryCardWidth), secondaryColumn...))
	}
	segments.AddContent(common.NewBlocksContent(frameStyle(), columns...))

	var footer []string
	switch session.Realm {
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

	return segments, nil
}

func vehicleBlocksWidth(blocks []prepare.StatsBlock[session.BlockData], sessionStyle, careerStyle, labelStyle, containerStyle common.Style) (map[string]float64, float64) {
	presetBlockWidth := make(map[string]float64)
	var contentWidth, maxBlockWidth float64
	for _, block := range blocks {
		var width float64
		{
			size := common.MeasureString(block.Data.Session.String(), *sessionStyle.Font)
			width = common.Max(width, size.TotalWidth+sessionStyle.PaddingX*2)
		}
		{
			size := common.MeasureString(block.Data.Career.String(), *careerStyle.Font)
			width = common.Max(width, size.TotalWidth+careerStyle.PaddingX*2)
		}
		{
			size := common.MeasureString(block.Label, *labelStyle.Font)
			width = common.Max(width, size.TotalWidth+labelStyle.PaddingX*2)
		}
		w := width + float64(iconSize)
		maxBlockWidth = common.Max(maxBlockWidth, w)
		presetBlockWidth[block.Tag.String()] = common.Max(presetBlockWidth[block.Tag.String()], w) // add space for comparison icons
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
		if block.Tag == prepare.TagWN8 || block.Tag == prepare.TagRankedRating {
			tierNameSize := common.MeasureString(common.GetWN8TierName(block.Value.Float()), *overviewSpecialRatingLabelStyle(nil).Font)
			tierNameWithPadding := tierNameSize.TotalWidth + overviewSpecialRatingPillStyle(nil).PaddingX*2
			presetBlockWidth[block.Tag.String()] = common.Max(presetBlockWidth[block.Tag.String()], specialRatingIconSize, tierNameWithPadding)
		}
	}
	return presetBlockWidth, contentWidth
}

func makeVehicleCard(vehicle session.VehicleCard, blockSizes map[string]float64, cardWidth float64) common.Block {
	var content []common.Block
	for _, block := range vehicle.Blocks {
		content = append(content,
			common.NewBlocksContent(statsBlockStyle(blockSizes[block.Tag.String()]),
				common.NewTextContent(vehicleBlockStyle.session, block.Data.Session.String()),
				common.NewTextContent(vehicleBlockStyle.career, block.Data.Career.String()),
			),
		)
	}
	return common.NewBlocksContent(vehicleCardStyle(cardWidth),
		common.NewBlocksContent(vehicleCardTitleContainerStyle(cardWidth),
			common.NewTextContent(vehicleCardTitleStyle, vehicle.Title), // vehicle name
			common.NewTextContent(vehicleCardTitleStyle, vehicle.Meta),  // tier
		),
		common.NewBlocksContent(vehicleBlocksRowStyle(0), content...),
	)
}

func makeVehicleLegendCard(reference session.VehicleCard, blockSizes map[string]float64, cardWidth float64) common.Block {
	var content []common.Block
	for _, block := range reference.Blocks {
		style := statsBlockStyle(blockSizes[block.Tag.String()])
		style.BackgroundColor = common.DefaultCardColor
		style.BorderRadius = common.BorderRadiusSM
		style.PaddingY = 5
		content = append(content,
			common.NewBlocksContent(style,
				common.NewTextContent(vehicleBlockStyle.label, block.Label),
			),
		)
	}
	style := vehicleCardStyle(cardWidth)
	style.BackgroundColor = nil
	style.PaddingY = 0
	style.PaddingX = 0
	return common.NewBlocksContent(style, common.NewBlocksContent(vehicleBlocksRowStyle(0), content...))
}

func makeOverviewCard(card session.OverviewCard, blockSizes map[string]float64, cardWidth float64) common.Block {
	var content []common.Block
	for _, column := range card.Blocks {
		var columnContent []common.Block
		for _, block := range column {
			var col common.Block
			if block.Tag == prepare.TagWN8 || block.Tag == prepare.TagRankedRating {
				col = makeSpecialRatingColumn(block, blockSizes[block.Tag.String()])
			} else {
				col = common.NewBlocksContent(statsBlockStyle(blockSizes[block.Tag.String()]),
					common.NewTextContent(vehicleBlockStyle.session, block.Data.Session.String()),
					// common.NewTextContent(vehicleBlockStyle.career, block.Data.Career.String()),
					common.NewTextContent(vehicleBlockStyle.label, block.Label),
				)
			}
			columnContent = append(columnContent, col)
		}
		content = append(content, common.NewBlocksContent(overviewColumnStyle(0), columnContent...))
	}
	return common.NewBlocksContent(overviewCardStyle(cardWidth), content...)
}

func makeSpecialRatingColumn(block prepare.StatsBlock[session.BlockData], width float64) common.Block {
	switch block.Tag {
	case prepare.TagWN8:
		ratingColors := common.GetWN8Colors(block.Value.Float())
		if block.Value.Float() <= 0 {
			ratingColors.Content = common.TextAlt
			ratingColors.Background = common.TextAlt
		}

		var column []common.Block
		iconTop := common.AftermathLogo(ratingColors.Background, common.DefaultLogoOptions())
		column = append(column, common.NewImageContent(common.Style{Width: specialRatingIconSize, Height: specialRatingIconSize}, iconTop))

		column = append(column, common.NewBlocksContent(overviewColumnStyle(width),
			common.NewTextContent(vehicleBlockStyle.session, block.Data.Session.String()),
			common.NewBlocksContent(
				overviewSpecialRatingPillStyle(ratingColors.Background),
				common.NewTextContent(overviewSpecialRatingLabelStyle(ratingColors.Content), common.GetWN8TierName(block.Value.Float())),
			),
		))
		return common.NewBlocksContent(specialRatingColumnStyle, column...)

	default:
		return common.NewBlocksContent(statsBlockStyle(width),
			common.NewTextContent(vehicleBlockStyle.session, block.Data.Session.String()),
			// common.NewTextContent(vehicleBlockStyle.career, block.Data.Career.String()),
			common.NewTextContent(vehicleBlockStyle.label, block.Label),
		)
	}
}

// func uniqueBlockWN8(stats prepare.StatsBlock[period.BlockData]) common.Block {
// 	var blocks []common.Block

// 	valueStyle, labelStyle := style.block(stats)
// 	valueBlock := common.NewTextContent(valueStyle, stats.Value.String())

// 	ratingColors := common.GetWN8Colors(stats.Value.Float())
// 	if stats.Value.Float() <= 0 {
// 		ratingColors.Content = common.TextAlt
// 		ratingColors.Background = common.TextAlt
// 	}

// 	iconTop := common.AftermathLogo(ratingColors.Background, common.DefaultLogoOptions())
// 	iconBlockTop := common.NewImageContent(common.Style{Width: float64(iconTop.Bounds().Dx()), Height: float64(iconTop.Bounds().Dy())}, iconTop)

// 	style.blockContainer.Gap = 10
// 	blocks = append(blocks, common.NewBlocksContent(style.blockContainer, iconBlockTop, valueBlock))

// 	if stats.Value.Float() >= 0 {
// 		labelStyle.FontColor = ratingColors.Content
// 		blocks = append(blocks, common.NewBlocksContent(common.Style{
// 			PaddingY:        5,
// 			PaddingX:        10,
// 			BorderRadius:    15,
// 			BackgroundColor: ratingColors.Background,
// 		}, common.NewTextContent(labelStyle, common.GetWN8TierName(stats.Value.Float()))))
// 	}

// 	return common.NewBlocksContent(common.Style{Direction: common.DirectionVertical, AlignItems: common.AlignItemsCenter, Gap: 10, PaddingY: 5}, blocks...)
// }
