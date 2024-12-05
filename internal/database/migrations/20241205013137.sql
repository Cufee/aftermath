-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_account_snapshot" table
CREATE TABLE `new_account_snapshot` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `last_battle_time` datetime NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` blob NOT NULL DEFAULT '', `regular_battles` integer NOT NULL, `regular_frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_account_id_snapshot_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "account_snapshot" to new temporary table "new_account_snapshot"
INSERT INTO `new_account_snapshot` (`id`, `created_at`, `updated_at`, `type`, `last_battle_time`, `reference_id`, `rating_battles`, `rating_frame`, `regular_battles`, `regular_frame`, `account_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `last_battle_time`, `reference_id`, `rating_battles`, IFNULL(`rating_frame`, '') AS `rating_frame`, `regular_battles`, IFNULL(`regular_frame`, '') AS `regular_frame`, `account_id` FROM `account_snapshot`;
-- Drop "account_snapshot" table after copying rows
DROP TABLE `account_snapshot`;
-- Rename temporary table "new_account_snapshot" to "account_snapshot"
ALTER TABLE `new_account_snapshot` RENAME TO `account_snapshot`;
-- Create index "account_snapshot_id_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_id_idx` ON `account_snapshot` (`id`);
-- Create index "account_snapshot_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_created_at_idx` ON `account_snapshot` (`created_at`);
-- Create index "account_snapshot_type_account_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "account_snapshot_type_account_id_reference_id_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_reference_id_idx` ON `account_snapshot` (`type`, `account_id`, `reference_id`);
-- Create index "account_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_reference_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `reference_id`, `created_at`);
-- Create "new_app_configuration" table
CREATE TABLE `new_app_configuration` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `key` text NOT NULL, `value` blob NOT NULL DEFAULT '', `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "app_configuration" to new temporary table "new_app_configuration"
INSERT INTO `new_app_configuration` (`id`, `created_at`, `updated_at`, `key`, `value`, `metadata`) SELECT `id`, `created_at`, `updated_at`, `key`, IFNULL(`value`, '') AS `value`, IFNULL(`metadata`, '') AS `metadata` FROM `app_configuration`;
-- Drop "app_configuration" table after copying rows
DROP TABLE `app_configuration`;
-- Rename temporary table "new_app_configuration" to "app_configuration"
ALTER TABLE `new_app_configuration` RENAME TO `app_configuration`;
-- Create index "app_configuration_id_idx" to table: "app_configuration"
CREATE INDEX `app_configuration_id_idx` ON `app_configuration` (`id`);
-- Create index "app_configuration_key_idx" to table: "app_configuration"
CREATE UNIQUE INDEX `app_configuration_key_idx` ON `app_configuration` (`key`);
-- Create "new_auth_nonce" table
CREATE TABLE `new_auth_nonce` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `active` bool NOT NULL, `expires_at` datetime NOT NULL, `identifier` text NOT NULL, `public_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "auth_nonce" to new temporary table "new_auth_nonce"
INSERT INTO `new_auth_nonce` (`id`, `created_at`, `updated_at`, `active`, `expires_at`, `identifier`, `public_id`, `metadata`) SELECT `id`, `created_at`, `updated_at`, `active`, `expires_at`, `identifier`, `public_id`, IFNULL(`metadata`, '') AS `metadata` FROM `auth_nonce`;
-- Drop "auth_nonce" table after copying rows
DROP TABLE `auth_nonce`;
-- Rename temporary table "new_auth_nonce" to "auth_nonce"
ALTER TABLE `new_auth_nonce` RENAME TO `auth_nonce`;
-- Create index "auth_nonce_public_id_idx" to table: "auth_nonce"
CREATE UNIQUE INDEX `auth_nonce_public_id_idx` ON `auth_nonce` (`public_id`);
-- Create index "auth_nonce_public_id_active_expires_at_idx" to table: "auth_nonce"
CREATE INDEX `auth_nonce_public_id_active_expires_at_idx` ON `auth_nonce` (`public_id`, `active`, `expires_at`);
-- Create "new_clan" table
CREATE TABLE `new_clan` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `tag` text NOT NULL, `name` text NOT NULL, `emblem_id` text NULL DEFAULT '', `members` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "clan" to new temporary table "new_clan"
INSERT INTO `new_clan` (`id`, `created_at`, `updated_at`, `tag`, `name`, `emblem_id`, `members`) SELECT `id`, `created_at`, `updated_at`, `tag`, `name`, `emblem_id`, IFNULL(`members`, '') AS `members` FROM `clan`;
-- Drop "clan" table after copying rows
DROP TABLE `clan`;
-- Rename temporary table "new_clan" to "clan"
ALTER TABLE `new_clan` RENAME TO `clan`;
-- Create index "clan_id_idx" to table: "clan"
CREATE INDEX `clan_id_idx` ON `clan` (`id`);
-- Create index "clan_tag_idx" to table: "clan"
CREATE INDEX `clan_tag_idx` ON `clan` (`tag`);
-- Create index "clan_name_idx" to table: "clan"
CREATE INDEX `clan_name_idx` ON `clan` (`name`);
-- Create "new_cron_task" table
CREATE TABLE `new_cron_task` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `targets` blob NOT NULL DEFAULT '', `status` text NOT NULL, `scheduled_after` datetime NOT NULL, `last_run` datetime NOT NULL, `tries_left` integer NOT NULL DEFAULT 0, `logs` blob NOT NULL DEFAULT '', `data` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "cron_task" to new temporary table "new_cron_task"
INSERT INTO `new_cron_task` (`id`, `created_at`, `updated_at`, `type`, `reference_id`, `targets`, `status`, `scheduled_after`, `last_run`, `tries_left`, `logs`, `data`) SELECT `id`, `created_at`, `updated_at`, `type`, `reference_id`, IFNULL(`targets`, '') AS `targets`, `status`, `scheduled_after`, `last_run`, `tries_left`, IFNULL(`logs`, '') AS `logs`, IFNULL(`data`, '') AS `data` FROM `cron_task`;
-- Drop "cron_task" table after copying rows
DROP TABLE `cron_task`;
-- Rename temporary table "new_cron_task" to "cron_task"
ALTER TABLE `new_cron_task` RENAME TO `cron_task`;
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
-- Create "new_session" table
CREATE TABLE `new_session` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `expires_at` datetime NOT NULL, `public_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `session_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "session" to new temporary table "new_session"
INSERT INTO `new_session` (`id`, `created_at`, `updated_at`, `expires_at`, `public_id`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `expires_at`, `public_id`, IFNULL(`metadata`, '') AS `metadata`, `user_id` FROM `session`;
-- Drop "session" table after copying rows
DROP TABLE `session`;
-- Rename temporary table "new_session" to "session"
ALTER TABLE `new_session` RENAME TO `session`;
-- Create index "session_public_id_idx" to table: "session"
CREATE UNIQUE INDEX `session_public_id_idx` ON `session` (`public_id`);
-- Create index "session_public_id_expires_at_idx" to table: "session"
CREATE INDEX `session_public_id_expires_at_idx` ON `session` (`public_id`, `expires_at`);
-- Create "new_user" table
CREATE TABLE `new_user` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `username` text NOT NULL DEFAULT '', `permissions` text NOT NULL DEFAULT '', `feature_flags` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "user" to new temporary table "new_user"
INSERT INTO `new_user` (`id`, `created_at`, `updated_at`, `username`, `permissions`, `feature_flags`) SELECT `id`, `created_at`, `updated_at`, `username`, `permissions`, IFNULL(`feature_flags`, '') AS `feature_flags` FROM `user`;
-- Drop "user" table after copying rows
DROP TABLE `user`;
-- Rename temporary table "new_user" to "user"
ALTER TABLE `new_user` RENAME TO `user`;
-- Create index "user_id_idx" to table: "user"
CREATE INDEX `user_id_idx` ON `user` (`id`);
-- Create index "user_username_idx" to table: "user"
CREATE INDEX `user_username_idx` ON `user` (`username`);
-- Create "new_vehicle" table
CREATE TABLE `new_vehicle` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `tier` integer NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "vehicle" to new temporary table "new_vehicle"
INSERT INTO `new_vehicle` (`id`, `created_at`, `updated_at`, `tier`, `localized_names`) SELECT `id`, `created_at`, `updated_at`, `tier`, IFNULL(`localized_names`, '') AS `localized_names` FROM `vehicle`;
-- Drop "vehicle" table after copying rows
DROP TABLE `vehicle`;
-- Rename temporary table "new_vehicle" to "vehicle"
ALTER TABLE `new_vehicle` RENAME TO `vehicle`;
-- Create index "vehicle_id_idx" to table: "vehicle"
CREATE INDEX `vehicle_id_idx` ON `vehicle` (`id`);
-- Create "new_vehicle_average" table
CREATE TABLE `new_vehicle_average` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `data` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "vehicle_average" to new temporary table "new_vehicle_average"
INSERT INTO `new_vehicle_average` (`id`, `created_at`, `updated_at`, `data`) SELECT `id`, `created_at`, `updated_at`, IFNULL(`data`, '') AS `data` FROM `vehicle_average`;
-- Drop "vehicle_average" table after copying rows
DROP TABLE `vehicle_average`;
-- Rename temporary table "new_vehicle_average" to "vehicle_average"
ALTER TABLE `new_vehicle_average` RENAME TO `vehicle_average`;
-- Create index "vehicle_average_id_idx" to table: "vehicle_average"
CREATE INDEX `vehicle_average_id_idx` ON `vehicle_average` (`id`);
-- Create "new_vehicle_snapshot" table
CREATE TABLE `new_vehicle_snapshot` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` datetime NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_snapshot_account_id_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "vehicle_snapshot" to new temporary table "new_vehicle_snapshot"
INSERT INTO `new_vehicle_snapshot` (`id`, `created_at`, `updated_at`, `type`, `vehicle_id`, `reference_id`, `battles`, `last_battle_time`, `frame`, `account_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `vehicle_id`, `reference_id`, `battles`, `last_battle_time`, IFNULL(`frame`, '') AS `frame`, `account_id` FROM `vehicle_snapshot`;
-- Drop "vehicle_snapshot" table after copying rows
DROP TABLE `vehicle_snapshot`;
-- Rename temporary table "new_vehicle_snapshot" to "vehicle_snapshot"
ALTER TABLE `new_vehicle_snapshot` RENAME TO `vehicle_snapshot`;
-- Create index "vehicle_snapshot_id_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_id_idx` ON `vehicle_snapshot` (`id`);
-- Create index "vehicle_snapshot_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_created_at_idx` ON `vehicle_snapshot` (`created_at`);
-- Create index "vehicle_snapshot_vehicle_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_vehicle_id_created_at_idx` ON `vehicle_snapshot` (`vehicle_id`, `created_at`);
-- Create index "vehicle)snapshot_account_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle)snapshot_account_id_created_at_idx` ON `vehicle_snapshot` (`account_id`, `created_at`);
-- Create index "vehicle_snapshot_account_id_type_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_account_id_type_created_at_idx` ON `vehicle_snapshot` (`account_id`, `type`, `created_at`);
-- Create "new_game_map" table
CREATE TABLE `new_game_map` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `game_modes` blob NOT NULL DEFAULT '', `supremacy_points` integer NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "game_map" to new temporary table "new_game_map"
INSERT INTO `new_game_map` (`id`, `created_at`, `updated_at`, `game_modes`, `supremacy_points`, `localized_names`) SELECT `id`, `created_at`, `updated_at`, IFNULL(`game_modes`, '') AS `game_modes`, `supremacy_points`, IFNULL(`localized_names`, '') AS `localized_names` FROM `game_map`;
-- Drop "game_map" table after copying rows
DROP TABLE `game_map`;
-- Rename temporary table "new_game_map" to "game_map"
ALTER TABLE `new_game_map` RENAME TO `game_map`;
-- Create index "game_map_id_idx" to table: "game_map"
CREATE INDEX `game_map_id_idx` ON `game_map` (`id`);
-- Create "new_widget_settings" table
CREATE TABLE `new_widget_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `reference_id` text NOT NULL, `title` text NULL, `session_from` datetime NULL, `metadata` blob NOT NULL DEFAULT '', `styles` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, `session_reference_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `widget_setting_user_id_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "widget_settings" to new temporary table "new_widget_settings"
INSERT INTO `new_widget_settings` (`id`, `created_at`, `updated_at`, `reference_id`, `title`, `session_from`, `metadata`, `styles`, `user_id`, `session_reference_id`) SELECT `id`, `created_at`, `updated_at`, `reference_id`, `title`, `session_from`, IFNULL(`metadata`, '') AS `metadata`, IFNULL(`styles`, '') AS `styles`, `user_id`, `session_reference_id` FROM `widget_settings`;
-- Drop "widget_settings" table after copying rows
DROP TABLE `widget_settings`;
-- Rename temporary table "new_widget_settings" to "widget_settings"
ALTER TABLE `new_widget_settings` RENAME TO `widget_settings`;
-- Create index "widget_settings_id" to table: "widget_settings"
CREATE INDEX `widget_settings_id` ON `widget_settings` (`id`);
-- Create "new_game_mode" table
CREATE TABLE `new_game_mode` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Copy rows from old table "game_mode" to new temporary table "new_game_mode"
INSERT INTO `new_game_mode` (`id`, `created_at`, `updated_at`, `localized_names`) SELECT `id`, `created_at`, `updated_at`, IFNULL(`localized_names`, '') AS `localized_names` FROM `game_mode`;
-- Drop "game_mode" table after copying rows
DROP TABLE `game_mode`;
-- Rename temporary table "new_game_mode" to "game_mode"
ALTER TABLE `new_game_mode` RENAME TO `game_mode`;
-- Create index "game_mode_id_idx" to table: "game_mode"
CREATE INDEX `game_mode_id_idx` ON `game_mode` (`id`);
-- Create "new_moderation_request" table
CREATE TABLE `new_moderation_request` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `moderator_comment` text NULL, `context` text NULL, `reference_id` text NOT NULL, `action_reason` text NULL, `action_status` text NOT NULL, `data` blob NOT NULL DEFAULT '', `requestor_id` text NOT NULL, `moderator_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `moderation_request_moderator_id_user_id_fk` FOREIGN KEY (`moderator_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `moderation_request_requestor_id_user_id_fk` FOREIGN KEY (`requestor_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "moderation_request" to new temporary table "new_moderation_request"
INSERT INTO `new_moderation_request` (`id`, `created_at`, `updated_at`, `moderator_comment`, `context`, `reference_id`, `action_reason`, `action_status`, `data`, `requestor_id`, `moderator_id`) SELECT `id`, `created_at`, `updated_at`, `moderator_comment`, `context`, `reference_id`, `action_reason`, `action_status`, IFNULL(`data`, '') AS `data`, `requestor_id`, `moderator_id` FROM `moderation_request`;
-- Drop "moderation_request" table after copying rows
DROP TABLE `moderation_request`;
-- Rename temporary table "new_moderation_request" to "moderation_request"
ALTER TABLE `new_moderation_request` RENAME TO `moderation_request`;
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
-- Create "new_user_restriction" table
CREATE TABLE `new_user_restriction` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `expires_at` datetime NOT NULL, `type` text NOT NULL, `restriction` text NOT NULL, `public_reason` text NOT NULL, `moderator_comment` text NOT NULL, `events` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_restriction_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Copy rows from old table "user_restriction" to new temporary table "new_user_restriction"
INSERT INTO `new_user_restriction` (`id`, `created_at`, `updated_at`, `expires_at`, `type`, `restriction`, `public_reason`, `moderator_comment`, `events`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `expires_at`, `type`, `restriction`, `public_reason`, `moderator_comment`, IFNULL(`events`, '') AS `events`, `user_id` FROM `user_restriction`;
-- Drop "user_restriction" table after copying rows
DROP TABLE `user_restriction`;
-- Rename temporary table "new_user_restriction" to "user_restriction"
ALTER TABLE `new_user_restriction` RENAME TO `user_restriction`;
-- Create index "user_restriction_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_id_idx` ON `user_restriction` (`id`);
-- Create index "user_restriction_user_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_user_id_idx` ON `user_restriction` (`user_id`);
-- Create index "user_restriction_expires_at_user_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_expires_at_user_id_idx` ON `user_restriction` (`expires_at`, `user_id`);
-- Create "new_user_content" table
CREATE TABLE `new_user_content` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `value` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_content_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "user_content" to new temporary table "new_user_content"
INSERT INTO `new_user_content` (`id`, `created_at`, `updated_at`, `type`, `reference_id`, `value`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `reference_id`, `value`, IFNULL(`metadata`, '') AS `metadata`, `user_id` FROM `user_content`;
-- Drop "user_content" table after copying rows
DROP TABLE `user_content`;
-- Rename temporary table "new_user_content" to "user_content"
ALTER TABLE `new_user_content` RENAME TO `user_content`;
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
-- Create "new_discord_interaction" table
CREATE TABLE `new_discord_interaction` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `result` text NOT NULL, `event_id` text NOT NULL, `guild_id` text NOT NULL, `snowflake` text NOT NULL DEFAULT '', `channel_id` text NOT NULL, `message_id` text NOT NULL, `type` text NOT NULL, `locale` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `discord_interaction_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "discord_interaction" to new temporary table "new_discord_interaction"
INSERT INTO `new_discord_interaction` (`id`, `created_at`, `updated_at`, `result`, `event_id`, `guild_id`, `snowflake`, `channel_id`, `message_id`, `type`, `locale`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `result`, `event_id`, `guild_id`, `snowflake`, `channel_id`, `message_id`, `type`, `locale`, IFNULL(`metadata`, '') AS `metadata`, `user_id` FROM `discord_interaction`;
-- Drop "discord_interaction" table after copying rows
DROP TABLE `discord_interaction`;
-- Rename temporary table "new_discord_interaction" to "discord_interaction"
ALTER TABLE `new_discord_interaction` RENAME TO `discord_interaction`;
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
-- Create "new_user_connection" table
CREATE TABLE `new_user_connection` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `verified` bool NOT NULL DEFAULT false, `selected` bool NOT NULL DEFAULT false, `reference_id` text NOT NULL, `permissions` text NULL DEFAULT '', `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_connection_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "user_connection" to new temporary table "new_user_connection"
INSERT INTO `new_user_connection` (`id`, `created_at`, `updated_at`, `type`, `verified`, `selected`, `reference_id`, `permissions`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `verified`, `selected`, `reference_id`, `permissions`, IFNULL(`metadata`, '') AS `metadata`, `user_id` FROM `user_connection`;
-- Drop "user_connection" table after copying rows
DROP TABLE `user_connection`;
-- Rename temporary table "new_user_connection" to "user_connection"
ALTER TABLE `new_user_connection` RENAME TO `user_connection`;
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
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
