package client

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	render "github.com/cufee/aftermath/internal/stats/render/session/v1"
	"github.com/rs/zerolog/log"
)

func (c *client) SessionCards(ctx context.Context, accountId string, from time.Time, o ...RequestOption) (prepare.Cards, Metadata, error) {
	var opts = requestOptions{}
	for _, apply := range o {
		apply(&opts)
	}

	meta := Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

	stop := meta.Timer("database#GetAccountByID")
	_, err := c.database.GetAccountByID(ctx, accountId)
	stop()
	if err != nil {
		if database.IsNotFound(err) {
			_, err := c.fetchClient.Account(ctx, accountId) // this will cache the account
			if err != nil {
				return prepare.Cards{}, meta, err
			}
			go func(id string) {
				// record a session in the background
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
				defer cancel()

				_, err := logic.RecordAccountSnapshots(ctx, c.wargaming, c.database, c.wargaming.RealmFromAccountID(id), false, id)
				if err != nil {
					log.Err(err).Str("accountId", id).Msg("failed to record account snapshot")
				}
			}(accountId)
			return prepare.Cards{}, meta, ErrAccountNotTracked
		}
		return prepare.Cards{}, meta, err
	}

	printer, err := localization.NewPrinter("stats", c.locale)
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	stop = meta.Timer("fetchClient#SessionStats")
	if from.IsZero() {
		from = time.Now()
	}
	session, career, err := c.fetchClient.SessionStats(ctx, accountId, from, opts.FetchOpts()...)
	stop()
	if err != nil {
		if errors.Is(fetch.ErrSessionNotFound, err) {
			go func(id string) {
				// record a session in the background
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
				defer cancel()

				_, err := logic.RecordAccountSnapshots(ctx, c.wargaming, c.database, c.wargaming.RealmFromAccountID(id), false, id)
				if err != nil {
					log.Err(err).Str("accountId", id).Msg("failed to record account snapshot")
				}
			}(accountId)
		}
		return prepare.Cards{}, meta, err
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

	glossary, err := c.database.GetVehicles(ctx, vehicles)
	if err != nil {
		return prepare.Cards{}, meta, err
	}
	stop()

	stop = meta.Timer("prepare#NewCards")
	cards, err := prepare.NewCards(session, career, glossary, common.WithPrinter(printer, c.locale))
	stop()
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	return cards, meta, nil
}

func (c *client) SessionImage(ctx context.Context, accountId string, from time.Time, o ...RequestOption) (Image, Metadata, error) {
	var opts = requestOptions{}
	for _, apply := range o {
		apply(&opts)
	}

	cards, meta, err := c.SessionCards(ctx, accountId, from, o...)
	if err != nil {
		return nil, meta, err
	}

	stop := meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(meta.Stats["session"], meta.Stats["career"], cards, nil, opts.RenderOpts()...)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return &imageImp{image}, meta, err
}
