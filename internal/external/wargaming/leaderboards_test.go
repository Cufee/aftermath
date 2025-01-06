package wargaming

import (
	"context"
	"testing"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/matryer/is"
)

func TestRatingLeaderboardClient(t *testing.T) {
	is := is.New(t)

	client, err := NewRatingLeaderboardClient()
	is.NoErr(err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	season, err := client.CurrentSeason(ctx, types.RealmNorthAmerica)
	is.NoErr(err)
	is.True(len(season.Leagues) > 0)

	ctx2, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	players, err := client.LeagueTop(ctx2, types.RealmNorthAmerica, season.Leagues[0].ID)
	is.NoErr(err)
	is.True(len(players) > 0)

	ctx3, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	position, err := client.PlayerPosition(ctx3, types.RealmNorthAmerica, players[0].AccountID, 1)
	is.NoErr(err)
	is.True(position.AccountID == players[0].AccountID)
	is.True(position.Position == 1)
}
