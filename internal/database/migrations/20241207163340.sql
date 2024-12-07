-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "accounts" table
ALTER TABLE `accounts` RENAME TO `account`;
-- Create index "account_id_idx" to table: "account"
CREATE INDEX `account_id_idx` ON `account` (`id`);
-- Create index "account_id_last_battle_time_idx" to table: "account"
CREATE INDEX `account_id_last_battle_time_idx` ON `account` (`id`, `last_battle_time`);
-- Create index "account_realm_idx" to table: "account"
CREATE INDEX `account_realm_idx` ON `account` (`realm`);
-- Create index "account_realm_last_battle_time_idx" to table: "account"
CREATE INDEX `account_realm_last_battle_time_idx` ON `account` (`realm`, `last_battle_time`);
-- Create index "account_clan_id_idx" to table: "account"
CREATE INDEX `account_clan_id_idx` ON `account` (`clan_id`);

-- Drop "account_snapshots" table
ALTER TABLE `account_snapshots` RENAME TO `account_snapshot`;
-- Create index "account_snapshot_id_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_id_idx` ON `account_snapshot` (`id`);
-- Create index "account_snapshot_type_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_idx` ON `account_snapshot` (`type`);
-- Create index "account_snapshot_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_created_at_idx` ON `account_snapshot` (`created_at`);
-- Create index "account_snapshot_type_account_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "account_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_reference_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `reference_id`, `created_at`);
-- Create index "account_snapshot_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_account_id_reference_id_created_at_idx` ON `account_snapshot` (`account_id`, `reference_id`, `created_at`);

-- Drop "app_configurations" table
ALTER TABLE `app_configurations` RENAME TO `app_configuration`;
-- Create index "app_configuration_id_idx" to table: "app_configuration"
CREATE INDEX `app_configuration_id_idx` ON `app_configuration` (`id`);
-- Create index "app_configuration_key_idx" to table: "app_configuration"
CREATE UNIQUE INDEX `app_configuration_key_idx` ON `app_configuration` (`key`);

-- Drop "application_commands" table
ALTER TABLE `application_commands` RENAME TO `application_command`;
-- Create index "application_command_id_idx" to table: "application_command"
CREATE INDEX `application_command_id_idx` ON `application_command` (`id`);
-- Create index "application_command_options_hash_idx" to table: "application_command"
CREATE INDEX `application_command_options_hash_idx` ON `application_command` (`options_hash`);

-- Drop "auth_nonces" table
ALTER TABLE `auth_nonces` RENAME TO `auth_nonce`;
-- Create index "auth_nonce_public_id_idx" to table: "auth_nonce"
CREATE UNIQUE INDEX `auth_nonce_public_id_idx` ON `auth_nonce` (`public_id`);
-- Create index "auth_nonce_public_id_active_expires_at_idx" to table: "auth_nonce"
CREATE INDEX `auth_nonce_public_id_active_expires_at_idx` ON `auth_nonce` (`public_id`, `active`, `expires_at`);

-- Drop "clans" table
ALTER TABLE `clans` RENAME TO `clan`;
-- Create index "clan_id_idx" to table: "clan"
CREATE INDEX `clan_id_idx` ON `clan` (`id`);
-- Create index "clan_tag_idx" to table: "clan"
CREATE INDEX `clan_tag_idx` ON `clan` (`tag`);
-- Create index "clan_name_idx" to table: "clan"
CREATE INDEX `clan_name_idx` ON `clan` (`name`);

-- Drop "cron_tasks" table
ALTER TABLE `cron_tasks` RENAME TO `cron_task`;
-- Create index "cron_task_id_idx" to table: "cron_task"
CREATE INDEX `cron_task_id_idx` ON `cron_task` (`id`);
-- Create index "cron_task_reference_id_idx" to table: "cron_task"
CREATE INDEX `cron_task_reference_id_idx` ON `cron_task` (`reference_id`);
-- Create index "cron_task_status_last_run_idx" to table: "cron_task"
CREATE INDEX `cron_task_status_last_run_idx` ON `cron_task` (`status`, `last_run`);
-- Create index "cron_task_status_created_at_idx" to table: "cron_task"
CREATE INDEX `cron_task_status_created_at_idx` ON `cron_task` (`status`, `created_at`);
-- Create index "cron_task_status_scheduled_after_idx" to table: "cron_task"
CREATE INDEX `cron_task_status_scheduled_after_idx` ON `cron_task` (`status`, `scheduled_after`);

-- Drop "sessions" table
ALTER TABLE `sessions` RENAME TO `session`;
-- Create index "session_public_id_idx" to table: "session"
CREATE UNIQUE INDEX `session_public_id_idx` ON `session` (`public_id`);
-- Create index "session_public_id_expires_at_idx" to table: "session"
CREATE INDEX `session_public_id_expires_at_idx` ON `session` (`public_id`, `expires_at`);

-- Drop "users" table
ALTER TABLE `users` RENAME TO `user`;
-- Create index "user_id_idx" to table: "user"
CREATE INDEX `user_id_idx` ON `user` (`id`);
-- Create index "user_username_idx" to table: "user"
CREATE INDEX `user_username_idx` ON `user` (`username`);

-- Drop "user_connections" table
ALTER TABLE `user_connections` RENAME TO `user_connection`;
-- Create index "user_connection_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_id_idx` ON `user_connection` (`id`);
-- Create index "user_connection_user_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_user_id_idx` ON `user_connection` (`user_id`);
-- Create index "user_connection_type_user_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_type_user_id_idx` ON `user_connection` (`type`, `user_id`);
-- Create index "user_connection_reference_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_reference_id_idx` ON `user_connection` (`reference_id`);
-- Create index "user_connection_type_reference_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_type_reference_id_idx` ON `user_connection` (`type`, `reference_id`);
-- Create index "user_connection_reference_id_user_id_type_idx" to table: "user_connection"
CREATE UNIQUE INDEX `user_connection_reference_id_user_id_type_idx` ON `user_connection` (`reference_id`, `user_id`, `type`);

-- Drop "user_subscriptions" table
ALTER TABLE `user_subscriptions` RENAME TO `user_subscription`;
-- Create index "user_subscription_id_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_id_idx` ON `user_subscription` (`id`);
-- Create index "user_subscription_user_id_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_user_id_idx` ON `user_subscription` (`user_id`);
-- Create index "user_subscription_type_user_id_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_type_user_id_idx` ON `user_subscription` (`type`, `user_id`);
-- Create index "user_subscription_expires_at_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_expires_at_idx` ON `user_subscription` (`expires_at`);
-- Create index "user_subscription_expires_at_user_id_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_expires_at_user_id_idx` ON `user_subscription` (`expires_at`, `user_id`);

-- Drop "vehicles" table
ALTER TABLE `vehicles` RENAME TO `vehicle`;
-- Create index "vehicle_id_idx" to table: "vehicle"
CREATE INDEX `vehicle_id_idx` ON `vehicle` (`id`);

-- Drop "vehicle_averages" table
ALTER TABLE `vehicle_averages` RENAME TO `vehicle_average`;
-- Create index "vehicle_average_id_idx" to table: "vehicle_average"
CREATE INDEX `vehicle_average_id_idx` ON `vehicle_average` (`id`);

-- Drop "vehicle_snapshots" table
ALTER TABLE `vehicle_snapshots` RENAME TO `vehicle_snapshot`;
-- Create index "vehicle_snapshot_id_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_id_idx` ON `vehicle_snapshot` (`id`);
-- Create index "vehicle_snapshot_type_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_type_idx` ON `vehicle_snapshot` (`type`);
-- Create index "vehicle_snapshot_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_created_at_idx` ON `vehicle_snapshot` (`created_at`);
-- Create index "vehicle_snapshot_type_account_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_type_account_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "vehicle_snapshot_type_account_id_vehicle_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_type_account_id_vehicle_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `vehicle_id`, `created_at`);
-- Create index "vehicle_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `vehicle_id`, `reference_id`, `created_at`);
-- Create index "vehicle_snapshot_vehicle_id_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_vehicle_id_reference_id_created_at_idx` ON `vehicle_snapshot` (`vehicle_id`, `reference_id`, `created_at`);

-- Drop "leaderboard_scores" table
DROP TABLE `leaderboard_scores`;

-- Drop "game_maps" table
ALTER TABLE `game_maps` RENAME TO `game_map`;
-- Create index "game_map_id_idx" to table: "game_map"
CREATE INDEX `game_map_id_idx` ON `game_map` (`id`);

-- Create "new_widget_settings" table
CREATE TABLE `new_widget_settings` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `reference_id` text NOT NULL, `title` text NULL, `session_from` text NULL, `metadata` blob NOT NULL DEFAULT '', `styles` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, `session_reference_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `widget_setting_user_id_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "widget_settings" to new temporary table "new_widget_settings"
INSERT INTO `new_widget_settings` (`id`, `created_at`, `updated_at`, `reference_id`, `title`, `session_from`, `metadata`, `styles`, `user_id`, `session_reference_id`) SELECT `id`, `created_at`, `updated_at`, `reference_id`, `title`, `session_from`, IFNULL(`metadata`, '') AS `metadata`, IFNULL(`styles`, '') AS `styles`, `user_id`, `session_reference_id` FROM `widget_settings`;
-- Drop "widget_settings" table after copying rows
DROP TABLE `widget_settings`;
-- Rename temporary table "new_widget_settings" to "widget_settings"
ALTER TABLE `new_widget_settings` RENAME TO `widget_settings`;
-- Create index "widget_settings_id" to table: "widget_settings"
CREATE INDEX `widget_settings_id` ON `widget_settings` (`id`);

-- Drop "game_modes" table
ALTER TABLE `game_modes` RENAME TO `game_mode`;
-- Create index "game_mode_id_idx" to table: "game_mode"
CREATE INDEX `game_mode_id_idx` ON `game_mode` (`id`);

-- Drop "moderation_requests" table
ALTER TABLE `moderation_requests` RENAME TO `moderation_request`;
-- Create index "moderation_request_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_id_idx` ON `moderation_request` (`id`);
-- Create index "moderation_request_reference_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_reference_id_idx` ON `moderation_request` (`reference_id`);
-- Create index "moderation_request_requestor_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_requestor_id_idx` ON `moderation_request` (`requestor_id`);
-- Create index "moderation_request_moderator_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_moderator_id_idx` ON `moderation_request` (`moderator_id`);
-- Create index "moderation_request_requestor_id_reference_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_requestor_id_reference_id_idx` ON `moderation_request` (`requestor_id`, `reference_id`);
-- Create index "moderation_request_requestor_id_reference_id_action_status_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_requestor_id_reference_id_action_status_idx` ON `moderation_request` (`requestor_id`, `reference_id`, `action_status`);

-- Drop "user_restrictions" table
ALTER TABLE `user_restrictions` RENAME TO `user_restriction`;
-- Create index "user_restriction_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_id_idx` ON `user_restriction` (`id`);
-- Create index "user_restriction_user_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_user_id_idx` ON `user_restriction` (`user_id`);
-- Create index "user_restriction_expires_at_user_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_expires_at_user_id_idx` ON `user_restriction` (`expires_at`, `user_id`);

-- Drop "user_contents" table
ALTER TABLE `user_contents` RENAME TO `user_content`;
-- Create index "user_content_id_idx" to table: "user_content"
CREATE INDEX `user_content_id_idx` ON `user_content` (`id`);
-- Create index "user_content_user_id_idx" to table: "user_content"
CREATE INDEX `user_content_user_id_idx` ON `user_content` (`user_id`);
-- Create index "user_content_reference_id_idx" to table: "user_content"
CREATE INDEX `user_content_reference_id_idx` ON `user_content` (`reference_id`);
-- Create index "user_content_type_user_id_idx" to table: "user_content"
CREATE UNIQUE INDEX `user_content_type_user_id_idx` ON `user_content` (`type`, `user_id`);
-- Create index "user_content_type_reference_id_idx" to table: "user_content"
CREATE UNIQUE INDEX `user_content_type_reference_id_idx` ON `user_content` (`type`, `reference_id`);

-- Drop "discord_interactions" table
ALTER TABLE `discord_interactions` RENAME TO `discord_interaction`;
-- Create index "discord_interaction_id_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_id_idx` ON `discord_interaction` (`id`);
-- Create index "discord_interaction_user_id_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_user_id_idx` ON `discord_interaction` (`user_id`);
-- Create index "discord_interaction_snowflake_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_snowflake_idx` ON `discord_interaction` (`snowflake`);
-- Create index "discord_interaction_created_at_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_created_at_idx` ON `discord_interaction` (`created_at`);
-- Create index "discord_interaction_user_id_type_created_at_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_user_id_type_created_at_idx` ON `discord_interaction` (`user_id`, `type`, `created_at`);
-- Create index "discord_interaction_channel_id_type_created_at_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_channel_id_type_created_at_idx` ON `discord_interaction` (`channel_id`, `type`, `created_at`);

-- Create "account_achievements_snapshot" table
CREATE TABLE `account_achievements_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `reference_id` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_id_account_achievements_snapshot_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "account_achievements_snapshot_id_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_id_idx` ON `account_achievements_snapshot` (`id`);
-- Create index "account_achievements_snapshot_type_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_type_idx` ON `account_achievements_snapshot` (`type`);
-- Create index "account_achievements_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_created_at_idx` ON `account_achievements_snapshot` (`created_at`);
-- Create index "account_achievements_snapshot_type_account_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_type_account_id_created_at_idx` ON `account_achievements_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "account_achievements_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_type_account_id_reference_id_created_at_idx` ON `account_achievements_snapshot` (`type`, `account_id`, `reference_id`, `created_at`);
-- Create index "account_achievements_snapshot_account_id_reference_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_account_id_reference_id_created_at_idx` ON `account_achievements_snapshot` (`account_id`, `reference_id`, `created_at`);

-- Create "vehicle_achievements_snapshot" table
CREATE TABLE `vehicle_achievements_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_achievements_snapshot_account_id_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "vehicle_achievements_snapshot_id_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_id_idx` ON `vehicle_achievements_snapshot` (`id`);
-- Create index "vehicle_achievements_snapshot_type_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_type_idx` ON `vehicle_achievements_snapshot` (`type`);
-- Create index "vehicle_achievements_snapshot_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_created_at_idx` ON `vehicle_achievements_snapshot` (`created_at`);
-- Create index "vehicle_achievements_snapshot_type_account_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_type_account_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "vehicle_achievements_snapshot_type_account_id_vehicle_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_type_account_id_vehicle_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `vehicle_id`, `created_at`);
-- Create index "vehicle_achievements_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `vehicle_id`, `reference_id`, `created_at`);
-- Create index "vehicle_achievements_snapshot_vehicle_id_reference_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_vehicle_id_reference_id_created_at_idx` ON `vehicle_achievements_snapshot` (`vehicle_id`, `reference_id`, `created_at`);

-- Create "manual_migration" table
CREATE TABLE `manual_migration` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `key` text NOT NULL, `finished` bool NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "manual_migration_id_idx" to table: "manual_migration"
CREATE INDEX `manual_migration_id_idx` ON `manual_migration` (`id`);
-- Create index "manual_migration_key_idx" to table: "manual_migration"
CREATE UNIQUE INDEX `manual_migration_key_idx` ON `manual_migration` (`key`);

-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
