package wargaming

import (
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/rs/zerolog"
)

type Client interface {
	client.Client
}

func NewClientFromEnv(primaryAppId string, primaryAppRps int, requestTimeout time.Duration, proxyHostList string) (Client, error) {
	return client.NewEmbeddedClient(primaryAppId, primaryAppRps, proxyHostList, requestTimeout, client.WithLogLevel(zerolog.WarnLevel))
}
