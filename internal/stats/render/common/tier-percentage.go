package common

import (
	"fmt"
	"image/color"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/frame"
)

func NewTierPercentageCard(style Style, vehicles map[string]frame.VehicleStatsFrame, glossary map[int]database.GlossaryVehicle) Block {
	var blocks []Block
	var elements int = 10

	backgroundSharePrimary := DefaultCardColor
	backgroundShareSecondary := color.RGBA{120, 120, 120, 120}

	for i := range elements {
		shade := backgroundSharePrimary
		if i%2 == 0 {
			shade = backgroundShareSecondary
		}

		blocks = append(blocks, NewBlocksContent(Style{
			BackgroundColor: shade,
			Width:           style.Width / float64(elements),
			JustifyContent:  JustifyContentCenter,
		}, NewTextContent(Style{Font: &FontMedium, FontColor: TextPrimary}, fmt.Sprint(i))))
	}

	return NewBlocksContent(style, blocks...)

}
