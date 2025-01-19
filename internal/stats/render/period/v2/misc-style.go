package period

import (
	"image/color"

	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/facepaint/style"
)

var (
	clanTagBackgroundColor = color.NRGBA{60, 60, 60, 100}
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

	GrowHorizontal: true,
	Gap:            20,
}

var styledPlayerNameCard = style.Style{
	Direction:      style.DirectionHorizontal,
	AlignItems:     style.AlignItemsCenter,
	JustifyContent: style.JustifyContentSpaceAround,

	GrowHorizontal: true,
	// GrowVertical:   true,
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
