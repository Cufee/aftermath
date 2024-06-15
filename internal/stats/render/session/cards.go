package session

import (
	"strings"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare/session"
	"github.com/cufee/aftermath/internal/stats/render/common"

	"github.com/cufee/aftermath/internal/stats/render"
)

func cardsToSegments(session, career fetch.AccountStatsOverPeriod, cards session.Cards, subs []database.UserSubscription, opts render.Options) (render.Segments, error) {
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
		//
	}
	// rating overview
	if session.RatingBattles.Battles > 0 {
		//
	}
	// rating vehicle cards go on the primary block - only show if there are no highlights
	if len(cards.Unrated.Highlights) == 0 {
		for _, card := range cards.Rating.Vehicles {
			// [title]	[session]
			titleSize := common.MeasureString(card.Title, *ratingVehicleCardTitleStyle.Font)
			presetBlockWidth, contentWidth := vehicleBlocksWidth(card, ratingVehicleBlockStyle.session, ratingVehicleBlockStyle.career, ratingVehicleBlockStyle.label, ratingVehicleBlockStyle.container)
			// add the gap and card padding, the gap here accounts for title being inline with content
			contentWidth += ratingVehicleBlocksRowStyle(0).Gap*float64(len(card.Blocks)) + ratingVehicleCardStyle(0).PaddingX*2 + titleSize.TotalWidth

			primaryCardWidth = common.Max(primaryCardWidth, contentWidth)
			for key, width := range presetBlockWidth {
				primaryCardBlockSizes[key] = common.Max(primaryCardBlockSizes[key], width)
			}
		}
	}
	// highlighted vehicles go on the primary block
	{
		for _, card := range cards.Unrated.Highlights {
			// [card label]	[session]
			// [title]  	 [label]
			labelSize := common.MeasureString(card.Meta, *vehicleCardTitleStyle.Font)
			titleSize := common.MeasureString(card.Title, *vehicleCardTitleStyle.Font)
			presetBlockWidth, contentWidth := vehicleBlocksWidth(card, vehicleBlockStyle.session, vehicleBlockStyle.career, vehicleBlockStyle.label, vehicleBlockStyle.container)
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
			// [label  ] ...
			presetBlockWidth, contentWidth := vehicleBlocksWidth(card, vehicleBlockStyle.session, vehicleBlockStyle.career, vehicleBlockStyle.label, vehicleBlockStyle.container)
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

	// Header card
	if headerCard, headerCardExists := common.NewHeaderCard(totalFrameWidth, subs, opts.PromoText); headerCardExists {
		segments.AddHeader(headerCard)
	}

	// // User Subscription Badge and promo text
	// if addPromoText && opts.PromoText != nil {
	// 	// Users without a subscription get promo text
	// 	var textBlocks []common.Block
	// 	for _, text := range opts.PromoText {
	// 		textBlocks = append(textBlocks, render.NewTextContent(promoTextStyle, text))
	// 	}
	// 	cards = append(cards, common.NewBlocksContent(render.Style{
	// 		Direction:  render.DirectionVertical,
	// 		AlignItems: render.AlignItemsCenter,
	// 	},
	// 		textBlocks...,
	// 	))
	// }
	// if badges, _ := badges.SubscriptionsBadges(player.Subscriptions); len(badges) > 0 {
	// 	cards = append(cards, render.NewBlocksContent(render.Style{Direction: render.DirectionHorizontal, AlignItems: render.AlignItemsCenter, Gap: 10},
	// 		badges...,
	// 	))
	// }

	primaryColumn = append(primaryColumn,
		common.NewPlayerTitleCard(common.DefaultPlayerTitleStyle(playerNameCardStyle(primaryCardWidth)), session.Account.Nickname, session.Account.ClanTag, subs),
	)

	// // Rating Cards
	// if len(player.Cards.Rating) > 0 {
	// 	ratingGroup, err := makeCardsGroup(player.Cards.Rating, cardWidth, cardBlockSizes)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	cards = append(cards, ratingGroup)
	// }

	// // Unrated Cards
	// if len(player.Cards.Unrated) > 0 {
	// 	unratedGroup, err := makeCardsGroup(player.Cards.Unrated, cardWidth, cardBlockSizes)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	cards = append(cards, unratedGroup)
	// }

	columns := []common.Block{common.NewBlocksContent(columnStyle(primaryCardWidth), primaryColumn...)}
	if len(secondaryColumn) > 0 {
		columns = append(columns, common.NewBlocksContent(columnStyle(secondaryCardWidth), secondaryColumn...))
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

func vehicleBlocksWidth(card session.VehicleCard, sessionStyle, careerStyle, labelStyle, containerStyle common.Style) (map[string]float64, float64) {
	presetBlockWidth := make(map[string]float64)
	var contentWidth, maxBlockWidth float64
	for _, block := range card.Blocks {
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
