-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_users" table
CREATE TABLE `new_users` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `username` text NOT NULL DEFAULT (''), `permissions` text NOT NULL DEFAULT (''), `feature_flags` json NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `created_at`, `updated_at`, `permissions`, `feature_flags`) SELECT `id`, `created_at`, `updated_at`, `permissions`, `feature_flags` FROM `users`;
-- Drop "users" table after copying rows
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "user_id" to table: "users"
CREATE INDEX `user_id` ON `users` (`id`);
-- Create index "user_username" to table: "users"
CREATE INDEX `user_username` ON `users` (`username`);
-- Create "auth_nonces" table
CREATE TABLE `auth_nonces` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `active` bool NOT NULL, `expires_at` datetime NOT NULL, `identifier` text NOT NULL, `public_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "auth_nonces_public_id_key" to table: "auth_nonces"
CREATE UNIQUE INDEX `auth_nonces_public_id_key` ON `auth_nonces` (`public_id`);
-- Create index "authnonce_public_id_active_expires_at" to table: "auth_nonces"
CREATE INDEX `authnonce_public_id_active_expires_at` ON `auth_nonces` (`public_id`, `active`, `expires_at`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
