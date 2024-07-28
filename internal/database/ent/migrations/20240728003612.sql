-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_discord_interactions" table
CREATE TABLE `new_discord_interactions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `result` text NOT NULL, `event_id` text NOT NULL, `guild_id` text NOT NULL, `snowflake` text NOT NULL DEFAULT (''), `channel_id` text NOT NULL, `message_id` text NOT NULL, `type` text NOT NULL, `locale` text NOT NULL, `metadata` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `discord_interactions_users_discord_interactions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "discord_interactions" to new temporary table "new_discord_interactions"
INSERT INTO `new_discord_interactions` (`id`, `created_at`, `updated_at`, `result`, `event_id`, `guild_id`, `channel_id`, `message_id`, `type`, `locale`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `result`, `event_id`, `guild_id`, `channel_id`, `message_id`, `type`, `locale`, `metadata`, `user_id` FROM `discord_interactions`;
-- Drop "discord_interactions" table after copying rows
DROP TABLE `discord_interactions`;
-- Rename temporary table "new_discord_interactions" to "discord_interactions"
ALTER TABLE `new_discord_interactions` RENAME TO `discord_interactions`;
-- Create index "discordinteraction_id" to table: "discord_interactions"
CREATE INDEX `discordinteraction_id` ON `discord_interactions` (`id`);
-- Create index "discordinteraction_user_id" to table: "discord_interactions"
CREATE INDEX `discordinteraction_user_id` ON `discord_interactions` (`user_id`);
-- Create index "discordinteraction_snowflake" to table: "discord_interactions"
CREATE INDEX `discordinteraction_snowflake` ON `discord_interactions` (`snowflake`);
-- Create index "discordinteraction_created_at" to table: "discord_interactions"
CREATE INDEX `discordinteraction_created_at` ON `discord_interactions` (`created_at`);
-- Create index "discordinteraction_user_id_type_created_at" to table: "discord_interactions"
CREATE INDEX `discordinteraction_user_id_type_created_at` ON `discord_interactions` (`user_id`, `type`, `created_at`);
-- Create index "discordinteraction_channel_id_type_created_at" to table: "discord_interactions"
CREATE INDEX `discordinteraction_channel_id_type_created_at` ON `discord_interactions` (`channel_id`, `type`, `created_at`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
