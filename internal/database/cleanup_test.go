package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func TestSnapshotCleanup(t *testing.T) {
	is := is.New(t)

	client := MustTestClient(t)

	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.VehicleSnapshot.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.VehicleSnapshot.TableName()))

	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.AccountSnapshot.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.AccountSnapshot.TableName()))

	accountID := "a-TestSnapshotCleanup"
	_, err := client.UpsertAccounts(context.Background(), &models.Account{ID: accountID, Realm: "test", Nickname: "test_account"})
	assert.NoError(t, err, "failed to upsert an account")

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", table.Account.TableName(), accountID))

	t.Run("delete expired snapshot when a new one is present", func(t *testing.T) {
		is := is.New(t)

		id1, id2 := "s-1", "s-2"
		defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id IN ('%s', '%s');", table.VehicleSnapshot.TableName(), id1, id2))

		expiredTime := time.Now().Add(time.Hour * -1)
		expired := models.VehicleSnapshot{
			ID:          id1,
			VehicleID:   "v1",
			AccountID:   accountID,
			ReferenceID: "r1",

			Type:           models.SnapshotTypeDaily,
			CreatedAt:      expiredTime,
			LastBattleTime: expiredTime,
		}

		notExpiredTime := time.Now()
		notExpired := models.VehicleSnapshot{
			ID:          id2,
			VehicleID:   "v1",
			AccountID:   accountID,
			ReferenceID: "r1",

			Type:           models.SnapshotTypeDaily,
			CreatedAt:      notExpiredTime,
			LastBattleTime: notExpiredTime,
		}

		err := client.CreateVehicleSnapshots(context.Background(), &expired, &notExpired)
		is.NoErr(err)
		_, err = client.DeleteExpiredSnapshots(context.Background(), notExpiredTime.Add(time.Minute*-1))
		is.NoErr(err)
		{
			survived, err := client.GetVehicleSnapshots(context.Background(), accountID, nil, models.SnapshotTypeDaily)
			is.NoErr(err)
			is.True(len(survived) == 1)
			is.True(survived[0].ID == notExpired.ID)
		}
		{
			deleted, err := client.GetVehicleSnapshots(context.Background(), accountID, nil, models.SnapshotTypeDaily, WithCreatedBefore(expiredTime.Add(time.Minute)))
			is.NoErr(err)
			is.True(len(deleted) == 0)
		}
	})

	t.Run("do not delete expired snapshot when a new one is not present", func(t *testing.T) {
		is := is.New(t)

		id1, id2 := "s-3", "s-4"
		defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id IN ('%s', '%s');", table.VehicleSnapshot.TableName(), id1, id2))

		expiredTime := time.Now().Add(time.Hour * -1)
		expired := models.VehicleSnapshot{
			ID:          id1,
			VehicleID:   "v2",
			AccountID:   accountID,
			ReferenceID: "r2",

			Type:           models.SnapshotTypeDaily,
			CreatedAt:      expiredTime,
			LastBattleTime: expiredTime,
		}

		err := client.CreateVehicleSnapshots(context.Background(), &expired)
		is.NoErr(err)
		_, err = client.DeleteExpiredSnapshots(context.Background(), expiredTime.Add(time.Hour))
		is.NoErr(err)

		survived, err := client.GetVehicleSnapshots(context.Background(), accountID, nil, models.SnapshotTypeDaily)
		is.NoErr(err)
		is.True(len(survived) == 1)
		is.True(survived[0].ID == expired.ID)
	})
}
