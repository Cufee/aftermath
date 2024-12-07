-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Rename table "cron_tasks" to "cron_task"
ALTER TABLE `cron_tasks` RENAME TO `cron_task`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
