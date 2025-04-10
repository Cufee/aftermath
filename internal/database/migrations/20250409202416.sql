-- Create "discord_ad_run" table
CREATE TABLE "discord_ad_run" ("id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY, "created_at" text NOT NULL, "updated_at" text NOT NULL, "campaign_id" text NOT NULL, "content_id" text NOT NULL, "guild_id" text NOT NULL, "channel_id" text NOT NULL, "message_id" text NOT NULL, "locale" text NOT NULL, "tags" text NOT NULL, "metadata" bytea NOT NULL DEFAULT '\x', PRIMARY KEY ("id"));
-- Create index "discord_ad_run_campaign_id_idx" to table: "discord_ad_run"
CREATE INDEX "discord_ad_run_campaign_id_idx" ON "discord_ad_run" ("campaign_id");
-- Create index "discord_ad_run_channel_id_idx" to table: "discord_ad_run"
CREATE INDEX "discord_ad_run_channel_id_idx" ON "discord_ad_run" ("channel_id");
-- Create index "discord_ad_run_content_id_idx" to table: "discord_ad_run"
CREATE INDEX "discord_ad_run_content_id_idx" ON "discord_ad_run" ("content_id");
-- Create index "discord_ad_run_created_at_channel_id_idx" to table: "discord_ad_run"
CREATE INDEX "discord_ad_run_created_at_channel_id_idx" ON "discord_ad_run" ("created_at", "channel_id");
-- Create index "discord_ad_run_created_at_idx" to table: "discord_ad_run"
CREATE INDEX "discord_ad_run_created_at_idx" ON "discord_ad_run" ("created_at");
-- Create index "discord_ad_run_guild_id_idx" to table: "discord_ad_run"
CREATE INDEX "discord_ad_run_guild_id_idx" ON "discord_ad_run" ("guild_id");
-- Create index "discord_ad_run_message_id_idx" to table: "discord_ad_run"
CREATE INDEX "discord_ad_run_message_id_idx" ON "discord_ad_run" ("message_id");
