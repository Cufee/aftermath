-- Add column "tries_left" to table: "cron_tasks"
ALTER TABLE `cron_tasks` ADD COLUMN `tries_left` integer NOT NULL;
