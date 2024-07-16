package client

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	period "github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	session "github.com/cufee/aftermath/internal/stats/prepare/session/v1"
	"golang.org/x/text/language"
)

var _ Client = &client{}

type client struct {
	fetchClient fetch.Client
	wargaming   wargaming.Client
	database    database.Client
	locale      language.Tag
}

type Client interface {
	PeriodCards(ctx context.Context, accountId string, from time.Time, opts ...RequestOption) (period.Cards, Metadata, error)
	PeriodImage(ctx context.Context, accountId string, from time.Time, opts ...RequestOption) (Image, Metadata, error)

	SessionCards(ctx context.Context, accountId string, from time.Time, opts ...RequestOption) (session.Cards, Metadata, error)
	SessionImage(ctx context.Context, accountId string, from time.Time, opts ...RequestOption) (Image, Metadata, error)

	ReplayCards(ctx context.Context, replayURL string, o ...RequestOption) (replay.Cards, Metadata, error)
	ReplayImage(ctx context.Context, replayURL string, o ...RequestOption) (Image, Metadata, error)
}

func NewClient(fetch fetch.Client, database database.Client, wargaming wargaming.Client, locale language.Tag) Client {
	return &client{fetch, wargaming, database, locale}
}
