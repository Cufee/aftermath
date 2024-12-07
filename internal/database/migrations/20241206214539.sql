-- Create "account_achievements_snapshot" table
CREATE TABLE `account_achievements_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `reference_id` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_id_account_achievements_snapshot_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "account_achievements_snapshot_id_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_id_idx` ON `account_achievements_snapshot` (`id`);
-- Create index "account_achievements_snapshot_type_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_type_idx` ON `account_achievements_snapshot` (`type`);
-- Create index "account_achievements_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_created_at_idx` ON `account_achievements_snapshot` (`created_at`);
-- Create index "account_achievements_snapshot_type_account_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_type_account_id_created_at_idx` ON `account_achievements_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "account_achievements_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_type_account_id_reference_id_created_at_idx` ON `account_achievements_snapshot` (`type`, `account_id`, `reference_id`, `created_at`);
-- Create index "account_achievements_snapshot_account_id_reference_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_account_id_reference_id_created_at_idx` ON `account_achievements_snapshot` (`account_id`, `reference_id`, `created_at`);
-- Create "vehicle_achievements_snapshot" table
CREATE TABLE `vehicle_achievements_snapshot` (`id` text NOT NULL, `created_at` text NOT NULL, `type` text NOT NULL, `vehicle_id` text NOT NULL, `reference_id` text NOT NULL, `battles` integer NOT NULL, `last_battle_time` text NOT NULL, `frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `vehicle_achievements_snapshot_account_id_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "vehicle_achievements_snapshot_id_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_id_idx` ON `vehicle_achievements_snapshot` (`id`);
-- Create index "vehicle_achievements_snapshot_type_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_type_idx` ON `vehicle_achievements_snapshot` (`type`);
-- Create index "vehicle_achievements_snapshot_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_created_at_idx` ON `vehicle_achievements_snapshot` (`created_at`);
-- Create index "vehicle_achievements_snapshot_type_account_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_type_account_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `created_at`);
-- Create index "vehicle_achievements_snapshot_type_account_id_vehicle_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_type_account_id_vehicle_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `vehicle_id`, `created_at`);
-- Create index "vehicle_achievements_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx` ON `vehicle_achievements_snapshot` (`type`, `account_id`, `vehicle_id`, `reference_id`, `created_at`);
-- Create index "vehicle_achievements_snapshot_vehicle_id_reference_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_vehicle_id_reference_id_created_at_idx` ON `vehicle_achievements_snapshot` (`vehicle_id`, `reference_id`, `created_at`);
