package wargaming

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/client"
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

	return client.NewEmbeddedClient(primaryAppId, primaryRps, proxyHostList, timeout)
}
