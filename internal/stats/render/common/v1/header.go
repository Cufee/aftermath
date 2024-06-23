package common

import (
	"github.com/cufee/aftermath/internal/database/models"
)

func NewHeaderCard(width float64, subscriptions []models.UserSubscription, promoText []string) (Block, bool) {
	var cards []Block

	var addPromoText = len(promoText) > 0
	for _, sub := range subscriptions {
		switch sub.Type {
		case models.SubscriptionTypePro, models.SubscriptionTypePlus, models.SubscriptionTypeDeveloper:
			addPromoText = false
		}
		if !addPromoText {
			break
		}
	}

	if addPromoText {
		// Users without a subscription get promo text
		var textBlocks []Block
		for _, text := range promoText {
			textBlocks = append(textBlocks, NewTextContent(Style{Font: &FontMedium, FontColor: TextPrimary}, text))
		}
		cards = append(cards, NewBlocksContent(Style{
			Direction:  DirectionVertical,
			AlignItems: AlignItemsCenter,
		},
			textBlocks...,
		))
	}

	// User Subscription Badge and promo text
	if badges, _ := SubscriptionsBadges(subscriptions); len(badges) > 0 {
		cards = append(cards, NewBlocksContent(Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, Gap: 10},
			badges...,
		))
	}

	if len(cards) < 1 {
		return Block{}, false
	}

	return NewBlocksContent(Style{Direction: DirectionVertical, AlignItems: AlignItemsCenter, JustifyContent: JustifyContentCenter, Gap: 10, Width: width}, cards...), true
}
