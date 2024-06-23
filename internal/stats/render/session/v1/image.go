package session

import (
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/internal/stats/render"
)

func CardsToImage(session, career fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts ...render.Option) (image.Image, error) {
	o := render.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	segments, err := cardsToSegments(session, career, cards, subs, o)
	if err != nil {
		return nil, err
	}

	return segments.Render(opts...)
}
