package session

import (
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/facepaint/style"
)

const (
	debugVehicleCards = false

	vehicleIconSizeWN8 = 20.0
)

type vehicleCardStyle struct {
	card         style.StyleOptions
	titleWrapper style.StyleOptions
	titleText    func() style.StyleOptions

	stats style.StyleOptions
	value func(float64) *style.Style
}

var styledVehicleLegendPillWrapper = style.NewStyle(style.Parent(style.Style{
	Direction:      style.DirectionHorizontal,
	JustifyContent: style.JustifyContentSpaceAround,
	Gap:            5,
}))

func styledVehicleLegendPill() *style.Style {
	return &style.Style{
		Debug: debugVehicleCards,

		Color: common.TextAlt,
		Font:  common.FontSmall(),

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

	titleWrapper: style.NewStyle(style.Parent(style.Style{
		Debug: debugVehicleCards,

		GrowHorizontal: true,
		Gap:            10,
		JustifyContent: style.JustifyContentSpaceBetween,
	})),
	titleText: func() style.StyleOptions {
		return style.NewStyle(style.Parent(style.Style{
			Color: common.TextSecondary,
			Font:  common.FontMedium(),
		}))
	},

	stats: style.NewStyle(style.Parent(style.Style{
		Debug: debugVehicleCards,

		Direction:      style.DirectionHorizontal,
		JustifyContent: style.JustifyContentSpaceBetween,
		GrowHorizontal: true,
		Gap:            10,
	})),
	value: func(width float64) *style.Style {
		return &style.Style{
			Width:          width,
			Color:          common.TextPrimary,
			Font:           common.FontLarge(),
			GrowHorizontal: true,
			JustifyContent: style.JustifyContentCenter,
		}
	},
}
