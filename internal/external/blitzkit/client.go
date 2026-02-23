package blitzkit

import (
	"context"
	"net/http"
	"time"

	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/pkg/errors"
)

//go:generate go tool github.com/bufbuild/buf/cmd/buf generate --template buf.gen.yaml --path averages.proto

var ErrServiceUnavailable = errors.New("blitzkit averages unavailable")

type Client interface {
	CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error)
}

var _ Client = &client{} // just a marker to see if it is implemented correctly

type client struct {
	http           http.Client
	requestTimeout time.Duration
}

func NewClient(requestTimeout time.Duration) (client, error) {
	return client{
		requestTimeout: requestTimeout,
		http:           http.Client{Timeout: requestTimeout},
	}, nil
}
