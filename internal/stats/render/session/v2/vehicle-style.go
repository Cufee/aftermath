package session

import (
	"image/color"

	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/facepaint/style"
)

const (
	debugVehicleCards = false

	vehicleIconSizeWN8 = 14.0
)

var (
	iconBackgroundColorVehicle = color.NRGBA{40, 40, 40, 80}
)

type vehicleCardStyle struct {
	card             style.StyleOptions
	titleIconWrapper style.StyleOptions
	titleWrapper     style.StyleOptions
	titleText        func() style.StyleOptions

	stats        style.StyleOptions
	value        func() *style.Style
	valueWrapper func(float64) *style.Style
}

var styledVehicleLegendPillWrapper = style.NewStyle(style.Parent(style.Style{
	Direction:      style.DirectionHorizontal,
	JustifyContent: style.JustifyContentSpaceBetween,
	Gap:            5,
}))

func styledVehicleLegendPill(width float64) style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		Debug: debugVehicleCards,
		Width: width,

		JustifyContent: style.JustifyContentCenter,
	}))
}

func styledVehicleLegendPillText() *style.Style {
	return &style.Style{
		Color:           common.TextAlt,
		Font:            common.FontSmall(),
		BackgroundColor: common.DefaultCardColor,
		BlurBackground:  cardBackgroundBlur,

		BorderRadiusTopLeft:     common.BorderRadiusSM,
		BorderRadiusTopRight:    common.BorderRadiusSM,
		BorderRadiusBottomLeft:  common.BorderRadiusSM,
		BorderRadiusBottomRight: common.BorderRadiusSM,

		PaddingLeft:   15,
		PaddingRight:  15,
		PaddingTop:    5,
		PaddingBottom: 5,
	}
}

var styledVehicleCard = vehicleCardStyle{
	card: style.NewStyle(style.Parent(style.Style{
		Debug: debugVehicleCards,

		Direction: style.DirectionVertical,

		BackgroundColor: common.DefaultCardColor,
		BlurBackground:  cardBackgroundBlur,

		BorderRadiusTopLeft:     common.BorderRadiusLG,
		BorderRadiusTopRight:    common.BorderRadiusLG,
		BorderRadiusBottomLeft:  common.BorderRadiusLG,
		BorderRadiusBottomRight: common.BorderRadiusLG,

		GrowHorizontal: true,
		Gap:            5,

		PaddingLeft:   cardPaddingX / 1.5,
		PaddingRight:  cardPaddingX / 1.5,
		PaddingTop:    cardPaddingY / 2,
		PaddingBottom: cardPaddingY / 2,
	})),

	titleIconWrapper: style.NewStyle(style.Parent(style.Style{
		BorderRadiusTopLeft:     common.BorderRadiusXS,
		BorderRadiusTopRight:    common.BorderRadiusXS,
		BorderRadiusBottomLeft:  common.BorderRadiusXS,
		BorderRadiusBottomRight: common.BorderRadiusXS,

		BackgroundColor: iconBackgroundColorVehicle,

		BlurBackground: cardBackgroundBlur,

		PaddingLeft:   7,
		PaddingRight:  8,
		PaddingTop:    7,
		PaddingBottom: 7,
	})),
	titleWrapper: style.NewStyle(style.Parent(style.Style{
		Debug: debugVehicleCards,

		GrowHorizontal: true,
		Gap:            10,
	})),
	titleText: func() style.StyleOptions {
		return style.NewStyle(style.Parent(style.Style{
			Color:          common.TextSecondary,
			Font:           common.FontMedium(),
			GrowHorizontal: true,
		}))
	},

	stats: style.NewStyle(style.Parent(style.Style{
		Debug: debugVehicleCards,

		Direction:      style.DirectionHorizontal,
		JustifyContent: style.JustifyContentSpaceBetween,
		GrowHorizontal: true,
		Gap:            10,
	})),
	value: func() *style.Style {
		return &style.Style{
			Color:          common.TextPrimary,
			Font:           common.FontLarge(),
			JustifyContent: style.JustifyContentCenter,
		}
	},
	valueWrapper: func(width float64) *style.Style {
		return &style.Style{
			Width:          width,
			JustifyContent: style.JustifyContentCenter,
			AlignItems:     style.AlignItemsCenter,
		}
	},
}
