package sakura

import (
	"image/color"
	"math"

	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/facepaint/style"
)

var (
	textPrimary   = color.NRGBA{255, 231, 222, 255}
	textSecondary = color.NRGBA{233, 177, 205, 255}
	textAlt       = color.NRGBA{195, 130, 158, 255}

	cardBackgroundBlur  = 20.0
	cardBackgroundColor = color.NRGBA{66, 13, 33, 120}
)

var styledBackground = style.NewStyle(
	style.SetBorderRadius(common.BorderRadius2XL),
	style.SetBlur(1),
	style.SetPosition(style.PositionAbsolute),
	style.SetZIndex(math.MinInt),
)

type footerStyle struct {
	container style.StyleOptions
	pill      func() style.StyleOptions
}

var styledFooter = footerStyle{
	container: style.NewStyle(style.Parent(style.Style{
		Direction:  style.DirectionHorizontal,
		AlignItems: style.AlignItemsCenter,
		Gap:        5,
	})),
	pill: func() style.StyleOptions {
		return style.NewStyle(style.Parent(style.Style{
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
		}))
	},
}

var styledOuterFrame = style.NewStyle(style.Parent(style.Style{
	Debug: false,

	Direction:  style.DirectionVertical,
	AlignItems: style.AlignItemsCenter,
	Gap:        5,
}))

var styledInnerFrame = style.NewStyle(style.Parent(style.Style{
	Debug: false,

	Direction:  style.DirectionVertical,
	AlignItems: style.AlignItemsCenter,
	Gap:        10,
}))

var styledContentFrame = style.NewStyle(style.Parent(style.Style{
	Debug: false,

	Direction: style.DirectionVertical,
	Gap:       10,

	PaddingLeft:   30,
	PaddingRight:  30,
	PaddingTop:    30,
	PaddingBottom: 30,
}))
