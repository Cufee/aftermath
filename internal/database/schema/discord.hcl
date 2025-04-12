table "application_command" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = text
  }
  column "updated_at" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  column "version" {
    null = false
    type = text
  }
  column "options_hash" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "application_command_id_idx" {
    columns = [column.id]
  }
  index "application_command_options_hash_idx" {
    columns = [column.options_hash]
  }
}

table "discord_interaction" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = text
  }
  column "updated_at" {
    null = false
    type = text
  }
  column "result" {
    null = false
    type = text
  }
  column "event_id" {
    null = false
    type = text
  }
  column "guild_id" {
    null = false
    type = text
  }
  column "snowflake" {
    null    = false
    type    = text
    default = ""
  }
  column "channel_id" {
    null = false
    type = text
  }
  column "message_id" {
    null = false
    type = text
  }
  column "type" {
    null = false
    type = text
  }
  column "locale" {
    null = false
    type = text
  }
  column "metadata" {
    null = false
    type = bytea
    default = ""
  }
  column "user_id" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "discord_interaction_user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "discord_interaction_id_idx" {
    columns = [column.id]
  }
  index "discord_interaction_user_id_idx" {
    columns = [column.user_id]
  }
  index "discord_interaction_snowflake_idx" {
    columns = [column.snowflake]
  }
  index "discord_interaction_created_at_idx" {
    columns = [column.created_at]
  }
  index "discord_interaction_user_id_type_created_at_idx" {
    columns = [column.user_id, column.type, column.created_at]
  }
  index "discord_interaction_channel_id_type_created_at_idx" {
    columns = [column.channel_id, column.type, column.created_at]
  }
}

table "discord_ad_run" {
  schema = schema.main
  column "id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      increment = 1
      start = 0
    }
  }
  column "created_at" {
    null = false
    type = text
  }
  column "updated_at" {
    null = false
    type = text
  }

  column "campaign_id" {
      null = false
      type = text
  }
  column "content_id" {
      null = false
      type = text
  }

  column "guild_id" {
    null = false
    type = text
  }
  column "channel_id" {
    null = false
    type = text
  }
  column "message_id" {
    null = true
    type = text
  }

  column "locale" {
    null = false
    type = text
  }
  column "tags" {
      null = false
      type = text
  }
  column "metadata" {
    null = false
    type = bytea
    default = ""
  }

  primary_key {
    columns = [column.id]
  }

  index "discord_ad_run_created_at_idx" {
    columns = [column.created_at]
  }
  index "discord_ad_run_created_at_channel_id_idx" {
    columns = [column.created_at, column.channel_id]
  }
  index "discord_ad_run_campaign_id_idx" {
    columns = [column.campaign_id]
  }
  index "discord_ad_run_content_id_idx" {
    columns = [column.content_id]
  }
  index "discord_ad_run_guild_id_idx" {
    columns = [column.guild_id]
  }
  index "discord_ad_run_channel_id_idx" {
    columns = [column.channel_id]
  }
  index "discord_ad_run_message_id_idx" {
    columns = [column.message_id]
  }
}
