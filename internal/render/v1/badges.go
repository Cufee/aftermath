package render

import (
	"slices"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/assets"
)

type subscriptionPillStyle struct {
	Text      Style
	Icon      Style
	Container Style
}

type subscriptionHeader struct {
	Name  string
	Icon  string
	Style subscriptionPillStyle
}

func (sub subscriptionHeader) Block() (Block, error) {
	if tierImage, ok := assets.GetLoadedImage(sub.Icon); ok {
		content := []Block{NewImageContent(sub.Style.Icon, tierImage)}
		if sub.Name != "" {
			content = append(content, NewTextContent(sub.Style.Text, sub.Name))
		}
		return NewBlocksContent(sub.Style.Container, content...), nil
	}
	return Block{}, errors.New("tier icon not found")
}

func SubscriptionsBadges(subscriptions []models.UserSubscription) ([]Block, error) {
	slices.SortFunc(subscriptions, func(i, j models.UserSubscription) int {
		return subscriptionWeight[j.Type] - subscriptionWeight[i.Type]
	})

	var badges []Block
	// Moderator role group
	for _, subscription := range subscriptions {
		var header *subscriptionHeader
		switch subscription.Type {
		case models.SubscriptionTypeDeveloper:
			header = subscriptionDeveloper
		case models.SubscriptionTypeServerModerator:
			header = subscriptionServerModerator
		case models.SubscriptionTypeContentModerator:
			header = subscriptionContentModerator
		}

		if header != nil {
			block, err := header.Block()
			if err != nil {
				return nil, err
			}
			badges = append(badges, block)
			break
		}
	}
	// Community role group
	for _, subscription := range subscriptions {
		var header *subscriptionHeader
		switch subscription.Type {
		case models.SubscriptionTypeContentTranslator:
			header = subscriptionTranslator
		case models.SubscriptionTypeThumbsCounter:
			if count, _ := subscription.Meta["count"].(float64); count > 0 {
				header = subscriptionThumbsUp(count)
			}
		}

		if header != nil {
			block, err := header.Block()
			if err != nil {
				return nil, err
			}
			badges = append(badges, block)
			break
		}
	}
	// Paid member group
	for _, subscription := range subscriptions {
		var header *subscriptionHeader
		switch subscription.Type {
		case models.SubscriptionTypePro:
			header = userSubscriptionPro
		case models.SubscriptionTypePlus:
			header = userSubscriptionPlus
		case models.SubscriptionTypeServerBooster:
			header = subscriptionServerBooster
		case models.SubscriptionTypeSupporter:
			header = userSubscriptionSupporter
		}

		if header != nil {
			block, err := header.Block()
			if err != nil {
				return nil, err
			}
			badges = append(badges, block)
			break
		}
	}

	return badges, nil
}

func ClanSubscriptionsBadges(subscriptions []models.UserSubscription) *subscriptionHeader {
	var headers []*subscriptionHeader

	for _, subscription := range subscriptions {
		switch subscription.Type {
		case models.SubscriptionTypeProClan:
			headers = append(headers, clanSubscriptionPro)
		case models.SubscriptionTypeVerifiedClan:
			headers = append(headers, clanSubscriptionVerified)
		}
	}

	if len(headers) > 0 {
		return headers[0]
	}

	return nil
}
