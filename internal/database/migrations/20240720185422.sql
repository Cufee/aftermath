-- Create "game_modes" table
CREATE TABLE `game_modes` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `localized_names` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "gamemode_id" to table: "game_modes"
CREATE INDEX `gamemode_id` ON `game_modes` (`id`);
