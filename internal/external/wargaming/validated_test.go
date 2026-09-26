package wargaming

import (
	"context"
	"errors"
	"testing"

	"github.com/cufee/am-wg-proxy-next/v2/types"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
)

type vehicleResponseClient struct {
	Client
	data       []types.VehicleStatsFrame
	err        error
	accountID  string
	vehicleIDs []string
	options    types.Options
}

func (c *vehicleResponseClient) AccountVehicles(_ context.Context, _ types.Realm, id string, vehicles []string, opts ...types.Option) ([]types.VehicleStatsFrame, error) {
	c.accountID, c.vehicleIDs, c.options = id, vehicles, types.GetOptions(opts...)
	return c.data, c.err
}

func TestValidatedClientFiltersMissingBattleTimes(t *testing.T) {
	base := &vehicleResponseClient{data: []types.VehicleStatsFrame{
		{TankID: 7953, Stats: types.StatsFrame{Battles: 98, Wins: 61}},
		{TankID: 2, LastBattleTime: -1, Stats: types.StatsFrame{Battles: 1}},
		{TankID: 3, LastBattleTime: 1790447568, Stats: types.StatsFrame{Battles: 1}},
		{TankID: 4},
		{TankID: 5, LastBattleTime: 1790447568, Stats: types.StatsFrame{Battles: 5, DamageDealt: 10000}},
	}}
	counter := IncompleteVehicleStats.WithLabelValues(types.RealmEurope.String(), "618150601")
	var before, after dto.Metric
	require.NoError(t, counter.Write(&before))
	got, err := (&validatedClient{base}).AccountVehicles(context.Background(), types.RealmEurope, "618150601", []string{"7953", "3"}, types.WithFields("all"))
	require.NoError(t, err)
	require.Equal(t, base.data[2:], got)
	require.Equal(t, 7953, base.data[0].TankID, "filter must not mutate upstream buffers")
	require.Equal(t, "618150601", base.accountID)
	require.Equal(t, []string{"7953", "3"}, base.vehicleIDs)
	require.Equal(t, []string{"all"}, base.options.Fields)
	require.NoError(t, counter.Write(&after))
	require.Equal(t, float64(2), after.GetCounter().GetValue()-before.GetCounter().GetValue())
}

func TestValidatedClientPreservesRequestErrors(t *testing.T) {
	want := errors.New("upstream unavailable")
	base := &vehicleResponseClient{err: want, data: []types.VehicleStatsFrame{{TankID: 7953}}}
	got, err := (&validatedClient{base}).AccountVehicles(context.Background(), types.RealmEurope, "618150601", nil)
	require.ErrorIs(t, err, want)
	require.Nil(t, got)
}
