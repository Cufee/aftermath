-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_account_snapshot" table
CREATE TABLE `new_account_snapshot` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `last_battle_time` datetime NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` blob NOT NULL DEFAULT '', `regular_battles` integer NOT NULL, `regular_frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "account_snapshot" to new temporary table "new_account_snapshot"
INSERT INTO `new_account_snapshot` (`id`, `created_at`, `updated_at`, `type`, `last_battle_time`, `reference_id`, `rating_battles`, `rating_frame`, `regular_battles`, `regular_frame`, `account_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `last_battle_time`, `reference_id`, `rating_battles`, `rating_frame`, `regular_battles`, `regular_frame`, `account_id` FROM `account_snapshot`;
-- Drop "account_snapshot" table after copying rows
DROP TABLE `account_snapshot`;
-- Rename temporary table "new_account_snapshot" to "account_snapshot"
ALTER TABLE `new_account_snapshot` RENAME TO `account_snapshot`;
-- Create index "account_snapshot_id_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_id_idx` ON `account_snapshot` (`id`);
-- Create index "account_snapshot_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_created_at_idx` ON `account_snapshot` (`created_at`);
-- Create index "account_snapshot_type_account_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "account_snapshot_type_account_id_reference_id_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_reference_id_idx` ON `account_snapshot` (`type`, `account_id`, `reference_id`);
-- Create index "account_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_reference_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `reference_id`, `created_at`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
