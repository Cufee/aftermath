table "user" {
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
  column "username" {
    null    = false
    type    = text
    default = ""
  }
  column "permissions" {
    null    = false
    type    = text
    default = ""
  }
  column "feature_flags" {
    null = false
    type = bytea
    default = ""
  }
  primary_key {
    columns = [column.id]
  }
  index "user_id_idx" {
    columns = [column.id]
  }
  index "user_username_idx" {
    columns = [column.username]
  }
}

table "user_connection" {
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
  column "type" {
    null = false
    type = text
  }
  column "verified" {
    null = false
    type = bool
    default = false
  }
  column "selected" {
    null = false
    type = bool
    default = false
  }
  column "reference_id" {
    null = false
    type = text
  }
  column "permissions" {
    null    = true
    type    = text
    default = ""
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
  foreign_key "user_connection_user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "user_connection_id_idx" {
    columns = [column.id]
  }
  index "user_connection_user_id_idx" {
    columns = [column.user_id]
  }
  index "user_connection_type_user_id_idx" {
    columns = [column.type, column.user_id]
  }
  index "user_connection_reference_id_idx" {
    columns = [column.reference_id]
  }
  index "user_connection_type_reference_id_idx" {
    columns = [column.type, column.reference_id]
  }
  index "user_connection_reference_id_user_id_type_idx" {
    unique  = true
    columns = [column.reference_id, column.user_id, column.type]
  }
}

table "user_subscription" {
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
  column "type" {
    null = false
    type = text
  }
  column "expires_at" {
    null = false
    type = text
  }
  column "permissions" {
    null = false
    type = text
  }
  column "reference_id" {
    null = false
    type = text
  }
  column "user_id" {
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
  foreign_key "user_subscription_user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "user_subscription_id_idx" {
    columns = [column.id]
  }
  index "user_subscription_user_id_idx" {
    columns = [column.user_id]
  }
  index "user_subscription_type_user_id_idx" {
    columns = [column.type, column.user_id]
  }
  index "user_subscription_expires_at_idx" {
    columns = [column.expires_at]
  }
  index "user_subscription_expires_at_user_id_idx" {
    columns = [column.expires_at, column.user_id]
  }
}

table "user_content" {
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
  column "type" {
    null = false
    type = text
  }
  column "reference_id" {
    null = false
    type = text
  }
  column "value" {
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
  foreign_key "user_content_user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "user_content_id_idx" {
    columns = [column.id]
  }
  index "user_content_user_id_idx" {
    columns = [column.user_id]
  }
  index "user_content_reference_id_idx" {
    columns = [column.reference_id]
  }
  index "user_content_type_user_id_idx" {
    unique  = true
    columns = [column.type, column.user_id]
  }
  index "user_content_type_reference_id_idx" {
    unique  = true
    columns = [column.type, column.reference_id]
  }
}


table "user_restriction" {
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
  column "expires_at" {
    null = false
    type = text
  }
  column "type" {
    null = false
    type = text
  }
  column "restriction" {
    null = false
    type = text
  }
  column "public_reason" {
    null = false
    type = text
  }
  column "moderator_comment" {
    null = false
    type = text
  }
  column "events" {
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
  foreign_key "user_restriction_user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "user_restriction_id_idx" {
    columns = [column.id]
  }
  index "user_restriction_user_id_idx" {
    columns = [column.user_id]
  }
  index "user_restriction_expires_at_user_id_idx" {
    columns = [column.expires_at, column.user_id]
  }
}

table "moderation_request" {
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
  column "moderator_comment" {
    null = true
    type = text
  }
  column "context" {
    null = true
    type = text
  }
  column "reference_id" {
    null = false
    type = text
  }
  column "action_reason" {
    null = true
    type = text
  }
  column "action_status" {
    null = false
    type = text
  }
  column "data" {
    null = false
    type = bytea
    default = ""
  }
  column "requestor_id" {
    null = false
    type = text
  }
  column "moderator_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "moderation_request_moderator_id_user_id_fk" {
    columns     = [column.moderator_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "moderation_request_requestor_id_user_id_fk" {
    columns     = [column.requestor_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "moderation_request_id_idx" {
    columns = [column.id]
  }
  index "moderation_request_reference_id_idx" {
    columns = [column.reference_id]
  }
  index "moderation_request_requestor_id_idx" {
    columns = [column.requestor_id]
  }
  index "moderation_request_moderator_id_idx" {
    columns = [column.moderator_id]
  }
  index "moderation_request_requestor_id_reference_id_idx" {
    columns = [column.requestor_id, column.reference_id]
  }
  index "moderation_request_requestor_id_reference_id_action_status_idx" {
    columns = [column.requestor_id, column.reference_id, column.action_status]
  }
}

table "widget_settings" {
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
  column "reference_id" {
    null = false
    type = text
  }
  column "title" {
    null = true
    type = text
  }
  column "session_from" {
    null = true
    type = text
  }
  column "metadata" {
    null = false
    type = bytea
    default = ""
  }
  column "styles" {
    null = false
    type = bytea
    default = ""
  }
  column "user_id" {
    null = false
    type = text
  }
  column "session_reference_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "widget_setting_user_id_user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "widget_settings_id" {
    columns = [column.id]
  }
}