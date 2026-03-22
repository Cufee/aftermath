package blitzkit

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
)

//go:generate go tool github.com/bufbuild/buf/cmd/buf generate --template buf.gen.yaml --path averages.proto

var ErrServiceUnavailable = errors.New("blitzkit averages unavailable")

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
