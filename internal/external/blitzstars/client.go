package blitzstars

import (
	"context"
	"net/http"
	"time"
)

type Client interface {
	AccountTankHistories(ctx context.Context, accountId string) (map[int][]TankHistoryEntry, error)
}

// var _ Client = &client{} // just a marker to see if it is implemented correctly

type client struct {
	http           http.Client
	apiURL         string
	requestTimeout time.Duration
}

func NewClient(apiURL string, requestTimeout time.Duration) *client {
	return &client{
		apiURL:         apiURL,
		requestTimeout: requestTimeout,
		http:           http.Client{Timeout: requestTimeout},
	}
}
