package stats

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/aftermath/internal/stats/render"
	"golang.org/x/text/language"
)

var _ Renderer = &renderer{}

type renderer struct {
	fetchClient fetch.Client
	wargaming   wargaming.Client
	database    database.Client
	locale      language.Tag
}

type Renderer interface {
	Period(ctx context.Context, accountId string, from time.Time, opts ...render.Option) (Image, Metadata, error)
	Session(ctx context.Context, accountId string, from time.Time, opts ...render.Option) (Image, Metadata, error)

	// Replay(accountId string, from time.Time) (image.Image, error)
}

func NewRenderer(fetch fetch.Client, database database.Client, wargaming wargaming.Client, locale language.Tag) *renderer {
	return &renderer{fetch, wargaming, database, locale}
}
