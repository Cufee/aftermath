package client

import (
	"context"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
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

	printer, err := localization.NewPrinterWithFallback("stats", r.locale)
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	// cache account and record session snapshots
	go func(id string) {
		_, err = r.database.GetAccountByID(ctx, accountId)
		if !database.IsNotFound(err) {
			// account was found or some other error happened - no need to do anything here
			return
		}
		// record a session in the background
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		_, err := logic.RecordAccountSnapshots(ctx, r.wargaming, r.database, r.wargaming.RealmFromAccountID(id), false, logic.WithReference(id, opts.referenceID))
		if err != nil {
			log.Err(err).Str("accountId", id).Msg("failed to record account snapshot")
		}
	}(accountId)

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
	if opts.vehicleID != "" && !slices.Contains(vehicles, opts.vehicleID) {
		vehicles = append(vehicles, opts.vehicleID)
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

	printer, err := localization.NewPrinterWithFallback("stats", r.locale)
	if err != nil {
		return nil, meta, err
	}

	stop := meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(meta.Stats["period"], cards, nil, opts.RenderOpts(printer)...)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return &imageImp{image}, meta, err
}
