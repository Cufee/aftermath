-- Create index "account_id" to table: "accounts"
CREATE INDEX `account_id` ON `accounts` (`id`);
-- Create index "accountsnapshot_id" to table: "account_snapshots"
CREATE INDEX `accountsnapshot_id` ON `account_snapshots` (`id`);
-- Create index "achievementssnapshot_id" to table: "achievements_snapshots"
CREATE INDEX `achievementssnapshot_id` ON `achievements_snapshots` (`id`);
-- Create index "appconfiguration_id" to table: "app_configurations"
CREATE INDEX `appconfiguration_id` ON `app_configurations` (`id`);
-- Create index "applicationcommand_id" to table: "application_commands"
CREATE INDEX `applicationcommand_id` ON `application_commands` (`id`);
-- Create index "clan_id" to table: "clans"
CREATE INDEX `clan_id` ON `clans` (`id`);
-- Create index "crontask_id" to table: "cron_tasks"
CREATE INDEX `crontask_id` ON `cron_tasks` (`id`);
-- Create index "user_id" to table: "users"
CREATE INDEX `user_id` ON `users` (`id`);
-- Create index "userconnection_id" to table: "user_connections"
CREATE INDEX `userconnection_id` ON `user_connections` (`id`);
-- Create index "userconnection_reference_id_user_id_type" to table: "user_connections"
CREATE UNIQUE INDEX `userconnection_reference_id_user_id_type` ON `user_connections` (`reference_id`, `user_id`, `type`);
-- Create index "usercontent_id" to table: "user_contents"
CREATE INDEX `usercontent_id` ON `user_contents` (`id`);
-- Create index "usersubscription_id" to table: "user_subscriptions"
CREATE INDEX `usersubscription_id` ON `user_subscriptions` (`id`);
-- Create index "vehicle_id" to table: "vehicles"
CREATE INDEX `vehicle_id` ON `vehicles` (`id`);
-- Create index "vehicleaverage_id" to table: "vehicle_averages"
CREATE INDEX `vehicleaverage_id` ON `vehicle_averages` (`id`);
-- Create index "vehiclesnapshot_id" to table: "vehicle_snapshots"
CREATE INDEX `vehiclesnapshot_id` ON `vehicle_snapshots` (`id`);
