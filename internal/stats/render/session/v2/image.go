package session

import (
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"golang.org/x/sync/errgroup"
)

type vehicleWN8 struct {
	id      string
	wn8     frame.Value
	sortKey int
}

// CardsToImages renders one PNG per page when the session has many vehicles.
func CardsToImages(session, career fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts ...common.Option) ([]image.Image, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	pages, err := generateSessionPages(session, career, cards, subs, o)
	if err != nil {
		return nil, err
	}

	out := make([]image.Image, len(pages))
	var g errgroup.Group
	for i, page := range pages {
		g.Go(func() error {
			img, err := page.Render()
			if err != nil {
				return err
			}
			out[i] = img
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return out, nil
}
