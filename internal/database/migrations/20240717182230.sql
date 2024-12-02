-- Create "widget_settings" table
CREATE TABLE `widget_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `reference_id` text NOT NULL, `title` text NULL, `session_from` datetime NULL, `metadata` json NOT NULL, `styles` json NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `widget_settings_users_widgets` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "widgetsettings_id" to table: "widget_settings"
CREATE INDEX `widgetsettings_id` ON `widget_settings` (`id`);
