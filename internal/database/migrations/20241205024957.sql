-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_account" table
CREATE TABLE `new_account` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `last_battle_time` datetime NOT NULL, `account_created_at` datetime NOT NULL, `realm` text NOT NULL, `nickname` text NOT NULL, `private` bool NOT NULL DEFAULT false, `clan_id` text NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "account" to new temporary table "new_account"
INSERT INTO `new_account` (`id`, `created_at`, `updated_at`, `last_battle_time`, `account_created_at`, `realm`, `nickname`, `private`, `clan_id`) SELECT `id`, `created_at`, `updated_at`, `last_battle_time`, `account_created_at`, `realm`, `nickname`, `private`, `clan_id` FROM `account`;
-- Drop "account" table after copying rows
DROP TABLE `account`;
-- Rename temporary table "new_account" to "account"
ALTER TABLE `new_account` RENAME TO `account`;
-- Create index "account_id_idx" to table: "account"
CREATE INDEX `account_id_idx` ON `account` (`id`);
-- Create index "account_id_last_battle_time_idx" to table: "account"
CREATE INDEX `account_id_last_battle_time_idx` ON `account` (`id`, `last_battle_time`);
-- Create index "account_realm_idx" to table: "account"
CREATE INDEX `account_realm_idx` ON `account` (`realm`);
-- Create index "account_realm_last_battle_time_idx" to table: "account"
CREATE INDEX `account_realm_last_battle_time_idx` ON `account` (`realm`, `last_battle_time`);
-- Create index "account_clan_id_idx" to table: "account"
CREATE INDEX `account_clan_id_idx` ON `account` (`clan_id`);
-- Create "new_account_snapshot" table
CREATE TABLE `new_account_snapshot` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `last_battle_time` datetime NOT NULL, `reference_id` text NOT NULL, `rating_battles` integer NOT NULL, `rating_frame` blob NOT NULL DEFAULT '', `regular_battles` integer NOT NULL, `regular_frame` blob NOT NULL DEFAULT '', `account_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `account_account_id_snapshot_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
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
