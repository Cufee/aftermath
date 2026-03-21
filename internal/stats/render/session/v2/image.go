package session

import (
	"image"
	"strconv"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
)

func CardsToImage(session, career fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts ...common.Option) (image.Image, error) {
	o := common.DefaultOptions()
	for _, apply := range opts {
		apply(&o)
	}

	block, err := generateCards(session, career, cards, subs, o)
	if err != nil {
		return nil, err
	}

	rendered, err := block.Render()
	if err != nil {
		return nil, err
	}

	if o.Theme.ForegroundOverlay != nil {
		seed, _ := strconv.Atoi(career.Account.ID)
		rendered = o.Theme.ForegroundOverlay(rendered, rendered.Bounds(), seed)
	}

	return rendered, nil
}
