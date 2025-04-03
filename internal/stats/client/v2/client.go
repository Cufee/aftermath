package client

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/stats/client/common"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	period "github.com/cufee/aftermath/internal/stats/prepare/period/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/replay/v1"
	"github.com/cufee/aftermath/internal/stats/prepare/session/v1"
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
	PeriodCards(ctx context.Context, accountId string, from time.Time, opts ...common.RequestOption) (period.Cards, common.Metadata, error)
	PeriodImage(ctx context.Context, accountId string, from time.Time, opts ...common.RequestOption) (common.Image, common.Metadata, error)

	SessionCards(ctx context.Context, accountId string, from time.Time, opts ...common.RequestOption) (session.Cards, common.Metadata, error)
	SessionImage(ctx context.Context, accountId string, from time.Time, opts ...common.RequestOption) (common.Image, common.Metadata, error)
	EmptySessionCards(ctx context.Context, accountId string) (session.Cards, common.Metadata, error)

	ReplayCards(ctx context.Context, replayURL string, o ...common.RequestOption) (replay.Cards, common.Metadata, error)
	ReplayImage(ctx context.Context, replayURL string, o ...common.RequestOption) (common.Image, common.Metadata, error)
}

func NewClient(fetch fetch.Client, database database.Client, wargaming wargaming.Client, locale language.Tag) Client {
	return &client{fetch, wargaming, database, locale}
}
