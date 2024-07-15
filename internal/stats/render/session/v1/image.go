package session

import (
	"image"
	"image/color"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/fogleman/gg"
)

func CardsToImage(session, career fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts ...common.Option) (image.Image, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	segments, err := cardsToSegments(session, career, cards, subs, o)
	if err != nil {
		return nil, err
	}

	if o.Background != nil {
		var accentColors []color.Color
		for _, vehicle := range session.RegularBattles.Vehicles {
			c := common.GetWN8Colors(vehicle.WN8().Float()).Background
			if _, _, _, a := c.RGBA(); a > 0 {
				accentColors = append(accentColors, c)
			}
		}
		if len(accentColors) < 1 {
			accentColors = common.DefaultLogoColorOptions
		}

		patternSeed, _ := strconv.Atoi(session.Account.ID)
		if patternSeed == 0 {
			patternSeed = int(time.Now().Unix())
		}
		ctx := gg.NewContextForImage(o.Background)
		overlay := common.NewBrandedBackground(ctx.Width(), ctx.Height(), accentColors, patternSeed)
		ctx.DrawImage(overlay, 0, 0)
		o.Background = ctx.Image()
	}

	return segments.Render(func(opt *common.Options) { opt.Background = o.Background })
}
