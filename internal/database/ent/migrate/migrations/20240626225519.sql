-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_user_contents" table
CREATE TABLE `new_user_contents` (`id` text NOT NULL, `created_at` integer NOT NULL, `updated_at` integer NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `value` text NOT NULL, `metadata` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_contents_users_content` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "user_contents" to new temporary table "new_user_contents"
INSERT INTO `new_user_contents` (`id`, `created_at`, `updated_at`, `type`, `reference_id`, `value`, `metadata`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `reference_id`, `value`, `metadata`, `user_id` FROM `user_contents`;
-- Drop "user_contents" table after copying rows
DROP TABLE `user_contents`;
-- Rename temporary table "new_user_contents" to "user_contents"
ALTER TABLE `new_user_contents` RENAME TO `user_contents`;
-- Create index "usercontent_id" to table: "user_contents"
CREATE INDEX `usercontent_id` ON `user_contents` (`id`);
-- Create index "usercontent_user_id" to table: "user_contents"
CREATE INDEX `usercontent_user_id` ON `user_contents` (`user_id`);
-- Create index "usercontent_type_user_id" to table: "user_contents"
CREATE INDEX `usercontent_type_user_id` ON `user_contents` (`type`, `user_id`);
-- Create index "usercontent_reference_id" to table: "user_contents"
CREATE INDEX `usercontent_reference_id` ON `user_contents` (`reference_id`);
-- Create index "usercontent_type_reference_id" to table: "user_contents"
CREATE INDEX `usercontent_type_reference_id` ON `user_contents` (`type`, `reference_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
