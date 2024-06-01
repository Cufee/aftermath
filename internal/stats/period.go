package stats

import (
	"context"
	"image"
	"time"

	"github.com/cufee/aftermath/internal/stats/fetch"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/period"
	render "github.com/cufee/aftermath/internal/stats/render/period"
)

type renderer struct {
	fetchClient fetch.Client
}

func (r *renderer) Period(ctx context.Context, accountId string, from time.Time) (image.Image, Metadata, error) {
	meta := Metadata{}

	stop := meta.Timer("fetchClient#PeriodStats")
	stats, err := r.fetchClient.PeriodStats(ctx, accountId, from)
	stop()
	if err != nil {
		return nil, meta, err
	}
	meta.Stats = stats

	stop = meta.Timer("prepare#NewCards")
	cards, err := prepare.NewCards(stats, nil)
	stop()
	if err != nil {
		return nil, meta, err
	}

	stop = meta.Timer("render#CardsToImage")
	image, err := render.CardsToImage(stats, cards, nil)
	stop()
	if err != nil {
		return nil, meta, err
	}

	return image, meta, err
}
