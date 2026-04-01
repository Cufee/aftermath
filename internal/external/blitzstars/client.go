package blitzstars

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var ErrServiceUnavailable = errors.New("blitz stars unavailable")

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
