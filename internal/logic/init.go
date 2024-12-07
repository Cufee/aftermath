package logic

import (
	"context"
	"os"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/pkg/errors"
)

func ApplyInitOptions(db database.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if id := os.Getenv("INIT_GLOBAL_ADMIN_USER"); id != "" {
		_, err := db.GetOrCreateUserByID(ctx, id, database.AddPermissions(permissions.GlobalAdmin))
		if err != nil {
			return errors.Wrap(err, "failed to upsert a user with permissions")
		}
	}
	return nil
}
