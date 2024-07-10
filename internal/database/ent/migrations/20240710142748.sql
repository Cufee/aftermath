-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "achievements_snapshots" table
DROP TABLE `achievements_snapshots`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
