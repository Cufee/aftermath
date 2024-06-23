package renderer

import (
	"context"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	options "github.com/cufee/aftermath/internal/stats/render"
	render "github.com/cufee/aftermath/internal/stats/render/period/v1"
)

func (r *renderer) Period(ctx context.Context, accountId string, from time.Time, opts ...options.Option) (Image, Metadata, error) {
	meta := Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

	printer, err := localization.NewPrinter("stats", r.locale)
	if err != nil {
		return nil, meta, err
	}

	stop := meta.Timer("fetchClient#PeriodStats")
	stats, err := r.fetchClient.PeriodStats(ctx, accountId, from, fetch.WithWN8())
	stop()
	if err != nil {
		return nil, meta, err
	}
	meta.Stats["period"] = stats

	stop = meta.Timer("prepare#GetVehicles")
	var vehicles []string
	for id := range stats.RegularBattles.Vehicles {
		vehicles = append(vehicles, id)
	}
	for id := range stats.RatingBattles.Vehicles {
		if !slices.Contains(vehicles, id) {
			vehicles = append(vehicles, id)
		}
	}

	glossary, err := r.database.GetVehicles(ctx, vehicles)
	if err != nil {
		return nil, meta, err
	}
	stop()

	stop = meta.Timer("prepare#NewCards")
	cards, err := prepare.NewCards(stats, glossary, common.WithPrinter(printer, r.locale))
	stop()
	if err != nil {
		return nil, meta, err
	}

	stop = meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(stats, cards, nil, opts...)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return &imageImp{image}, meta, err
}
