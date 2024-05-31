package stats

import (
	"image"

	"github.com/cufee/aftermath/internal/stats/fetch"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/period"
	render "github.com/cufee/aftermath/internal/stats/render/period"
)

func PeriodCards(stats fetch.AccountStatsOverPeriod, opts ...prepare.Option) (prepare.Cards, error) {
	return prepare.NewCards(stats, nil, opts...)
}

func PeriodImage(stats fetch.AccountStatsOverPeriod, opts ...prepare.Option) (image.Image, error) {
	cards, err := PeriodCards(stats, opts...)
	if err != nil {
		return nil, err
	}

	return render.CardsToImage(stats, cards, nil)
}
