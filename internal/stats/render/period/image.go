package period

import (
	"image"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare/period"
	"github.com/cufee/aftermath/internal/stats/render"
)

func CardsToImage(stats fetch.AccountStatsOverPeriod, cards period.Cards, subs []database.UserSubscription, opts ...render.Option) (image.Image, error) {
	o := render.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	segments, err := generateCards(stats, cards, subs, o)
	if err != nil {
		return nil, err
	}

	return segments.Render(opts...)
}
