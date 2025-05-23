package blitzstars

import (
	"context"
	"net/http"
	"time"

	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/pkg/errors"
)

var ErrServiceUnavailable = errors.New("blitz stars unavailable")

type Client interface {
	CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error)
}

// var _ Client = &client{} // just a marker to see if it is implemented correctly

type client struct {
	http           http.Client
	apiURL         string
	requestTimeout time.Duration
}

func NewClient(apiURL string, requestTimeout time.Duration) (client, error) {
	return client{
		apiURL:         apiURL,
		requestTimeout: requestTimeout,
		http:           http.Client{Timeout: requestTimeout},
	}, nil
}
