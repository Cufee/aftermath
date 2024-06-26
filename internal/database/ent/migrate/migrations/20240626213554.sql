-- Create "discord_interactions" table
CREATE TABLE `discord_interactions` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `command` text NOT NULL, `reference_id` text NOT NULL, `type` text NOT NULL, `locale` text NOT NULL, `options` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `discord_interactions_users_discord_interactions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
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
