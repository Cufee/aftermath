package sakura

import (
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
)

func CardsToImage(session, career fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts ...common.Option) (image.Image, error) {
	err := lazyLoadAssets()
	if err != nil {
		return nil, err
	}

	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	// generate cards
	block, err := generateCards(session, career, cards, subs, o)
	if err != nil {
		return nil, err
	}

	// render
	return block.Render()
}
