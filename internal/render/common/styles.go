package common

import "github.com/cufee/facepaint/style"

var (
	CardPaddingX = 35.0
	CardPaddingY = 30.0
)

// HighlightCardStyle defines the styles for a highlight/vehicle card with title and stats.
type HighlightCardStyle struct {
	Card         style.Style
	TitleWrapper style.Style
	TitleLabel   func() *style.Style
	TitleVehicle func() *style.Style
	StatsWrapper style.Style
	Stats        style.Style
	BlockValue   func() *style.Style
	BlockLabel   func() *style.Style
}

func NewHighlightCardStyle(theme Theme) HighlightCardStyle {
	return HighlightCardStyle{
		Card: ApplyTheme(style.Style{
			Direction:  style.DirectionHorizontal,
			AlignItems: style.AlignItemsCenter,

			GrowHorizontal: true,
			Gap:            20,

			PaddingLeft:   CardPaddingX / 1.5,
			PaddingRight:  CardPaddingX / 1.5,
			PaddingTop:    CardPaddingY / 1.5,
			PaddingBottom: CardPaddingY / 1.5,
		}, theme.Card),
		TitleWrapper: style.Style{
			GrowHorizontal: true,
			Direction:      style.DirectionVertical,
		},
		TitleLabel: func() *style.Style {
			s := ApplyTheme(style.Style{Font: FontSmall()}, theme.TextSecondary())
			return &s
		},
		TitleVehicle: func() *style.Style {
			s := ApplyTheme(style.Style{Font: FontMedium()}, theme.TextPrimary())
			return &s
		},
		Stats: style.Style{
			Direction:      style.DirectionVertical,
			AlignItems:     style.AlignItemsCenter,
			JustifyContent: style.JustifyContentCenter,
		},
		StatsWrapper: style.Style{
			Direction:  style.DirectionHorizontal,
			AlignItems: style.AlignItemsCenter,
			Gap:        10,
		},
		BlockValue: func() *style.Style {
			s := ApplyTheme(style.Style{Font: FontMedium()}, theme.TextPrimary())
			return &s
		},
		BlockLabel: func() *style.Style {
			s := ApplyTheme(style.Style{Font: FontSmall()}, theme.TextAlt())
			return &s
		},
	}
}

// PlayerNameWrapperStyle returns the style for the player name header card, themed.
func PlayerNameWrapperStyle(theme Theme) style.Style {
	return ApplyTheme(style.Style{
		Direction:  style.DirectionHorizontal,
		AlignItems: style.AlignItemsCenter,

		PaddingLeft:   5,
		PaddingRight:  5,
		PaddingTop:    5,
		PaddingBottom: 5,

		Height: 50,

		GrowHorizontal: true,
		Gap:            20,
	}, theme.Card)
}

// ClanTagCardStyle returns the style for the clan tag pill, themed.
func ClanTagCardStyle(theme Theme) style.Style {
	return ApplyTheme(style.Style{
		Direction:      style.DirectionHorizontal,
		AlignItems:     style.AlignItemsCenter,
		JustifyContent: style.JustifyContentSpaceAround,

		GrowVertical: true,

		PaddingLeft:   12,
		PaddingRight:  12,
		PaddingTop:    10,
		PaddingBottom: 10,
	}, theme.ClanTag)
}

// FooterPillStyle returns the style for a footer text pill, themed.
func FooterPillStyle(theme Theme) style.Style {
	return ApplyTheme(style.Style{
		PaddingLeft:   10,
		PaddingRight:  10,
		PaddingTop:    5,
		PaddingBottom: 5,
	}, theme.Footer)
}

// PlayerNameTextStyle returns the text style for the player name.
func PlayerNameTextStyle(theme Theme) style.Style {
	return ApplyTheme(style.Style{Font: FontMedium()}, theme.TextPrimary())
}

// ClanTagTextStyle returns the text style for the clan tag.
func ClanTagTextStyle(theme Theme) style.Style {
	return ApplyTheme(style.Style{Font: FontSmall()}, theme.TextSecondary())
}

// FinalFrameStyle returns the style for the outermost frame, themed.
// A theme can add padding, background color, or border radius around the entire image.
func FinalFrameStyle(theme Theme) style.Style {
	return ApplyTheme(style.Style{
		Direction:  style.DirectionVertical,
		AlignItems: style.AlignItemsCenter,
		Gap:        5,
	}, theme.Frame)
}

// Pure layout styles that have no theme dependency.
var (
	PlayerNameCardLayout = style.Style{
		Direction:      style.DirectionHorizontal,
		AlignItems:     style.AlignItemsCenter,
		JustifyContent: style.JustifyContentSpaceAround,

		GrowHorizontal: true,
	}

	CardsBackgroundStyle = style.NewStyle(
		style.SetBorderRadius(BorderRadius2XL),
		style.SetBlur(DefaultBackgroundBlur),
		style.SetPosition(style.PositionAbsolute),
		style.SetZIndex(-99),
	)

	FooterWrapperLayout = style.Style{
		Direction:  style.DirectionHorizontal,
		AlignItems: style.AlignItemsCenter,
		Gap:        5,
	}
)
