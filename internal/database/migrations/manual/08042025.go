package manual

import (
	"context"
	"fmt"
	"os"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/gen/public/model"
	"github.com/cufee/aftermath/internal/database/gen/public/table"
	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func migration_08042025(ctx context.Context, client database.Client) error {
	shouldRun, cleanup, err := startMigration(ctx, client, "migration_08042025")
	if err != nil {
		return err
	}
	if !shouldRun {
		log.Debug().Str("key", "migration_08042025").Msg("migration already complete")
		return nil
	}
	dbPath := os.Getenv("DATABASE_PATH_08042025")
	if dbPath == "" {
		return nil
	}
	log.Debug().Str("key", "migration_08042025").Msg("running migration")

	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	{ // migrate account
		var count int
		row := db.QueryRow("SELECT COUNT(*) from account")
		err = row.Scan(&count)
		if err != nil {
			return errors.Wrap(err, "failed to get accounts count")
		}
		for i := 0; i < count; i += 1000 {
			accounts := make([]*model.Account, 0, 1000)
			err := db.Select(&accounts, fmt.Sprintf("SELECT * FROM account LIMIT 1000 OFFSET %d;", i))
			if len(accounts) > 0 {
				stmt := table.Account.INSERT(table.Account.AllColumns).MODELS(accounts).ON_CONFLICT(table.Account.ID).DO_NOTHING()
				_, err = stmt.ExecContext(ctx, client.Unsafe())
				if err != nil {
					return errors.Wrap(err, "failed to insert accounts")
				}
			}
			if len(accounts) < 1000 {
				break
			}
		}
	}

	{ // migrate app_configuration
		var records []*model.AppConfiguration
		err := db.Select(&records, "SELECT * from app_configuration")
		if err != nil {
			return errors.Wrap(err, "failed to get app_configuration")
		}
		if len(records) > 0 {
			stmt := table.AppConfiguration.INSERT(table.AppConfiguration.AllColumns).MODELS(records).ON_CONFLICT(table.AppConfiguration.ID).DO_NOTHING()
			_, err = stmt.ExecContext(ctx, client.Unsafe())
			if err != nil {
				return errors.Wrap(err, "failed to insert app configurations")
			}
		}
	}

	{ // migrate user
		var records []*model.User
		err := db.Select(&records, "SELECT * from user")
		if err != nil && !database.IsNotFound(err) {
			return errors.Wrap(err, "failed to get users")
		}
		if len(records) > 0 {
			stmt := table.User.INSERT(table.User.AllColumns).MODELS(records).ON_CONFLICT(table.User.ID).DO_NOTHING()
			_, err = stmt.ExecContext(ctx, client.Unsafe())
			if err != nil {
				return errors.Wrap(err, "failed to insert users")
			}
		}
	}

	{ // migrate user_connection
		var records []*model.UserConnection
		err := db.Select(&records, "SELECT * from user_connection")
		if err != nil && !database.IsNotFound(err) {
			return errors.Wrap(err, "failed to get user connections")
		}
		if len(records) > 0 {
			stmt := table.UserConnection.INSERT(table.UserConnection.AllColumns).MODELS(records).ON_CONFLICT(table.UserConnection.ReferenceID, table.UserConnection.UserID, table.UserConnection.Type).DO_NOTHING()
			_, err = stmt.ExecContext(ctx, client.Unsafe())
			if err != nil {
				return errors.Wrap(err, "failed to insert user connections")
			}
		}
	}

	{ // migrate user_subscription
		var records []*model.UserSubscription
		err := db.Select(&records, "SELECT * from user_subscription")
		if err != nil && !database.IsNotFound(err) {
			return errors.Wrap(err, "failed to get user subscriptions")
		}
		if len(records) > 0 {
			stmt := table.UserSubscription.INSERT(table.UserSubscription.AllColumns).MODELS(records).ON_CONFLICT(table.UserSubscription.ID).DO_NOTHING()
			_, err = stmt.ExecContext(ctx, client.Unsafe())
			if err != nil {
				return errors.Wrap(err, "failed to insert user subscriptions")
			}
		}
	}

	{ // migrate user_content
		var records []*model.UserContent
		err := db.Select(&records, "SELECT * from user_content")
		if err != nil && !database.IsNotFound(err) {
			return errors.Wrap(err, "failed to get user content")
		}
		if len(records) > 0 {
			stmt := table.UserContent.INSERT(table.UserContent.AllColumns).MODELS(records).ON_CONFLICT(table.UserContent.UserID, table.UserContent.Type).DO_NOTHING()
			_, err = stmt.ExecContext(ctx, client.Unsafe())
			if err != nil {
				return errors.Wrap(err, "failed to insert user content")
			}
		}
	}

	{ // migrate user_restriction
		var records []*model.UserRestriction
		err := db.Select(&records, "SELECT * from user_restriction")
		if err != nil && !database.IsNotFound(err) {
			return errors.Wrap(err, "failed to get user restrictions")
		}
		if len(records) > 0 {
			stmt := table.UserRestriction.INSERT(table.UserRestriction.AllColumns).MODELS(records).ON_CONFLICT(table.UserRestriction.ID).DO_NOTHING()
			_, err = stmt.ExecContext(ctx, client.Unsafe())
			if err != nil {
				return errors.Wrap(err, "failed to insert user restrictions")
			}
		}
	}

	{ // migrate moderation_request
		var records []*model.ModerationRequest
		err := db.Select(&records, "SELECT * from moderation_request")
		if err != nil && !database.IsNotFound(err) {
			return errors.Wrap(err, "failed to get moderation requests")
		}
		if len(records) > 0 {
			stmt := table.ModerationRequest.INSERT(table.ModerationRequest.AllColumns).MODELS(records).ON_CONFLICT(table.ModerationRequest.ID).DO_NOTHING()
			_, err = stmt.ExecContext(ctx, client.Unsafe())
			if err != nil {
				return errors.Wrap(err, "failed to insert moderation requests")
			}
		}
	}

	{ // migrate widget_settings
		var records []*model.WidgetSettings
		err := db.Select(&records, "SELECT * from widget_settings")
		if err != nil && !database.IsNotFound(err) {
			return errors.Wrap(err, "failed to get widget settings")
		}
		if len(records) > 0 {
			stmt := table.WidgetSettings.INSERT(table.WidgetSettings.AllColumns).MODELS(records).ON_CONFLICT(table.WidgetSettings.ID).DO_NOTHING()
			_, err = stmt.ExecContext(ctx, client.Unsafe())
			if err != nil {
				return errors.Wrap(err, "failed to insert widget settings")
			}
		}
	}

	// cleanup(ctx)
	_ = cleanup
	return nil
}

func migration_08042025_4(ctx context.Context, client database.Client) error {
	shouldRun, cleanup, err := startMigration(ctx, client, "migration_08042025_4")
	if err != nil {
		return err
	}
	if !shouldRun {
		log.Debug().Str("key", "migration_08042025_4").Msg("migration already complete")
		return nil
	}
	dbPath := os.Getenv("DATABASE_PATH_08042025")
	if dbPath == "" {
		return errors.New("missing required env variable")
	}
	log.Debug().Str("key", "migration_08042025_4").Msg("running migration")

	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	{ // migrate account_snapshot
		var count int
		row := db.QueryRow("SELECT COUNT(*) from account_snapshot")
		err = row.Scan(&count)
		if err != nil {
			return errors.Wrap(err, "failed to get account snapshots count")
		}
		for i := 0; i < count; i += 500 {
			records := make([]*model.AccountSnapshot, 0, 500)
			err := db.Select(&records, fmt.Sprintf("SELECT * FROM account_snapshot LIMIT 500 OFFSET %d;", i))
			if len(records) > 0 {
				stmt := table.AccountSnapshot.INSERT(table.AccountSnapshot.AllColumns).MODELS(records).ON_CONFLICT(table.AccountSnapshot.ID).DO_NOTHING()
				_, err = stmt.ExecContext(ctx, client.Unsafe())
				if err != nil {
					return errors.Wrap(err, "failed to insert account snapshots")
				}
			}
			if len(records) < 500 {
				break
			}
		}
	}

	{ // migrate vehicle_snapshot
		var count int
		row := db.QueryRow("SELECT COUNT(*) from vehicle_snapshot")
		err = row.Scan(&count)
		if err != nil {
			return errors.Wrap(err, "failed to get vehicle snapshots count")
		}
		for i := 0; i < count; i += 500 {
			records := make([]*model.VehicleSnapshot, 0, 500)
			err := db.Select(&records, fmt.Sprintf("SELECT * FROM vehicle_snapshot LIMIT 500 OFFSET %d;", i))
			if len(records) > 0 {
				stmt := table.VehicleSnapshot.INSERT(table.VehicleSnapshot.AllColumns).MODELS(records).ON_CONFLICT(table.VehicleSnapshot.ID).DO_NOTHING()
				_, err = stmt.ExecContext(ctx, client.Unsafe())
				if err != nil {
					return errors.Wrap(err, "failed to insert vehicle snapshots")
				}
			}
			if len(records) < 500 {
				break
			}
		}
	}

	cleanup(ctx)
	return nil
}
