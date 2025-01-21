package sakura

import (
	"github.com/cufee/facepaint/style"
)

type nameStyleType struct{}

var nameStyle nameStyleType

func (nameStyleType) container() style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		JustifyContent: style.JustifyContentCenter,
		AlignItems:     style.AlignItemsCenter,
		Gap:            20,

		BackgroundColor: cardBackgroundColor,
		BlurBackground:  cardBackgroundBlur,
	}),
		style.SetPaddingX(20),
		style.SetPaddingY(10),
		style.SetBorderRadius(25),
	)
}

func (nameStyleType) nickname() style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		Font:  FontLarge(),
		Color: textPrimary,
	}))
}
func (nameStyleType) nicknameContainer() style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		//
	}))
}

func (nameStyleType) tag() style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		Font:  FontXL(),
		Color: textSecondary,
	}),
	)
}
func (nameStyleType) tagContainer() style.StyleOptions {
	return style.NewStyle(style.Parent(style.Style{
		//
	}))
}
