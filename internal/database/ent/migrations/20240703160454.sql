-- Create "sessions" table
CREATE TABLE `sessions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `expires_at` datetime NOT NULL, `public_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "sessions_public_id_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_public_id_key` ON `sessions` (`public_id`);
-- Create index "session_public_id_expires_at" to table: "sessions"
CREATE INDEX `session_public_id_expires_at` ON `sessions` (`public_id`, `expires_at`);
