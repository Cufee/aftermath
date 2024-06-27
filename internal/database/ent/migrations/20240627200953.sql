-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_cron_tasks" table
CREATE TABLE `new_cron_tasks` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `targets` json NOT NULL, `status` text NOT NULL, `scheduled_after` datetime NOT NULL, `last_run` datetime NOT NULL, `tries_left` integer NOT NULL DEFAULT (0), `logs` json NOT NULL, `data` json NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "cron_tasks" to new temporary table "new_cron_tasks"
INSERT INTO `new_cron_tasks` (`id`, `created_at`, `updated_at`, `type`, `reference_id`, `targets`, `status`, `scheduled_after`, `last_run`, `logs`, `data`) SELECT `id`, `created_at`, `updated_at`, `type`, `reference_id`, `targets`, `status`, `scheduled_after`, `last_run`, `logs`, `data` FROM `cron_tasks`;
-- Drop "cron_tasks" table after copying rows
DROP TABLE `cron_tasks`;
-- Rename temporary table "new_cron_tasks" to "cron_tasks"
ALTER TABLE `new_cron_tasks` RENAME TO `cron_tasks`;
-- Create index "crontask_id" to table: "cron_tasks"
CREATE INDEX `crontask_id` ON `cron_tasks` (`id`);
-- Create index "crontask_reference_id" to table: "cron_tasks"
CREATE INDEX `crontask_reference_id` ON `cron_tasks` (`reference_id`);
-- Create index "crontask_status_last_run" to table: "cron_tasks"
CREATE INDEX `crontask_status_last_run` ON `cron_tasks` (`status`, `last_run`);
-- Create index "crontask_status_created_at" to table: "cron_tasks"
CREATE INDEX `crontask_status_created_at` ON `cron_tasks` (`status`, `created_at`);
-- Create index "crontask_status_scheduled_after" to table: "cron_tasks"
CREATE INDEX `crontask_status_scheduled_after` ON `cron_tasks` (`status`, `scheduled_after`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
