package replay

import (
	"image"
	"image/color"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/disintegration/imaging"
)

func CardsToImage(replay fetch.Replay, cards replay.Cards, opts ...common.Option) (image.Image, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	segments, err := generateCards(replay, cards, o)
	if err != nil {
		return nil, err
	}

	if o.Background != nil {
		var accentColors []color.Color
		var wn8Values []float32
		for _, player := range append(replay.Teams.Allies, replay.Teams.Enemies...) {
			wn8Values = append(wn8Values, player.Performance.WN8().Float())
		}
		for _, value := range wn8Values {
			c := common.GetWN8Colors(value).Background
			if _, _, _, a := c.RGBA(); a > 0 {
				accentColors = append(accentColors, c)
			}
		}
		if len(accentColors) < 1 {
			accentColors = common.DefaultLogoColorOptions
		}

		patternSeed, _ := strconv.Atoi(replay.Protagonist.ID)
		if patternSeed == 0 {
			patternSeed = int(time.Now().Unix())
		}
		overlay := common.DefaultBrandedOverlay(accentColors, patternSeed)
		o.Background = imaging.PasteCenter(o.Background, imaging.Fill(overlay, o.Background.Bounds().Dx(), o.Background.Bounds().Dy(), imaging.Center, imaging.Linear))
	}

	return segments.Render(func(op *common.Options) { op.Background = o.Background })
}
