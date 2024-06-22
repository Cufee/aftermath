// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "last_battle_time", Type: field.TypeInt},
		{Name: "account_created_at", Type: field.TypeInt},
		{Name: "realm", Type: field.TypeString, Size: 5},
		{Name: "nickname", Type: field.TypeString},
		{Name: "private", Type: field.TypeBool, Default: false},
		{Name: "clan_id", Type: field.TypeString, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_clans_accounts",
				Columns:    []*schema.Column{AccountsColumns[8]},
				RefColumns: []*schema.Column{ClansColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "account_id_last_battle_time",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[0], AccountsColumns[3]},
			},
			{
				Name:    "account_realm",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[5]},
			},
			{
				Name:    "account_realm_last_battle_time",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[5], AccountsColumns[3]},
			},
			{
				Name:    "account_clan_id",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[8]},
			},
		},
	}
	// AccountSnapshotsColumns holds the columns for the "account_snapshots" table.
	AccountSnapshotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"live", "daily"}},
		{Name: "last_battle_time", Type: field.TypeInt},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "rating_battles", Type: field.TypeInt},
		{Name: "rating_frame", Type: field.TypeJSON},
		{Name: "regular_battles", Type: field.TypeInt},
		{Name: "regular_frame", Type: field.TypeJSON},
		{Name: "account_id", Type: field.TypeString},
	}
	// AccountSnapshotsTable holds the schema information for the "account_snapshots" table.
	AccountSnapshotsTable = &schema.Table{
		Name:       "account_snapshots",
		Columns:    AccountSnapshotsColumns,
		PrimaryKey: []*schema.Column{AccountSnapshotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_snapshots_accounts_snapshots",
				Columns:    []*schema.Column{AccountSnapshotsColumns[10]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "accountsnapshot_created_at",
				Unique:  false,
				Columns: []*schema.Column{AccountSnapshotsColumns[1]},
			},
			{
				Name:    "accountsnapshot_type_account_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{AccountSnapshotsColumns[3], AccountSnapshotsColumns[10], AccountSnapshotsColumns[1]},
			},
			{
				Name:    "accountsnapshot_type_account_id_reference_id",
				Unique:  false,
				Columns: []*schema.Column{AccountSnapshotsColumns[3], AccountSnapshotsColumns[10], AccountSnapshotsColumns[5]},
			},
			{
				Name:    "accountsnapshot_type_account_id_reference_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{AccountSnapshotsColumns[3], AccountSnapshotsColumns[10], AccountSnapshotsColumns[5], AccountSnapshotsColumns[1]},
			},
		},
	}
	// AchievementsSnapshotsColumns holds the columns for the "achievements_snapshots" table.
	AchievementsSnapshotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"live", "daily"}},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "battles", Type: field.TypeInt},
		{Name: "last_battle_time", Type: field.TypeInt},
		{Name: "data", Type: field.TypeJSON},
		{Name: "account_id", Type: field.TypeString},
	}
	// AchievementsSnapshotsTable holds the schema information for the "achievements_snapshots" table.
	AchievementsSnapshotsTable = &schema.Table{
		Name:       "achievements_snapshots",
		Columns:    AchievementsSnapshotsColumns,
		PrimaryKey: []*schema.Column{AchievementsSnapshotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "achievements_snapshots_accounts_achievement_snapshots",
				Columns:    []*schema.Column{AchievementsSnapshotsColumns[8]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "achievementssnapshot_created_at",
				Unique:  false,
				Columns: []*schema.Column{AchievementsSnapshotsColumns[1]},
			},
			{
				Name:    "achievementssnapshot_account_id_reference_id",
				Unique:  false,
				Columns: []*schema.Column{AchievementsSnapshotsColumns[8], AchievementsSnapshotsColumns[4]},
			},
			{
				Name:    "achievementssnapshot_account_id_reference_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{AchievementsSnapshotsColumns[8], AchievementsSnapshotsColumns[4], AchievementsSnapshotsColumns[1]},
			},
		},
	}
	// AppConfigurationsColumns holds the columns for the "app_configurations" table.
	AppConfigurationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "value", Type: field.TypeJSON},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
	}
	// AppConfigurationsTable holds the schema information for the "app_configurations" table.
	AppConfigurationsTable = &schema.Table{
		Name:       "app_configurations",
		Columns:    AppConfigurationsColumns,
		PrimaryKey: []*schema.Column{AppConfigurationsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appconfiguration_key",
				Unique:  false,
				Columns: []*schema.Column{AppConfigurationsColumns[3]},
			},
		},
	}
	// ApplicationCommandsColumns holds the columns for the "application_commands" table.
	ApplicationCommandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "version", Type: field.TypeString},
		{Name: "options_hash", Type: field.TypeString},
	}
	// ApplicationCommandsTable holds the schema information for the "application_commands" table.
	ApplicationCommandsTable = &schema.Table{
		Name:       "application_commands",
		Columns:    ApplicationCommandsColumns,
		PrimaryKey: []*schema.Column{ApplicationCommandsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "applicationcommand_options_hash",
				Unique:  false,
				Columns: []*schema.Column{ApplicationCommandsColumns[5]},
			},
		},
	}
	// ClansColumns holds the columns for the "clans" table.
	ClansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "tag", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "emblem_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "members", Type: field.TypeJSON},
	}
	// ClansTable holds the schema information for the "clans" table.
	ClansTable = &schema.Table{
		Name:       "clans",
		Columns:    ClansColumns,
		PrimaryKey: []*schema.Column{ClansColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "clan_tag",
				Unique:  false,
				Columns: []*schema.Column{ClansColumns[3]},
			},
			{
				Name:    "clan_name",
				Unique:  false,
				Columns: []*schema.Column{ClansColumns[4]},
			},
		},
	}
	// CronTasksColumns holds the columns for the "cron_tasks" table.
	CronTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "type", Type: field.TypeString},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "targets", Type: field.TypeJSON},
		{Name: "status", Type: field.TypeString},
		{Name: "scheduled_after", Type: field.TypeInt},
		{Name: "last_run", Type: field.TypeInt},
		{Name: "logs", Type: field.TypeJSON},
		{Name: "data", Type: field.TypeJSON},
	}
	// CronTasksTable holds the schema information for the "cron_tasks" table.
	CronTasksTable = &schema.Table{
		Name:       "cron_tasks",
		Columns:    CronTasksColumns,
		PrimaryKey: []*schema.Column{CronTasksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "crontask_reference_id",
				Unique:  false,
				Columns: []*schema.Column{CronTasksColumns[4]},
			},
			{
				Name:    "crontask_status_last_run",
				Unique:  false,
				Columns: []*schema.Column{CronTasksColumns[6], CronTasksColumns[8]},
			},
			{
				Name:    "crontask_status_created_at",
				Unique:  false,
				Columns: []*schema.Column{CronTasksColumns[6], CronTasksColumns[1]},
			},
			{
				Name:    "crontask_status_scheduled_after",
				Unique:  false,
				Columns: []*schema.Column{CronTasksColumns[6], CronTasksColumns[7]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "permissions", Type: field.TypeString, Default: ""},
		{Name: "feature_flags", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserConnectionsColumns holds the columns for the "user_connections" table.
	UserConnectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"wargaming"}},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "permissions", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "user_id", Type: field.TypeString},
	}
	// UserConnectionsTable holds the schema information for the "user_connections" table.
	UserConnectionsTable = &schema.Table{
		Name:       "user_connections",
		Columns:    UserConnectionsColumns,
		PrimaryKey: []*schema.Column{UserConnectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_connections_users_connections",
				Columns:    []*schema.Column{UserConnectionsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userconnection_user_id",
				Unique:  false,
				Columns: []*schema.Column{UserConnectionsColumns[7]},
			},
			{
				Name:    "userconnection_type_user_id",
				Unique:  false,
				Columns: []*schema.Column{UserConnectionsColumns[3], UserConnectionsColumns[7]},
			},
			{
				Name:    "userconnection_reference_id",
				Unique:  false,
				Columns: []*schema.Column{UserConnectionsColumns[4]},
			},
			{
				Name:    "userconnection_type_reference_id",
				Unique:  false,
				Columns: []*schema.Column{UserConnectionsColumns[3], UserConnectionsColumns[4]},
			},
		},
	}
	// UserContentsColumns holds the columns for the "user_contents" table.
	UserContentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"clan-background-image", "personal-background-image"}},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "value", Type: field.TypeJSON},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "user_id", Type: field.TypeString},
	}
	// UserContentsTable holds the schema information for the "user_contents" table.
	UserContentsTable = &schema.Table{
		Name:       "user_contents",
		Columns:    UserContentsColumns,
		PrimaryKey: []*schema.Column{UserContentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_contents_users_content",
				Columns:    []*schema.Column{UserContentsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usercontent_user_id",
				Unique:  false,
				Columns: []*schema.Column{UserContentsColumns[7]},
			},
			{
				Name:    "usercontent_type_user_id",
				Unique:  false,
				Columns: []*schema.Column{UserContentsColumns[3], UserContentsColumns[7]},
			},
			{
				Name:    "usercontent_reference_id",
				Unique:  false,
				Columns: []*schema.Column{UserContentsColumns[4]},
			},
			{
				Name:    "usercontent_type_reference_id",
				Unique:  false,
				Columns: []*schema.Column{UserContentsColumns[3], UserContentsColumns[4]},
			},
		},
	}
	// UserSubscriptionsColumns holds the columns for the "user_subscriptions" table.
	UserSubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"aftermath-pro", "aftermath-pro-clan", "aftermath-plus", "supporter", "verified-clan", "server-moderator", "content-moderator", "developer", "server-booster", "content-translator"}},
		{Name: "expires_at", Type: field.TypeInt},
		{Name: "permissions", Type: field.TypeString},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// UserSubscriptionsTable holds the schema information for the "user_subscriptions" table.
	UserSubscriptionsTable = &schema.Table{
		Name:       "user_subscriptions",
		Columns:    UserSubscriptionsColumns,
		PrimaryKey: []*schema.Column{UserSubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_subscriptions_users_subscriptions",
				Columns:    []*schema.Column{UserSubscriptionsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usersubscription_user_id",
				Unique:  false,
				Columns: []*schema.Column{UserSubscriptionsColumns[7]},
			},
			{
				Name:    "usersubscription_type_user_id",
				Unique:  false,
				Columns: []*schema.Column{UserSubscriptionsColumns[3], UserSubscriptionsColumns[7]},
			},
			{
				Name:    "usersubscription_expires_at",
				Unique:  false,
				Columns: []*schema.Column{UserSubscriptionsColumns[4]},
			},
			{
				Name:    "usersubscription_expires_at_user_id",
				Unique:  false,
				Columns: []*schema.Column{UserSubscriptionsColumns[4], UserSubscriptionsColumns[7]},
			},
		},
	}
	// VehiclesColumns holds the columns for the "vehicles" table.
	VehiclesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "tier", Type: field.TypeInt},
		{Name: "localized_names", Type: field.TypeJSON},
	}
	// VehiclesTable holds the schema information for the "vehicles" table.
	VehiclesTable = &schema.Table{
		Name:       "vehicles",
		Columns:    VehiclesColumns,
		PrimaryKey: []*schema.Column{VehiclesColumns[0]},
	}
	// VehicleAveragesColumns holds the columns for the "vehicle_averages" table.
	VehicleAveragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "data", Type: field.TypeJSON},
	}
	// VehicleAveragesTable holds the schema information for the "vehicle_averages" table.
	VehicleAveragesTable = &schema.Table{
		Name:       "vehicle_averages",
		Columns:    VehicleAveragesColumns,
		PrimaryKey: []*schema.Column{VehicleAveragesColumns[0]},
	}
	// VehicleSnapshotsColumns holds the columns for the "vehicle_snapshots" table.
	VehicleSnapshotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeInt},
		{Name: "updated_at", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"live", "daily"}},
		{Name: "vehicle_id", Type: field.TypeString},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "battles", Type: field.TypeInt},
		{Name: "last_battle_time", Type: field.TypeInt},
		{Name: "frame", Type: field.TypeJSON},
		{Name: "account_id", Type: field.TypeString},
	}
	// VehicleSnapshotsTable holds the schema information for the "vehicle_snapshots" table.
	VehicleSnapshotsTable = &schema.Table{
		Name:       "vehicle_snapshots",
		Columns:    VehicleSnapshotsColumns,
		PrimaryKey: []*schema.Column{VehicleSnapshotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vehicle_snapshots_accounts_vehicle_snapshots",
				Columns:    []*schema.Column{VehicleSnapshotsColumns[9]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "vehiclesnapshot_created_at",
				Unique:  false,
				Columns: []*schema.Column{VehicleSnapshotsColumns[1]},
			},
			{
				Name:    "vehiclesnapshot_vehicle_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{VehicleSnapshotsColumns[4], VehicleSnapshotsColumns[1]},
			},
			{
				Name:    "vehiclesnapshot_account_id_created_at",
				Unique:  false,
				Columns: []*schema.Column{VehicleSnapshotsColumns[9], VehicleSnapshotsColumns[1]},
			},
			{
				Name:    "vehiclesnapshot_account_id_type_created_at",
				Unique:  false,
				Columns: []*schema.Column{VehicleSnapshotsColumns[9], VehicleSnapshotsColumns[3], VehicleSnapshotsColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AccountSnapshotsTable,
		AchievementsSnapshotsTable,
		AppConfigurationsTable,
		ApplicationCommandsTable,
		ClansTable,
		CronTasksTable,
		UsersTable,
		UserConnectionsTable,
		UserContentsTable,
		UserSubscriptionsTable,
		VehiclesTable,
		VehicleAveragesTable,
		VehicleSnapshotsTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = ClansTable
	AccountSnapshotsTable.ForeignKeys[0].RefTable = AccountsTable
	AchievementsSnapshotsTable.ForeignKeys[0].RefTable = AccountsTable
	UserConnectionsTable.ForeignKeys[0].RefTable = UsersTable
	UserContentsTable.ForeignKeys[0].RefTable = UsersTable
	UserSubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	VehicleSnapshotsTable.ForeignKeys[0].RefTable = AccountsTable
}
