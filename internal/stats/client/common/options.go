package common

import (
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	common "github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
	"golang.org/x/text/language"
)

type requestOptions struct {
	snapshotType       models.SnapshotType
	backgroundImage    image.Image
	backgroundURL      string
	backgroundIsCustom bool
	referenceID        string
	promoText          []string
	withWN8            bool
	VehicleIDs         []string
	Subscriptions      []models.UserSubscription

	vehicleTags    []prepare.Tag
	ratingColumns  []prepare.TagColumn[string]
	unratedColumns []prepare.TagColumn[string]
}

func (o requestOptions) ReferenceID() string {
	return o.referenceID
}

type RequestOption func(o *requestOptions)
type RequestOptions []RequestOption

func (o RequestOptions) Options() requestOptions {
	var opts = requestOptions{}
	for _, apply := range o {
		apply(&opts)
	}
	return opts
}

func WithSubscriptions(subs []models.UserSubscription) RequestOption {
	return func(o *requestOptions) { o.Subscriptions = subs }
}
func WithWN8() RequestOption {
	return func(o *requestOptions) { o.withWN8 = true }
}
func WithVehicleIDs(vid ...string) RequestOption {
	return func(o *requestOptions) { o.VehicleIDs = vid }
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
func WithBackgroundURL(url string, isCustom bool) RequestOption {
	return func(o *requestOptions) { o.backgroundURL = url; o.backgroundIsCustom = isCustom }
}
func WithBackground(image image.Image, isCustom bool) RequestOption {
	return func(o *requestOptions) { o.backgroundImage = image; o.backgroundIsCustom = isCustom }
}

func (o requestOptions) RenderOpts(printer func(string) string) []common.Option {
	var copts []common.Option
	if o.promoText != nil {
		copts = append(copts, common.WithPromoText(o.promoText...))
	}
	if o.VehicleIDs != nil {
		copts = append(copts, common.WithVehicleIDs(o.VehicleIDs...))
	}
	if printer != nil {
		copts = append(copts, common.WithPrinter(printer))
	}
	if o.backgroundImage != nil {
		copts = append(copts, common.WithBackground(o.backgroundImage, o.backgroundIsCustom))
	} else if o.backgroundURL != "" {
		copts = append(copts, common.WithBackgroundURL(o.backgroundURL, o.backgroundIsCustom))
	} else {
		copts = append(copts, common.WithBackgroundURL("static://bg-default", false))
	}
	return copts
}

func (o requestOptions) PrepareOpts(printer func(string) string, locale language.Tag) []prepare.Option {
	var popts []prepare.Option
	popts = append(popts, prepare.WithPrinter(printer, locale))
	if o.vehicleTags != nil {
		popts = append(popts, prepare.WithVehicleTags(o.vehicleTags...))
	}
	if o.ratingColumns != nil {
		popts = append(popts, prepare.WithRatingColumns(o.ratingColumns...))
	}
	if o.unratedColumns != nil {
		popts = append(popts, prepare.WithUnratedColumns(o.unratedColumns...))
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
	if o.VehicleIDs != nil {
		opts = append(opts, fetch.WithVehicleIDs(o.VehicleIDs...))
	}
	return opts
}
