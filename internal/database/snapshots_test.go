package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func TestVehicleSnapshots(t *testing.T) {
	client := MustTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.VehicleSnapshot.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.VehicleSnapshot.TableName()))

	accountID := "a-TestVehicleSnapshots"
	err := client.UpsertAccounts(ctx, &models.Account{ID: accountID, Realm: "test", Nickname: "test_account"})
	assert.NoError(t, err, "failed to upsert an account")

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", table.Account.TableName(), accountID))

	createdAtVehicle1 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	createdAtVehicle2 := time.Date(2023, 8, 1, 0, 0, 0, 100, time.UTC)
	createdAtVehicle3 := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

	createdAtVehicle4 := time.Date(2023, 9, 1, 0, 0, 0, 100, time.UTC)
	createdAtVehicle5 := time.Date(2023, 9, 2, 0, 0, 0, 0, time.UTC)

	vehicle1 := models.VehicleSnapshot{
		VehicleID:   "v1",
		AccountID:   accountID,
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle1,
		LastBattleTime: createdAtVehicle1,
	}
	vehicle2 := models.VehicleSnapshot{
		VehicleID:   "v1",
		AccountID:   accountID,
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle2,
		LastBattleTime: createdAtVehicle2,
	}
	vehicle3 := models.VehicleSnapshot{
		VehicleID:   "v1",
		AccountID:   accountID,
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle3,
		LastBattleTime: createdAtVehicle3,
	}
	vehicle4 := models.VehicleSnapshot{
		VehicleID:   "v4",
		AccountID:   accountID,
		ReferenceID: "r2",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle4,
		LastBattleTime: createdAtVehicle4,
	}
	vehicle5 := models.VehicleSnapshot{
		VehicleID:   "v5",
		AccountID:   accountID,
		ReferenceID: "r2",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle4,
		LastBattleTime: createdAtVehicle4,
	}
	vehicle6 := models.VehicleSnapshot{
		VehicleID:   "v5",
		AccountID:   accountID,
		ReferenceID: "r2",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle5,
		LastBattleTime: createdAtVehicle5,
	}

	{ // create snapshots
		snapshots := []*models.VehicleSnapshot{&vehicle1, &vehicle2, &vehicle3, &vehicle4, &vehicle5, &vehicle6}
		err := client.CreateVehicleSnapshots(ctx, snapshots...)
		assert.NoError(t, err, "create vehicle snapshot should not error")
	}
	t.Run("vehicles need to be ordered by createdAt DESC when queried with created before", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, accountID, nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle3), WithReferenceIDIn("r1"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 1, "should return exactly 1 snapshot")
		assert.True(t, vehicles[0].CreatedAt.Equal(createdAtVehicle2), "wrong vehicle snapshot returned\nvehicles:%#v\nexpected:%#v", vehicles, createdAtVehicle2)
	})
	t.Run("only 1 vehicle is returned per ID", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, accountID, nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle5.Add(time.Hour)), WithReferenceIDIn("r2"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 2, "should return exactly 2 snapshot")
		assert.NotEqual(t, vehicles[0].ID, vehicles[1].ID, "each vehicle id should only be returned once\nvehicles:%#v", vehicles)
	})
	t.Run("get a vehicle with a specific id", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, accountID, []string{"v5"}, models.SnapshotTypeDaily, WithReferenceIDIn("r2"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 1, "should return exactly 1 snapshot")
		assert.NotEqual(t, vehicles[0].ID, "v5", "incorrect vehicle returned\nvehicles:%#v", vehicles)
	})
	t.Run("intentionally 0 result query", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, accountID, nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle1), WithReferenceIDIn("r1"))
		assert.NoError(t, err, "no results from a raw query does not trigger an error")
		assert.Len(t, vehicles, 0, "return should have no results\nvehicles:%#v", vehicles)
	})
	t.Run("get vehicle last battle time", func(t *testing.T) {
		vehicles, err := client.GetVehiclesLastBattleTimes(ctx, accountID, nil, models.SnapshotTypeDaily, WithReferenceIDIn("r2"))
		assert.NoError(t, err, "no results from a raw query does not trigger an error")
		assert.Len(t, vehicles, 2, "return should have exactly 2 results\nvehicles:%#v", vehicles)
		assert.Equal(t, vehicles[vehicle4.VehicleID], vehicle4.LastBattleTime, "vehicle4 last battle time check")
		assert.Equal(t, vehicles[vehicle6.VehicleID], vehicle6.LastBattleTime, "vehicle6 last battle time check")

	})
}

func TestAccountSnapshots(t *testing.T) {
	is := is.New(t)

	client := MustTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.AccountSnapshot.TableName()))

	accountID := "a-TestAccountSnapshots"
	accountID2 := "a-TestAccountSnapshots-2"
	err := client.UpsertAccounts(ctx, &models.Account{ID: accountID, Realm: "test", Nickname: "test_account"}, &models.Account{ID: accountID2, Realm: "test", Nickname: "test_account"})
	is.NoErr(err)

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id IN ('%s', '%s');", table.Account.TableName(), accountID, accountID2))

	createdAt1 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	createdAt2 := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
	createdAt3 := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

	snapshot1 := models.AccountSnapshot{
		AccountID:   accountID,
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAt1,
		LastBattleTime: createdAt1,
	}
	snapshot2 := models.AccountSnapshot{
		AccountID:   accountID,
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAt2,
		LastBattleTime: createdAt2,
	}
	snapshot3 := models.AccountSnapshot{
		AccountID:   accountID,
		ReferenceID: "r2",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAt3,
		LastBattleTime: createdAt3,
	}
	snapshot4 := models.AccountSnapshot{
		AccountID:   accountID2,
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAt3,
		LastBattleTime: createdAt3,
	}

	{
		err := client.CreateAccountSnapshots(context.Background(), &snapshot1, &snapshot2, &snapshot3, &snapshot4)
		is.NoErr(err)
	}
	t.Run("test get last battles", func(t *testing.T) {
		is := is.New(t)

		times, err := client.GetAccountLastBattleTimes(context.Background(), []string{accountID}, snapshot1.Type)
		is.NoErr(err)
		is.True(len(times) == 1)
		is.True(times[accountID].Unix() == createdAt3.Unix())
	})

	t.Run("test get with created before/after", func(t *testing.T) {
		is := is.New(t)

		snapshots, err := client.GetAccountSnapshots(context.Background(), []string{accountID}, snapshot1.Type, WithCreatedBefore(createdAt2))
		is.NoErr(err)
		is.True(len(snapshots) == 1)
		is.True(snapshots[0].CreatedAt.Unix() == createdAt1.Unix())
	})

	t.Run("test get with reference id", func(t *testing.T) {
		is := is.New(t)

		snapshots, err := client.GetAccountSnapshots(context.Background(), []string{accountID}, snapshot4.Type, WithReferenceIDIn("r2"))
		is.NoErr(err)
		is.True(len(snapshots) == 1)
		is.True(snapshots[0].CreatedAt.Unix() == createdAt3.Unix())
	})
}
