package common

import (
	"errors"
	"image/color"
	"slices"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/render/assets"
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

var (
	subscriptionWeight = map[database.SubscriptionType]int{
		database.SubscriptionTypeDeveloper: 999,
		// Moderators
		database.SubscriptionTypeServerModerator:  99,
		database.SubscriptionTypeContentModerator: 98,
		// Paid
		database.SubscriptionTypePro:     89,
		database.SubscriptionTypeProClan: 88,
		database.SubscriptionTypePlus:    79,
		//
		database.SubscriptionTypeSupporter:     29,
		database.SubscriptionTypeServerBooster: 28,
		//
		database.SubscriptionTypeVerifiedClan: 19,
	}

	// Personal
	userSubscriptionSupporter = &subscriptionHeader{
		Name: "Supporter",
		Icon: "images/icons/fire",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColor, BorderRadius: 15, PaddingX: 7, PaddingY: 5, Height: 32},
			Icon:      Style{Width: 16, Height: 16, BackgroundColor: TextSubscriptionPlus},
			Text:      Style{Font: &FontSmall, FontColor: TextSecondary, PaddingX: 5},
		},
	}
	userSubscriptionPlus = &subscriptionHeader{
		Name: "Aftermath+",
		Icon: "images/icons/star",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColor, BorderRadius: 15, PaddingX: 5, PaddingY: 5, Height: 32},
			Icon:      Style{Width: 24, Height: 24, BackgroundColor: TextSubscriptionPlus},
			Text:      Style{Font: &FontSmall, FontColor: TextSecondary, PaddingX: 5},
		},
	}
	userSubscriptionPro = &subscriptionHeader{
		Name: "Aftermath Pro",
		Icon: "images/icons/star",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColor, BorderRadius: 15, PaddingX: 5, PaddingY: 5, Height: 32},
			Icon:      Style{Width: 24, Height: 24, BackgroundColor: TextSubscriptionPremium},
			Text:      Style{Font: &FontSmall, FontColor: TextSecondary, PaddingX: 5},
		},
	}
	// Clans
	clanSubscriptionVerified = &subscriptionHeader{
		Icon: "images/icons/verify",
		Style: subscriptionPillStyle{
			Icon:      Style{Width: 28, Height: 28, BackgroundColor: TextAlt},
			Container: Style{Direction: DirectionHorizontal},
		},
	}
	clanSubscriptionPro = &subscriptionHeader{
		Icon: "images/icons/star-multiple",
		Style: subscriptionPillStyle{
			Icon:      Style{Width: 28, Height: 28, BackgroundColor: TextAlt},
			Container: Style{Direction: DirectionHorizontal},
		},
	}

	// Community
	subscriptionDeveloper = &subscriptionHeader{
		Name: "Developer",
		Icon: "images/icons/github",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: color.RGBA{64, 32, 128, 180}, BorderRadius: 15, PaddingX: 6, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20, BackgroundColor: TextPrimary},
			Text:      Style{Font: &FontSmall, FontColor: TextPrimary, PaddingX: 5},
		},
	}
	subscriptionServerModerator = &subscriptionHeader{
		Name: "Community Moderator",
		Icon: "images/icons/logo-128",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColor, BorderRadius: 15, PaddingX: 7, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20},
			Text:      Style{Font: &FontSmall, FontColor: TextSecondary, PaddingX: 2},
		},
	}
	subscriptionContentModerator = &subscriptionHeader{
		Name: "Moderator",
		Icon: "images/icons/logo-128",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColor, BorderRadius: 15, PaddingX: 7, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20},
			Text:      Style{Font: &FontSmall, FontColor: TextSecondary, PaddingX: 2},
		},
	}
	subscriptionServerBooster = &subscriptionHeader{
		Name: "Booster",
		Icon: "images/icons/discord-booster",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColor, BorderRadius: 15, PaddingX: 10, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20},
			Text:      Style{Font: &FontSmall, FontColor: TextSecondary},
		},
	}
	subscriptionTranslator = &subscriptionHeader{
		Name: "Translator",
		Icon: "images/icons/translator",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColor, BorderRadius: 15, PaddingX: 10, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20, BackgroundColor: TextPrimary},
			Text:      Style{Font: &FontSmall, FontColor: TextSecondary},
		},
	}
)

func SubscriptionsBadges(subscriptions []database.UserSubscription) ([]Block, error) {
	slices.SortFunc(subscriptions, func(i, j database.UserSubscription) int {
		return subscriptionWeight[j.Type] - subscriptionWeight[i.Type]
	})

	var badges []Block
	for _, subscription := range subscriptions {
		var header *subscriptionHeader
		switch subscription.Type {
		case database.SubscriptionTypeDeveloper:
			header = subscriptionDeveloper
		case database.SubscriptionTypeServerModerator:
			header = subscriptionServerModerator
		case database.SubscriptionTypeContentModerator:
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
	for _, subscription := range subscriptions {
		var header *subscriptionHeader
		switch subscription.Type {
		case database.SubscriptionTypeContentTranslator:
			header = subscriptionTranslator
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
	for _, subscription := range subscriptions {
		var header *subscriptionHeader
		switch subscription.Type {
		case database.SubscriptionTypePro:
			header = userSubscriptionPro
		case database.SubscriptionTypePlus:
			header = userSubscriptionPlus
		case database.SubscriptionTypeServerBooster:
			header = subscriptionServerBooster
		case database.SubscriptionTypeSupporter:
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

func ClanSubscriptionsBadges(subscriptions []database.UserSubscription) *subscriptionHeader {
	var headers []*subscriptionHeader

	for _, subscription := range subscriptions {
		switch subscription.Type {
		case database.SubscriptionTypeProClan:
			headers = append(headers, clanSubscriptionPro)
		case database.SubscriptionTypeVerifiedClan:
			headers = append(headers, clanSubscriptionVerified)
		}
	}

	if len(headers) > 0 {
		return headers[0]
	}

	return nil
}
