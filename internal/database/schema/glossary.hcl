table "vehicle" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "tier" {
    null = false
    type = integer
  }
  column "localized_names" {
    null = false
    type = blob
    default = ""
  }
  primary_key {
    columns = [column.id]
  }
  index "vehicle_id_idx" {
    columns = [column.id]
  }
}

table "vehicle_average" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "data" {
    null = false
    type = blob
    default = ""
  }
  primary_key {
    columns = [column.id]
  }
  index "vehicle_average_id_idx" {
    columns = [column.id]
  }
}

table "game_map" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "game_modes" {
    null = false
    type = blob
    default = ""
  }
  column "supremacy_points" {
    null = false
    type = integer
  }
  column "localized_names" {
    null = false
    type = blob
    default = ""
  }
  primary_key {
    columns = [column.id]
  }
  index "game_map_id_idx" {
    columns = [column.id]
  }
}

table "game_mode" {
  schema = schema.main
  column "id" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "localized_names" {
    null = false
    type = blob
    default = ""
  }
  primary_key {
    columns = [column.id]
  }
  index "game_mode_id_idx" {
    columns = [column.id]
  }
}