-- create "manual_migration" table
CREATE TABLE IF NOT EXISTS "manual_migration" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "key" text NOT NULL,
  "finished" boolean NOT NULL,
  "metadata" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "manual_migration_id_idx" to table: "manual_migration"
CREATE INDEX IF NOT EXISTS "manual_migration_id_idx" ON "manual_migration" ("id");
-- create index "manual_migration_key_idx" to table: "manual_migration"
CREATE UNIQUE INDEX IF NOT EXISTS "manual_migration_key_idx" ON "manual_migration" ("key");
-- create "game_mode" table
CREATE TABLE IF NOT EXISTS "game_mode" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "localized_names" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "game_mode_id_idx" to table: "game_mode"
CREATE INDEX IF NOT EXISTS "game_mode_id_idx" ON "game_mode" ("id");
-- create "vehicle" table
CREATE TABLE IF NOT EXISTS "vehicle" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "tier" integer NOT NULL,
  "localized_names" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "vehicle_id_idx" to table: "vehicle"
CREATE INDEX IF NOT EXISTS "vehicle_id_idx" ON "vehicle" ("id");
-- create "app_configuration" table
CREATE TABLE IF NOT EXISTS "app_configuration" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "key" text NOT NULL,
  "value" bytea NOT NULL DEFAULT '\x',
  "metadata" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "app_configuration_id_idx" to table: "app_configuration"
CREATE INDEX IF NOT EXISTS "app_configuration_id_idx" ON "app_configuration" ("id");
-- create index "app_configuration_key_idx" to table: "app_configuration"
CREATE UNIQUE INDEX IF NOT EXISTS "app_configuration_key_idx" ON "app_configuration" ("key");
-- create "application_command" table
CREATE TABLE IF NOT EXISTS "application_command" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "name" text NOT NULL,
  "version" text NOT NULL,
  "options_hash" text NOT NULL,
  PRIMARY KEY ("id")
);
-- create index "application_command_id_idx" to table: "application_command"
CREATE INDEX IF NOT EXISTS "application_command_id_idx" ON "application_command" ("id");
-- create index "application_command_options_hash_idx" to table: "application_command"
CREATE INDEX IF NOT EXISTS "application_command_options_hash_idx" ON "application_command" ("options_hash");
-- create "auth_nonce" table
CREATE TABLE IF NOT EXISTS "auth_nonce" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "active" boolean NOT NULL,
  "expires_at" text NOT NULL,
  "identifier" text NOT NULL,
  "public_id" text NOT NULL,
  "metadata" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "auth_nonce_public_id_active_expires_at_idx" to table: "auth_nonce"
CREATE INDEX IF NOT EXISTS "auth_nonce_public_id_active_expires_at_idx" ON "auth_nonce" ("public_id", "active", "expires_at");
-- create index "auth_nonce_public_id_idx" to table: "auth_nonce"
CREATE UNIQUE INDEX IF NOT EXISTS "auth_nonce_public_id_idx" ON "auth_nonce" ("public_id");
-- create "clan" table
CREATE TABLE IF NOT EXISTS "clan" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "tag" text NOT NULL,
  "name" text NOT NULL,
  "emblem_id" text NULL DEFAULT '',
  "members" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "clan_id_idx" to table: "clan"
CREATE INDEX IF NOT EXISTS "clan_id_idx" ON "clan" ("id");
-- create index "clan_name_idx" to table: "clan"
CREATE INDEX IF NOT EXISTS "clan_name_idx" ON "clan" ("name");
-- create index "clan_tag_idx" to table: "clan"
CREATE INDEX IF NOT EXISTS "clan_tag_idx" ON "clan" ("tag");
-- create "cron_task" table
CREATE TABLE IF NOT EXISTS "cron_task" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "type" text NOT NULL,
  "reference_id" text NOT NULL,
  "targets" bytea NOT NULL DEFAULT '\x',
  "status" text NOT NULL,
  "scheduled_after" text NOT NULL,
  "last_run" text NOT NULL,
  "tries_left" integer NOT NULL DEFAULT 0,
  "logs" bytea NOT NULL DEFAULT '\x',
  "data" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "cron_task_id_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS "cron_task_id_idx" ON "cron_task" ("id");
-- create index "cron_task_reference_id_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS "cron_task_reference_id_idx" ON "cron_task" ("reference_id");
-- create index "cron_task_status_created_at_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS "cron_task_status_created_at_idx" ON "cron_task" ("status", "created_at");
-- create index "cron_task_status_last_run_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS "cron_task_status_last_run_idx" ON "cron_task" ("status", "last_run");
-- create index "cron_task_status_scheduled_after_idx" to table: "cron_task"
CREATE INDEX IF NOT EXISTS "cron_task_status_scheduled_after_idx" ON "cron_task" ("status", "scheduled_after");
-- create "discord_ad_run" table
CREATE TABLE IF NOT EXISTS "discord_ad_run" (
  "id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "campaign_id" text NOT NULL,
  "content_id" text NOT NULL,
  "guild_id" text NOT NULL,
  "channel_id" text NOT NULL,
  "message_id" text NULL,
  "locale" text NOT NULL,
  "tags" text NOT NULL,
  "metadata" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "discord_ad_run_campaign_id_idx" to table: "discord_ad_run"
CREATE INDEX IF NOT EXISTS "discord_ad_run_campaign_id_idx" ON "discord_ad_run" ("campaign_id");
-- create index "discord_ad_run_channel_id_idx" to table: "discord_ad_run"
CREATE INDEX IF NOT EXISTS "discord_ad_run_channel_id_idx" ON "discord_ad_run" ("channel_id");
-- create index "discord_ad_run_content_id_idx" to table: "discord_ad_run"
CREATE INDEX IF NOT EXISTS "discord_ad_run_content_id_idx" ON "discord_ad_run" ("content_id");
-- create index "discord_ad_run_created_at_channel_id_idx" to table: "discord_ad_run"
CREATE INDEX IF NOT EXISTS "discord_ad_run_created_at_channel_id_idx" ON "discord_ad_run" ("created_at", "channel_id");
-- create index "discord_ad_run_created_at_idx" to table: "discord_ad_run"
CREATE INDEX IF NOT EXISTS "discord_ad_run_created_at_idx" ON "discord_ad_run" ("created_at");
-- create index "discord_ad_run_guild_id_idx" to table: "discord_ad_run"
CREATE INDEX IF NOT EXISTS "discord_ad_run_guild_id_idx" ON "discord_ad_run" ("guild_id");
-- create index "discord_ad_run_message_id_idx" to table: "discord_ad_run"
CREATE INDEX IF NOT EXISTS "discord_ad_run_message_id_idx" ON "discord_ad_run" ("message_id");
-- create "account" table
CREATE TABLE IF NOT EXISTS "account" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "last_battle_time" text NOT NULL,
  "account_created_at" text NOT NULL,
  "realm" text NOT NULL,
  "nickname" text NOT NULL,
  "private" boolean NOT NULL DEFAULT false,
  "clan_id" text NULL,
  PRIMARY KEY ("id")
);
-- create index "account_clan_id_idx" to table: "account"
CREATE INDEX IF NOT EXISTS "account_clan_id_idx" ON "account" ("clan_id");
-- create index "account_id_idx" to table: "account"
CREATE INDEX IF NOT EXISTS "account_id_idx" ON "account" ("id");
-- create index "account_id_last_battle_time_idx" to table: "account"
CREATE INDEX IF NOT EXISTS "account_id_last_battle_time_idx" ON "account" ("id", "last_battle_time");
-- create index "account_realm_idx" to table: "account"
CREATE INDEX IF NOT EXISTS "account_realm_idx" ON "account" ("realm");
-- create index "account_realm_last_battle_time_idx" to table: "account"
CREATE INDEX IF NOT EXISTS "account_realm_last_battle_time_idx" ON "account" ("realm", "last_battle_time");
-- create "game_map" table
CREATE TABLE IF NOT EXISTS "game_map" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "game_modes" bytea NOT NULL DEFAULT '\x',
  "supremacy_points" integer NOT NULL,
  "localized_names" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "game_map_id_idx" to table: "game_map"
CREATE INDEX IF NOT EXISTS "game_map_id_idx" ON "game_map" ("id");
-- create "vehicle_average" table
CREATE TABLE IF NOT EXISTS "vehicle_average" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "data" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id")
);
-- create index "vehicle_average_id_idx" to table: "vehicle_average"
CREATE INDEX IF NOT EXISTS "vehicle_average_id_idx" ON "vehicle_average" ("id");
-- create "account_achievements_snapshot" table
CREATE TABLE IF NOT EXISTS "account_achievements_snapshot" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "type" text NOT NULL,
  "battles" integer NOT NULL,
  "last_battle_time" text NOT NULL,
  "reference_id" text NOT NULL,
  "frame" bytea NOT NULL DEFAULT '\x',
  "account_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "account_id_account_achievements_snapshot_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "account" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "account_achievements_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "account_achievements_created_at_idx" ON "account_achievements_snapshot" ("created_at");
-- create index "account_achievements_snapshot_account_id_reference_id_created_a" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_account_id_reference_id_created_a" ON "account_achievements_snapshot" ("account_id", "reference_id", "created_at");
-- create index "account_achievements_snapshot_account_id_type_reference_id_crea" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_account_id_type_reference_id_crea" ON "account_achievements_snapshot" ("account_id", "type", "reference_id", "created_at");
-- create index "account_achievements_snapshot_id_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_id_idx" ON "account_achievements_snapshot" ("id");
-- create index "account_achievements_snapshot_type_account_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_type_account_id_created_at_idx" ON "account_achievements_snapshot" ("type", "account_id", "created_at");
-- create index "account_achievements_snapshot_type_account_id_reference_id_crea" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_type_account_id_reference_id_crea" ON "account_achievements_snapshot" ("type", "account_id", "reference_id", "created_at");
-- create index "account_achievements_snapshot_type_idx" to table: "account_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "account_achievements_snapshot_type_idx" ON "account_achievements_snapshot" ("type");
-- create "account_snapshot" table
CREATE TABLE IF NOT EXISTS "account_snapshot" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "type" text NOT NULL,
  "last_battle_time" text NOT NULL,
  "reference_id" text NOT NULL,
  "rating_battles" integer NOT NULL,
  "rating_frame" bytea NOT NULL DEFAULT '\x',
  "regular_battles" integer NOT NULL,
  "regular_frame" bytea NOT NULL DEFAULT '\x',
  "account_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "account_account_id_snapshot_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "account" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "account_snapshot_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS "account_snapshot_account_id_reference_id_created_at_idx" ON "account_snapshot" ("account_id", "reference_id", "created_at");
-- create index "account_snapshot_account_id_type_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS "account_snapshot_account_id_type_reference_id_created_at_idx" ON "account_snapshot" ("account_id", "type", "reference_id", "created_at");
-- create index "account_snapshot_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS "account_snapshot_created_at_idx" ON "account_snapshot" ("created_at");
-- create index "account_snapshot_id_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS "account_snapshot_id_idx" ON "account_snapshot" ("id");
-- create index "account_snapshot_type_account_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS "account_snapshot_type_account_id_created_at_idx" ON "account_snapshot" ("type", "account_id", "created_at");
-- create index "account_snapshot_type_account_id_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS "account_snapshot_type_account_id_reference_id_created_at_idx" ON "account_snapshot" ("type", "account_id", "reference_id", "created_at");
-- create index "account_snapshot_type_idx" to table: "account_snapshot"
CREATE INDEX IF NOT EXISTS "account_snapshot_type_idx" ON "account_snapshot" ("type");
-- create "user" table
CREATE TABLE IF NOT EXISTS "user" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "username" text NOT NULL DEFAULT '',
  "permissions" text NOT NULL DEFAULT '',
  "feature_flags" bytea NOT NULL DEFAULT '\x',
  "automod_verified" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
-- create index "user_id_idx" to table: "user"
CREATE INDEX IF NOT EXISTS "user_id_idx" ON "user" ("id");
-- create index "user_username_idx" to table: "user"
CREATE INDEX IF NOT EXISTS "user_username_idx" ON "user" ("username");
-- create "discord_interaction" table
CREATE TABLE IF NOT EXISTS "discord_interaction" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "result" text NOT NULL,
  "event_id" text NOT NULL,
  "guild_id" text NOT NULL,
  "snowflake" text NOT NULL DEFAULT '',
  "channel_id" text NOT NULL,
  "message_id" text NOT NULL,
  "type" text NOT NULL,
  "locale" text NOT NULL,
  "metadata" bytea NOT NULL DEFAULT '\x',
  "user_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "discord_interaction_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "discord_interaction_channel_id_type_created_at_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS "discord_interaction_channel_id_type_created_at_idx" ON "discord_interaction" ("channel_id", "type", "created_at");
-- create index "discord_interaction_created_at_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS "discord_interaction_created_at_idx" ON "discord_interaction" ("created_at");
-- create index "discord_interaction_id_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS "discord_interaction_id_idx" ON "discord_interaction" ("id");
-- create index "discord_interaction_snowflake_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS "discord_interaction_snowflake_idx" ON "discord_interaction" ("snowflake");
-- create index "discord_interaction_user_id_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS "discord_interaction_user_id_idx" ON "discord_interaction" ("user_id");
-- create index "discord_interaction_user_id_type_created_at_idx" to table: "discord_interaction"
CREATE INDEX IF NOT EXISTS "discord_interaction_user_id_type_created_at_idx" ON "discord_interaction" ("user_id", "type", "created_at");
-- create "moderation_request" table
CREATE TABLE IF NOT EXISTS "moderation_request" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "moderator_comment" text NULL,
  "context" text NULL,
  "reference_id" text NOT NULL,
  "action_reason" text NULL,
  "action_status" text NOT NULL,
  "data" bytea NOT NULL DEFAULT '\x',
  "requestor_id" text NOT NULL,
  "moderator_id" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "moderation_request_moderator_id_user_id_fk" FOREIGN KEY ("moderator_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "moderation_request_requestor_id_user_id_fk" FOREIGN KEY ("requestor_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "moderation_request_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS "moderation_request_id_idx" ON "moderation_request" ("id");
-- create index "moderation_request_moderator_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS "moderation_request_moderator_id_idx" ON "moderation_request" ("moderator_id");
-- create index "moderation_request_reference_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS "moderation_request_reference_id_idx" ON "moderation_request" ("reference_id");
-- create index "moderation_request_requestor_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS "moderation_request_requestor_id_idx" ON "moderation_request" ("requestor_id");
-- create index "moderation_request_requestor_id_reference_id_action_status_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS "moderation_request_requestor_id_reference_id_action_status_idx" ON "moderation_request" ("requestor_id", "reference_id", "action_status");
-- create index "moderation_request_requestor_id_reference_id_idx" to table: "moderation_request"
CREATE INDEX IF NOT EXISTS "moderation_request_requestor_id_reference_id_idx" ON "moderation_request" ("requestor_id", "reference_id");
-- create "session" table
CREATE TABLE IF NOT EXISTS "session" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "expires_at" text NOT NULL,
  "public_id" text NOT NULL,
  "metadata" bytea NOT NULL DEFAULT '\x',
  "user_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "session_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "session_public_id_expires_at_idx" to table: "session"
CREATE INDEX IF NOT EXISTS "session_public_id_expires_at_idx" ON "session" ("public_id", "expires_at");
-- create index "session_public_id_idx" to table: "session"
CREATE UNIQUE INDEX IF NOT EXISTS "session_public_id_idx" ON "session" ("public_id");
-- create "user_connection" table
CREATE TABLE IF NOT EXISTS "user_connection" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "type" text NOT NULL,
  "verified" boolean NOT NULL DEFAULT false,
  "selected" boolean NOT NULL DEFAULT false,
  "reference_id" text NOT NULL,
  "permissions" text NULL DEFAULT '',
  "metadata" bytea NOT NULL DEFAULT '\x',
  "user_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_connection_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "user_connection_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS "user_connection_id_idx" ON "user_connection" ("id");
-- create index "user_connection_reference_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS "user_connection_reference_id_idx" ON "user_connection" ("reference_id");
-- create index "user_connection_reference_id_user_id_type_idx" to table: "user_connection"
CREATE UNIQUE INDEX IF NOT EXISTS "user_connection_reference_id_user_id_type_idx" ON "user_connection" ("reference_id", "user_id", "type");
-- create index "user_connection_type_reference_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS "user_connection_type_reference_id_idx" ON "user_connection" ("type", "reference_id");
-- create index "user_connection_type_user_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS "user_connection_type_user_id_idx" ON "user_connection" ("type", "user_id");
-- create index "user_connection_user_id_idx" to table: "user_connection"
CREATE INDEX IF NOT EXISTS "user_connection_user_id_idx" ON "user_connection" ("user_id");
-- create "user_content" table
CREATE TABLE IF NOT EXISTS "user_content" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "type" text NOT NULL,
  "reference_id" text NOT NULL,
  "value" bytea NOT NULL,
  "metadata" bytea NOT NULL DEFAULT '\x',
  "user_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_content_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "user_content_id_idx" to table: "user_content"
CREATE INDEX IF NOT EXISTS "user_content_id_idx" ON "user_content" ("id");
-- create index "user_content_reference_id_idx" to table: "user_content"
CREATE INDEX IF NOT EXISTS "user_content_reference_id_idx" ON "user_content" ("reference_id");
-- create index "user_content_type_reference_id_idx" to table: "user_content"
CREATE UNIQUE INDEX IF NOT EXISTS "user_content_type_reference_id_idx" ON "user_content" ("type", "reference_id");
-- create index "user_content_type_user_id_idx" to table: "user_content"
CREATE UNIQUE INDEX IF NOT EXISTS "user_content_type_user_id_idx" ON "user_content" ("type", "user_id");
-- create index "user_content_user_id_idx" to table: "user_content"
CREATE INDEX IF NOT EXISTS "user_content_user_id_idx" ON "user_content" ("user_id");
-- create "user_restriction" table
CREATE TABLE IF NOT EXISTS "user_restriction" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "expires_at" text NOT NULL,
  "type" text NOT NULL,
  "restriction" text NOT NULL,
  "public_reason" text NOT NULL,
  "moderator_comment" text NOT NULL,
  "events" bytea NOT NULL DEFAULT '\x',
  "user_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_restriction_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "user_restriction_expires_at_user_id_idx" to table: "user_restriction"
CREATE INDEX IF NOT EXISTS "user_restriction_expires_at_user_id_idx" ON "user_restriction" ("expires_at", "user_id");
-- create index "user_restriction_id_idx" to table: "user_restriction"
CREATE INDEX IF NOT EXISTS "user_restriction_id_idx" ON "user_restriction" ("id");
-- create index "user_restriction_user_id_idx" to table: "user_restriction"
CREATE INDEX IF NOT EXISTS "user_restriction_user_id_idx" ON "user_restriction" ("user_id");
-- create "user_subscription" table
CREATE TABLE IF NOT EXISTS "user_subscription" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "type" text NOT NULL,
  "expires_at" text NOT NULL,
  "permissions" text NOT NULL,
  "reference_id" text NOT NULL,
  "user_id" text NOT NULL,
  "metadata" bytea NOT NULL DEFAULT '\x',
  PRIMARY KEY ("id"),
  CONSTRAINT "user_subscription_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "user_subscription_expires_at_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS "user_subscription_expires_at_idx" ON "user_subscription" ("expires_at");
-- create index "user_subscription_expires_at_user_id_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS "user_subscription_expires_at_user_id_idx" ON "user_subscription" ("expires_at", "user_id");
-- create index "user_subscription_id_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS "user_subscription_id_idx" ON "user_subscription" ("id");
-- create index "user_subscription_type_user_id_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS "user_subscription_type_user_id_idx" ON "user_subscription" ("type", "user_id");
-- create index "user_subscription_user_id_idx" to table: "user_subscription"
CREATE INDEX IF NOT EXISTS "user_subscription_user_id_idx" ON "user_subscription" ("user_id");
-- create "vehicle_achievements_snapshot" table
CREATE TABLE IF NOT EXISTS "vehicle_achievements_snapshot" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "type" text NOT NULL,
  "vehicle_id" text NOT NULL,
  "reference_id" text NOT NULL,
  "battles" integer NOT NULL,
  "last_battle_time" text NOT NULL,
  "frame" bytea NOT NULL DEFAULT '\x',
  "account_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "vehicle_achievements_snapshot_account_id_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "account" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "vehicle_achievements_snapshot_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_created_at_idx" ON "vehicle_achievements_snapshot" ("created_at");
-- create index "vehicle_achievements_snapshot_id_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_id_idx" ON "vehicle_achievements_snapshot" ("id");
-- create index "vehicle_achievements_snapshot_type_account_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_type_account_id_created_at_idx" ON "vehicle_achievements_snapshot" ("type", "account_id", "created_at");
-- create index "vehicle_achievements_snapshot_type_account_id_vehicle_id_create" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_type_account_id_vehicle_id_create" ON "vehicle_achievements_snapshot" ("type", "account_id", "vehicle_id", "created_at");
-- create index "vehicle_achievements_snapshot_type_account_id_vehicle_id_refere" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_type_account_id_vehicle_id_refere" ON "vehicle_achievements_snapshot" ("type", "account_id", "vehicle_id", "reference_id", "created_at");
-- create index "vehicle_achievements_snapshot_type_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_type_idx" ON "vehicle_achievements_snapshot" ("type");
-- create index "vehicle_achievements_snapshot_vehicle_id_reference_id_created_a" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_vehicle_id_reference_id_created_a" ON "vehicle_achievements_snapshot" ("vehicle_id", "reference_id", "created_at");
-- create index "vehicle_achievements_snapshot_vehicle_id_type_reference_id_crea" to table: "vehicle_achievements_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_achievements_snapshot_vehicle_id_type_reference_id_crea" ON "vehicle_achievements_snapshot" ("vehicle_id", "type", "reference_id", "created_at");
-- create "vehicle_snapshot" table
CREATE TABLE IF NOT EXISTS "vehicle_snapshot" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "type" text NOT NULL,
  "vehicle_id" text NOT NULL,
  "reference_id" text NOT NULL,
  "battles" integer NOT NULL,
  "last_battle_time" text NOT NULL,
  "frame" bytea NOT NULL DEFAULT '\x',
  "account_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "vehicle_snapshot_account_id_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "account" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "vehicle_snapshot_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_snapshot_created_at_idx" ON "vehicle_snapshot" ("created_at");
-- create index "vehicle_snapshot_id_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_snapshot_id_idx" ON "vehicle_snapshot" ("id");
-- create index "vehicle_snapshot_type_account_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_snapshot_type_account_id_created_at_idx" ON "vehicle_snapshot" ("type", "account_id", "created_at");
-- create index "vehicle_snapshot_type_account_id_vehicle_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_snapshot_type_account_id_vehicle_id_created_at_idx" ON "vehicle_snapshot" ("type", "account_id", "vehicle_id", "created_at");
-- create index "vehicle_snapshot_type_account_id_vehicle_id_reference_id_create" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_snapshot_type_account_id_vehicle_id_reference_id_create" ON "vehicle_snapshot" ("type", "account_id", "vehicle_id", "reference_id", "created_at");
-- create index "vehicle_snapshot_type_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_snapshot_type_idx" ON "vehicle_snapshot" ("type");
-- create index "vehicle_snapshot_vehicle_id_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_snapshot_vehicle_id_reference_id_created_at_idx" ON "vehicle_snapshot" ("vehicle_id", "reference_id", "created_at");
-- create index "vehicle_snapshot_vehicle_id_type_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX IF NOT EXISTS "vehicle_snapshot_vehicle_id_type_reference_id_created_at_idx" ON "vehicle_snapshot" ("vehicle_id", "type", "reference_id", "created_at");
-- create "widget_settings" table
CREATE TABLE IF NOT EXISTS "widget_settings" (
  "id" text NOT NULL,
  "created_at" text NOT NULL,
  "updated_at" text NOT NULL,
  "reference_id" text NOT NULL,
  "title" text NULL,
  "session_from" text NULL,
  "metadata" bytea NOT NULL DEFAULT '\x',
  "styles" bytea NOT NULL DEFAULT '\x',
  "user_id" text NOT NULL,
  "session_reference_id" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "widget_setting_user_id_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "widget_settings_id" to table: "widget_settings"
CREATE INDEX IF NOT EXISTS "widget_settings_id" ON "widget_settings" ("id");
