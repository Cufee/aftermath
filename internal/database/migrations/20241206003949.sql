-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_vehicle_snapshot" table
CREATE TABLE `new_vehicle_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_snapshot_account_id_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "vehicle_snapshot" to new temporary table "new_vehicle_snapshot"
INSERT INTO `new_vehicle_snapshot` (`id`, `created_at`, `type`, `vehicle_id`, `reference_id`, `battles`, `last_battle_time`, `frame`, `account_id`) SELECT `id`, `created_at`, `type`, `vehicle_id`, `reference_id`, `battles`, `last_battle_time`, `frame`, `account_id` FROM `vehicle_snapshot`;
-- Drop "vehicle_snapshot" table after copying rows
DROP TABLE `vehicle_snapshot`;
-- Rename temporary table "new_vehicle_snapshot" to "vehicle_snapshot"
ALTER TABLE `new_vehicle_snapshot` RENAME TO `vehicle_snapshot`;
-- Create index "vehicle_snapshot_id_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_id_idx` ON `vehicle_snapshot` (`id`);
-- Create index "vehicle_snapshot_type_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_type_idx` ON `vehicle_snapshot` (`type`);
-- Create index "vehicle_snapshot_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_created_at_idx` ON `vehicle_snapshot` (`created_at`);
-- Create index "vehicle_snapshot_type_account_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_type_account_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "vehicle_snapshot_type_account_id_vehicle_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_type_account_id_vehicle_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `vehicle_id`, `created_at`);
-- Create index "vehicle_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx` ON `vehicle_snapshot` (`type`, `account_id`, `vehicle_id`, `reference_id`, `created_at`);
-- Create index "vehicle_snapshot_vehicle_id_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_vehicle_id_reference_id_created_at_idx` ON `vehicle_snapshot` (`vehicle_id`, `reference_id`, `created_at`);
-- Create "new_account_snapshot" table
CREATE TABLE `new_account_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `last_battle_time` text NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` blob NOT NULL DEFAULT '', `regular_battles` integer NOT NULL, `regular_frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_account_id_snapshot_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "account_snapshot" to new temporary table "new_account_snapshot"
INSERT INTO `new_account_snapshot` (`id`, `created_at`, `type`, `last_battle_time`, `reference_id`, `rating_battles`, `rating_frame`, `regular_battles`, `regular_frame`, `account_id`) SELECT `id`, `created_at`, `type`, `last_battle_time`, `reference_id`, `rating_battles`, `rating_frame`, `regular_battles`, `regular_frame`, `account_id` FROM `account_snapshot`;
-- Drop "account_snapshot" table after copying rows
DROP TABLE `account_snapshot`;
-- Rename temporary table "new_account_snapshot" to "account_snapshot"
ALTER TABLE `new_account_snapshot` RENAME TO `account_snapshot`;
-- Create index "account_snapshot_id_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_id_idx` ON `account_snapshot` (`id`);
-- Create index "account_snapshot_type_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_idx` ON `account_snapshot` (`type`);
-- Create index "account_snapshot_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_created_at_idx` ON `account_snapshot` (`created_at`);
-- Create index "account_snapshot_type_account_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "account_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_type_account_id_reference_id_created_at_idx` ON `account_snapshot` (`type`, `account_id`, `reference_id`, `created_at`);
-- Create index "account_snapshot_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_account_id_reference_id_created_at_idx` ON `account_snapshot` (`account_id`, `reference_id`, `created_at`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
