package period

import (
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/period/v1"
)

func CardsToImage(stats fetch.AccountStatsOverPeriod, cards period.Cards, subs []models.UserSubscription, opts ...common.Option) (image.Image, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	cardsBlock, err := generateCards(stats, cards, subs, o)
	if err != nil {
		return nil, err
	}

	return cardsBlock.Render()
}
