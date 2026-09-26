package wargaming

import (
	"context"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
)

type Client interface {
	client.Client
}

var IncompleteVehicleStats = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "wargaming_incomplete_vehicle_stats_total",
	Help: "Number of tank entries with battles but no last battle time filtered from Wargaming responses (not unique tanks).",
}, []string{"realm", "account_id"})

type validatedClient struct{ Client }

func (c *validatedClient) AccountVehicles(ctx context.Context, realm types.Realm, id string, vehicles []string, opts ...types.Option) ([]types.VehicleStatsFrame, error) {
	data, err := c.Client.AccountVehicles(ctx, realm, id, vehicles, opts...)
	if err != nil {
		return nil, err
	}
	valid := make([]types.VehicleStatsFrame, 0, len(data))
	for _, vehicle := range data {
		// WG placeholders have lifetime battles/wins but no usable detailed stats.
		if vehicle.Stats.Battles > 0 && vehicle.LastBattleTime <= 0 {
			IncompleteVehicleStats.WithLabelValues(realm.String(), id).Inc()
			continue
		}
		valid = append(valid, vehicle)
	}
	return valid, nil
}

func NewClientFromEnv(primaryAppId string, primaryAppRps int, requestTimeout time.Duration, proxyHostList, userAgent string) (Client, error) {
	c, err := client.NewEmbeddedClient(
		primaryAppId,
		primaryAppRps,
		proxyHostList,
		requestTimeout,
		client.WithLogLevel(zerolog.WarnLevel),
		client.WithUserAgent(userAgent),
	)
	if err != nil {
		return nil, err
	}
	return &validatedClient{c}, nil
}
