package common

import "github.com/cufee/am-wg-proxy-next/v2/types"

func NewFooterCard(text string) Block {
	backgroundColor := DefaultCardColorNoAlpha
	backgroundColor.A = 120
	return NewBlocksContent(Style{
		JustifyContent:  JustifyContentCenter,
		AlignItems:      AlignItemsCenter,
		Direction:       DirectionVertical,
		PaddingX:        12.5,
		PaddingY:        5,
		BackgroundColor: backgroundColor,
		BorderRadius:    BorderRadiusSM,
		// Debug:           true,
	}, NewTextContent(Style{Font: FontSmall(), FontColor: TextSecondary}, text))
}

func RealmLabel(realm types.Realm) string {
	switch realm {
	case types.RealmNorthAmerica:
		return "North America"
	case types.RealmEurope:
		return "Europe"
	case types.RealmAsia:
		return "Asia"
	case types.RealmChina:
		return "China"
	case types.RealmRussia:
		return "Russia"
	}
	return "Classified"
}
