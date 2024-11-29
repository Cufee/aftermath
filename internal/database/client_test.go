package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func TestConcurrentWrites(t *testing.T) {
	is := is.New(t)

	client, err := NewSQLiteClient(filepath.Join(os.Getenv("DATABASE_PATH"), os.Getenv("DATABASE_NAME")), WithDebug())
	assert.NoError(t, err, "new client should not error")
	defer client.Disconnect()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	client.db.VehicleSnapshot.Delete().Where().Exec(ctx)
	defer client.db.VehicleSnapshot.Delete().Exec(ctx)

	createdAtVehicle1 := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)

	vehicle1 := models.VehicleSnapshot{
		Type:           models.SnapshotTypeDaily,
		CreatedAt:      createdAtVehicle1,
		LastBattleTime: createdAtVehicle1,
	}

	var fn []func()
	for i := range 1000 {
		id := fmt.Sprint(i)
		v := vehicle1
		v.VehicleID = id
		v.AccountID = id
		v.ReferenceID = id
		fn = append(fn, func() {
			_, err := client.UpsertAccounts(ctx, &models.Account{ID: id, Realm: "test", Nickname: "test_account"})
			is.NoErr(err)

			err = client.CreateAccountVehicleSnapshots(ctx, id, &v)
			is.NoErr(err)
		})
	}

	var wg sync.WaitGroup
	for _, fn := range fn {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}

	wg.Wait()
}
