-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_account_snapshots" table
CREATE TABLE `new_account_snapshots` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `last_battle_time` datetime NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` json NOT NULL, `regular_battles` integer NOT NULL, `regular_frame` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_snapshots_accounts_account_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "account_snapshots" to new temporary table "new_account_snapshots"
INSERT INTO `new_account_snapshots` (`id`, `created_at`, `updated_at`, `type`, `last_battle_time`, `reference_id`, `rating_battles`, `rating_frame`, `regular_battles`, `regular_frame`, `account_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `last_battle_time`, `reference_id`, `rating_battles`, `rating_frame`, `regular_battles`, `regular_frame`, `account_id` FROM `account_snapshots`;
-- Drop "account_snapshots" table after copying rows
DROP TABLE `account_snapshots`;
-- Rename temporary table "new_account_snapshots" to "account_snapshots"
ALTER TABLE `new_account_snapshots` RENAME TO `account_snapshots`;
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
-- Create "new_achievements_snapshots" table
CREATE TABLE `new_achievements_snapshots` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` datetime NOT NULL, `data` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `achievements_snapshots_accounts_achievement_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "achievements_snapshots" to new temporary table "new_achievements_snapshots"
INSERT INTO `new_achievements_snapshots` (`id`, `created_at`, `updated_at`, `type`, `reference_id`, `battles`, `last_battle_time`, `data`, `account_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `reference_id`, `battles`, `last_battle_time`, `data`, `account_id` FROM `achievements_snapshots`;
-- Drop "achievements_snapshots" table after copying rows
DROP TABLE `achievements_snapshots`;
-- Rename temporary table "new_achievements_snapshots" to "achievements_snapshots"
ALTER TABLE `new_achievements_snapshots` RENAME TO `achievements_snapshots`;
-- Create index "achievementssnapshot_id" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_id` ON `achievements_snapshots` (`id`);
-- Create index "achievementssnapshot_created_at" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_created_at` ON `achievements_snapshots` (`created_at`);
-- Create index "achievementssnapshot_account_id_reference_id" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_account_id_reference_id` ON `achievements_snapshots` (`account_id`, `reference_id`);
-- Create index "achievementssnapshot_account_id_reference_id_created_at" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_account_id_reference_id_created_at` ON `achievements_snapshots` (`account_id`, `reference_id`, `created_at`);
-- Create "new_discord_interactions" table
CREATE TABLE `new_discord_interactions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `command` text NOT NULL, `reference_id` text NOT NULL, `type` text NOT NULL, `locale` text NOT NULL, `options` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `discord_interactions_users_discord_interactions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "discord_interactions" to new temporary table "new_discord_interactions"
INSERT INTO `new_discord_interactions` (`id`, `created_at`, `updated_at`, `command`, `reference_id`, `type`, `locale`, `options`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `command`, `reference_id`, `type`, `locale`, `options`, `user_id` FROM `discord_interactions`;
-- Drop "discord_interactions" table after copying rows
DROP TABLE `discord_interactions`;
-- Rename temporary table "new_discord_interactions" to "discord_interactions"
ALTER TABLE `new_discord_interactions` RENAME TO `discord_interactions`;
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
-- Create "new_user_connections" table
CREATE TABLE `new_user_connections` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `permissions` text NULL DEFAULT (''), `metadata` json NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_connections_users_connections` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "user_connections" to new temporary table "new_user_connections"
INSERT INTO `new_user_connections` (`id`, `created_at`, `updated_at`, `type`, `reference_id`, `permissions`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `reference_id`, `permissions`, `metadata`, `user_id` FROM `user_connections`;
-- Drop "user_connections" table after copying rows
DROP TABLE `user_connections`;
-- Rename temporary table "new_user_connections" to "user_connections"
ALTER TABLE `new_user_connections` RENAME TO `user_connections`;
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
-- Create "new_user_contents" table
CREATE TABLE `new_user_contents` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `value` text NOT NULL, `metadata` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_contents_users_content` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "user_contents" to new temporary table "new_user_contents"
INSERT INTO `new_user_contents` (`id`, `created_at`, `updated_at`, `type`, `reference_id`, `value`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `reference_id`, `value`, `metadata`, `user_id` FROM `user_contents`;
-- Drop "user_contents" table after copying rows
DROP TABLE `user_contents`;
-- Rename temporary table "new_user_contents" to "user_contents"
ALTER TABLE `new_user_contents` RENAME TO `user_contents`;
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
-- Create "new_user_subscriptions" table
CREATE TABLE `new_user_subscriptions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `expires_at` datetime NOT NULL, `permissions` text NOT NULL, `reference_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_subscriptions_users_subscriptions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "user_subscriptions" to new temporary table "new_user_subscriptions"
INSERT INTO `new_user_subscriptions` (`id`, `created_at`, `updated_at`, `type`, `expires_at`, `permissions`, `reference_id`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `expires_at`, `permissions`, `reference_id`, `user_id` FROM `user_subscriptions`;
-- Drop "user_subscriptions" table after copying rows
DROP TABLE `user_subscriptions`;
-- Rename temporary table "new_user_subscriptions" to "user_subscriptions"
ALTER TABLE `new_user_subscriptions` RENAME TO `user_subscriptions`;
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
-- Create "new_vehicle_snapshots" table
CREATE TABLE `new_vehicle_snapshots` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` datetime NOT NULL, `frame` json NOT NULL, `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_snapshots_accounts_vehicle_snapshots` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "vehicle_snapshots" to new temporary table "new_vehicle_snapshots"
INSERT INTO `new_vehicle_snapshots` (`id`, `created_at`, `updated_at`, `type`, `vehicle_id`, `reference_id`, `battles`, `last_battle_time`, `frame`, `account_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `vehicle_id`, `reference_id`, `battles`, `last_battle_time`, `frame`, `account_id` FROM `vehicle_snapshots`;
-- Drop "vehicle_snapshots" table after copying rows
DROP TABLE `vehicle_snapshots`;
-- Rename temporary table "new_vehicle_snapshots" to "vehicle_snapshots"
ALTER TABLE `new_vehicle_snapshots` RENAME TO `vehicle_snapshots`;
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
-- Create "new_sessions" table
CREATE TABLE `new_sessions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `expires_at` datetime NOT NULL, `public_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "sessions" to new temporary table "new_sessions"
INSERT INTO `new_sessions` (`id`, `created_at`, `updated_at`, `expires_at`, `public_id`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `expires_at`, `public_id`, `user_id` FROM `sessions`;
-- Drop "sessions" table after copying rows
DROP TABLE `sessions`;
-- Rename temporary table "new_sessions" to "sessions"
ALTER TABLE `new_sessions` RENAME TO `sessions`;
-- Create index "sessions_public_id_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_public_id_key` ON `sessions` (`public_id`);
-- Create index "session_public_id_expires_at" to table: "sessions"
CREATE INDEX `session_public_id_expires_at` ON `sessions` (`public_id`, `expires_at`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
