package period

import (
	"image"
	"image/color"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/disintegration/imaging"
)

func CardsToImage(stats fetch.AccountStatsOverPeriod, cards period.Cards, subs []models.UserSubscription, opts ...common.Option) (image.Image, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	segments, err := generateCards(stats, cards, subs, o)
	if err != nil {
		return nil, err
	}

	if o.Background != nil {
		var accentColors []color.Color
		for _, vehicle := range stats.RegularBattles.Vehicles {
			c := common.GetWN8Colors(vehicle.WN8().Float()).Background
			if _, _, _, a := c.RGBA(); a > 0 {
				accentColors = append(accentColors, c)
			}
		}
		if len(accentColors) < 1 {
			accentColors = common.DefaultLogoColorOptions
		}

		patternSeed, _ := strconv.Atoi(stats.Account.ID)
		if patternSeed == 0 {
			patternSeed = int(time.Now().Unix())
		}
		overlay := common.DefaultBrandedOverlay(accentColors, patternSeed)
		o.Background = imaging.OverlayCenter(o.Background, imaging.Fill(overlay, o.Background.Bounds().Dx(), o.Background.Bounds().Dy(), imaging.Center, imaging.Linear), 100)
	}

	return segments.Render(func(op *common.Options) { op.Background = o.Background })
}
