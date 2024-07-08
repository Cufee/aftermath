-- Create "leaderboard_scores" table
CREATE TABLE `leaderboard_scores` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `score` real NOT NULL, `reference_id` text NOT NULL, `leaderboard_id` text NOT NULL, `meta` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "leaderboardscore_id" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_id` ON `leaderboard_scores` (`id`);
-- Create index "leaderboardscore_created_at" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_created_at` ON `leaderboard_scores` (`created_at`);
-- Create index "leaderboardscore_created_at_type" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_created_at_type` ON `leaderboard_scores` (`created_at`, `type`);
-- Create index "leaderboardscore_score_type" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_score_type` ON `leaderboard_scores` (`score`, `type`);
-- Create index "leaderboardscore_leaderboard_id_type" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_type` ON `leaderboard_scores` (`leaderboard_id`, `type`);
-- Create index "leaderboardscore_leaderboard_id_score_type" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_score_type` ON `leaderboard_scores` (`leaderboard_id`, `score`, `type`);
-- Create index "leaderboardscore_leaderboard_id_reference_id_type" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_reference_id_type` ON `leaderboard_scores` (`leaderboard_id`, `reference_id`, `type`);
-- Create index "leaderboardscore_leaderboard_id_reference_id_score_type" to table: "leaderboard_scores"
CREATE INDEX `leaderboardscore_leaderboard_id_reference_id_score_type` ON `leaderboard_scores` (`leaderboard_id`, `reference_id`, `score`, `type`);
