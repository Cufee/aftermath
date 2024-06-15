package stats

import (
	"context"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/prepare/common"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/session"
	options "github.com/cufee/aftermath/internal/stats/render"
	render "github.com/cufee/aftermath/internal/stats/render/session"
	"github.com/rs/zerolog/log"
)

func (r *renderer) Session(ctx context.Context, accountId string, from time.Time, opts ...options.Option) (Image, Metadata, error) {
	meta := Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

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
	session, career, err := r.fetchClient.SessionStats(ctx, accountId, from, fetch.WithWN8())
	stop()
	if err != nil {
		return nil, meta, err
	}
	meta.Stats["career"] = career
	meta.Stats["session"] = session

	stop = meta.Timer("prepare#GetVehicles")
	var vehicles []string
	for id := range session.RegularBattles.Vehicles {
		vehicles = append(vehicles, id)
	}
	for id := range session.RatingBattles.Vehicles {
		if !slices.Contains(vehicles, id) {
			vehicles = append(vehicles, id)
		}
	}
	for id := range career.RegularBattles.Vehicles {
		if !slices.Contains(vehicles, id) {
			vehicles = append(vehicles, id)
		}
	}
	for id := range career.RatingBattles.Vehicles {
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
	cards, err := prepare.NewCards(session, career, glossary, common.WithPrinter(printer, r.locale))
	stop()
	if err != nil {
		return nil, meta, err
	}

	stop = meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(session, career, cards, nil, opts...)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return &imageImp{image}, meta, err
}
