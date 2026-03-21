package spring2026

import (
	"image/color"

	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/facepaint/style"
)

var (
	cardColor       = color.NRGBA{30, 10, 20, 180}
	clanTagColor    = color.NRGBA{60, 25, 40, 120}
	textPrimary     = color.NRGBA{255, 240, 245, 255}
	textSecondary   = color.NRGBA{220, 190, 200, 255}
	textAlt         = color.NRGBA{170, 140, 155, 255}
	footerCardColor = color.NRGBA{30, 10, 20, 180}
)

func Theme() common.Theme {
	noBlur := 0.0
	return common.Theme{
		Background:     backgroundImage,
		BackgroundBlur: &noBlur,
		Frame: style.NewStyle(func(s *style.Style) {
			s.PaddingLeft = 15
			s.PaddingRight = 15
			s.PaddingTop = 15
			s.PaddingBottom = 15
		}),
		Card: style.NewStyle(
			style.SetBorderRadius(common.BorderRadiusLG),
			func(s *style.Style) {
				s.BackgroundColor = cardColor
				s.BlurBackground = 10.0
			},
		),
		ClanTag: style.NewStyle(
			style.SetBorderRadius(common.BorderRadiusMD),
			func(s *style.Style) {
				s.BackgroundColor = clanTagColor
			},
		),
		Footer: style.NewStyle(
			style.SetBorderRadius(common.BorderRadiusSM),
			func(s *style.Style) {
				s.BackgroundColor = footerCardColor
				s.Color = textAlt
				s.Font = common.FontSmall()
			},
		),
		TextPrimary: func() style.StyleOptions {
			return style.NewStyle(func(s *style.Style) {
				s.Color = textPrimary
			})
		},
		TextSecondary: func() style.StyleOptions {
			return style.NewStyle(func(s *style.Style) {
				s.Color = textSecondary
			})
		},
		TextAlt: func() style.StyleOptions {
			return style.NewStyle(func(s *style.Style) {
				s.Color = textAlt
			})
		},
		BackgroundOverlay: makeBackgroundOverlay(),
		ForegroundOverlay: makeForegroundOverlay(processedPetals),
	}
}
