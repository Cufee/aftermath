package client

import (
	"context"
	"slices"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/logic"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	render "github.com/cufee/aftermath/internal/stats/render/replay/v1"
)

func (r *client) ReplayCards(ctx context.Context, replayURL string, o ...RequestOption) (prepare.Cards, Metadata, error) {
	var opts = requestOptions{}
	for _, apply := range o {
		apply(&opts)
	}

	meta := newMeta()

	printer, err := localization.NewPrinterWithFallback("stats", r.locale)
	if err != nil {
		return prepare.Cards{}, meta, err
	}

	stop := meta.Timer("fetchClient#PeriodStats")
	replay, err := r.fetchClient.ReplayRemote(ctx, replayURL)
	stop()
	if err != nil {
		return prepare.Cards{}, meta, err
	}
	meta.Replay = replay

	stop = meta.Timer("prepare#GetVehicles")
	var vehicles []string
	for _, player := range append(replay.Teams.Allies, replay.Teams.Enemies...) {
		if id := player.VehicleID; !slices.Contains(vehicles, id) {
			vehicles = append(vehicles, id)
		}
	}

	glossary, err := r.database.GetVehicles(ctx, vehicles)
	if err != nil {
		return prepare.Cards{}, meta, err
	}
	stop()

	gameModeNames, err := r.database.GetGameModeNames(ctx, replay.GameMode.String())
	if err != nil && !database.IsNotFound(err) {
		return prepare.Cards{}, meta, nil
	}

	stop = meta.Timer("prepare#NewCards")
	cards, err := prepare.NewCards(replay, glossary, gameModeNames, opts.PrepareOpts(printer, r.locale)...)
	stop()

	return cards, meta, err
}

func (r *client) ReplayImage(ctx context.Context, replayURL string, o ...RequestOption) (Image, Metadata, error) {
	var opts = requestOptions{}
	for _, apply := range o {
		apply(&opts)
	}

	cards, meta, err := r.ReplayCards(ctx, replayURL, o...)
	if err != nil {
		return nil, meta, err
	}

	printer, err := localization.NewPrinterWithFallback("stats", r.locale)
	if err != nil {
		return nil, meta, err
	}

	if img, _, err := logic.GetBackgroundImageFromRef(ctx, r.database, meta.Replay.Protagonist.ID); err == nil {
		opts.backgroundImage = img
	}

	stop := meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(meta.Replay, cards, opts.RenderOpts(printer)...)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return &imageImp{image}, meta, err
}
