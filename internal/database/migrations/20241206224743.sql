-- Create "manual_migration" table
CREATE TABLE `manual_migration` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `key` text NOT NULL, `finished` bool NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create index "manual_migration_id_idx" to table: "manual_migration"
CREATE INDEX `manual_migration_id_idx` ON `manual_migration` (`id`);
-- Create index "manual_migration_key_idx" to table: "manual_migration"
CREATE UNIQUE INDEX `manual_migration_key_idx` ON `manual_migration` (`key`);
-- Clear tables with recent timestamp changes
DELETE FROM `session`;
DELETE FROM `auth_nonce`;
DELETE FROM `discord_interaction`;