-- Create "ad_events" table
CREATE TABLE `ad_events` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `user_id` text NOT NULL, `guild_id` text NOT NULL, `channel_id` text NOT NULL, `locale` text NOT NULL, `message_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "adevent_id" to table: "ad_events"
CREATE INDEX `adevent_id` ON `ad_events` (`id`);
-- Create index "adevent_user_id_guild_id_channel_id_created_at" to table: "ad_events"
CREATE INDEX `adevent_user_id_guild_id_channel_id_created_at` ON `ad_events` (`user_id`, `guild_id`, `channel_id`, `created_at`);
-- Create "ad_messages" table
CREATE TABLE `ad_messages` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `enabled` bool NOT NULL, `weight` integer NOT NULL, `chance` real NOT NULL, `message` json NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- Create index "admessage_id" to table: "ad_messages"
CREATE INDEX `admessage_id` ON `ad_messages` (`id`);
-- Create index "admessage_weight_enabled" to table: "ad_messages"
CREATE INDEX `admessage_weight_enabled` ON `ad_messages` (`weight`, `enabled`);
