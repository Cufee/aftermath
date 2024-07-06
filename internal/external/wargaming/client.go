package wargaming

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/rs/zerolog"
)

type Client interface {
	client.Client
}

func NewClientFromEnv(primaryAppId, primaryAppRps, requestTimeout, proxyHostList string) (Client, error) {
	primaryRps, err := strconv.Atoi(primaryAppRps)
	if err != nil {
		return nil, fmt.Errorf("primaryAppRps invalid: %w", err)

	}

	timeoutInt, err := strconv.Atoi(requestTimeout)
	if err != nil {
		return nil, fmt.Errorf("requestTimeout invalid: %w", err)
	}
	timeout := time.Second * time.Duration(timeoutInt)

	return client.NewEmbeddedClient(primaryAppId, primaryRps, proxyHostList, timeout, client.WithLogLevel(zerolog.WarnLevel))
}

var PublicAppID = os.Getenv("WG_AUTH_APP_ID")
