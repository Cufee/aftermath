-- reverse: create index "widget_settings_id" to table: "widget_settings"
DROP INDEX "widget_settings_id";
-- reverse: create "widget_settings" table
DROP TABLE "widget_settings";
-- reverse: create index "vehicle_snapshot_vehicle_id_type_reference_id_created_at_idx" to table: "vehicle_snapshot"
DROP INDEX "vehicle_snapshot_vehicle_id_type_reference_id_created_at_idx";
-- reverse: create index "vehicle_snapshot_vehicle_id_reference_id_created_at_idx" to table: "vehicle_snapshot"
DROP INDEX "vehicle_snapshot_vehicle_id_reference_id_created_at_idx";
-- reverse: create index "vehicle_snapshot_type_idx" to table: "vehicle_snapshot"
DROP INDEX "vehicle_snapshot_type_idx";
-- reverse: create index "vehicle_snapshot_type_account_id_vehicle_id_reference_id_create" to table: "vehicle_snapshot"
DROP INDEX "vehicle_snapshot_type_account_id_vehicle_id_reference_id_create";
-- reverse: create index "vehicle_snapshot_type_account_id_vehicle_id_created_at_idx" to table: "vehicle_snapshot"
DROP INDEX "vehicle_snapshot_type_account_id_vehicle_id_created_at_idx";
-- reverse: create index "vehicle_snapshot_type_account_id_created_at_idx" to table: "vehicle_snapshot"
DROP INDEX "vehicle_snapshot_type_account_id_created_at_idx";
-- reverse: create index "vehicle_snapshot_id_idx" to table: "vehicle_snapshot"
DROP INDEX "vehicle_snapshot_id_idx";
-- reverse: create index "vehicle_snapshot_created_at_idx" to table: "vehicle_snapshot"
DROP INDEX "vehicle_snapshot_created_at_idx";
-- reverse: create "vehicle_snapshot" table
DROP TABLE "vehicle_snapshot";
-- reverse: create index "vehicle_achievements_snapshot_vehicle_id_type_reference_id_crea" to table: "vehicle_achievements_snapshot"
DROP INDEX "vehicle_achievements_snapshot_vehicle_id_type_reference_id_crea";
-- reverse: create index "vehicle_achievements_snapshot_vehicle_id_reference_id_created_a" to table: "vehicle_achievements_snapshot"
DROP INDEX "vehicle_achievements_snapshot_vehicle_id_reference_id_created_a";
-- reverse: create index "vehicle_achievements_snapshot_type_idx" to table: "vehicle_achievements_snapshot"
DROP INDEX "vehicle_achievements_snapshot_type_idx";
-- reverse: create index "vehicle_achievements_snapshot_type_account_id_vehicle_id_refere" to table: "vehicle_achievements_snapshot"
DROP INDEX "vehicle_achievements_snapshot_type_account_id_vehicle_id_refere";
-- reverse: create index "vehicle_achievements_snapshot_type_account_id_vehicle_id_create" to table: "vehicle_achievements_snapshot"
DROP INDEX "vehicle_achievements_snapshot_type_account_id_vehicle_id_create";
-- reverse: create index "vehicle_achievements_snapshot_type_account_id_created_at_idx" to table: "vehicle_achievements_snapshot"
DROP INDEX "vehicle_achievements_snapshot_type_account_id_created_at_idx";
-- reverse: create index "vehicle_achievements_snapshot_id_idx" to table: "vehicle_achievements_snapshot"
DROP INDEX "vehicle_achievements_snapshot_id_idx";
-- reverse: create index "vehicle_achievements_snapshot_created_at_idx" to table: "vehicle_achievements_snapshot"
DROP INDEX "vehicle_achievements_snapshot_created_at_idx";
-- reverse: create "vehicle_achievements_snapshot" table
DROP TABLE "vehicle_achievements_snapshot";
-- reverse: create index "user_subscription_user_id_idx" to table: "user_subscription"
DROP INDEX "user_subscription_user_id_idx";
-- reverse: create index "user_subscription_type_user_id_idx" to table: "user_subscription"
DROP INDEX "user_subscription_type_user_id_idx";
-- reverse: create index "user_subscription_id_idx" to table: "user_subscription"
DROP INDEX "user_subscription_id_idx";
-- reverse: create index "user_subscription_expires_at_user_id_idx" to table: "user_subscription"
DROP INDEX "user_subscription_expires_at_user_id_idx";
-- reverse: create index "user_subscription_expires_at_idx" to table: "user_subscription"
DROP INDEX "user_subscription_expires_at_idx";
-- reverse: create "user_subscription" table
DROP TABLE "user_subscription";
-- reverse: create index "user_restriction_user_id_idx" to table: "user_restriction"
DROP INDEX "user_restriction_user_id_idx";
-- reverse: create index "user_restriction_id_idx" to table: "user_restriction"
DROP INDEX "user_restriction_id_idx";
-- reverse: create index "user_restriction_expires_at_user_id_idx" to table: "user_restriction"
DROP INDEX "user_restriction_expires_at_user_id_idx";
-- reverse: create "user_restriction" table
DROP TABLE "user_restriction";
-- reverse: create index "user_content_user_id_idx" to table: "user_content"
DROP INDEX "user_content_user_id_idx";
-- reverse: create index "user_content_type_user_id_idx" to table: "user_content"
DROP INDEX "user_content_type_user_id_idx";
-- reverse: create index "user_content_type_reference_id_idx" to table: "user_content"
DROP INDEX "user_content_type_reference_id_idx";
-- reverse: create index "user_content_reference_id_idx" to table: "user_content"
DROP INDEX "user_content_reference_id_idx";
-- reverse: create index "user_content_id_idx" to table: "user_content"
DROP INDEX "user_content_id_idx";
-- reverse: create "user_content" table
DROP TABLE "user_content";
-- reverse: create index "user_connection_user_id_idx" to table: "user_connection"
DROP INDEX "user_connection_user_id_idx";
-- reverse: create index "user_connection_type_user_id_idx" to table: "user_connection"
DROP INDEX "user_connection_type_user_id_idx";
-- reverse: create index "user_connection_type_reference_id_idx" to table: "user_connection"
DROP INDEX "user_connection_type_reference_id_idx";
-- reverse: create index "user_connection_reference_id_user_id_type_idx" to table: "user_connection"
DROP INDEX "user_connection_reference_id_user_id_type_idx";
-- reverse: create index "user_connection_reference_id_idx" to table: "user_connection"
DROP INDEX "user_connection_reference_id_idx";
-- reverse: create index "user_connection_id_idx" to table: "user_connection"
DROP INDEX "user_connection_id_idx";
-- reverse: create "user_connection" table
DROP TABLE "user_connection";
-- reverse: create index "session_public_id_idx" to table: "session"
DROP INDEX "session_public_id_idx";
-- reverse: create index "session_public_id_expires_at_idx" to table: "session"
DROP INDEX "session_public_id_expires_at_idx";
-- reverse: create "session" table
DROP TABLE "session";
-- reverse: create index "moderation_request_requestor_id_reference_id_idx" to table: "moderation_request"
DROP INDEX "moderation_request_requestor_id_reference_id_idx";
-- reverse: create index "moderation_request_requestor_id_reference_id_action_status_idx" to table: "moderation_request"
DROP INDEX "moderation_request_requestor_id_reference_id_action_status_idx";
-- reverse: create index "moderation_request_requestor_id_idx" to table: "moderation_request"
DROP INDEX "moderation_request_requestor_id_idx";
-- reverse: create index "moderation_request_reference_id_idx" to table: "moderation_request"
DROP INDEX "moderation_request_reference_id_idx";
-- reverse: create index "moderation_request_moderator_id_idx" to table: "moderation_request"
DROP INDEX "moderation_request_moderator_id_idx";
-- reverse: create index "moderation_request_id_idx" to table: "moderation_request"
DROP INDEX "moderation_request_id_idx";
-- reverse: create "moderation_request" table
DROP TABLE "moderation_request";
-- reverse: create index "discord_interaction_user_id_type_created_at_idx" to table: "discord_interaction"
DROP INDEX "discord_interaction_user_id_type_created_at_idx";
-- reverse: create index "discord_interaction_user_id_idx" to table: "discord_interaction"
DROP INDEX "discord_interaction_user_id_idx";
-- reverse: create index "discord_interaction_snowflake_idx" to table: "discord_interaction"
DROP INDEX "discord_interaction_snowflake_idx";
-- reverse: create index "discord_interaction_id_idx" to table: "discord_interaction"
DROP INDEX "discord_interaction_id_idx";
-- reverse: create index "discord_interaction_created_at_idx" to table: "discord_interaction"
DROP INDEX "discord_interaction_created_at_idx";
-- reverse: create index "discord_interaction_channel_id_type_created_at_idx" to table: "discord_interaction"
DROP INDEX "discord_interaction_channel_id_type_created_at_idx";
-- reverse: create "discord_interaction" table
DROP TABLE "discord_interaction";
-- reverse: create index "user_username_idx" to table: "user"
DROP INDEX "user_username_idx";
-- reverse: create index "user_id_idx" to table: "user"
DROP INDEX "user_id_idx";
-- reverse: create "user" table
DROP TABLE "user";
-- reverse: create index "account_snapshot_type_idx" to table: "account_snapshot"
DROP INDEX "account_snapshot_type_idx";
-- reverse: create index "account_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_snapshot"
DROP INDEX "account_snapshot_type_account_id_reference_id_created_at_idx";
-- reverse: create index "account_snapshot_type_account_id_created_at_idx" to table: "account_snapshot"
DROP INDEX "account_snapshot_type_account_id_created_at_idx";
-- reverse: create index "account_snapshot_id_idx" to table: "account_snapshot"
DROP INDEX "account_snapshot_id_idx";
-- reverse: create index "account_snapshot_created_at_idx" to table: "account_snapshot"
DROP INDEX "account_snapshot_created_at_idx";
-- reverse: create index "account_snapshot_account_id_type_reference_id_created_at_idx" to table: "account_snapshot"
DROP INDEX "account_snapshot_account_id_type_reference_id_created_at_idx";
-- reverse: create index "account_snapshot_account_id_reference_id_created_at_idx" to table: "account_snapshot"
DROP INDEX "account_snapshot_account_id_reference_id_created_at_idx";
-- reverse: create "account_snapshot" table
DROP TABLE "account_snapshot";
-- reverse: create index "account_achievements_snapshot_type_idx" to table: "account_achievements_snapshot"
DROP INDEX "account_achievements_snapshot_type_idx";
-- reverse: create index "account_achievements_snapshot_type_account_id_reference_id_crea" to table: "account_achievements_snapshot"
DROP INDEX "account_achievements_snapshot_type_account_id_reference_id_crea";
-- reverse: create index "account_achievements_snapshot_type_account_id_created_at_idx" to table: "account_achievements_snapshot"
DROP INDEX "account_achievements_snapshot_type_account_id_created_at_idx";
-- reverse: create index "account_achievements_snapshot_id_idx" to table: "account_achievements_snapshot"
DROP INDEX "account_achievements_snapshot_id_idx";
-- reverse: create index "account_achievements_snapshot_account_id_type_reference_id_crea" to table: "account_achievements_snapshot"
DROP INDEX "account_achievements_snapshot_account_id_type_reference_id_crea";
-- reverse: create index "account_achievements_snapshot_account_id_reference_id_created_a" to table: "account_achievements_snapshot"
DROP INDEX "account_achievements_snapshot_account_id_reference_id_created_a";
-- reverse: create index "account_achievements_created_at_idx" to table: "account_achievements_snapshot"
DROP INDEX "account_achievements_created_at_idx";
-- reverse: create "account_achievements_snapshot" table
DROP TABLE "account_achievements_snapshot";
-- reverse: create index "vehicle_average_id_idx" to table: "vehicle_average"
DROP INDEX "vehicle_average_id_idx";
-- reverse: create "vehicle_average" table
DROP TABLE "vehicle_average";
-- reverse: create index "game_map_id_idx" to table: "game_map"
DROP INDEX "game_map_id_idx";
-- reverse: create "game_map" table
DROP TABLE "game_map";
-- reverse: create index "account_realm_last_battle_time_idx" to table: "account"
DROP INDEX "account_realm_last_battle_time_idx";
-- reverse: create index "account_realm_idx" to table: "account"
DROP INDEX "account_realm_idx";
-- reverse: create index "account_id_last_battle_time_idx" to table: "account"
DROP INDEX "account_id_last_battle_time_idx";
-- reverse: create index "account_id_idx" to table: "account"
DROP INDEX "account_id_idx";
-- reverse: create index "account_clan_id_idx" to table: "account"
DROP INDEX "account_clan_id_idx";
-- reverse: create "account" table
DROP TABLE "account";
-- reverse: create index "discord_ad_run_message_id_idx" to table: "discord_ad_run"
DROP INDEX "discord_ad_run_message_id_idx";
-- reverse: create index "discord_ad_run_guild_id_idx" to table: "discord_ad_run"
DROP INDEX "discord_ad_run_guild_id_idx";
-- reverse: create index "discord_ad_run_created_at_idx" to table: "discord_ad_run"
DROP INDEX "discord_ad_run_created_at_idx";
-- reverse: create index "discord_ad_run_created_at_channel_id_idx" to table: "discord_ad_run"
DROP INDEX "discord_ad_run_created_at_channel_id_idx";
-- reverse: create index "discord_ad_run_content_id_idx" to table: "discord_ad_run"
DROP INDEX "discord_ad_run_content_id_idx";
-- reverse: create index "discord_ad_run_channel_id_idx" to table: "discord_ad_run"
DROP INDEX "discord_ad_run_channel_id_idx";
-- reverse: create index "discord_ad_run_campaign_id_idx" to table: "discord_ad_run"
DROP INDEX "discord_ad_run_campaign_id_idx";
-- reverse: create "discord_ad_run" table
DROP TABLE "discord_ad_run";
-- reverse: create index "cron_task_status_scheduled_after_idx" to table: "cron_task"
DROP INDEX "cron_task_status_scheduled_after_idx";
-- reverse: create index "cron_task_status_last_run_idx" to table: "cron_task"
DROP INDEX "cron_task_status_last_run_idx";
-- reverse: create index "cron_task_status_created_at_idx" to table: "cron_task"
DROP INDEX "cron_task_status_created_at_idx";
-- reverse: create index "cron_task_reference_id_idx" to table: "cron_task"
DROP INDEX "cron_task_reference_id_idx";
-- reverse: create index "cron_task_id_idx" to table: "cron_task"
DROP INDEX "cron_task_id_idx";
-- reverse: create "cron_task" table
DROP TABLE "cron_task";
-- reverse: create index "clan_tag_idx" to table: "clan"
DROP INDEX "clan_tag_idx";
-- reverse: create index "clan_name_idx" to table: "clan"
DROP INDEX "clan_name_idx";
-- reverse: create index "clan_id_idx" to table: "clan"
DROP INDEX "clan_id_idx";
-- reverse: create "clan" table
DROP TABLE "clan";
-- reverse: create index "auth_nonce_public_id_idx" to table: "auth_nonce"
DROP INDEX "auth_nonce_public_id_idx";
-- reverse: create index "auth_nonce_public_id_active_expires_at_idx" to table: "auth_nonce"
DROP INDEX "auth_nonce_public_id_active_expires_at_idx";
-- reverse: create "auth_nonce" table
DROP TABLE "auth_nonce";
-- reverse: create index "application_command_options_hash_idx" to table: "application_command"
DROP INDEX "application_command_options_hash_idx";
-- reverse: create index "application_command_id_idx" to table: "application_command"
DROP INDEX "application_command_id_idx";
-- reverse: create "application_command" table
DROP TABLE "application_command";
-- reverse: create index "app_configuration_key_idx" to table: "app_configuration"
DROP INDEX "app_configuration_key_idx";
-- reverse: create index "app_configuration_id_idx" to table: "app_configuration"
DROP INDEX "app_configuration_id_idx";
-- reverse: create "app_configuration" table
DROP TABLE "app_configuration";
-- reverse: create index "vehicle_id_idx" to table: "vehicle"
DROP INDEX "vehicle_id_idx";
-- reverse: create "vehicle" table
DROP TABLE "vehicle";
-- reverse: create index "game_mode_id_idx" to table: "game_mode"
DROP INDEX "game_mode_id_idx";
-- reverse: create "game_mode" table
DROP TABLE "game_mode";
-- reverse: create index "manual_migration_key_idx" to table: "manual_migration"
DROP INDEX "manual_migration_key_idx";
-- reverse: create index "manual_migration_id_idx" to table: "manual_migration"
DROP INDEX "manual_migration_id_idx";
-- reverse: create "manual_migration" table
DROP TABLE "manual_migration";
