package stats

import (
	"context"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/period"
	options "github.com/cufee/aftermath/internal/stats/render"
	render "github.com/cufee/aftermath/internal/stats/render/period"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *renderer) Session(ctx context.Context, accountId string, from time.Time, opts ...options.Option) (Image, Metadata, error) {
	meta := Metadata{}
	stop := meta.Timer("database#GetAccountByID")
	_, err := r.database.GetAccountByID(ctx, accountId)
	stop()
	if err != nil {
		if db.IsErrNotFound(err) {
			_, err := r.fetchClient.Account(ctx, accountId) // this will cache the account
			if err != nil {
				return nil, meta, err
			}
			go func() {
				// record a session in the background
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
				defer cancel()

				_, err := logic.RecordAccountSnapshots(ctx, r.wargaming, r.database, r.wargaming.RealmFromAccountID(accountId), false, accountId)
				if err != nil {
					log.Err(err).Str("accountId", accountId).Msg("failed to record account snapshot")
				}
			}()
			return nil, meta, ErrAccountNotTracked
		}
		return nil, meta, err
	}

	printer, err := localization.NewPrinter("stats", r.locale)
	if err != nil {
		return nil, meta, err
	}

	stop = meta.Timer("fetchClient#SessionStats")
	stats, err := r.fetchClient.SessionStats(ctx, accountId, from, fetch.WithWN8())
	stop()
	if errors.Is(err, fetch.ErrSessionNotFound) && time.Since(from).Hours()/24 <= 90 {
		// we dont have a session, but one might be available from blitzstars
		stats, err = r.fetchClient.PeriodStats(ctx, accountId, from, fetch.WithWN8())
		// the error will be checked below
	}
	if errors.Is(err, fetch.ErrSessionNotFound) {
		// Get account info and return a blank session
		current, err := r.fetchClient.CurrentStats(ctx, accountId)
		if err != nil {
			return nil, meta, err
		}
		current.PeriodEnd = time.Now()
		current.PeriodStart = current.PeriodEnd
		current.RatingBattles = fetch.StatsWithVehicles{}
		current.RegularBattles = fetch.StatsWithVehicles{}
		stats = current
	}
	if err != nil {
		return nil, meta, err
	}
	meta.Stats = stats

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
	cards, err := prepare.NewCards(stats, glossary, prepare.WithPrinter(printer, r.locale))
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
