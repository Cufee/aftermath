package stats

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/localization"
	"github.com/cufee/aftermath/internal/stats/fetch"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/period"
	render "github.com/cufee/aftermath/internal/stats/render/period"
	"golang.org/x/text/language"
)

type renderer struct {
	fetchClient fetch.Client
	locale      language.Tag
}

func (r *renderer) Period(ctx context.Context, accountId string, from time.Time) (Image, Metadata, error) {
	meta := Metadata{}

	printer, err := localization.NewPrinter("stats", r.locale)
	if err != nil {
		return nil, meta, err
	}

	stop := meta.Timer("fetchClient#PeriodStats")
	stats, err := r.fetchClient.PeriodStats(ctx, accountId, from)
	stop()
	if err != nil {
		return nil, meta, err
	}
	meta.Stats = stats

	stop = meta.Timer("prepare#NewCards")
	cards, err := prepare.NewCards(stats, nil, prepare.WithPrinter(printer))
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

	return &imageImp{image}, meta, err
}
