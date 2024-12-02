-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "discord_interactions" table
DROP TABLE `discord_interactions`;
-- Create "discord_interactions" table
CREATE TABLE `discord_interactions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `result` text NOT NULL, `event_id` text NOT NULL, `guild_id` text NOT NULL, `channel_id` text NOT NULL, `message_id` text NOT NULL, `type` text NOT NULL, `locale` text NOT NULL, `metadata` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `discord_interactions_users_discord_interactions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "discordinteraction_id" to table: "discord_interactions"
CREATE INDEX `discordinteraction_id` ON `discord_interactions` (`id`);
-- Create index "discordinteraction_user_id" to table: "discord_interactions"
CREATE INDEX `discordinteraction_user_id` ON `discord_interactions` (`user_id`);
-- Create index "discordinteraction_created_at" to table: "discord_interactions"
CREATE INDEX `discordinteraction_created_at` ON `discord_interactions` (`created_at`);
-- Create index "discordinteraction_user_id_type_created_at" to table: "discord_interactions"
CREATE INDEX `discordinteraction_user_id_type_created_at` ON `discord_interactions` (`user_id`, `type`, `created_at`);
-- Create index "discordinteraction_channel_id_type_created_at" to table: "discord_interactions"
CREATE INDEX `discordinteraction_channel_id_type_created_at` ON `discord_interactions` (`channel_id`, `type`, `created_at`);
-- Drop "ad_events" table
DROP TABLE `ad_events`;
-- Drop "ad_messages" table
DROP TABLE `ad_messages`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
