package period

import (
	"image"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare/period"
	"github.com/cufee/aftermath/internal/stats/render/common"
)

func CardsToImage(stats fetch.AccountStatsOverPeriod, cards period.Cards, subs []database.UserSubscription, opts ...Option) (image.Image, error) {
	o := options{}
	for _, apply := range opts {
		apply(&o)
	}

	renderedCards, err := generateCards(stats, cards, subs, o)
	if err != nil {
		return nil, err
	}

	allCards := common.NewBlocksContent(
		common.Style{
			Direction:  common.DirectionVertical,
			AlignItems: common.AlignItemsCenter,
			PaddingX:   20,
			PaddingY:   20,
			Gap:        10,
			// Debug:      true,
		}, renderedCards...)

	cardsImage, err := allCards.Render()
	if err != nil {
		return nil, err
	}

	return cardsImage, nil
}
