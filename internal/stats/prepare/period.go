package prepare

import (
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare/period"
)

func Period(stats fetch.AccountStatsOverPeriod, opts ...period.Option) (period.Cards, error) {
	return period.NewCards(stats, nil, opts...)
}
