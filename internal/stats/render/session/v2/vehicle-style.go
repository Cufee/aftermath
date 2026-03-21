package session

import (
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/facepaint/style"
)

const (
	debugVehicleCards = false

	vehicleIconSizeWN8 = 16.0
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

func newVehicleLegendPillWrapper() style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		Direction:      style.DirectionHorizontal,
		JustifyContent: style.JustifyContentSpaceBetween,
		Gap:            5,
	}))
}

func newVehicleLegendPill(width float64) style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		Debug: debugVehicleCards,
		Width: width,

		JustifyContent: style.JustifyContentCenter,
	}))
}

func newVehicleLegendPillText(theme common.Theme) *style.Style {
	pillCard := theme.Card.Chain(style.SetBorderRadius(common.BorderRadiusSM))
	s := common.ApplyTheme(style.Style{
		Font: common.FontSmall(),

		PaddingLeft:   15,
		PaddingRight:  15,
		PaddingTop:    5,
		PaddingBottom: 5,
	}, pillCard)
	themed := common.ApplyTheme(s, theme.TextAlt())
	return &themed
}

func newVehicleCardStyle(theme common.Theme) vehicleCardStyle {
	return vehicleCardStyle{
		card: style.NewStyle(style.Parent(common.ApplyTheme(style.Style{
			Debug: debugVehicleCards,

			Direction: style.DirectionVertical,

			GrowHorizontal: true,
			Gap:            5,

			PaddingLeft:   common.CardPaddingX / 1.5,
			PaddingRight:  common.CardPaddingX / 1.5,
			PaddingTop:    common.CardPaddingY / 2,
			PaddingBottom: common.CardPaddingY / 2,
		}, theme.Card))),

		titleIconWrapper: style.NewStyle(style.Parent(style.Style{})),
		titleWrapper: style.NewStyle(style.Parent(style.Style{
			Debug:      debugVehicleCards,
			AlignItems: style.AlignItemsCenter,

			GrowHorizontal: true,
			Gap:            10,
		})),
		titleText: func() style.StyleOptions {
			return style.NewStyle(style.Parent(common.ApplyTheme(style.Style{
				Font:           common.FontMedium(),
				GrowHorizontal: true,
			}, theme.TextSecondary())))
		},

		stats: style.NewStyle(style.Parent(style.Style{
			Debug: debugVehicleCards,

			Direction:      style.DirectionHorizontal,
			JustifyContent: style.JustifyContentSpaceBetween,
			GrowHorizontal: true,
			Gap:            10,
		})),
		value: func() *style.Style {
			s := common.ApplyTheme(style.Style{
				Font:           common.FontLarge(),
				JustifyContent: style.JustifyContentCenter,
			}, theme.TextPrimary())
			return &s
		},
		valueWrapper: func(width float64) *style.Style {
			return &style.Style{
				Width:          width,
				JustifyContent: style.JustifyContentCenter,
				AlignItems:     style.AlignItemsCenter,
			}
		},
	}
}
