package client

import (
	"context"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	render "github.com/cufee/aftermath/internal/stats/render/period/v1"
)

func (r *client) PeriodCards(ctx context.Context, accountId string, from time.Time, o ...RequestOption) (prepare.Cards, Metadata, error) {
	var opts = requestOptions{}
	for _, apply := range o {
		apply(&opts)
	}

	meta := Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

	printer, err := localization.NewPrinter("stats", r.locale)
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	stop := meta.Timer("fetchClient#PeriodStats")
	stats, err := r.fetchClient.PeriodStats(ctx, accountId, from, opts.FetchOpts()...)
	stop()
	if err != nil {
		return prepare.Cards{}, meta, err
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
		return prepare.Cards{}, meta, err
	}
	stop()

	stop = meta.Timer("prepare#NewCards")
	cards, err := prepare.NewCards(stats, glossary, opts.PrepareOpts(printer, r.locale)...)
	stop()

	return cards, meta, err
}

func (r *client) PeriodImage(ctx context.Context, accountId string, from time.Time, o ...RequestOption) (Image, Metadata, error) {
	var opts = requestOptions{}
	for _, apply := range o {
		apply(&opts)
	}

	cards, meta, err := r.PeriodCards(ctx, accountId, from, o...)
	if err != nil {
		return nil, meta, err
	}

	stop := meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(meta.Stats["period"], cards, nil, opts.RenderOpts()...)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return &imageImp{image}, meta, err
}
