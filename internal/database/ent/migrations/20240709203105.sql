-- Create "accounts" table
CREATE TABLE `accounts` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `last_battle_time` datetime NOT NULL, `account_created_at` datetime NOT NULL, `realm` text NOT NULL, `nickname` text NOT NULL, `private` bool NOT NULL DEFAULT (false), `clan_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `accounts_clans_accounts` FOREIGN KEY (`clan_id`) REFERENCES `clans` (`id`) ON DELETE SET NULL);
-- Create index "account_id" to table: "accounts"
CREATE INDEX `account_id` ON `accounts` (`id`);
-- Create index "account_id_last_battle_time" to table: "accounts"
CREATE INDEX `account_id_last_battle_time` ON `accounts` (`id`, `last_battle_time`);
-- Create index "account_realm" to table: "accounts"
CREATE INDEX `account_realm` ON `accounts` (`realm`);
-- Create index "account_realm_last_battle_time" to table: "accounts"
CREATE INDEX `account_realm_last_battle_time` ON `accounts` (`realm`, `last_battle_time`);
-- Create index "account_clan_id" to table: "accounts"
CREATE INDEX `account_clan_id` ON `accounts` (`clan_id`);
-- Create "account_snapshots" table
CREATE TABLE `account_snapshots` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `last_battle_time` datetime NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` json NOT NULL, `regular_battles` integer NOT NULL, `regular_frame` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_snapshots_accounts_account_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE);
-- Create index "accountsnapshot_id" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_id` ON `account_snapshots` (`id`);
-- Create index "accountsnapshot_created_at" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_created_at` ON `account_snapshots` (`created_at`);
-- Create index "accountsnapshot_type_account_id_created_at" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_type_account_id_created_at` ON `account_snapshots` (`type`, `account_id`, `created_at`);
-- Create index "accountsnapshot_type_account_id_reference_id" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_type_account_id_reference_id` ON `account_snapshots` (`type`, `account_id`, `reference_id`);
-- Create index "accountsnapshot_type_account_id_reference_id_created_at" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_type_account_id_reference_id_created_at` ON `account_snapshots` (`type`, `account_id`, `reference_id`, `created_at`);
-- Create "achievements_snapshots" table
CREATE TABLE `achievements_snapshots` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` datetime NOT NULL, `data` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `achievements_snapshots_accounts_achievement_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE);
-- Create index "achievementssnapshot_id" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_id` ON `achievements_snapshots` (`id`);
-- Create index "achievementssnapshot_created_at" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_created_at` ON `achievements_snapshots` (`created_at`);
-- Create index "achievementssnapshot_account_id_reference_id" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_account_id_reference_id` ON `achievements_snapshots` (`account_id`, `reference_id`);
-- Create index "achievementssnapshot_account_id_reference_id_created_at" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_account_id_reference_id_created_at` ON `achievements_snapshots` (`account_id`, `reference_id`, `created_at`);
-- Create "app_configurations" table
CREATE TABLE `app_configurations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `key` text NOT NULL, `value` json NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- Create index "app_configurations_key_key" to table: "app_configurations"
CREATE UNIQUE INDEX `app_configurations_key_key` ON `app_configurations` (`key`);
-- Create index "appconfiguration_id" to table: "app_configurations"
CREATE INDEX `appconfiguration_id` ON `app_configurations` (`id`);
-- Create index "appconfiguration_key" to table: "app_configurations"
CREATE INDEX `appconfiguration_key` ON `app_configurations` (`key`);
-- Create "application_commands" table
CREATE TABLE `application_commands` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `name` text NOT NULL, `version` text NOT NULL, `options_hash` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "applicationcommand_id" to table: "application_commands"
CREATE INDEX `applicationcommand_id` ON `application_commands` (`id`);
-- Create index "applicationcommand_options_hash" to table: "application_commands"
CREATE INDEX `applicationcommand_options_hash` ON `application_commands` (`options_hash`);
-- Create "auth_nonces" table
CREATE TABLE `auth_nonces` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `active` bool NOT NULL, `expires_at` datetime NOT NULL, `identifier` text NOT NULL, `public_id` text NOT NULL, `metadata` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "auth_nonces_public_id_key" to table: "auth_nonces"
CREATE UNIQUE INDEX `auth_nonces_public_id_key` ON `auth_nonces` (`public_id`);
-- Create index "authnonce_public_id_active_expires_at" to table: "auth_nonces"
CREATE INDEX `authnonce_public_id_active_expires_at` ON `auth_nonces` (`public_id`, `active`, `expires_at`);
-- Create "clans" table
CREATE TABLE `clans` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `tag` text NOT NULL, `name` text NOT NULL, `emblem_id` text NULL DEFAULT (''), `members` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "clan_id" to table: "clans"
CREATE INDEX `clan_id` ON `clans` (`id`);
-- Create index "clan_tag" to table: "clans"
CREATE INDEX `clan_tag` ON `clans` (`tag`);
-- Create index "clan_name" to table: "clans"
CREATE INDEX `clan_name` ON `clans` (`name`);
-- Create "cron_tasks" table
CREATE TABLE `cron_tasks` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `targets` json NOT NULL, `status` text NOT NULL, `scheduled_after` datetime NOT NULL, `last_run` datetime NOT NULL, `tries_left` integer NOT NULL DEFAULT (0), `logs` json NOT NULL, `data` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "crontask_id" to table: "cron_tasks"
CREATE INDEX `crontask_id` ON `cron_tasks` (`id`);
-- Create index "crontask_reference_id" to table: "cron_tasks"
CREATE INDEX `crontask_reference_id` ON `cron_tasks` (`reference_id`);
-- Create index "crontask_status_last_run" to table: "cron_tasks"
CREATE INDEX `crontask_status_last_run` ON `cron_tasks` (`status`, `last_run`);
-- Create index "crontask_status_created_at" to table: "cron_tasks"
CREATE INDEX `crontask_status_created_at` ON `cron_tasks` (`status`, `created_at`);
-- Create index "crontask_status_scheduled_after" to table: "cron_tasks"
CREATE INDEX `crontask_status_scheduled_after` ON `cron_tasks` (`status`, `scheduled_after`);
-- Create "discord_interactions" table
CREATE TABLE `discord_interactions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `command` text NOT NULL, `reference_id` text NOT NULL, `type` text NOT NULL, `locale` text NOT NULL, `options` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `discord_interactions_users_discord_interactions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "discordinteraction_id" to table: "discord_interactions"
CREATE INDEX `discordinteraction_id` ON `discord_interactions` (`id`);
-- Create index "discordinteraction_command" to table: "discord_interactions"
CREATE INDEX `discordinteraction_command` ON `discord_interactions` (`command`);
-- Create index "discordinteraction_user_id" to table: "discord_interactions"
CREATE INDEX `discordinteraction_user_id` ON `discord_interactions` (`user_id`);
-- Create index "discordinteraction_user_id_type" to table: "discord_interactions"
CREATE INDEX `discordinteraction_user_id_type` ON `discord_interactions` (`user_id`, `type`);
-- Create index "discordinteraction_reference_id" to table: "discord_interactions"
CREATE INDEX `discordinteraction_reference_id` ON `discord_interactions` (`reference_id`);
-- Create "leaderboard_scores" table
CREATE TABLE `leaderboard_scores` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `score` real NOT NULL, `account_id` text NOT NULL, `reference_id` text NOT NULL, `leaderboard_id` text NOT NULL, `meta` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "leaderboardscore_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_id` ON `leaderboard_scores` (`id`);
-- Create index "leaderboardscore_created_at" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_created_at` ON `leaderboard_scores` (`created_at`);
-- Create index "leaderboardscore_account_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_account_id` ON `leaderboard_scores` (`account_id`);
-- Create index "leaderboardscore_reference_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_reference_id` ON `leaderboard_scores` (`reference_id`);
-- Create index "leaderboardscore_leaderboard_id_type_account_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_type_account_id` ON `leaderboard_scores` (`leaderboard_id`, `type`, `account_id`);
-- Create index "leaderboardscore_leaderboard_id_type_reference_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_type_reference_id` ON `leaderboard_scores` (`leaderboard_id`, `type`, `reference_id`);
-- Create "sessions" table
CREATE TABLE `sessions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `expires_at` datetime NOT NULL, `public_id` text NOT NULL, `metadata` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "sessions_public_id_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_public_id_key` ON `sessions` (`public_id`);
-- Create index "session_public_id_expires_at" to table: "sessions"
CREATE INDEX `session_public_id_expires_at` ON `sessions` (`public_id`, `expires_at`);
-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `username` text NOT NULL DEFAULT (''), `permissions` text NOT NULL DEFAULT (''), `feature_flags` json NULL, PRIMARY KEY (`id`));
-- Create index "user_id" to table: "users"
CREATE INDEX `user_id` ON `users` (`id`);
-- Create index "user_username" to table: "users"
CREATE INDEX `user_username` ON `users` (`username`);
-- Create "user_connections" table
CREATE TABLE `user_connections` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `permissions` text NULL DEFAULT (''), `metadata` json NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_connections_users_connections` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "userconnection_id" to table: "user_connections"
CREATE INDEX `userconnection_id` ON `user_connections` (`id`);
-- Create index "userconnection_user_id" to table: "user_connections"
CREATE INDEX `userconnection_user_id` ON `user_connections` (`user_id`);
-- Create index "userconnection_type_user_id" to table: "user_connections"
CREATE INDEX `userconnection_type_user_id` ON `user_connections` (`type`, `user_id`);
-- Create index "userconnection_reference_id" to table: "user_connections"
CREATE INDEX `userconnection_reference_id` ON `user_connections` (`reference_id`);
-- Create index "userconnection_type_reference_id" to table: "user_connections"
CREATE INDEX `userconnection_type_reference_id` ON `user_connections` (`type`, `reference_id`);
-- Create index "userconnection_reference_id_user_id_type" to table: "user_connections"
CREATE UNIQUE INDEX `userconnection_reference_id_user_id_type` ON `user_connections` (`reference_id`, `user_id`, `type`);
-- Create "user_contents" table
CREATE TABLE `user_contents` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `value` text NOT NULL, `metadata` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_contents_users_content` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "usercontent_id" to table: "user_contents"
CREATE INDEX `usercontent_id` ON `user_contents` (`id`);
-- Create index "usercontent_user_id" to table: "user_contents"
CREATE INDEX `usercontent_user_id` ON `user_contents` (`user_id`);
-- Create index "usercontent_type_user_id" to table: "user_contents"
CREATE INDEX `usercontent_type_user_id` ON `user_contents` (`type`, `user_id`);
-- Create index "usercontent_reference_id" to table: "user_contents"
CREATE INDEX `usercontent_reference_id` ON `user_contents` (`reference_id`);
-- Create index "usercontent_type_reference_id" to table: "user_contents"
CREATE INDEX `usercontent_type_reference_id` ON `user_contents` (`type`, `reference_id`);
-- Create "user_subscriptions" table
CREATE TABLE `user_subscriptions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `expires_at` datetime NOT NULL, `permissions` text NOT NULL, `reference_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_subscriptions_users_subscriptions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "usersubscription_id" to table: "user_subscriptions"
CREATE INDEX `usersubscription_id` ON `user_subscriptions` (`id`);
-- Create index "usersubscription_user_id" to table: "user_subscriptions"
CREATE INDEX `usersubscription_user_id` ON `user_subscriptions` (`user_id`);
-- Create index "usersubscription_type_user_id" to table: "user_subscriptions"
CREATE INDEX `usersubscription_type_user_id` ON `user_subscriptions` (`type`, `user_id`);
-- Create index "usersubscription_expires_at" to table: "user_subscriptions"
CREATE INDEX `usersubscription_expires_at` ON `user_subscriptions` (`expires_at`);
-- Create index "usersubscription_expires_at_user_id" to table: "user_subscriptions"
CREATE INDEX `usersubscription_expires_at_user_id` ON `user_subscriptions` (`expires_at`, `user_id`);
-- Create "vehicles" table
CREATE TABLE `vehicles` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `tier` integer NOT NULL, `localized_names` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "vehicle_id" to table: "vehicles"
CREATE INDEX `vehicle_id` ON `vehicles` (`id`);
-- Create "vehicle_averages" table
CREATE TABLE `vehicle_averages` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `data` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "vehicleaverage_id" to table: "vehicle_averages"
CREATE INDEX `vehicleaverage_id` ON `vehicle_averages` (`id`);
-- Create "vehicle_snapshots" table
CREATE TABLE `vehicle_snapshots` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` datetime NOT NULL, `frame` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_snapshots_accounts_vehicle_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE);
-- Create index "vehiclesnapshot_id" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_id` ON `vehicle_snapshots` (`id`);
-- Create index "vehiclesnapshot_created_at" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_created_at` ON `vehicle_snapshots` (`created_at`);
-- Create index "vehiclesnapshot_vehicle_id_created_at" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_vehicle_id_created_at` ON `vehicle_snapshots` (`vehicle_id`, `created_at`);
-- Create index "vehiclesnapshot_account_id_created_at" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_account_id_created_at` ON `vehicle_snapshots` (`account_id`, `created_at`);
-- Create index "vehiclesnapshot_account_id_type_created_at" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_account_id_type_created_at` ON `vehicle_snapshots` (`account_id`, `type`, `created_at`);
