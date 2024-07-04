package client

import (
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
)

type requestOptions struct {
	snapshotType  models.SnapshotType
	backgroundURL string
	referenceID   string
	promoText     []string
	withWN8       bool
}

type RequestOption func(o *requestOptions)

func WithWN8() RequestOption {
	return func(o *requestOptions) { o.withWN8 = true }
}
func WithReferenceID(refID string) RequestOption {
	return func(o *requestOptions) { o.referenceID = refID }
}
func WithPromoText(text ...string) RequestOption {
	return func(o *requestOptions) { o.promoText = append(o.promoText, text...) }
}
func WithType(t models.SnapshotType) RequestOption {
	return func(o *requestOptions) { o.snapshotType = t }
}
func WithBackgroundURL(url string) RequestOption {
	return func(o *requestOptions) { o.backgroundURL = url }
}

func (o requestOptions) RenderOpts() []common.Option {
	var copts []common.Option
	if o.promoText != nil {
		copts = append(copts, common.WithPromoText(o.promoText...))
	}
	copts = append(copts, common.WithBackground(o.backgroundURL))
	return copts
}

func (o requestOptions) FetchOpts() []fetch.StatsOption {
	var opts []fetch.StatsOption
	if o.snapshotType != "" {
		opts = append(opts, fetch.WithType(o.snapshotType))
	}
	if o.referenceID != "" {
		opts = append(opts, fetch.WithReferenceID(o.referenceID))
	}
	if o.withWN8 {
		opts = append(opts, fetch.WithWN8())
	}
	return opts
}