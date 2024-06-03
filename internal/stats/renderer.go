package stats

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/stats/fetch"
	"golang.org/x/text/language"
)

var _ Renderer = &renderer{}

type Renderer interface {
	Period(ctx context.Context, accountId string, from time.Time) (Image, Metadata, error)

	// Replay(accountId string, from time.Time) (image.Image, error)
	// Session(accountId string, from time.Time) (image.Image, error)
}

func NewRenderer(fetch fetch.Client, locale language.Tag) *renderer {
	return &renderer{fetch, locale}
}
