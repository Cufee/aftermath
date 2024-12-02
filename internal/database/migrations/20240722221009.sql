-- Create "user_restrictions" table
CREATE TABLE `user_restrictions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `expires_at` datetime NOT NULL, `type` text NOT NULL, `restriction` text NOT NULL, `public_reason` text NOT NULL, `moderator_comment` text NOT NULL, `events` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_restrictions_users_restrictions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "userrestriction_id" to table: "user_restrictions"
CREATE INDEX `userrestriction_id` ON `user_restrictions` (`id`);
-- Create index "userrestriction_user_id" to table: "user_restrictions"
CREATE INDEX `userrestriction_user_id` ON `user_restrictions` (`user_id`);
-- Create index "userrestriction_expires_at_user_id" to table: "user_restrictions"
CREATE INDEX `userrestriction_expires_at_user_id` ON `user_restrictions` (`expires_at`, `user_id`);
