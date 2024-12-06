table "account" {
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
  column "last_battle_time" {
    null = false
    type = text
  }
  column "account_created_at" {
    null = false
    type = text
  }
  column "realm" {
    null = false
    type = text
  }
  column "nickname" {
    null = false
    type = text
  }
  column "private" {
    null    = false
    type    = bool
    default = false
  }
  column "clan_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "account_id_idx" {
    columns = [column.id]
  }
  index "account_id_last_battle_time_idx" {
    columns = [column.id, column.last_battle_time]
  }
  index "account_realm_idx" {
    columns = [column.realm]
  }
  index "account_realm_last_battle_time_idx" {
    columns = [column.realm, column.last_battle_time]
  }
  index "account_clan_id_idx" {
    columns = [column.clan_id]
  }
}

table "clan" {
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
  column "tag" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  column "emblem_id" {
    null    = true
    type    = text
    default = ""
  }
  column "members" {
    null = false
    type = blob
    default = ""
  }
  primary_key {
    columns = [column.id]
  }
  index "clan_id_idx" {
    columns = [column.id]
  }
  index "clan_tag_idx" {
    columns = [column.tag]
  }
  index "clan_name_idx" {
    columns = [column.name]
  }
}

table "account_snapshot" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = text
  }
  column "type" {
    null = false
    type = text
  }
  column "last_battle_time" {
    null = false
    type = text
  }
  column "reference_id" {
    null = false
    type = text
  }
  column "rating_battles" {
    null = false
    type = integer
  }
  column "rating_frame" {
    null = false
    type = blob
    default = ""
  }
  column "regular_battles" {
    null = false
    type = integer
  }
  column "regular_frame" {
    null = false
    type = blob
    default = ""
  }
  column "account_id" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "account_account_id_snapshot_account_id_fk" {
    columns     = [column.account_id]
    ref_columns = [table.account.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "account_snapshot_id_idx" {
    columns = [column.id]
  }
  index "account_snapshot_type_idx" {
    columns = [column.type]
  }
  index "account_snapshot_created_at_idx" {
    columns = [column.created_at]
  }
  index "account_snapshot_type_account_id_created_at_idx" {
    columns = [column.type, column.account_id, column.created_at]
  }
  index "account_snapshot_type_account_id_reference_id_created_at_idx" {
    columns = [column.type, column.account_id, column.reference_id, column.created_at]
  }
  index "account_snapshot_account_id_reference_id_created_at_idx" {
    columns = [column.account_id, column.reference_id, column.created_at]
  }
}

table "vehicle_snapshot" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = text
  }
  column "type" {
    null = false
    type = text
  }
  column "vehicle_id" {
    null = false
    type = text
  }
  column "reference_id" {
    null = false
    type = text
  }
  column "battles" {
    null = false
    type = integer
  }
  column "last_battle_time" {
    null = false
    type = text
  }
  column "frame" {
    null = false
    type = blob
    default = ""
  }
  column "account_id" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "vehicle_snapshot_account_id_account_id_fk" {
    columns     = [column.account_id]
    ref_columns = [table.account.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "vehicle_snapshot_id_idx" {
    columns = [column.id]
  }
  index "vehicle_snapshot_type_idx" {
    columns = [column.type]
  }
  index "vehicle_snapshot_created_at_idx" {
    columns = [column.created_at]
  }
  index "vehicle_snapshot_type_account_id_created_at_idx" {
    columns = [column.type, column.account_id, column.created_at]
  }
  index "vehicle_snapshot_type_account_id_vehicle_id_created_at_idx" {
    columns = [column.type, column.account_id, column.vehicle_id, column.created_at]
  }
  index "vehicle_snapshot_type_account_id_vehicle_id_reference_id_created_at_idx" {
    columns = [column.type, column.account_id, column.vehicle_id, column.reference_id, column.created_at]
  }
  index "vehicle_snapshot_vehicle_id_reference_id_created_at_idx" {
    columns = [column.vehicle_id, column.reference_id, column.created_at]
  }
}