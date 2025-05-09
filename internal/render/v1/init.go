package render

import (
	"fmt"
	"image/color"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/assets"
)

var DiscordBackgroundColor = color.NRGBA{49, 51, 56, 255}

var (
	TextPrimary   = color.NRGBA{255, 255, 255, 255}
	TextSecondary = color.NRGBA{204, 204, 204, 255}
	TextAlt       = color.NRGBA{150, 150, 150, 255}

	TextSubscriptionPlus    = color.NRGBA{72, 167, 250, 255}
	TextSubscriptionPremium = color.NRGBA{255, 223, 0, 255}

	DefaultCardColor        = color.NRGBA{10, 10, 10, 150}
	DefaultCardColorNoAlpha = color.NRGBA{10, 10, 10, 255}
	ClanTagBackgroundColor  = color.NRGBA{10, 10, 10, 100}

	ColorAftermathRed    = color.NRGBA{255, 0, 120, 255}
	ColorAftermathBlue   = color.NRGBA{72, 167, 250, 255}
	ColorAftermathYellow = color.NRGBA{255, 223, 0, 255}

	BorderRadiusXL = 30.0
	BorderRadiusLG = 25.0
	BorderRadiusMD = 20.0
	BorderRadiusSM = 15.0
	BorderRadiusXS = 10.0
)

var (
	subscriptionWeight = map[models.SubscriptionType]int{
		models.SubscriptionTypeDeveloper: 999,
		// Moderators
		models.SubscriptionTypeServerModerator:  99,
		models.SubscriptionTypeContentModerator: 98,
		// Paid
		models.SubscriptionTypePro:     89,
		models.SubscriptionTypeProClan: 88,
		models.SubscriptionTypePlus:    79,
		//
		models.SubscriptionTypeSupporter:     29,
		models.SubscriptionTypeServerBooster: 28,
		//
		models.SubscriptionTypeVerifiedClan: 19,
	}
)

var (
	userSubscriptionSupporter, userSubscriptionPlus, userSubscriptionPro, clanSubscriptionVerified, clanSubscriptionPro, subscriptionDeveloper, subscriptionServerModerator, subscriptionContentModerator, subscriptionServerBooster, subscriptionTranslator *subscriptionHeader
)

var defaultFont []byte
var fonts = make(map[float64]*Font)

func Font2XL() *Font {
	return getFont(36)
}
func FontXL() *Font {
	return getFont(32)
}
func FontLarge() *Font {
	return getFont(24)
}
func FontMedium() *Font {
	return getFont(18)
}
func FontSmall() *Font {
	return getFont(14)
}

func getFont(size float64) *Font {
	if fonts[size] == nil {
		fonts[size] = NewFont(defaultFont, size)
	}
	return fonts[size]
}

func InitLoadedAssets() error {
	fontData, ok := assets.GetLoadedFontFace("default")
	if !ok {
		return errors.New("default font not found")
	}
	defaultFont = fontData

	// cache all fonts
	Font2XL()
	FontXL()
	FontLarge()
	FontMedium()
	FontSmall()

	// Personal
	userSubscriptionSupporter = &subscriptionHeader{
		Name: "Supporter",
		Icon: "images/icons/fire",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColorNoAlpha, BorderRadius: 15, PaddingX: 7, PaddingY: 5, Height: 32},
			Icon:      Style{Width: 16, Height: 16, BackgroundColor: TextSubscriptionPlus},
			Text:      Style{Font: FontSmall(), FontColor: TextSecondary, PaddingX: 5},
		},
	}
	userSubscriptionPlus = &subscriptionHeader{
		Name: "Aftermath+",
		Icon: "images/icons/star",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColorNoAlpha, BorderRadius: 15, PaddingX: 5, PaddingY: 5, Height: 32},
			Icon:      Style{Width: 24, Height: 24, BackgroundColor: TextSubscriptionPlus},
			Text:      Style{Font: FontSmall(), FontColor: TextSecondary, PaddingX: 5},
		},
	}
	userSubscriptionPro = &subscriptionHeader{
		Name: "Aftermath Pro",
		Icon: "images/icons/star",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColorNoAlpha, BorderRadius: 15, PaddingX: 5, PaddingY: 5, Height: 32},
			Icon:      Style{Width: 24, Height: 24, BackgroundColor: TextSubscriptionPremium},
			Text:      Style{Font: FontSmall(), FontColor: TextSecondary, PaddingX: 5},
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
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: color.NRGBA{64, 32, 128, 180}, BorderRadius: 15, PaddingX: 6, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20, BackgroundColor: TextPrimary},
			Text:      Style{Font: FontSmall(), FontColor: TextPrimary, PaddingX: 5},
		},
	}
	subscriptionServerModerator = &subscriptionHeader{
		Name: "Community Moderator",
		Icon: "images/icons/logo-128",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColorNoAlpha, BorderRadius: 15, PaddingX: 7, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20},
			Text:      Style{Font: FontSmall(), FontColor: TextSecondary, PaddingX: 2},
		},
	}
	subscriptionContentModerator = &subscriptionHeader{
		Name: "Moderator",
		Icon: "images/icons/logo-128",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColorNoAlpha, BorderRadius: 15, PaddingX: 7, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20},
			Text:      Style{Font: FontSmall(), FontColor: TextSecondary, PaddingX: 2},
		},
	}
	subscriptionServerBooster = &subscriptionHeader{
		Name: "Booster",
		Icon: "images/icons/discord-booster",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColorNoAlpha, BorderRadius: 15, PaddingX: 10, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20},
			Text:      Style{Font: FontSmall(), FontColor: TextSecondary},
		},
	}
	subscriptionTranslator = &subscriptionHeader{
		Name: "Translator",
		Icon: "images/icons/translator",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColorNoAlpha, BorderRadius: 15, PaddingX: 10, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20, BackgroundColor: TextPrimary},
			Text:      Style{Font: FontSmall(), FontColor: TextSecondary},
		},
	}
	return nil
}

func subscriptionThumbsUp(count float64) *subscriptionHeader {
	return &subscriptionHeader{
		Name: fmt.Sprintf("%0.0f", count),
		Icon: "images/icons/thumbs-up",
		Style: subscriptionPillStyle{
			Container: Style{Direction: DirectionHorizontal, AlignItems: AlignItemsCenter, BackgroundColor: DefaultCardColorNoAlpha, BorderRadius: 15, PaddingX: 10, PaddingY: 5, Gap: 5, Height: 32},
			Icon:      Style{Width: 20, Height: 20, BackgroundColor: TextPrimary},
			Text:      Style{Font: FontSmall(), FontColor: TextSecondary},
		}}
}
