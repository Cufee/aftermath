package database

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

/*
DATABASE_PATH, DATABASE_NAME need to be set, and migrations need to be applied
*/
func TestGetVehicleSnapshots(t *testing.T) {
	client, err := NewSQLiteClient(filepath.Join(os.Getenv("DATABASE_PATH"), os.Getenv("DATABASE_NAME")), WithDebug())
	assert.NoError(t, err, "new client should not error")
	defer client.Disconnect()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	client.db.VehicleSnapshot.Delete().Where().Exec(ctx)
	defer client.db.VehicleSnapshot.Delete().Exec(ctx)

	_, err = client.UpsertAccounts(ctx, []models.Account{{ID: "a1", Realm: "test", Nickname: "test_account"}})
	assert.NoError(t, err, "failed to upsert an account")

	createdAtVehicle1 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	createdAtVehicle2 := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
	createdAtVehicle3 := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

	createdAtVehicle4 := time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC)
	createdAtVehicle5 := time.Date(2023, 9, 2, 0, 0, 0, 0, time.UTC)

	vehicle1 := models.VehicleSnapshot{
		VehicleID:   "v1",
		AccountID:   "a1",
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle1,
		LastBattleTime: createdAtVehicle1,
	}
	vehicle2 := models.VehicleSnapshot{
		VehicleID:   "v1",
		AccountID:   "a1",
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle2,
		LastBattleTime: createdAtVehicle2,
	}
	vehicle3 := models.VehicleSnapshot{
		VehicleID:   "v1",
		AccountID:   "a1",
		ReferenceID: "r1",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle3,
		LastBattleTime: createdAtVehicle3,
	}
	vehicle4 := models.VehicleSnapshot{
		VehicleID:   "v4",
		AccountID:   "a1",
		ReferenceID: "r2",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle4,
		LastBattleTime: createdAtVehicle4,
	}
	vehicle5 := models.VehicleSnapshot{
		VehicleID:   "v5",
		AccountID:   "a1",
		ReferenceID: "r2",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle5,
		LastBattleTime: createdAtVehicle5,
	}
	vehicle6 := models.VehicleSnapshot{
		VehicleID:   "v5",
		AccountID:   "a1",
		ReferenceID: "r2",

		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle5,
		LastBattleTime: createdAtVehicle5,
	}

	{ // create snapshots
		snapshots := []models.VehicleSnapshot{vehicle1, vehicle2, vehicle3, vehicle4, vehicle5, vehicle6}
		aErr, err := client.CreateAccountVehicleSnapshots(ctx, "a1", snapshots...)
		assert.NoError(t, err, "create vehicle snapshot should not error")
		assert.Nil(t, aErr, "insert returned some errors")
	}
	{ // when we check created after, vehicles need to be ordered by createdAt ASC, so we expect to get vehicle2 back
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", nil, models.SnapshotTypeDaily, WithCreatedAfter(createdAtVehicle1), WithReferenceIDIn("r1"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 1, "should return exactly 1 snapshot")
		assert.True(t, vehicles[0].CreatedAt.Equal(createdAtVehicle2), "wrong vehicle snapshot returned\nvehicles:%#v\nexpected:%#v", vehicles, createdAtVehicle2)
	}
	{ // when we check created before, vehicles need to be ordered by createdAt DESC, so we expect to get vehicle2 back
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle3), WithReferenceIDIn("r1"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 1, "should return exactly 1 snapshot")
		assert.True(t, vehicles[0].CreatedAt.Equal(createdAtVehicle2), "wrong vehicle snapshot returned\nvehicles:%#v\nexpected:%#v", vehicles, createdAtVehicle2)
	}
	{ // make sure only 1 vehicle is returned per ID
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle5.Add(time.Hour)), WithReferenceIDIn("r2"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 2, "should return exactly 2 snapshot")
		assert.NotEqual(t, vehicles[0].ID, vehicles[1].ID, "each vehicle id should only be returned once\nvehicles:%#v", vehicles)
	}
	{ // get a vehicle with a specific id
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", []string{"v5"}, models.SnapshotTypeDaily, WithReferenceIDIn("r2"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 1, "should return exactly 1 snapshot")
		assert.NotEqual(t, vehicles[0].ID, "v5", "incorrect vehicle returned\nvehicles:%#v", vehicles)
	}
	{ // this should return no result
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle1), WithReferenceIDIn("r1"))
		assert.NoError(t, err, "no results from a raw query does not trigger an error")
		assert.Len(t, vehicles, 0, "return should have no results\nvehicles:%#v", vehicles)
	}
	{ // this should return 3 results
		lastBattles, err := client.GetVehicleLastBattleTimes(ctx, "a1", nil, models.SnapshotTypeDaily)
		assert.NoError(t, err, "no results from a raw query does not trigger an error")
		assert.Len(t, lastBattles, 3, "return should have 3 results\nvehicles:%#v", lastBattles)
		assert.Equal(t, createdAtVehicle3, lastBattles["v1"])
		assert.Equal(t, createdAtVehicle4, lastBattles["v4"])
		assert.Equal(t, createdAtVehicle5, lastBattles["v5"])
	}
}
