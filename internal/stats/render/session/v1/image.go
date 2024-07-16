package session

import (
	"image"
	"image/color"
	"slices"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

type vehicleWN8 struct {
	id      string
	wn8     frame.Value
	sortKey int
}

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
		var values []vehicleWN8
		for _, vehicle := range session.RegularBattles.Vehicles {
			if wn8 := vehicle.WN8(); !frame.InvalidValue.Equals(wn8) {
				values = append(values, vehicleWN8{vehicle.VehicleID, wn8, int(vehicle.LastBattleTime.Unix())})
			}
		}
		slices.SortFunc(values, func(a, b vehicleWN8) int { return b.sortKey - a.sortKey })
		if len(values) >= 10 {
			values = values[:9]
		}

		var accentColors []color.Color
		for _, value := range values {
			c := common.GetWN8Colors(value.wn8.Float()).Background
			if _, _, _, a := c.RGBA(); a > 0 {
				accentColors = append(accentColors, c)
			}
		}

		patternSeed, _ := strconv.Atoi(session.Account.ID)
		if patternSeed == 0 {
			patternSeed = int(time.Now().Unix())
		}
		o.Background = common.AddDefaultBrandedOverlay(o.Background, accentColors, patternSeed)
	}

	return segments.Render(func(opt *common.Options) { opt.Background = o.Background })
}
