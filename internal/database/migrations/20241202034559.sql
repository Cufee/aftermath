-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Alter "app_configurations" table
ALTER TABLE `app_configurations` RENAME TO `app_configuration`;
-- Alter "application_commands" table
ALTER TABLE `application_commands` RENAME TO `application_command`;
-- Alter "auth_nonces" table
ALTER TABLE `auth_nonces` RENAME TO `auth_nonce`;
-- Alter "sessions" table
ALTER TABLE `sessions` RENAME TO `session`;
-- Alter "users" table
ALTER TABLE `users` RENAME TO `user`;
-- Alter "user_connections" table
ALTER TABLE `user_connections` RENAME TO `user_connection`;
-- Alter "user_subscriptions" table
ALTER TABLE `user_subscriptions` RENAME TO `user_subscription`;
-- Alter "vehicles" table
ALTER TABLE `vehicles` RENAME TO `vehicle`;
-- Alter "vehicle_averages" table
ALTER TABLE `vehicle_averages` RENAME TO `vehicle_average`;
-- Alter "game_maps" table
ALTER TABLE `game_maps` RENAME TO `game_map`;
-- Alter "game_modes" table
ALTER TABLE `game_modes` RENAME TO `game_mode`;
-- Alter "moderation_requests" table
ALTER TABLE `moderation_requests` RENAME TO `moderation_request`;
-- Alter "user_restrictions" table
ALTER TABLE `user_restrictions` RENAME TO `user_restriction`;
-- Alter "user_contents" table
ALTER TABLE `user_contents` RENAME TO `user_content`;
-- Alter "discord_interactions" table
ALTER TABLE `discord_interactions` RENAME TO `discord_interaction`;
-- Drop "leaderboard_scores" table
DROP TABLE `leaderboard_scores`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
