package common

import (
	"image"

	"github.com/cufee/facepaint/style"
)

type Theme struct {
	// Frame appearance (outermost container wrapping background + cards + footer)
	Frame style.StyleOptions
	// Card appearance (BackgroundColor, BlurBackground, BorderRadius)
	Card style.StyleOptions
	// Clan tag pill appearance
	ClanTag style.StyleOptions
	// Footer pill appearance
	Footer style.StyleOptions

	// Text styles (Color + Font per tier)
	TextPrimary   func() style.StyleOptions
	TextSecondary func() style.StyleOptions
	TextAlt       func() style.StyleOptions

	// Optional background image bundled with the theme.
	// Used as the default background when no explicit background is provided.
	Background image.Image

	// BackgroundOverlay is rendered behind cards, on top of the background image.
	// seed is derived from the account ID for deterministic patterns.
	BackgroundOverlay func(bounds image.Rectangle, seed int) image.Image
	// ForegroundOverlay is composited on the final rendered image, on top of everything.
	// seed is derived from the account ID for deterministic patterns.
	ForegroundOverlay func(rendered image.Image, frame image.Rectangle, seed int) image.Image
}

func DefaultTheme() Theme {
	return Theme{
		Frame: style.NewStyle(),
		Card: style.NewStyle(
			style.SetBorderRadius(BorderRadiusLG),
			func(s *style.Style) {
				s.BackgroundColor = DefaultCardColor
				s.BlurBackground = 20.0
			},
		),
		ClanTag: style.NewStyle(
			style.SetBorderRadius(BorderRadiusMD),
			func(s *style.Style) {
				s.BackgroundColor = ClanTagBackgroundColor
			},
		),
		Footer: style.NewStyle(
			style.SetBorderRadius(BorderRadiusSM),
			func(s *style.Style) {
				s.BackgroundColor = DefaultCardColor
				s.Color = TextAlt
				s.Font = FontSmall()
			},
		),
		TextPrimary: func() style.StyleOptions {
			return style.NewStyle(func(s *style.Style) {
				s.Color = TextPrimary
			})
		},
		TextSecondary: func() style.StyleOptions {
			return style.NewStyle(func(s *style.Style) {
				s.Color = TextSecondary
			})
		},
		TextAlt: func() style.StyleOptions {
			return style.NewStyle(func(s *style.Style) {
				s.Color = TextAlt
			})
		},
	}
}

// ApplyTheme merges a layout style with theme appearance options.
// The layout is used as the base, then theme options override appearance fields on top.
func ApplyTheme(layout style.Style, appearance style.StyleOptions) style.Style {
	return style.NewStyle(style.Parent(layout)).Chain(appearance.Spread()...).Computed()
}
