-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_user_connection" table
CREATE TABLE `new_user_connection` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `verified` bool NOT NULL DEFAULT false, `selected` bool NOT NULL DEFAULT false, `reference_id` text NOT NULL, `permissions` text NULL DEFAULT '', `metadata` json NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_connection_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "user_connection" to new temporary table "new_user_connection"
INSERT INTO `new_user_connection` (`id`, `created_at`, `updated_at`, `type`, `verified`, `reference_id`, `permissions`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `verified`, `reference_id`, `permissions`, `metadata`, `user_id` FROM `user_connection`;
-- Drop "user_connection" table after copying rows
DROP TABLE `user_connection`;
-- Rename temporary table "new_user_connection" to "user_connection"
ALTER TABLE `new_user_connection` RENAME TO `user_connection`;
-- Create index "user_connection_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_id_idx` ON `user_connection` (`id`);
-- Create index "user_connection_user_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_user_id_idx` ON `user_connection` (`user_id`);
-- Create index "user_connection_type_user_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_type_user_id_idx` ON `user_connection` (`type`, `user_id`);
-- Create index "user_connection_reference_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_reference_id_idx` ON `user_connection` (`reference_id`);
-- Create index "user_connection_type_reference_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_type_reference_id_idx` ON `user_connection` (`type`, `reference_id`);
-- Create index "user_connection_reference_id_user_id_type_idx" to table: "user_connection"
CREATE UNIQUE INDEX `user_connection_reference_id_user_id_type_idx` ON `user_connection` (`reference_id`, `user_id`, `type`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
