table "auth_nonce" {
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
  column "active" {
    null = false
    type = bool
  }
  column "expires_at" {
    null = false
    type = text
  }
  column "identifier" {
    null = false
    type = text
  }
  column "public_id" {
    null = false
    type = text
  }
  column "metadata" {
    null = false
    type = blob
    default = ""
  }
  primary_key {
    columns = [column.id]
  }
  index "auth_nonce_public_id_idx" {
    unique  = true
    columns = [column.public_id]
  }
  index "auth_nonce_public_id_active_expires_at_idx" {
    columns = [column.public_id, column.active, column.expires_at]
  }
}

table "session" {
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
  column "public_id" {
    null = false
    type = text
  }
  column "metadata" {
    null = false
    type = blob
    default = ""
  }
  column "user_id" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "session_user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.user.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "session_public_id_idx" {
    unique  = true
    columns = [column.public_id]
  }
  index "session_public_id_expires_at_idx" {
    columns = [column.public_id, column.expires_at]
  }
}