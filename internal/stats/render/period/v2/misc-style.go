package period

import (
	"image/color"

	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/facepaint/style"
)

var (
	clanTagBackgroundColor = color.NRGBA{40, 40, 40, 100}
)

func styledPlayerName() style.Style {
	return style.Style{
		Color: common.TextPrimary,
		Font:  common.FontMedium(),
	}
}

func styledPlayerClanTag() style.Style {
	return style.Style{
		Color: common.TextSecondary,
		Font:  common.FontSmall(),
	}
}

var styledPlayerNameWrapper = style.Style{
	Direction:  style.DirectionHorizontal,
	AlignItems: style.AlignItemsCenter,

	BackgroundColor: common.DefaultCardColor,

	BorderRadiusTopLeft:     common.BorderRadiusLG,
	BorderRadiusTopRight:    common.BorderRadiusLG,
	BorderRadiusBottomLeft:  common.BorderRadiusLG,
	BorderRadiusBottomRight: common.BorderRadiusLG,

	PaddingLeft:   5,
	PaddingRight:  5,
	PaddingTop:    5,
	PaddingBottom: 5,

	Height: 50,

	GrowHorizontal: true,
	Gap:            20,
}

var styledPlayerNameCard = style.Style{
	Direction:      style.DirectionHorizontal,
	AlignItems:     style.AlignItemsCenter,
	JustifyContent: style.JustifyContentSpaceAround,

	GrowHorizontal: true,
}

var styledPlayerClanTagCard = style.Style{
	Direction:      style.DirectionHorizontal,
	AlignItems:     style.AlignItemsCenter,
	JustifyContent: style.JustifyContentSpaceAround,

	BackgroundColor: clanTagBackgroundColor,

	BorderRadiusTopLeft:     common.BorderRadiusMD,
	BorderRadiusTopRight:    common.BorderRadiusMD,
	BorderRadiusBottomLeft:  common.BorderRadiusMD,
	BorderRadiusBottomRight: common.BorderRadiusMD,

	GrowVertical: true,

	PaddingLeft:   12,
	PaddingRight:  12,
	PaddingTop:    10,
	PaddingBottom: 10,
}

var styledCardsFrame = style.Style{
	Debug: false,

	Direction:  style.DirectionVertical,
	AlignItems: style.AlignItemsCenter,
	Gap:        10,

	GrowHorizontal: true,

	PaddingLeft:   20,
	PaddingRight:  20,
	PaddingTop:    20,
	PaddingBottom: 20,
}

var styledFinalFrame = style.Style{
	Debug: false,

	Direction:  style.DirectionVertical,
	AlignItems: style.AlignItemsCenter,
	Gap:        5,
}

var styledCardsBackground = style.NewStyle(
	style.SetBorderRadius(common.BorderRadius2XL),
	style.SetBlur(common.DefaultBackgroundBlur),
	style.SetPosition(style.PositionAbsolute),
	style.SetZIndex(-99),
)

var styledFooterWrapper = style.Style{
	Direction:  style.DirectionHorizontal,
	AlignItems: style.AlignItemsCenter,
	Gap:        5,
}

func styledFooterCard() style.Style {
	return style.Style{
		Font:  common.FontSmall(),
		Color: common.TextAlt,

		BackgroundColor: common.DefaultCardColor,

		BorderRadiusTopLeft:     common.BorderRadiusSM,
		BorderRadiusTopRight:    common.BorderRadiusSM,
		BorderRadiusBottomLeft:  common.BorderRadiusSM,
		BorderRadiusBottomRight: common.BorderRadiusSM,

		PaddingLeft:   10,
		PaddingRight:  10,
		PaddingTop:    5,
		PaddingBottom: 5,
	}
}
