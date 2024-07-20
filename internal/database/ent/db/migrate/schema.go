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
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_battle_time", Type: field.TypeTime},
		{Name: "account_created_at", Type: field.TypeTime},
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
				Name:    "account_id",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[0]},
			},
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
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"live", "daily"}},
		{Name: "last_battle_time", Type: field.TypeTime},
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
				Symbol:     "account_snapshots_accounts_account_snapshots",
				Columns:    []*schema.Column{AccountSnapshotsColumns[10]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "accountsnapshot_id",
				Unique:  false,
				Columns: []*schema.Column{AccountSnapshotsColumns[0]},
			},
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
	// AppConfigurationsColumns holds the columns for the "app_configurations" table.
	AppConfigurationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
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
				Name:    "appconfiguration_id",
				Unique:  false,
				Columns: []*schema.Column{AppConfigurationsColumns[0]},
			},
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
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
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
				Name:    "applicationcommand_id",
				Unique:  false,
				Columns: []*schema.Column{ApplicationCommandsColumns[0]},
			},
			{
				Name:    "applicationcommand_options_hash",
				Unique:  false,
				Columns: []*schema.Column{ApplicationCommandsColumns[5]},
			},
		},
	}
	// AuthNoncesColumns holds the columns for the "auth_nonces" table.
	AuthNoncesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "identifier", Type: field.TypeString},
		{Name: "public_id", Type: field.TypeString, Unique: true},
		{Name: "metadata", Type: field.TypeJSON},
	}
	// AuthNoncesTable holds the schema information for the "auth_nonces" table.
	AuthNoncesTable = &schema.Table{
		Name:       "auth_nonces",
		Columns:    AuthNoncesColumns,
		PrimaryKey: []*schema.Column{AuthNoncesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "authnonce_public_id_active_expires_at",
				Unique:  false,
				Columns: []*schema.Column{AuthNoncesColumns[6], AuthNoncesColumns[3], AuthNoncesColumns[4]},
			},
		},
	}
	// ClansColumns holds the columns for the "clans" table.
	ClansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
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
				Name:    "clan_id",
				Unique:  false,
				Columns: []*schema.Column{ClansColumns[0]},
			},
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
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UPDATE_CLANS", "RECORD_SNAPSHOTS", "ACHIEVEMENT_LEADERBOARDS", "CLEANUP_DATABASE"}},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "targets", Type: field.TypeJSON},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"TASK_SCHEDULED", "TASK_IN_PROGRESS", "TASK_COMPLETE", "TASK_FAILED"}},
		{Name: "scheduled_after", Type: field.TypeTime},
		{Name: "last_run", Type: field.TypeTime},
		{Name: "tries_left", Type: field.TypeInt, Default: 0},
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
				Name:    "crontask_id",
				Unique:  false,
				Columns: []*schema.Column{CronTasksColumns[0]},
			},
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
	// DiscordInteractionsColumns holds the columns for the "discord_interactions" table.
	DiscordInteractionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "command", Type: field.TypeString},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"stats"}},
		{Name: "locale", Type: field.TypeString},
		{Name: "options", Type: field.TypeJSON},
		{Name: "user_id", Type: field.TypeString},
	}
	// DiscordInteractionsTable holds the schema information for the "discord_interactions" table.
	DiscordInteractionsTable = &schema.Table{
		Name:       "discord_interactions",
		Columns:    DiscordInteractionsColumns,
		PrimaryKey: []*schema.Column{DiscordInteractionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "discord_interactions_users_discord_interactions",
				Columns:    []*schema.Column{DiscordInteractionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "discordinteraction_id",
				Unique:  false,
				Columns: []*schema.Column{DiscordInteractionsColumns[0]},
			},
			{
				Name:    "discordinteraction_command",
				Unique:  false,
				Columns: []*schema.Column{DiscordInteractionsColumns[3]},
			},
			{
				Name:    "discordinteraction_user_id",
				Unique:  false,
				Columns: []*schema.Column{DiscordInteractionsColumns[8]},
			},
			{
				Name:    "discordinteraction_user_id_type",
				Unique:  false,
				Columns: []*schema.Column{DiscordInteractionsColumns[8], DiscordInteractionsColumns[5]},
			},
			{
				Name:    "discordinteraction_reference_id",
				Unique:  false,
				Columns: []*schema.Column{DiscordInteractionsColumns[4]},
			},
		},
	}
	// GameMapsColumns holds the columns for the "game_maps" table.
	GameMapsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "game_modes", Type: field.TypeJSON},
		{Name: "supremacy_points", Type: field.TypeInt},
		{Name: "localized_names", Type: field.TypeJSON},
	}
	// GameMapsTable holds the schema information for the "game_maps" table.
	GameMapsTable = &schema.Table{
		Name:       "game_maps",
		Columns:    GameMapsColumns,
		PrimaryKey: []*schema.Column{GameMapsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "gamemap_id",
				Unique:  false,
				Columns: []*schema.Column{GameMapsColumns[0]},
			},
		},
	}
	// GameModesColumns holds the columns for the "game_modes" table.
	GameModesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "localized_names", Type: field.TypeJSON},
	}
	// GameModesTable holds the schema information for the "game_modes" table.
	GameModesTable = &schema.Table{
		Name:       "game_modes",
		Columns:    GameModesColumns,
		PrimaryKey: []*schema.Column{GameModesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "gamemode_id",
				Unique:  false,
				Columns: []*schema.Column{GameModesColumns[0]},
			},
		},
	}
	// LeaderboardScoresColumns holds the columns for the "leaderboard_scores" table.
	LeaderboardScoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"hourly", "daily"}},
		{Name: "score", Type: field.TypeFloat32},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "leaderboard_id", Type: field.TypeString},
		{Name: "meta", Type: field.TypeJSON},
	}
	// LeaderboardScoresTable holds the schema information for the "leaderboard_scores" table.
	LeaderboardScoresTable = &schema.Table{
		Name:       "leaderboard_scores",
		Columns:    LeaderboardScoresColumns,
		PrimaryKey: []*schema.Column{LeaderboardScoresColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "leaderboardscore_id",
				Unique:  false,
				Columns: []*schema.Column{LeaderboardScoresColumns[0]},
			},
			{
				Name:    "leaderboardscore_created_at",
				Unique:  false,
				Columns: []*schema.Column{LeaderboardScoresColumns[1]},
			},
			{
				Name:    "leaderboardscore_reference_id",
				Unique:  false,
				Columns: []*schema.Column{LeaderboardScoresColumns[5]},
			},
			{
				Name:    "leaderboardscore_leaderboard_id_type_reference_id",
				Unique:  false,
				Columns: []*schema.Column{LeaderboardScoresColumns[6], LeaderboardScoresColumns[3], LeaderboardScoresColumns[5]},
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "public_id", Type: field.TypeString, Unique: true},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "user_id", Type: field.TypeString},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "session_public_id_expires_at",
				Unique:  false,
				Columns: []*schema.Column{SessionsColumns[4], SessionsColumns[3]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Default: ""},
		{Name: "permissions", Type: field.TypeString, Default: ""},
		{Name: "feature_flags", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[0]},
			},
			{
				Name:    "user_username",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// UserConnectionsColumns holds the columns for the "user_connections" table.
	UserConnectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
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
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userconnection_id",
				Unique:  false,
				Columns: []*schema.Column{UserConnectionsColumns[0]},
			},
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
			{
				Name:    "userconnection_reference_id_user_id_type",
				Unique:  true,
				Columns: []*schema.Column{UserConnectionsColumns[4], UserConnectionsColumns[7], UserConnectionsColumns[3]},
			},
		},
	}
	// UserContentsColumns holds the columns for the "user_contents" table.
	UserContentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"clan-background-image", "personal-background-image"}},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
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
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usercontent_id",
				Unique:  false,
				Columns: []*schema.Column{UserContentsColumns[0]},
			},
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
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"aftermath-pro", "aftermath-pro-clan", "aftermath-plus", "supporter", "verified-clan", "server-moderator", "content-moderator", "developer", "server-booster", "content-translator"}},
		{Name: "expires_at", Type: field.TypeTime},
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
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usersubscription_id",
				Unique:  false,
				Columns: []*schema.Column{UserSubscriptionsColumns[0]},
			},
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
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "tier", Type: field.TypeInt},
		{Name: "localized_names", Type: field.TypeJSON},
	}
	// VehiclesTable holds the schema information for the "vehicles" table.
	VehiclesTable = &schema.Table{
		Name:       "vehicles",
		Columns:    VehiclesColumns,
		PrimaryKey: []*schema.Column{VehiclesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vehicle_id",
				Unique:  false,
				Columns: []*schema.Column{VehiclesColumns[0]},
			},
		},
	}
	// VehicleAveragesColumns holds the columns for the "vehicle_averages" table.
	VehicleAveragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "data", Type: field.TypeJSON},
	}
	// VehicleAveragesTable holds the schema information for the "vehicle_averages" table.
	VehicleAveragesTable = &schema.Table{
		Name:       "vehicle_averages",
		Columns:    VehicleAveragesColumns,
		PrimaryKey: []*schema.Column{VehicleAveragesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vehicleaverage_id",
				Unique:  false,
				Columns: []*schema.Column{VehicleAveragesColumns[0]},
			},
		},
	}
	// VehicleSnapshotsColumns holds the columns for the "vehicle_snapshots" table.
	VehicleSnapshotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"live", "daily"}},
		{Name: "vehicle_id", Type: field.TypeString},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "battles", Type: field.TypeInt},
		{Name: "last_battle_time", Type: field.TypeTime},
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
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "vehiclesnapshot_id",
				Unique:  false,
				Columns: []*schema.Column{VehicleSnapshotsColumns[0]},
			},
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
	// WidgetSettingsColumns holds the columns for the "widget_settings" table.
	WidgetSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "reference_id", Type: field.TypeString},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "session_from", Type: field.TypeTime, Nullable: true},
		{Name: "session_reference_id", Type: field.TypeString, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "styles", Type: field.TypeJSON},
		{Name: "user_id", Type: field.TypeString},
	}
	// WidgetSettingsTable holds the schema information for the "widget_settings" table.
	WidgetSettingsTable = &schema.Table{
		Name:       "widget_settings",
		Columns:    WidgetSettingsColumns,
		PrimaryKey: []*schema.Column{WidgetSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "widget_settings_users_widgets",
				Columns:    []*schema.Column{WidgetSettingsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "widgetsettings_id",
				Unique:  false,
				Columns: []*schema.Column{WidgetSettingsColumns[0]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AccountSnapshotsTable,
		AppConfigurationsTable,
		ApplicationCommandsTable,
		AuthNoncesTable,
		ClansTable,
		CronTasksTable,
		DiscordInteractionsTable,
		GameMapsTable,
		GameModesTable,
		LeaderboardScoresTable,
		SessionsTable,
		UsersTable,
		UserConnectionsTable,
		UserContentsTable,
		UserSubscriptionsTable,
		VehiclesTable,
		VehicleAveragesTable,
		VehicleSnapshotsTable,
		WidgetSettingsTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = ClansTable
	AccountSnapshotsTable.ForeignKeys[0].RefTable = AccountsTable
	DiscordInteractionsTable.ForeignKeys[0].RefTable = UsersTable
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
	UserConnectionsTable.ForeignKeys[0].RefTable = UsersTable
	UserContentsTable.ForeignKeys[0].RefTable = UsersTable
	UserSubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	VehicleSnapshotsTable.ForeignKeys[0].RefTable = AccountsTable
	WidgetSettingsTable.ForeignKeys[0].RefTable = UsersTable
}
