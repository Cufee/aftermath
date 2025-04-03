package client

import (
	"context"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/stats/client/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	render "github.com/cufee/aftermath/internal/stats/render/period/v1"
)

func (r *client) PeriodCards(ctx context.Context, accountId string, from time.Time, o ...common.RequestOption) (prepare.Cards, common.Metadata, error) {
	opts := common.RequestOptions(o).Options()

	meta := common.Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

	printer, err := localization.NewPrinterWithFallback("stats", r.locale)
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	// cache account and record session snapshots
	go func(id, reference string) {
		_, err = r.database.GetAccountByID(ctx, id)
		if !database.IsNotFound(err) {
			// account was found or some other error happened - no need to do anything here
			return
		}
		recordAccountSnapshots(r.wargaming, r.database, id, reference)
	}(accountId, opts.ReferenceID())

	stop := meta.Timer("fetchClient#PeriodStats")
	stats, err := r.fetchClient.CurrentStats(ctx, accountId, opts.FetchOpts()...)
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

func (r *client) PeriodImage(ctx context.Context, accountId string, from time.Time, o ...common.RequestOption) (common.Image, common.Metadata, error) {
	opts := common.RequestOptions(o).Options()

	cards, meta, err := r.PeriodCards(ctx, accountId, from, o...)
	if err != nil {
		return nil, meta, err
	}

	printer, err := localization.NewPrinterWithFallback("stats", r.locale)
	if err != nil {
		return nil, meta, err
	}

	stop := meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(meta.Stats["period"], cards, opts.Subscriptions, opts.RenderOpts(printer)...)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return &imageImp{image}, meta, err
}
