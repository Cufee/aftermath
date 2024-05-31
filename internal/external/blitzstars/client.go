package blitzstars

import (
	"net/http"
	"time"
)

type Client struct {
	http           http.Client
	apiURL         string
	requestTimeout time.Duration
}

func NewClient(apiURL string, requestTimeout time.Duration) Client {
	return Client{
		apiURL:         apiURL,
		requestTimeout: requestTimeout,
		http:           http.Client{Timeout: requestTimeout},
	}
}
