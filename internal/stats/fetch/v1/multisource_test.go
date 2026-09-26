package fetch

import (
	"context"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/stretchr/testify/require"
)

func TestSessionTotalsWithCleanBaselineAndNewTank(t *testing.T) {
	c := multiSourceClient{database: totalsDB{}, wargaming: totalsWG{}, retriesPerRequest: 1}
	session, _, err := c.SessionStats(context.Background(), "123", time.Now())
	require.NoError(t, err)
	require.Equal(t, frame.ValueInt(4), session.RegularBattles.Battles)
	require.Equal(t, frame.ValueInt(8000), session.RegularBattles.DamageDealt)
	require.Equal(t, frame.ValueInt(2), session.RegularBattles.BattlesWon)
	require.Len(t, session.RegularBattles.Vehicles, 2)
	require.Equal(t, frame.ValueInt(2), session.RegularBattles.Vehicles["2"].Battles)
	require.Equal(t, frame.ValueInt(4000), session.RegularBattles.Vehicles["2"].DamageDealt)
	require.Equal(t, frame.ValueInt(2), session.RatingBattles.Battles)
	require.Equal(t, frame.ValueInt(1), session.RatingBattles.BattlesWon)
	require.Equal(t, frame.ValueInt(4000), session.RatingBattles.DamageDealt)
}

type totalsDB struct{ database.Client }

func (totalsDB) GetAccountSnapshots(context.Context, []string, models.SnapshotType, ...database.Query) ([]models.AccountSnapshot, error) {
	return []models.AccountSnapshot{{
		RegularBattles: frame.StatsFrame{Battles: 10, BattlesWon: 5, DamageDealt: 10000},
		RatingBattles:  frame.StatsFrame{Battles: 18, BattlesWon: 9, DamageDealt: 16000},
	}}, nil
}
func (totalsDB) GetVehicleSnapshots(context.Context, string, []string, models.SnapshotType, ...database.Query) ([]models.VehicleSnapshot, error) {
	return []models.VehicleSnapshot{
		{VehicleID: "1", Stats: frame.StatsFrame{Battles: 10, BattlesWon: 5, DamageDealt: 10000}},
	}, nil
}
func (totalsDB) UpsertAccounts(context.Context, ...*models.Account) error { return nil }

type totalsWG struct{ wargaming.Client }

func (totalsWG) RealmFromID(string) (types.Realm, bool) { return types.RealmEurope, true }
func (totalsWG) AccountByID(context.Context, types.Realm, string, ...types.Option) (types.ExtendedAccount, error) {
	var account types.ExtendedAccount
	account.Statistics.Rating = types.StatsFrame{Battles: 20, Wins: 10, DamageDealt: 20000}
	return account, nil
}
func (totalsWG) AccountClan(context.Context, types.Realm, string, ...types.Option) (types.ClanMember, error) {
	return types.ClanMember{}, nil
}
func (totalsWG) AccountVehicles(context.Context, types.Realm, string, []string, ...types.Option) ([]types.VehicleStatsFrame, error) {
	return []types.VehicleStatsFrame{
		{TankID: 1, Stats: types.StatsFrame{Battles: 12, Wins: 6, DamageDealt: 14000}},
		{TankID: 2, Stats: types.StatsFrame{Battles: 2, Wins: 1, DamageDealt: 4000}},
	}, nil
}
