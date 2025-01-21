package special

import (
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"github.com/cufee/aftermath/internal/stats/render/session/special/sakura"
	fallback "github.com/cufee/aftermath/internal/stats/render/session/v2"
)

const (
	Sakura = "sakura"
)

func SelectStyle(name string) func(session, career fetch.AccountStatsOverPeriod, cards session.Cards, subs []models.UserSubscription, opts ...common.Option) (image.Image, error) {
	switch name {
	default:
		return fallback.CardsToImage

	case Sakura:
		return sakura.CardsToImage
	}
}
