package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
)

func TestCronTasks(t *testing.T) {
	is := is.New(t)

	client := MustTestClient(t)
	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.CronTask.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.CronTask.TableName()))

	task := models.Task{
		Type:        models.TaskTypeDatabaseCleanup,
		ReferenceID: "r-1",
		Targets:     []string{"t-1"},
	}
	err := client.CreateTasks(context.Background(), task)
	is.NoErr(err)

	t.Run("", func(t *testing.T) {
		is := is.New(t)
		_ = is

	})
}
