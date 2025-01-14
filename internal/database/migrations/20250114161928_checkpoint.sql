-- atlas:checkpoint

-- Create "widget_settings" table
CREATE TABLE `widget_settings` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `reference_id` text NOT NULL, `title` text NULL, `session_from` text NULL, `metadata` blob NOT NULL DEFAULT '', `styles` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, `session_reference_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `widget_setting_user_id_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "widget_settings_id" to table: "widget_settings"
CREATE INDEX `widget_settings_id` ON `widget_settings` (`id`);
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
-- Create "account" table
CREATE TABLE `account` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `last_battle_time` text NOT NULL, `account_created_at` text NOT NULL, `realm` text NOT NULL, `nickname` text NOT NULL, `private` bool NOT NULL DEFAULT false, `clan_id` text NULL, PRIMARY KEY (`id`));
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
-- Create "account_snapshot" table
CREATE TABLE `account_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `last_battle_time` text NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` blob NOT NULL DEFAULT '', `regular_battles` integer NOT NULL, `regular_frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_account_id_snapshot_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
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
-- Create "app_configuration" table
CREATE TABLE `app_configuration` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `key` text NOT NULL, `value` blob NOT NULL DEFAULT '', `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "app_configuration_id_idx" to table: "app_configuration"
CREATE INDEX `app_configuration_id_idx` ON `app_configuration` (`id`);
-- Create index "app_configuration_key_idx" to table: "app_configuration"
CREATE UNIQUE INDEX `app_configuration_key_idx` ON `app_configuration` (`key`);
-- Create "application_command" table
CREATE TABLE `application_command` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `name` text NOT NULL, `version` text NOT NULL, `options_hash` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "application_command_id_idx" to table: "application_command"
CREATE INDEX `application_command_id_idx` ON `application_command` (`id`);
-- Create index "application_command_options_hash_idx" to table: "application_command"
CREATE INDEX `application_command_options_hash_idx` ON `application_command` (`options_hash`);
-- Create "auth_nonce" table
CREATE TABLE `auth_nonce` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `active` bool NOT NULL, `expires_at` text NOT NULL, `identifier` text NOT NULL, `public_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "auth_nonce_public_id_idx" to table: "auth_nonce"
CREATE UNIQUE INDEX `auth_nonce_public_id_idx` ON `auth_nonce` (`public_id`);
-- Create index "auth_nonce_public_id_active_expires_at_idx" to table: "auth_nonce"
CREATE INDEX `auth_nonce_public_id_active_expires_at_idx` ON `auth_nonce` (`public_id`, `active`, `expires_at`);
-- Create "clan" table
CREATE TABLE `clan` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `tag` text NOT NULL, `name` text NOT NULL, `emblem_id` text NULL DEFAULT '', `members` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "clan_id_idx" to table: "clan"
CREATE INDEX `clan_id_idx` ON `clan` (`id`);
-- Create index "clan_tag_idx" to table: "clan"
CREATE INDEX `clan_tag_idx` ON `clan` (`tag`);
-- Create index "clan_name_idx" to table: "clan"
CREATE INDEX `clan_name_idx` ON `clan` (`name`);
-- Create "cron_task" table
CREATE TABLE `cron_task` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `targets` blob NOT NULL DEFAULT '', `status` text NOT NULL, `scheduled_after` text NOT NULL, `last_run` text NOT NULL, `tries_left` integer NOT NULL DEFAULT 0, `logs` blob NOT NULL DEFAULT '', `data` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
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
-- Create "session" table
CREATE TABLE `session` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `expires_at` text NOT NULL, `public_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `session_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "session_public_id_idx" to table: "session"
CREATE UNIQUE INDEX `session_public_id_idx` ON `session` (`public_id`);
-- Create index "session_public_id_expires_at_idx" to table: "session"
CREATE INDEX `session_public_id_expires_at_idx` ON `session` (`public_id`, `expires_at`);
-- Create "user" table
CREATE TABLE `user` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `username` text NOT NULL DEFAULT '', `permissions` text NOT NULL DEFAULT '', `feature_flags` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "user_id_idx" to table: "user"
CREATE INDEX `user_id_idx` ON `user` (`id`);
-- Create index "user_username_idx" to table: "user"
CREATE INDEX `user_username_idx` ON `user` (`username`);
-- Create "user_connection" table
CREATE TABLE `user_connection` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `verified` bool NOT NULL DEFAULT false, `selected` bool NOT NULL DEFAULT false, `reference_id` text NOT NULL, `permissions` text NULL DEFAULT '', `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_connection_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
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
-- Create "vehicle" table
CREATE TABLE `vehicle` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `tier` integer NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "vehicle_id_idx" to table: "vehicle"
CREATE INDEX `vehicle_id_idx` ON `vehicle` (`id`);
-- Create "vehicle_average" table
CREATE TABLE `vehicle_average` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `data` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "vehicle_average_id_idx" to table: "vehicle_average"
CREATE INDEX `vehicle_average_id_idx` ON `vehicle_average` (`id`);
-- Create "vehicle_snapshot" table
CREATE TABLE `vehicle_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_snapshot_account_id_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
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
-- Create "game_map" table
CREATE TABLE `game_map` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `game_modes` blob NOT NULL DEFAULT '', `supremacy_points` integer NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "game_map_id_idx" to table: "game_map"
CREATE INDEX `game_map_id_idx` ON `game_map` (`id`);
-- Create "game_mode" table
CREATE TABLE `game_mode` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `localized_names` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "game_mode_id_idx" to table: "game_mode"
CREATE INDEX `game_mode_id_idx` ON `game_mode` (`id`);
-- Create "moderation_request" table
CREATE TABLE `moderation_request` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `moderator_comment` text NULL, `context` text NULL, `reference_id` text NOT NULL, `action_reason` text NULL, `action_status` text NOT NULL, `data` blob NOT NULL DEFAULT '', `requestor_id` text NOT NULL, `moderator_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `moderation_request_requestor_id_user_id_fk` FOREIGN KEY (`requestor_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `moderation_request_moderator_id_user_id_fk` FOREIGN KEY (`moderator_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
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
-- Create "user_restriction" table
CREATE TABLE `user_restriction` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `expires_at` text NOT NULL, `type` text NOT NULL, `restriction` text NOT NULL, `public_reason` text NOT NULL, `moderator_comment` text NOT NULL, `events` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_restriction_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "user_restriction_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_id_idx` ON `user_restriction` (`id`);
-- Create index "user_restriction_user_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_user_id_idx` ON `user_restriction` (`user_id`);
-- Create index "user_restriction_expires_at_user_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_expires_at_user_id_idx` ON `user_restriction` (`expires_at`, `user_id`);
-- Create "user_content" table
CREATE TABLE `user_content` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `value` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_content_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
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
-- Create "discord_interaction" table
CREATE TABLE `discord_interaction` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `result` text NOT NULL, `event_id` text NOT NULL, `guild_id` text NOT NULL, `snowflake` text NOT NULL DEFAULT '', `channel_id` text NOT NULL, `message_id` text NOT NULL, `type` text NOT NULL, `locale` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `discord_interaction_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
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
-- Create "user_subscription" table
CREATE TABLE `user_subscription` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `expires_at` text NOT NULL, `permissions` text NOT NULL, `reference_id` text NOT NULL, `user_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`), CONSTRAINT `user_subscription_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
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
