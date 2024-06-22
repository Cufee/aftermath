-- Create "accounts" table
CREATE TABLE `accounts` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `last_battle_time` integer NOT NULL, `account_created_at` integer NOT NULL, `realm` text NOT NULL, `nickname` text NOT NULL, `private` bool NOT NULL DEFAULT (false), `clan_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `accounts_clans_accounts` FOREIGN KEY (`clan_id`) REFERENCES `clans` (`id`) ON DELETE SET NULL);
-- Create index "account_id_last_battle_time" to table: "accounts"
CREATE INDEX `account_id_last_battle_time` ON `accounts` (`id`, `last_battle_time`);
-- Create index "account_realm" to table: "accounts"
CREATE INDEX `account_realm` ON `accounts` (`realm`);
-- Create index "account_realm_last_battle_time" to table: "accounts"
CREATE INDEX `account_realm_last_battle_time` ON `accounts` (`realm`, `last_battle_time`);
-- Create index "account_clan_id" to table: "accounts"
CREATE INDEX `account_clan_id` ON `accounts` (`clan_id`);
-- Create "account_snapshots" table
CREATE TABLE `account_snapshots` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `type` text NOT NULL, `last_battle_time` integer NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` json NOT NULL, `regular_battles` integer NOT NULL, `regular_frame` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_snapshots_accounts_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE NO ACTION);
-- Create index "accountsnapshot_created_at" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_created_at` ON `account_snapshots` (`created_at`);
-- Create index "accountsnapshot_type_account_id_created_at" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_type_account_id_created_at` ON `account_snapshots` (`type`, `account_id`, `created_at`);
-- Create index "accountsnapshot_type_account_id_reference_id" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_type_account_id_reference_id` ON `account_snapshots` (`type`, `account_id`, `reference_id`);
-- Create index "accountsnapshot_type_account_id_reference_id_created_at" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_type_account_id_reference_id_created_at` ON `account_snapshots` (`type`, `account_id`, `reference_id`, `created_at`);
-- Create "achievements_snapshots" table
CREATE TABLE `achievements_snapshots` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` integer NOT NULL, `data` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `achievements_snapshots_accounts_achievement_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE NO ACTION);
-- Create index "achievementssnapshot_created_at" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_created_at` ON `achievements_snapshots` (`created_at`);
-- Create index "achievementssnapshot_account_id_reference_id" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_account_id_reference_id` ON `achievements_snapshots` (`account_id`, `reference_id`);
-- Create index "achievementssnapshot_account_id_reference_id_created_at" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_account_id_reference_id_created_at` ON `achievements_snapshots` (`account_id`, `reference_id`, `created_at`);
-- Create "app_configurations" table
CREATE TABLE `app_configurations` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `key` text NOT NULL, `value` json NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- Create index "app_configurations_key_key" to table: "app_configurations"
CREATE UNIQUE INDEX `app_configurations_key_key` ON `app_configurations` (`key`);
-- Create index "appconfiguration_key" to table: "app_configurations"
CREATE INDEX `appconfiguration_key` ON `app_configurations` (`key`);
-- Create "application_commands" table
CREATE TABLE `application_commands` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `name` text NOT NULL, `version` text NOT NULL, `options_hash` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "application_commands_name_key" to table: "application_commands"
CREATE UNIQUE INDEX `application_commands_name_key` ON `application_commands` (`name`);
-- Create index "applicationcommand_options_hash" to table: "application_commands"
CREATE INDEX `applicationcommand_options_hash` ON `application_commands` (`options_hash`);
-- Create "clans" table
CREATE TABLE `clans` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `tag` text NOT NULL, `name` text NOT NULL, `emblem_id` text NULL DEFAULT (''), `members` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "clan_tag" to table: "clans"
CREATE INDEX `clan_tag` ON `clans` (`tag`);
-- Create index "clan_name" to table: "clans"
CREATE INDEX `clan_name` ON `clans` (`name`);
-- Create "cron_tasks" table
CREATE TABLE `cron_tasks` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `targets` json NOT NULL, `status` text NOT NULL, `scheduled_after` integer NOT NULL, `last_run` integer NOT NULL, `logs` json NOT NULL, `data` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "crontask_reference_id" to table: "cron_tasks"
CREATE INDEX `crontask_reference_id` ON `cron_tasks` (`reference_id`);
-- Create index "crontask_status_last_run" to table: "cron_tasks"
CREATE INDEX `crontask_status_last_run` ON `cron_tasks` (`status`, `last_run`);
-- Create index "crontask_status_created_at" to table: "cron_tasks"
CREATE INDEX `crontask_status_created_at` ON `cron_tasks` (`status`, `created_at`);
-- Create index "crontask_status_scheduled_after" to table: "cron_tasks"
CREATE INDEX `crontask_status_scheduled_after` ON `cron_tasks` (`status`, `scheduled_after`);
-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `permissions` text NOT NULL DEFAULT (''), `feature_flags` json NULL, PRIMARY KEY (`id`));
-- Create "user_connections" table
CREATE TABLE `user_connections` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `permissions` text NULL DEFAULT (''), `metadata` json NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_connections_users_connections` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "userconnection_user_id" to table: "user_connections"
CREATE INDEX `userconnection_user_id` ON `user_connections` (`user_id`);
-- Create index "userconnection_type_user_id" to table: "user_connections"
CREATE INDEX `userconnection_type_user_id` ON `user_connections` (`type`, `user_id`);
-- Create index "userconnection_reference_id" to table: "user_connections"
CREATE INDEX `userconnection_reference_id` ON `user_connections` (`reference_id`);
-- Create index "userconnection_type_reference_id" to table: "user_connections"
CREATE INDEX `userconnection_type_reference_id` ON `user_connections` (`type`, `reference_id`);
-- Create "user_contents" table
CREATE TABLE `user_contents` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `value` json NOT NULL, `metadata` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_contents_users_content` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "usercontent_user_id" to table: "user_contents"
CREATE INDEX `usercontent_user_id` ON `user_contents` (`user_id`);
-- Create index "usercontent_type_user_id" to table: "user_contents"
CREATE INDEX `usercontent_type_user_id` ON `user_contents` (`type`, `user_id`);
-- Create index "usercontent_reference_id" to table: "user_contents"
CREATE INDEX `usercontent_reference_id` ON `user_contents` (`reference_id`);
-- Create index "usercontent_type_reference_id" to table: "user_contents"
CREATE INDEX `usercontent_type_reference_id` ON `user_contents` (`type`, `reference_id`);
-- Create "user_subscriptions" table
CREATE TABLE `user_subscriptions` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `type` text NOT NULL, `expires_at` integer NOT NULL, `permissions` text NOT NULL, `reference_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_subscriptions_users_subscriptions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "usersubscription_user_id" to table: "user_subscriptions"
CREATE INDEX `usersubscription_user_id` ON `user_subscriptions` (`user_id`);
-- Create index "usersubscription_type_user_id" to table: "user_subscriptions"
CREATE INDEX `usersubscription_type_user_id` ON `user_subscriptions` (`type`, `user_id`);
-- Create index "usersubscription_expires_at" to table: "user_subscriptions"
CREATE INDEX `usersubscription_expires_at` ON `user_subscriptions` (`expires_at`);
-- Create index "usersubscription_expires_at_user_id" to table: "user_subscriptions"
CREATE INDEX `usersubscription_expires_at_user_id` ON `user_subscriptions` (`expires_at`, `user_id`);
-- Create "vehicles" table
CREATE TABLE `vehicles` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `tier` integer NOT NULL, `localized_names` json NOT NULL, PRIMARY KEY (`id`));
-- Create "vehicle_averages" table
CREATE TABLE `vehicle_averages` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `data` json NOT NULL, PRIMARY KEY (`id`));
-- Create "vehicle_snapshots" table
CREATE TABLE `vehicle_snapshots` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` integer NOT NULL, `frame` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_snapshots_accounts_vehicle_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE NO ACTION);
-- Create index "vehiclesnapshot_created_at" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_created_at` ON `vehicle_snapshots` (`created_at`);
-- Create index "vehiclesnapshot_vehicle_id_created_at" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_vehicle_id_created_at` ON `vehicle_snapshots` (`vehicle_id`, `created_at`);
-- Create index "vehiclesnapshot_account_id_created_at" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_account_id_created_at` ON `vehicle_snapshots` (`account_id`, `created_at`);
-- Create index "vehiclesnapshot_account_id_type_created_at" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_account_id_type_created_at` ON `vehicle_snapshots` (`account_id`, `type`, `created_at`);
