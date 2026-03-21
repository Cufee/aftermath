package session

import "github.com/cufee/facepaint/style"

var styledCardsSection = style.Style{
	Direction:  style.DirectionVertical,
	AlignItems: style.AlignItemsCenter,
	Gap:        10,
}

var styledCardsSectionsWrapper = style.Style{
	Direction: style.DirectionHorizontal,
	Gap:       20,
}

var styledStatsFrame = style.Style{
	Direction: style.DirectionVertical,
	Gap:       10,

	PaddingLeft:   30,
	PaddingRight:  30,
	PaddingTop:    30,
	PaddingBottom: 30,
}
