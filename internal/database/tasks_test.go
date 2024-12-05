package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/table"
)

func TestCronTasks(t *testing.T) {
	client := MustTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.CronTask.TableName()))
	//
	_ = ctx
	_ = client
}
