-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_leaderboard_scores" table
CREATE TABLE `new_leaderboard_scores` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `score` real NOT NULL, `reference_id` text NOT NULL, `leaderboard_id` text NOT NULL, `meta` json NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "leaderboard_scores" to new temporary table "new_leaderboard_scores"
INSERT INTO `new_leaderboard_scores` (`id`, `created_at`, `updated_at`, `type`, `score`, `reference_id`, `leaderboard_id`, `meta`) SELECT `id`, `created_at`, `updated_at`, `type`, `score`, `reference_id`, `leaderboard_id`, `meta` FROM `leaderboard_scores`;
-- Drop "leaderboard_scores" table after copying rows
DROP TABLE `leaderboard_scores`;
-- Rename temporary table "new_leaderboard_scores" to "leaderboard_scores"
ALTER TABLE `new_leaderboard_scores` RENAME TO `leaderboard_scores`;
-- Create index "leaderboardscore_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_id` ON `leaderboard_scores` (`id`);
-- Create index "leaderboardscore_created_at" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_created_at` ON `leaderboard_scores` (`created_at`);
-- Create index "leaderboardscore_reference_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_reference_id` ON `leaderboard_scores` (`reference_id`);
-- Create index "leaderboardscore_leaderboard_id_type_reference_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_type_reference_id` ON `leaderboard_scores` (`leaderboard_id`, `type`, `reference_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
