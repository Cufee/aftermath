-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_user_subscription" table
CREATE TABLE `new_user_subscription` (`id` text NOT NULL, `created_at` text NOT NULL, `updated_at` text NOT NULL, `type` text NOT NULL, `expires_at` text NOT NULL, `permissions` text NOT NULL, `reference_id` text NOT NULL, `user_id` text NOT NULL, `metadata` blob NOT NULL DEFAULT '', PRIMARY KEY (`id`), CONSTRAINT `user_subscription_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "user_subscription" to new temporary table "new_user_subscription"
INSERT INTO `new_user_subscription` (`id`, `created_at`, `updated_at`, `type`, `expires_at`, `permissions`, `reference_id`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `type`, `expires_at`, `permissions`, `reference_id`, `user_id` FROM `user_subscription`;
-- Drop "user_subscription" table after copying rows
DROP TABLE `user_subscription`;
-- Rename temporary table "new_user_subscription" to "user_subscription"
ALTER TABLE `new_user_subscription` RENAME TO `user_subscription`;
-- Create index "user_subscription_id_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_id_idx` ON `user_subscription` (`id`);
-- Create index "user_subscription_user_id_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_user_id_idx` ON `user_subscription` (`user_id`);
-- Create index "user_subscription_type_user_id_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_type_user_id_idx` ON `user_subscription` (`type`, `user_id`);
-- Create index "user_subscription_expires_at_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_expires_at_idx` ON `user_subscription` (`expires_at`);
-- Create index "user_subscription_expires_at_user_id_idx" to table: "user_subscription"
CREATE INDEX `user_subscription_expires_at_user_id_idx` ON `user_subscription` (`expires_at`, `user_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
