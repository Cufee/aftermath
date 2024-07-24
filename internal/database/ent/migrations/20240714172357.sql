-- Create "game_maps" table
CREATE TABLE `game_maps` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `game_modes` json NOT NULL, `supremacy_points` integer NOT NULL, `localized_names` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "gamemap_id" to table: "game_maps"
CREATE INDEX `gamemap_id` ON `game_maps` (`id`);
