-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "leaderboard_scores" table
DROP TABLE `leaderboard_scores`;
-- Create "leaderboard_scores" table
CREATE TABLE `leaderboard_scores` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `score` real NOT NULL, `account_id` text NOT NULL, `reference_id` text NOT NULL, `leaderboard_id` text NOT NULL, `meta` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "leaderboardscore_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_id` ON `leaderboard_scores` (`id`);
-- Create index "leaderboardscore_created_at" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_created_at` ON `leaderboard_scores` (`created_at`);
-- Create index "leaderboardscore_account_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_account_id` ON `leaderboard_scores` (`account_id`);
-- Create index "leaderboardscore_reference_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_reference_id` ON `leaderboard_scores` (`reference_id`);
-- Create index "leaderboardscore_leaderboard_id_type_account_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_type_account_id` ON `leaderboard_scores` (`leaderboard_id`, `type`, `account_id`);
-- Create index "leaderboardscore_leaderboard_id_type_reference_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_type_reference_id` ON `leaderboard_scores` (`leaderboard_id`, `type`, `reference_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
