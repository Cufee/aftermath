package client

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/stats/client/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	render "github.com/cufee/aftermath/internal/stats/render/session/v2"
)

func (c *client) SessionCards(ctx context.Context, accountId string, from time.Time, o ...common.RequestOption) (session.Cards, common.Metadata, error) {
	opts := common.RequestOptions(o).Options()

	meta := common.Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

	stop := meta.Timer("database#GetAccountByID")
	_, err := c.database.GetAccountByID(ctx, accountId)
	stop()
	if database.IsNotFound(err) {
		// record a session in the background
		go recordAccountSnapshots(c.wargaming, c.database, accountId, opts.ReferenceID())

		return session.Cards{}, meta, common.ErrAccountNotTracked
	}
	if err != nil {
		return session.Cards{}, meta, err
	}

	printer, err := localization.NewPrinterWithFallback("stats", c.locale)
	if err != nil {
		return session.Cards{}, meta, err
	}

	stop = meta.Timer("fetchClient#SessionStats")
	if from.IsZero() {
		from = time.Now()
	}
	sessionStats, careerStats, err := c.fetchClient.SessionStats(ctx, accountId, from, opts.FetchOpts()...)
	stop()
	if err != nil {
		if errors.Is(err, fetch.ErrSessionNotFound) {
			go recordAccountSnapshots(c.wargaming, c.database, accountId, opts.ReferenceID())
		}
		return session.Cards{}, meta, err
	}
	meta.Stats["career"] = careerStats
	meta.Stats["session"] = sessionStats

	stop = meta.Timer("prepare#GetVehicles")
	var vehicles []string
	for id := range sessionStats.RegularBattles.Vehicles {
		vehicles = append(vehicles, id)
	}
	for id := range sessionStats.RatingBattles.Vehicles {
		if !slices.Contains(vehicles, id) {
			vehicles = append(vehicles, id)
		}
	}
	for id := range careerStats.RegularBattles.Vehicles {
		if !slices.Contains(vehicles, id) {
			vehicles = append(vehicles, id)
		}
	}
	for id := range careerStats.RatingBattles.Vehicles {
		if !slices.Contains(vehicles, id) {
			vehicles = append(vehicles, id)
		}
	}
	if opts.VehicleID() != "" && !slices.Contains(vehicles, opts.VehicleID()) {
		vehicles = append(vehicles, opts.VehicleID())
	}

	glossary, err := c.database.GetVehicles(ctx, vehicles)
	if err != nil {
		return session.Cards{}, meta, err
	}
	stop()

	stop = meta.Timer("prepare#NewCards")

	cards, err := session.NewCards(sessionStats, careerStats, glossary, opts.PrepareOpts(printer, c.locale)...)
	stop()
	if err != nil {
		return session.Cards{}, meta, err
	}

	return cards, meta, nil
}

func (c *client) SessionImage(ctx context.Context, accountId string, from time.Time, o ...common.RequestOption) (common.Image, common.Metadata, error) {
	opts := common.RequestOptions(o).Options()

	cards, meta, err := c.SessionCards(ctx, accountId, from, o...)
	if err != nil {
		return nil, meta, err
	}

	printer, err := localization.NewPrinterWithFallback("stats", c.locale)
	if err != nil {
		return nil, meta, err
	}

	stop := meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(meta.Stats["session"], meta.Stats["career"], cards, opts.Subscriptions, opts.RenderOpts(printer)...)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return &imageImp{image}, meta, err
}

func (c *client) EmptySessionCards(ctx context.Context, accountId string) (session.Cards, common.Metadata, error) {
	meta := common.Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

	stop := meta.Timer("database#GetAccountByID")
	account, err := c.database.GetAccountByID(ctx, accountId)
	stop()
	if err != nil {
		if database.IsNotFound(err) {
			_, err := c.fetchClient.Account(ctx, accountId) // this will cache the account
			if err != nil {
				return session.Cards{}, meta, err
			}
			return session.Cards{}, meta, common.ErrAccountNotTracked
		}
		return session.Cards{}, meta, err
	}

	printer, err := localization.NewPrinterWithFallback("stats", c.locale)
	if err != nil {
		return session.Cards{}, meta, err
	}

	stop = meta.Timer("prepare#NewCards")
	cards, err := session.NewCards(fetch.AccountStatsOverPeriod{Account: account}, fetch.AccountStatsOverPeriod{Account: account}, nil, prepare.WithPrinter(printer, c.locale))
	stop()
	if err != nil {
		return session.Cards{}, meta, err
	}

	return cards, meta, nil
}
