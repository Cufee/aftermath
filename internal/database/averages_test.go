package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/matryer/is"
)

func TestAverages(t *testing.T) {
	client := MustTestClient(t)
	is := is.New(t)

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.VehicleAverage.TableName()))

	t.Run("create and get vehicle averages", func(t *testing.T) {
		is := is.New(t)

		averages := map[string]frame.StatsFrame{
			"id-1": {Battles: 420},
			"id-2": {Battles: 228},
		}

		errors, err := client.UpsertVehicleAverages(context.Background(), averages)
		is.NoErr(err)
		for _, err := range errors {
			is.NoErr(err)
		}

		found, err := client.GetVehicleAverages(context.Background(), []string{"id-1", "id-2"})
		is.NoErr(err)
		is.True(len(found) == 2)
		is.True(found["id-1"].Battles == 420)
		is.True(found["id-2"].Battles == 228)
	})
}
