package client

import (
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"github.com/cufee/aftermath/internal/stats/render/common/v1"
	"golang.org/x/text/language"
)

type requestOptions struct {
	snapshotType    models.SnapshotType
	backgroundImage image.Image
	backgroundURL   string
	referenceID     string
	promoText       []string
	vehicleID       string
	withWN8         bool
}

type RequestOption func(o *requestOptions)

func WithWN8() RequestOption {
	return func(o *requestOptions) { o.withWN8 = true }
}
func WithVehicleID(vid string) RequestOption {
	return func(o *requestOptions) { o.vehicleID = vid }
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
func WithBackground(image image.Image) RequestOption {
	return func(o *requestOptions) { o.backgroundImage = image }
}

func (o requestOptions) RenderOpts(printer func(string) string) []common.Option {
	var copts []common.Option
	if o.promoText != nil {
		copts = append(copts, common.WithPromoText(o.promoText...))
	}
	if o.vehicleID != "" {
		copts = append(copts, common.WithVehicleID(o.vehicleID))
	}
	if printer != nil {
		copts = append(copts, common.WithPrinter(printer))
	}
	if o.backgroundImage != nil {
		copts = append(copts, common.WithBackground(o.backgroundImage))
	} else if o.backgroundURL != "" {
		copts = append(copts, common.WithBackgroundURL(o.backgroundURL))
	} else {
		copts = append(copts, common.WithBackgroundURL("static://bg-default"))
	}
	return copts
}

func (o requestOptions) PrepareOpts(printer func(string) string, locale language.Tag) []prepare.Option {
	var popts []prepare.Option
	popts = append(popts, prepare.WithPrinter(printer, locale))
	if o.vehicleID != "" {
		popts = append(popts, prepare.WithVehicleID(o.vehicleID))
	}
	return popts
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
	if o.vehicleID != "" {
		opts = append(opts, fetch.WithVehicleID(o.vehicleID))
	}
	return opts
}
