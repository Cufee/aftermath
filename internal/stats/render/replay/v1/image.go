package replay

import (
	"image"
	"image/color"
	"slices"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/render/common"
	v1 "github.com/cufee/aftermath/internal/render/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/nao1215/imaging"
)

type playerWN8 struct {
	id      string
	wn8     frame.Value
	sortKey int
}

func CardsToImage(replay fetch.Replay, cards replay.Cards, opts ...common.Option) (image.Image, error) {
	segments, err := CardsToSegments(replay, cards, opts...)
	if err != nil {
		return nil, err
	}

	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	if o.Background != nil && !o.BackgroundIsCustom {
		var values []playerWN8
		for _, player := range append(replay.Teams.Allies, replay.Teams.Enemies...) {
			if wn8 := player.Performance.WN8(); !frame.InvalidValue.Equals(wn8) {
				values = append(values, playerWN8{player.VehicleID, wn8, int(player.TimeAlive.Float())})
			}
		}
		slices.SortFunc(values, func(a, b playerWN8) int { return b.sortKey - a.sortKey })

		var accentColors []color.Color
		c := common.GetWN8Colors(replay.Protagonist.Performance.WN8().Float()).Background
		if _, _, _, a := c.RGBA(); a > 0 {
			accentColors = append(accentColors, c)
		}

		patternSeed, _ := strconv.Atoi(replay.Protagonist.ID)
		if patternSeed == 0 {
			patternSeed = int(time.Now().Unix())
		}

		bounds, err := segments.ContentBounds(opts...)
		if err != nil {
			return nil, err
		}

		o.Background = imaging.Resize(o.Background, bounds.Dx(), bounds.Dy(), imaging.Gaussian)
		o.Background = common.AddDefaultBrandedOverlay(o.Background, accentColors, patternSeed, 0.35)
	}

	return segments.Render(func(op *common.Options) { op.Background = o.Background })
}

func CardsToSegments(replay fetch.Replay, cards replay.Cards, opts ...common.Option) (*v1.Segments, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	segments, err := generateCards(replay, cards, o.Printer)
	if err != nil {
		return nil, err
	}

	return &segments, nil
}
