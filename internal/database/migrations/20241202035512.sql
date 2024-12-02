-- Drop index "account_id" from table: "accounts"
DROP INDEX `account_id`;
-- Drop index "account_id_last_battle_time" from table: "accounts"
DROP INDEX `account_id_last_battle_time`;
-- Drop index "account_realm" from table: "accounts"
DROP INDEX `account_realm`;
-- Drop index "account_realm_last_battle_time" from table: "accounts"
DROP INDEX `account_realm_last_battle_time`;
-- Drop index "account_clan_id" from table: "accounts"
DROP INDEX `account_clan_id`;
-- Create index "account_id_idx" to table: "accounts"
CREATE INDEX `account_id_idx` ON `accounts` (`id`);
-- Create index "account_id_last_battle_time_idx" to table: "accounts"
CREATE INDEX `account_id_last_battle_time_idx` ON `accounts` (`id`, `last_battle_time`);
-- Create index "account_realm_idx" to table: "accounts"
CREATE INDEX `account_realm_idx` ON `accounts` (`realm`);
-- Create index "account_realm_last_battle_time_idx" to table: "accounts"
CREATE INDEX `account_realm_last_battle_time_idx` ON `accounts` (`realm`, `last_battle_time`);
-- Create index "account_clan_id_idx" to table: "accounts"
CREATE INDEX `account_clan_id_idx` ON `accounts` (`clan_id`);
-- Drop index "accountsnapshot_id" from table: "account_snapshots"
DROP INDEX `accountsnapshot_id`;
-- Drop index "accountsnapshot_created_at" from table: "account_snapshots"
DROP INDEX `accountsnapshot_created_at`;
-- Drop index "accountsnapshot_type_account_id_created_at" from table: "account_snapshots"
DROP INDEX `accountsnapshot_type_account_id_created_at`;
-- Drop index "accountsnapshot_type_account_id_reference_id" from table: "account_snapshots"
DROP INDEX `accountsnapshot_type_account_id_reference_id`;
-- Drop index "accountsnapshot_type_account_id_reference_id_created_at" from table: "account_snapshots"
DROP INDEX `accountsnapshot_type_account_id_reference_id_created_at`;
-- Create index "account_snapshot_id_idx" to table: "account_snapshots"
CREATE INDEX `account_snapshot_id_idx` ON `account_snapshots` (`id`);
-- Create index "account_snapshot_created_at_idx" to table: "account_snapshots"
CREATE INDEX `account_snapshot_created_at_idx` ON `account_snapshots` (`created_at`);
-- Create index "account_snapshot_type_account_id_created_at_idx" to table: "account_snapshots"
CREATE INDEX `account_snapshot_type_account_id_created_at_idx` ON `account_snapshots` (`type`, `account_id`, `created_at`);
-- Create index "account_snapshot_type_account_id_reference_id_idx" to table: "account_snapshots"
CREATE INDEX `account_snapshot_type_account_id_reference_id_idx` ON `account_snapshots` (`type`, `account_id`, `reference_id`);
-- Create index "account_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_snapshots"
CREATE INDEX `account_snapshot_type_account_id_reference_id_created_at_idx` ON `account_snapshots` (`type`, `account_id`, `reference_id`, `created_at`);
-- Drop index "app_configurations_key_key" from table: "app_configuration"
DROP INDEX `app_configurations_key_key`;
-- Drop index "appconfiguration_id" from table: "app_configuration"
DROP INDEX `appconfiguration_id`;
-- Drop index "appconfiguration_key" from table: "app_configuration"
DROP INDEX `appconfiguration_key`;
-- Create index "app_configuration_id_idx" to table: "app_configuration"
CREATE INDEX `app_configuration_id_idx` ON `app_configuration` (`id`);
-- Create index "app_configuration_key_idx" to table: "app_configuration"
CREATE UNIQUE INDEX `app_configuration_key_idx` ON `app_configuration` (`key`);
-- Drop index "applicationcommand_id" from table: "application_command"
DROP INDEX `applicationcommand_id`;
-- Drop index "applicationcommand_options_hash" from table: "application_command"
DROP INDEX `applicationcommand_options_hash`;
-- Create index "application_command_id_idx" to table: "application_command"
CREATE INDEX `application_command_id_idx` ON `application_command` (`id`);
-- Create index "application_command_options_hash_idx" to table: "application_command"
CREATE INDEX `application_command_options_hash_idx` ON `application_command` (`options_hash`);
-- Drop index "auth_nonces_public_id_key" from table: "auth_nonce"
DROP INDEX `auth_nonces_public_id_key`;
-- Drop index "authnonce_public_id_active_expires_at" from table: "auth_nonce"
DROP INDEX `authnonce_public_id_active_expires_at`;
-- Create index "auth_nonce_public_id_idx" to table: "auth_nonce"
CREATE UNIQUE INDEX `auth_nonce_public_id_idx` ON `auth_nonce` (`public_id`);
-- Create index "auth_nonce_public_id_active_expires_at_idx" to table: "auth_nonce"
CREATE INDEX `auth_nonce_public_id_active_expires_at_idx` ON `auth_nonce` (`public_id`, `active`, `expires_at`);
-- Drop index "clan_id" from table: "clans"
DROP INDEX `clan_id`;
-- Drop index "clan_tag" from table: "clans"
DROP INDEX `clan_tag`;
-- Drop index "clan_name" from table: "clans"
DROP INDEX `clan_name`;
-- Create index "clan_id_idx" to table: "clans"
CREATE INDEX `clan_id_idx` ON `clans` (`id`);
-- Create index "clan_tag_idx" to table: "clans"
CREATE INDEX `clan_tag_idx` ON `clans` (`tag`);
-- Create index "clan_name_idx" to table: "clans"
CREATE INDEX `clan_name_idx` ON `clans` (`name`);
-- Drop index "crontask_id" from table: "cron_tasks"
DROP INDEX `crontask_id`;
-- Drop index "crontask_reference_id" from table: "cron_tasks"
DROP INDEX `crontask_reference_id`;
-- Drop index "crontask_status_last_run" from table: "cron_tasks"
DROP INDEX `crontask_status_last_run`;
-- Drop index "crontask_status_created_at" from table: "cron_tasks"
DROP INDEX `crontask_status_created_at`;
-- Drop index "crontask_status_scheduled_after" from table: "cron_tasks"
DROP INDEX `crontask_status_scheduled_after`;
-- Create index "cron_task_id_idx" to table: "cron_tasks"
CREATE INDEX `cron_task_id_idx` ON `cron_tasks` (`id`);
-- Create index "cron_task_reference_id_idx" to table: "cron_tasks"
CREATE INDEX `cron_task_reference_id_idx` ON `cron_tasks` (`reference_id`);
-- Create index "cron_task_status_last_run_idx" to table: "cron_tasks"
CREATE INDEX `cron_task_status_last_run_idx` ON `cron_tasks` (`status`, `last_run`);
-- Create index "cron_task_status_created_at_idx" to table: "cron_tasks"
CREATE INDEX `cron_task_status_created_at_idx` ON `cron_tasks` (`status`, `created_at`);
-- Create index "cron_task_status_scheduled_after_idx" to table: "cron_tasks"
CREATE INDEX `cron_task_status_scheduled_after_idx` ON `cron_tasks` (`status`, `scheduled_after`);
-- Drop index "sessions_public_id_key" from table: "session"
DROP INDEX `sessions_public_id_key`;
-- Drop index "session_public_id_expires_at" from table: "session"
DROP INDEX `session_public_id_expires_at`;
-- Create index "session_public_id_idx" to table: "session"
CREATE UNIQUE INDEX `session_public_id_idx` ON `session` (`public_id`);
-- Create index "session_public_id_expires_at_idx" to table: "session"
CREATE INDEX `session_public_id_expires_at_idx` ON `session` (`public_id`, `expires_at`);
-- Drop index "user_id" from table: "user"
DROP INDEX `user_id`;
-- Drop index "user_username" from table: "user"
DROP INDEX `user_username`;
-- Create index "user_id_idx" to table: "user"
CREATE INDEX `user_id_idx` ON `user` (`id`);
-- Create index "user_username_idx" to table: "user"
CREATE INDEX `user_username_idx` ON `user` (`username`);
-- Drop index "userconnection_id" from table: "user_connection"
DROP INDEX `userconnection_id`;
-- Drop index "userconnection_user_id" from table: "user_connection"
DROP INDEX `userconnection_user_id`;
-- Drop index "userconnection_type_user_id" from table: "user_connection"
DROP INDEX `userconnection_type_user_id`;
-- Drop index "userconnection_reference_id" from table: "user_connection"
DROP INDEX `userconnection_reference_id`;
-- Drop index "userconnection_type_reference_id" from table: "user_connection"
DROP INDEX `userconnection_type_reference_id`;
-- Drop index "userconnection_reference_id_user_id_type" from table: "user_connection"
DROP INDEX `userconnection_reference_id_user_id_type`;
-- Create index "user_connection_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_id_idx` ON `user_connection` (`id`);
-- Create index "user_connection_user_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_user_id_idx` ON `user_connection` (`user_id`);
-- Create index "user_connection_type_user_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_type_user_id_idx` ON `user_connection` (`type`, `user_id`);
-- Create index "user_connection_reference_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_reference_id_idx` ON `user_connection` (`reference_id`);
-- Create index "user_connection_type_reference_id_idx" to table: "user_connection"
CREATE INDEX `user_connection_type_reference_id_idx` ON `user_connection` (`type`, `reference_id`);
-- Create index "user_connection_reference_id_user_id_type_idx" to table: "user_connection"
CREATE UNIQUE INDEX `user_connection_reference_id_user_id_type_idx` ON `user_connection` (`reference_id`, `user_id`, `type`);
-- Drop index "usersubscription_id" from table: "user_subscription"
DROP INDEX `usersubscription_id`;
-- Drop index "usersubscription_user_id" from table: "user_subscription"
DROP INDEX `usersubscription_user_id`;
-- Drop index "usersubscription_type_user_id" from table: "user_subscription"
DROP INDEX `usersubscription_type_user_id`;
-- Drop index "usersubscription_expires_at" from table: "user_subscription"
DROP INDEX `usersubscription_expires_at`;
-- Drop index "usersubscription_expires_at_user_id" from table: "user_subscription"
DROP INDEX `usersubscription_expires_at_user_id`;
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
-- Drop index "vehicle_id" from table: "vehicle"
DROP INDEX `vehicle_id`;
-- Create index "vehicle_id_idx" to table: "vehicle"
CREATE INDEX `vehicle_id_idx` ON `vehicle` (`id`);
-- Drop index "vehicleaverage_id" from table: "vehicle_average"
DROP INDEX `vehicleaverage_id`;
-- Create index "vehicle_average_id_idx" to table: "vehicle_average"
CREATE INDEX `vehicle_average_id_idx` ON `vehicle_average` (`id`);
-- Drop index "vehiclesnapshot_id" from table: "vehicle_snapshots"
DROP INDEX `vehiclesnapshot_id`;
-- Drop index "vehiclesnapshot_created_at" from table: "vehicle_snapshots"
DROP INDEX `vehiclesnapshot_created_at`;
-- Drop index "vehiclesnapshot_vehicle_id_created_at" from table: "vehicle_snapshots"
DROP INDEX `vehiclesnapshot_vehicle_id_created_at`;
-- Drop index "vehiclesnapshot_account_id_created_at" from table: "vehicle_snapshots"
DROP INDEX `vehiclesnapshot_account_id_created_at`;
-- Drop index "vehiclesnapshot_account_id_type_created_at" from table: "vehicle_snapshots"
DROP INDEX `vehiclesnapshot_account_id_type_created_at`;
-- Create index "vehicle_snapshot_id_idx" to table: "vehicle_snapshots"
CREATE INDEX `vehicle_snapshot_id_idx` ON `vehicle_snapshots` (`id`);
-- Create index "vehicle_snapshot_created_at_idx" to table: "vehicle_snapshots"
CREATE INDEX `vehicle_snapshot_created_at_idx` ON `vehicle_snapshots` (`created_at`);
-- Create index "vehicle_snapshot_vehicle_id_created_at_idx" to table: "vehicle_snapshots"
CREATE INDEX `vehicle_snapshot_vehicle_id_created_at_idx` ON `vehicle_snapshots` (`vehicle_id`, `created_at`);
-- Create index "vehicle)snapshot_account_id_created_at_idx" to table: "vehicle_snapshots"
CREATE INDEX `vehicle)snapshot_account_id_created_at_idx` ON `vehicle_snapshots` (`account_id`, `created_at`);
-- Create index "vehicle_snapshot_account_id_type_created_at_idx" to table: "vehicle_snapshots"
CREATE INDEX `vehicle_snapshot_account_id_type_created_at_idx` ON `vehicle_snapshots` (`account_id`, `type`, `created_at`);
-- Drop index "gamemap_id" from table: "game_map"
DROP INDEX `gamemap_id`;
-- Create index "game_map_id_idx" to table: "game_map"
CREATE INDEX `game_map_id_idx` ON `game_map` (`id`);
-- Drop index "widgetsettings_id" from table: "widget_settings"
DROP INDEX `widgetsettings_id`;
-- Create index "widget_settings_id" to table: "widget_settings"
CREATE INDEX `widget_settings_id` ON `widget_settings` (`id`);
-- Drop index "gamemode_id" from table: "game_mode"
DROP INDEX `gamemode_id`;
-- Create index "game_mode_id_idx" to table: "game_mode"
CREATE INDEX `game_mode_id_idx` ON `game_mode` (`id`);
-- Drop index "moderationrequest_id" from table: "moderation_request"
DROP INDEX `moderationrequest_id`;
-- Drop index "moderationrequest_reference_id" from table: "moderation_request"
DROP INDEX `moderationrequest_reference_id`;
-- Drop index "moderationrequest_requestor_id" from table: "moderation_request"
DROP INDEX `moderationrequest_requestor_id`;
-- Drop index "moderationrequest_moderator_id" from table: "moderation_request"
DROP INDEX `moderationrequest_moderator_id`;
-- Drop index "moderationrequest_requestor_id_reference_id" from table: "moderation_request"
DROP INDEX `moderationrequest_requestor_id_reference_id`;
-- Drop index "moderationrequest_requestor_id_reference_id_action_status" from table: "moderation_request"
DROP INDEX `moderationrequest_requestor_id_reference_id_action_status`;
-- Create index "moderation_request_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_id_idx` ON `moderation_request` (`id`);
-- Create index "moderation_request_reference_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_reference_id_idx` ON `moderation_request` (`reference_id`);
-- Create index "moderation_request_requestor_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_requestor_id_idx` ON `moderation_request` (`requestor_id`);
-- Create index "moderation_request_moderator_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_moderator_id_idx` ON `moderation_request` (`moderator_id`);
-- Create index "moderation_request_requestor_id_reference_id_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_requestor_id_reference_id_idx` ON `moderation_request` (`requestor_id`, `reference_id`);
-- Create index "moderation_request_requestor_id_reference_id_action_status_idx" to table: "moderation_request"
CREATE INDEX `moderation_request_requestor_id_reference_id_action_status_idx` ON `moderation_request` (`requestor_id`, `reference_id`, `action_status`);
-- Drop index "userrestriction_id" from table: "user_restriction"
DROP INDEX `userrestriction_id`;
-- Drop index "userrestriction_user_id" from table: "user_restriction"
DROP INDEX `userrestriction_user_id`;
-- Drop index "userrestriction_expires_at_user_id" from table: "user_restriction"
DROP INDEX `userrestriction_expires_at_user_id`;
-- Create index "user_restriction_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_id_idx` ON `user_restriction` (`id`);
-- Create index "user_restriction_user_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_user_id_idx` ON `user_restriction` (`user_id`);
-- Create index "user_restriction_expires_at_user_id_idx" to table: "user_restriction"
CREATE INDEX `user_restriction_expires_at_user_id_idx` ON `user_restriction` (`expires_at`, `user_id`);
-- Drop index "usercontent_id" from table: "user_content"
DROP INDEX `usercontent_id`;
-- Drop index "usercontent_user_id" from table: "user_content"
DROP INDEX `usercontent_user_id`;
-- Drop index "usercontent_reference_id" from table: "user_content"
DROP INDEX `usercontent_reference_id`;
-- Drop index "usercontent_type_user_id" from table: "user_content"
DROP INDEX `usercontent_type_user_id`;
-- Drop index "usercontent_type_reference_id" from table: "user_content"
DROP INDEX `usercontent_type_reference_id`;
-- Create index "user_content_id_idx" to table: "user_content"
CREATE INDEX `user_content_id_idx` ON `user_content` (`id`);
-- Create index "user_content_user_id_idx" to table: "user_content"
CREATE INDEX `user_content_user_id_idx` ON `user_content` (`user_id`);
-- Create index "user_content_reference_id_idx" to table: "user_content"
CREATE INDEX `user_content_reference_id_idx` ON `user_content` (`reference_id`);
-- Create index "user_content_type_user_id_idx" to table: "user_content"
CREATE UNIQUE INDEX `user_content_type_user_id_idx` ON `user_content` (`type`, `user_id`);
-- Create index "user_content_type_reference_id_idx" to table: "user_content"
CREATE UNIQUE INDEX `user_content_type_reference_id_idx` ON `user_content` (`type`, `reference_id`);
-- Drop index "discordinteraction_id" from table: "discord_interaction"
DROP INDEX `discordinteraction_id`;
-- Drop index "discordinteraction_user_id" from table: "discord_interaction"
DROP INDEX `discordinteraction_user_id`;
-- Drop index "discordinteraction_snowflake" from table: "discord_interaction"
DROP INDEX `discordinteraction_snowflake`;
-- Drop index "discordinteraction_created_at" from table: "discord_interaction"
DROP INDEX `discordinteraction_created_at`;
-- Drop index "discordinteraction_user_id_type_created_at" from table: "discord_interaction"
DROP INDEX `discordinteraction_user_id_type_created_at`;
-- Drop index "discordinteraction_channel_id_type_created_at" from table: "discord_interaction"
DROP INDEX `discordinteraction_channel_id_type_created_at`;
-- Create index "discord_interaction_id_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_id_idx` ON `discord_interaction` (`id`);
-- Create index "discord_interaction_user_id_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_user_id_idx` ON `discord_interaction` (`user_id`);
-- Create index "discord_interaction_snowflake_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_snowflake_idx` ON `discord_interaction` (`snowflake`);
-- Create index "discord_interaction_created_at_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_created_at_idx` ON `discord_interaction` (`created_at`);
-- Create index "discord_interaction_user_id_type_created_at_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_user_id_type_created_at_idx` ON `discord_interaction` (`user_id`, `type`, `created_at`);
-- Create index "discord_interaction_channel_id_type_created_at_idx" to table: "discord_interaction"
CREATE INDEX `discord_interaction_channel_id_type_created_at_idx` ON `discord_interaction` (`channel_id`, `type`, `created_at`);
