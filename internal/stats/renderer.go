package stats

import (
	"context"
	"image"
	"time"

	"github.com/cufee/aftermath/internal/stats/fetch"
	"golang.org/x/text/language"
)

type Renderer interface {
	Period(ctx context.Context, accountId string, from time.Time) (image.Image, Metadata, error)
	// Replay(accountId string, from time.Time) (image.Image, error)
	// Session(accountId string, from time.Time) (image.Image, error)
}

func NewRenderer(fetch fetch.Client, locale language.Tag) *renderer {
	return &renderer{fetch, locale}
}
