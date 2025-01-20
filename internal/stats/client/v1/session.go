package client

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	options "github.com/cufee/aftermath/internal/stats/client/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	render "github.com/cufee/aftermath/internal/stats/render/session/v1"
)

func (c *client) EmptySessionCards(ctx context.Context, accountId string) (prepare.Cards, options.Metadata, error) {
	meta := options.Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

	stop := meta.Timer("database#GetAccountByID")
	account, err := c.database.GetAccountByID(ctx, accountId)
	stop()
	if err != nil {
		if database.IsNotFound(err) {
			_, err := c.fetchClient.Account(ctx, accountId) // this will cache the account
			if err != nil {
				return prepare.Cards{}, meta, err
			}
			return prepare.Cards{}, meta, options.ErrAccountNotTracked
		}
		return prepare.Cards{}, meta, err
	}

	printer, err := localization.NewPrinterWithFallback("stats", c.locale)
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	stop = meta.Timer("prepare#NewCards")
	cards, err := prepare.NewCards(fetch.AccountStatsOverPeriod{Account: account}, fetch.AccountStatsOverPeriod{Account: account}, nil, common.WithPrinter(printer, c.locale))
	stop()
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	return cards, meta, nil
}

func (c *client) SessionCards(ctx context.Context, accountId string, from time.Time, o ...options.RequestOption) (prepare.Cards, options.Metadata, error) {
	opts := options.RequestOptions(o).Options()

	meta := options.Metadata{Stats: make(map[string]fetch.AccountStatsOverPeriod)}

	stop := meta.Timer("database#GetAccountByID")
	_, err := c.database.GetAccountByID(ctx, accountId)
	stop()
	if database.IsNotFound(err) {
		// record a session in the background
		go recordAccountSnapshots(c.wargaming, c.database, accountId, opts.ReferenceID())

		return prepare.Cards{}, meta, options.ErrAccountNotTracked
	}
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	printer, err := localization.NewPrinterWithFallback("stats", c.locale)
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
		if errors.Is(err, fetch.ErrSessionNotFound) {
			go recordAccountSnapshots(c.wargaming, c.database, accountId, opts.ReferenceID())
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
	if opts.VehicleID() != "" && !slices.Contains(vehicles, opts.VehicleID()) {
		vehicles = append(vehicles, opts.VehicleID())
	}

	glossary, err := c.database.GetVehicles(ctx, vehicles)
	if err != nil {
		return prepare.Cards{}, meta, err
	}
	stop()

	stop = meta.Timer("prepare#NewCards")

	cards, err := prepare.NewCards(session, career, glossary, opts.PrepareOpts(printer, c.locale)...)
	stop()
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	return cards, meta, nil
}

func (c *client) SessionImage(ctx context.Context, accountId string, from time.Time, o ...options.RequestOption) (options.Image, options.Metadata, error) {
	opts := options.RequestOptions(o).Options()

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
