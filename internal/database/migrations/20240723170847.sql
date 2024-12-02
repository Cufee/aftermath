-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "user_contents" table
DROP TABLE `user_contents`;
-- Create "new_user_contents" table
CREATE TABLE `user_contents` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `reference_id` text NOT NULL, `value` text NOT NULL, `metadata` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_contents_users_content` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "usercontent_id" to table: "user_contents"
CREATE INDEX `usercontent_id` ON `user_contents` (`id`);
-- Create index "usercontent_user_id" to table: "user_contents"
CREATE INDEX `usercontent_user_id` ON `user_contents` (`user_id`);
-- Create index "usercontent_reference_id" to table: "user_contents"
CREATE INDEX `usercontent_reference_id` ON `user_contents` (`reference_id`);
-- Create index "usercontent_type_user_id" to table: "user_contents"
CREATE UNIQUE INDEX `usercontent_type_user_id` ON `user_contents` (`type`, `user_id`);
-- Create index "usercontent_type_reference_id" to table: "user_contents"
CREATE UNIQUE INDEX `usercontent_type_reference_id` ON `user_contents` (`type`, `reference_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
