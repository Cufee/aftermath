package session

import (
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
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

	// Generate cards
	cardsBlock, err := generateCards(session, career, cards, subs, o)
	if err != nil {
		return nil, err
	}

	// Render
	return cardsBlock.Render()

}
