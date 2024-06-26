package common

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
