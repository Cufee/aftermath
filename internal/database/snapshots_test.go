package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/stretchr/testify/assert"
)

func TestVehicleSnapshots(t *testing.T) {
	client := MustTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_, err := client.UpsertAccounts(ctx, &models.Account{ID: "a1", Realm: "test", Nickname: "test_account"})
	assert.NoError(t, err, "failed to upsert an account")

	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.VehicleSnapshot.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.VehicleSnapshot.TableName()))

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
		snapshots := []*models.VehicleSnapshot{&vehicle1, &vehicle2, &vehicle3, &vehicle4, &vehicle5, &vehicle6}
		err := client.CreateVehicleSnapshots(ctx, snapshots...)
		assert.NoError(t, err, "create vehicle snapshot should not error")
	}
	t.Run("vehicles need to be ordered by createdAt ASC when queried with created after", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", nil, models.SnapshotTypeDaily, WithCreatedAfter(createdAtVehicle1.Add(time.Minute)), WithReferenceIDIn("r1"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 1, "should return exactly 1 snapshot")
		assert.True(t, vehicles[0].CreatedAt.Equal(createdAtVehicle2), "wrong vehicle snapshot returned\nvehicles:%#v\nexpected:%#v", vehicles, createdAtVehicle2)
	})
	t.Run("vehicles need to be ordered by createdAt DESC when queried with created before", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle3), WithReferenceIDIn("r1"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 1, "should return exactly 1 snapshot")
		assert.True(t, vehicles[0].CreatedAt.Equal(createdAtVehicle2), "wrong vehicle snapshot returned\nvehicles:%#v\nexpected:%#v", vehicles, createdAtVehicle2)
	})
	t.Run("only 1 vehicle is returned per ID", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle5.Add(time.Hour)), WithReferenceIDIn("r2"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 2, "should return exactly 2 snapshot")
		assert.NotEqual(t, vehicles[0].ID, vehicles[1].ID, "each vehicle id should only be returned once\nvehicles:%#v", vehicles)
	})
	t.Run("get a vehicle with a specific id", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", []string{"v5"}, models.SnapshotTypeDaily, WithReferenceIDIn("r2"))
		assert.NoError(t, err, "get vehicle snapshot error")
		assert.Len(t, vehicles, 1, "should return exactly 1 snapshot")
		assert.NotEqual(t, vehicles[0].ID, "v5", "incorrect vehicle returned\nvehicles:%#v", vehicles)
	})
	t.Run("intentionally 0 result query", func(t *testing.T) {
		vehicles, err := client.GetVehicleSnapshots(ctx, "a1", nil, models.SnapshotTypeDaily, WithCreatedBefore(createdAtVehicle1), WithReferenceIDIn("r1"))
		assert.NoError(t, err, "no results from a raw query does not trigger an error")
		assert.Len(t, vehicles, 0, "return should have no results\nvehicles:%#v", vehicles)
	})
}
