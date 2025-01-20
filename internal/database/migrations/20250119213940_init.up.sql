-- create "account" table
CREATE TABLE IF NOT EXISTS `account` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `last_battle_time` text NOT NULL, `account_created_at` text NOT NULL, `realm` text NOT NULL, `nickname` text NOT NULL, `private` bool NOT NULL DEFAULT false, `clan_id` text NULL, PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "account_id_idx" to table: "account"
CREATE INDEX IF NOT EXISTS `account_id_idx` ON `account` (`id`);
-- CREATE INDEX IF NOT EXISTS "account_id_last_battle_time_idx" to table: "account"
CREATE INDEX IF NOT EXISTS `account_id_last_battle_time_idx` ON `account` (`id`, `last_battle_time`);
-- CREATE INDEX IF NOT EXISTS "account_realm_idx" to table: "account"
CREATE INDEX IF NOT EXISTS `account_realm_idx` ON `account` (`realm`);
-- CREATE INDEX IF NOT EXISTS "account_realm_last_battle_time_idx" to table: "account"
CREATE INDEX IF NOT EXISTS `account_realm_last_battle_time_idx` ON `account` (`realm`, `last_battle_time`);
-- CREATE INDEX IF NOT EXISTS "account_clan_id_idx" to table: "account"
CREATE INDEX IF NOT EXISTS `account_clan_id_idx` ON `account` (`clan_id`);
-- create "clan" table
CREATE TABLE IF NOT EXISTS `clan` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `tag` text NOT NULL, `name` text NOT NULL, `emblem_id` text NULL DEFAULT '', `members` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "clan_id_idx" to table: "clan"
CREATE INDEX IF NOT EXISTS `clan_id_idx` ON `clan` (`id`);
-- CREATE INDEX IF NOT EXISTS "clan_tag_idx" to table: "clan"
CREATE INDEX IF NOT EXISTS `clan_tag_idx` ON `clan` (`tag`);
-- CREATE INDEX IF NOT EXISTS "clan_name_idx" to table: "clan"
CREATE INDEX IF NOT EXISTS `clan_name_idx` ON `clan` (`name`);
-- create "account_snapshot" table
CREATE TABLE IF NOT EXISTS `account_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `last_battle_time` text NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` blob NOT NULL DEFAULT '', `regular_battles` integer NOT NULL, `regular_frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_account_id_snapshot_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "account_snapshot_id_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS `account_snapshot_id_idx` ON `account_snapshot` (`id`);
-- CREATE INDEX IF NOT EXISTS "account_snapshot_type_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS `account_snapshot_type_idx` ON `account_snapshot` (`type`);
-- CREATE INDEX IF NOT EXISTS "account_snapshot_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS `account_snapshot_created_at_idx` ON `account_snapshot` (`created_at`);
-- CREATE INDEX IF NOT EXISTS "account_snapshot_type_account_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS `account_snapshot_type_account_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "account_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS `account_snapshot_type_account_id_reference_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `reference_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "account_snapshot_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS `account_snapshot_account_id_reference_id_created_at_idx` ON `account_snapshot` (`account_id`, `reference_id`, `created_at`);
-- create "account_achievements_snapshot" table
CREATE TABLE IF NOT EXISTS `account_achievements_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `reference_id` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_id_account_achievements_snapshot_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_id_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `account_achievements_snapshot_id_idx` ON `account_achievements_snapshot` (`id`);
-- CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_type_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `account_achievements_snapshot_type_idx` ON `account_achievements_snapshot` (`type`);
-- CREATE INDEX IF NOT EXISTS "account_achievements_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `account_achievements_created_at_idx` ON `account_achievements_snapshot` (`created_at`);
-- CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_type_account_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `account_achievements_snapshot_type_account_id_created_at_idx` ON `account_achievements_snapshot` (`type`, `account_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `account_achievements_snapshot_type_account_id_reference_id_created_at_idx` ON `account_achievements_snapshot` (`type`, `account_id`, `reference_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_account_id_reference_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `account_achievements_snapshot_account_id_reference_id_created_at_idx` ON `account_achievements_snapshot` (`account_id`, `reference_id`, `created_at`);
-- create "vehicle_snapshot" table
CREATE TABLE IF NOT EXISTS `vehicle_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_snapshot_account_id_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "vehicle_snapshot_id_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_snapshot_id_idx` ON `vehicle_snapshot` (`id`);
-- CREATE INDEX IF NOT EXISTS "vehicle_snapshot_type_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_snapshot_type_idx` ON `vehicle_snapshot` (`type`);
-- CREATE INDEX IF NOT EXISTS "vehicle_snapshot_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_snapshot_created_at_idx` ON `vehicle_snapshot` (`created_at`);
-- CREATE INDEX IF NOT EXISTS "vehicle_snapshot_type_account_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_snapshot_type_account_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "vehicle_snapshot_type_account_id_vehicle_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_snapshot_type_account_id_vehicle_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `vehicle_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "vehicle_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `vehicle_id`, `reference_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "vehicle_snapshot_vehicle_id_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_snapshot_vehicle_id_reference_id_created_at_idx` ON `vehicle_snapshot` (`vehicle_id`, `reference_id`, `created_at`);
-- create "vehicle_achievements_snapshot" table
CREATE TABLE IF NOT EXISTS `vehicle_achievements_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_achievements_snapshot_account_id_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_id_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_achievements_snapshot_id_idx` ON `vehicle_achievements_snapshot` (`id`);
-- CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_type_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_achievements_snapshot_type_idx` ON `vehicle_achievements_snapshot` (`type`);
-- CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_achievements_snapshot_created_at_idx` ON `vehicle_achievements_snapshot` (`created_at`);
-- CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_type_account_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_achievements_snapshot_type_account_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_type_account_id_vehicle_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_achievements_snapshot_type_account_id_vehicle_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `vehicle_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_achievements_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `vehicle_id`, `reference_id`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_vehicle_id_reference_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS `vehicle_achievements_snapshot_vehicle_id_reference_id_created_at_idx` ON `vehicle_achievements_snapshot` (`vehicle_id`, `reference_id`, `created_at`);
-- create "app_configuration" table
CREATE TABLE IF NOT EXISTS `app_configuration` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `key` text NOT NULL, `value` blob NOT NULL DEFAULT '', `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "app_configuration_id_idx" to table: "app_configuration"
CREATE INDEX IF NOT EXISTS `app_configuration_id_idx` ON `app_configuration` (`id`);
-- CREATE INDEX IF NOT EXISTS "app_configuration_key_idx" to table: "app_configuration"
CREATE UNIQUE INDEX IF NOT EXISTS `app_configuration_key_idx` ON `app_configuration` (`key`);
-- create "cron_task" table
CREATE TABLE IF NOT EXISTS `cron_task` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `targets` blob NOT NULL DEFAULT '', `status` text NOT NULL, `scheduled_after` text NOT NULL, `last_run` text NOT NULL, `tries_left` integer NOT NULL DEFAULT 0, `logs` blob NOT NULL DEFAULT '', `data` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "cron_task_id_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS `cron_task_id_idx` ON `cron_task` (`id`);
-- CREATE INDEX IF NOT EXISTS "cron_task_reference_id_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS `cron_task_reference_id_idx` ON `cron_task` (`reference_id`);
-- CREATE INDEX IF NOT EXISTS "cron_task_status_last_run_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS `cron_task_status_last_run_idx` ON `cron_task` (`status`, `last_run`);
-- CREATE INDEX IF NOT EXISTS "cron_task_status_created_at_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS `cron_task_status_created_at_idx` ON `cron_task` (`status`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "cron_task_status_scheduled_after_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS `cron_task_status_scheduled_after_idx` ON `cron_task` (`status`, `scheduled_after`);
-- create "manual_migration" table
CREATE TABLE IF NOT EXISTS `manual_migration` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `key` text NOT NULL, `finished` bool NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "manual_migration_id_idx" to table: "manual_migration"
CREATE INDEX IF NOT EXISTS `manual_migration_id_idx` ON `manual_migration` (`id`);
-- CREATE INDEX IF NOT EXISTS "manual_migration_key_idx" to table: "manual_migration"
CREATE UNIQUE INDEX IF NOT EXISTS `manual_migration_key_idx` ON `manual_migration` (`key`);
-- create "auth_nonce" table
CREATE TABLE IF NOT EXISTS `auth_nonce` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `active` bool NOT NULL, `expires_at` text NOT NULL, `identifier` text NOT NULL, `public_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "auth_nonce_public_id_idx" to table: "auth_nonce"
CREATE UNIQUE INDEX IF NOT EXISTS `auth_nonce_public_id_idx` ON `auth_nonce` (`public_id`);
-- CREATE INDEX IF NOT EXISTS "auth_nonce_public_id_active_expires_at_idx" to table: "auth_nonce"
CREATE INDEX IF NOT EXISTS `auth_nonce_public_id_active_expires_at_idx` ON `auth_nonce` (`public_id`, `active`, `expires_at`);
-- create "session" table
CREATE TABLE IF NOT EXISTS `session` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `expires_at` text NOT NULL, `public_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `session_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "session_public_id_idx" to table: "session"
CREATE UNIQUE INDEX IF NOT EXISTS `session_public_id_idx` ON `session` (`public_id`);
-- CREATE INDEX IF NOT EXISTS "session_public_id_expires_at_idx" to table: "session"
CREATE INDEX IF NOT EXISTS `session_public_id_expires_at_idx` ON `session` (`public_id`, `expires_at`);
-- create "application_command" table
CREATE TABLE IF NOT EXISTS `application_command` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `name` text NOT NULL, `version` text NOT NULL, `options_hash` text NOT NULL, PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "application_command_id_idx" to table: "application_command"
CREATE INDEX IF NOT EXISTS `application_command_id_idx` ON `application_command` (`id`);
-- CREATE INDEX IF NOT EXISTS "application_command_options_hash_idx" to table: "application_command"
CREATE INDEX IF NOT EXISTS `application_command_options_hash_idx` ON `application_command` (`options_hash`);
-- create "discord_interaction" table
CREATE TABLE IF NOT EXISTS `discord_interaction` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `result` text NOT NULL, `event_id` text NOT NULL, `guild_id` text NOT NULL, `snowflake` text NOT NULL DEFAULT '', `channel_id` text NOT NULL, `message_id` text NOT NULL, `type` text NOT NULL, `locale` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `discord_interaction_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "discord_interaction_id_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS `discord_interaction_id_idx` ON `discord_interaction` (`id`);
-- CREATE INDEX IF NOT EXISTS "discord_interaction_user_id_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS `discord_interaction_user_id_idx` ON `discord_interaction` (`user_id`);
-- CREATE INDEX IF NOT EXISTS "discord_interaction_snowflake_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS `discord_interaction_snowflake_idx` ON `discord_interaction` (`snowflake`);
-- CREATE INDEX IF NOT EXISTS "discord_interaction_created_at_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS `discord_interaction_created_at_idx` ON `discord_interaction` (`created_at`);
-- CREATE INDEX IF NOT EXISTS "discord_interaction_user_id_type_created_at_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS `discord_interaction_user_id_type_created_at_idx` ON `discord_interaction` (`user_id`, `type`, `created_at`);
-- CREATE INDEX IF NOT EXISTS "discord_interaction_channel_id_type_created_at_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS `discord_interaction_channel_id_type_created_at_idx` ON `discord_interaction` (`channel_id`, `type`, `created_at`);
-- create "vehicle" table
CREATE TABLE IF NOT EXISTS `vehicle` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `tier` integer NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "vehicle_id_idx" to table: "vehicle"
CREATE INDEX IF NOT EXISTS `vehicle_id_idx` ON `vehicle` (`id`);
-- create "vehicle_average" table
CREATE TABLE IF NOT EXISTS `vehicle_average` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `data` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "vehicle_average_id_idx" to table: "vehicle_average"
CREATE INDEX IF NOT EXISTS `vehicle_average_id_idx` ON `vehicle_average` (`id`);
-- create "game_map" table
CREATE TABLE IF NOT EXISTS `game_map` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `game_modes` blob NOT NULL DEFAULT '', `supremacy_points` integer NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "game_map_id_idx" to table: "game_map"
CREATE INDEX IF NOT EXISTS `game_map_id_idx` ON `game_map` (`id`);
-- create "game_mode" table
CREATE TABLE IF NOT EXISTS `game_mode` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "game_mode_id_idx" to table: "game_mode"
CREATE INDEX IF NOT EXISTS `game_mode_id_idx` ON `game_mode` (`id`);
-- create "user" table
CREATE TABLE IF NOT EXISTS `user` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `username` text NOT NULL DEFAULT '', `permissions` text NOT NULL DEFAULT '', `feature_flags` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- CREATE INDEX IF NOT EXISTS "user_id_idx" to table: "user"
CREATE INDEX IF NOT EXISTS `user_id_idx` ON `user` (`id`);
-- CREATE INDEX IF NOT EXISTS "user_username_idx" to table: "user"
CREATE INDEX IF NOT EXISTS `user_username_idx` ON `user` (`username`);
-- create "user_connection" table
CREATE TABLE IF NOT EXISTS `user_connection` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `verified` bool NOT NULL DEFAULT false, `selected` bool NOT NULL DEFAULT false, `reference_id` text NOT NULL, `permissions` text NULL DEFAULT '', `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_connection_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "user_connection_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS `user_connection_id_idx` ON `user_connection` (`id`);
-- CREATE INDEX IF NOT EXISTS "user_connection_user_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS `user_connection_user_id_idx` ON `user_connection` (`user_id`);
-- CREATE INDEX IF NOT EXISTS "user_connection_type_user_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS `user_connection_type_user_id_idx` ON `user_connection` (`type`, `user_id`);
-- CREATE INDEX IF NOT EXISTS "user_connection_reference_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS `user_connection_reference_id_idx` ON `user_connection` (`reference_id`);
-- CREATE INDEX IF NOT EXISTS "user_connection_type_reference_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS `user_connection_type_reference_id_idx` ON `user_connection` (`type`, `reference_id`);
-- CREATE INDEX IF NOT EXISTS "user_connection_reference_id_user_id_type_idx" to table: "user_connection"
CREATE UNIQUE INDEX IF NOT EXISTS `user_connection_reference_id_user_id_type_idx` ON `user_connection` (`reference_id`, `user_id`, `type`);
-- create "user_subscription" table
CREATE TABLE IF NOT EXISTS `user_subscription` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `expires_at` text NOT NULL, `permissions` text NOT NULL, `reference_id` text NOT NULL, `user_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`), CONSTRAINT `user_subscription_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "user_subscription_id_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS `user_subscription_id_idx` ON `user_subscription` (`id`);
-- CREATE INDEX IF NOT EXISTS "user_subscription_user_id_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS `user_subscription_user_id_idx` ON `user_subscription` (`user_id`);
-- CREATE INDEX IF NOT EXISTS "user_subscription_type_user_id_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS `user_subscription_type_user_id_idx` ON `user_subscription` (`type`, `user_id`);
-- CREATE INDEX IF NOT EXISTS "user_subscription_expires_at_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS `user_subscription_expires_at_idx` ON `user_subscription` (`expires_at`);
-- CREATE INDEX IF NOT EXISTS "user_subscription_expires_at_user_id_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS `user_subscription_expires_at_user_id_idx` ON `user_subscription` (`expires_at`, `user_id`);
-- create "user_content" table
CREATE TABLE IF NOT EXISTS `user_content` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `value` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_content_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "user_content_id_idx" to table: "user_content"
CREATE INDEX IF NOT EXISTS `user_content_id_idx` ON `user_content` (`id`);
-- CREATE INDEX IF NOT EXISTS "user_content_user_id_idx" to table: "user_content"
CREATE INDEX IF NOT EXISTS `user_content_user_id_idx` ON `user_content` (`user_id`);
-- CREATE INDEX IF NOT EXISTS "user_content_reference_id_idx" to table: "user_content"
CREATE INDEX IF NOT EXISTS `user_content_reference_id_idx` ON `user_content` (`reference_id`);
-- CREATE INDEX IF NOT EXISTS "user_content_type_user_id_idx" to table: "user_content"
CREATE UNIQUE INDEX IF NOT EXISTS `user_content_type_user_id_idx` ON `user_content` (`type`, `user_id`);
-- CREATE INDEX IF NOT EXISTS "user_content_type_reference_id_idx" to table: "user_content"
CREATE UNIQUE INDEX IF NOT EXISTS `user_content_type_reference_id_idx` ON `user_content` (`type`, `reference_id`);
-- create "user_restriction" table
CREATE TABLE IF NOT EXISTS `user_restriction` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `expires_at` text NOT NULL, `type` text NOT NULL, `restriction` text NOT NULL, `public_reason` text NOT NULL, `moderator_comment` text NOT NULL, `events` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_restriction_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- CREATE INDEX IF NOT EXISTS "user_restriction_id_idx" to table: "user_restriction"
CREATE INDEX IF NOT EXISTS `user_restriction_id_idx` ON `user_restriction` (`id`);
-- CREATE INDEX IF NOT EXISTS "user_restriction_user_id_idx" to table: "user_restriction"
CREATE INDEX IF NOT EXISTS `user_restriction_user_id_idx` ON `user_restriction` (`user_id`);
-- CREATE INDEX IF NOT EXISTS "user_restriction_expires_at_user_id_idx" to table: "user_restriction"
CREATE INDEX IF NOT EXISTS `user_restriction_expires_at_user_id_idx` ON `user_restriction` (`expires_at`, `user_id`);
-- create "moderation_request" table
CREATE TABLE IF NOT EXISTS `moderation_request` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `moderator_comment` text NULL, `context` text NULL, `reference_id` text NOT NULL, `action_reason` text NULL, `action_status` text NOT NULL, `data` blob NOT NULL DEFAULT '', `requestor_id` text NOT NULL, `moderator_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `moderation_request_moderator_id_user_id_fk` FOREIGN KEY (`moderator_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `moderation_request_requestor_id_user_id_fk` FOREIGN KEY (`requestor_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "moderation_request_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS `moderation_request_id_idx` ON `moderation_request` (`id`);
-- CREATE INDEX IF NOT EXISTS "moderation_request_reference_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS `moderation_request_reference_id_idx` ON `moderation_request` (`reference_id`);
-- CREATE INDEX IF NOT EXISTS "moderation_request_requestor_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS `moderation_request_requestor_id_idx` ON `moderation_request` (`requestor_id`);
-- CREATE INDEX IF NOT EXISTS "moderation_request_moderator_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS `moderation_request_moderator_id_idx` ON `moderation_request` (`moderator_id`);
-- CREATE INDEX IF NOT EXISTS "moderation_request_requestor_id_reference_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS `moderation_request_requestor_id_reference_id_idx` ON `moderation_request` (`requestor_id`, `reference_id`);
-- CREATE INDEX IF NOT EXISTS "moderation_request_requestor_id_reference_id_action_status_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS `moderation_request_requestor_id_reference_id_action_status_idx` ON `moderation_request` (`requestor_id`, `reference_id`, `action_status`);
-- create "widget_settings" table
CREATE TABLE IF NOT EXISTS `widget_settings` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `reference_id` text NOT NULL, `title` text NULL, `session_from` text NULL, `metadata` blob NOT NULL DEFAULT '', `styles` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, `session_reference_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `widget_setting_user_id_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- CREATE INDEX IF NOT EXISTS "widget_settings_id" to table: "widget_settings"
CREATE INDEX IF NOT EXISTS `widget_settings_id` ON `widget_settings` (`id`);
