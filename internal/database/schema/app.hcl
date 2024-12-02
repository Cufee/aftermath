table "app_configuration" {
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
  column "key" {
    null = false
    type = text
  }
  column "value" {
    null = false
    type = json
  }
  column "metadata" {
    null = true
    type = json
  }
  primary_key {
    columns = [column.id]
  }
  index "app_configuration_id_idx" {
    columns = [column.id]
  }
  index "app_configuration_key_idx" {
    unique  = true
    columns = [column.key]
  }
}

table "cron_tasks" {
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
  column "type" {
    null = false
    type = text
  }
  column "reference_id" {
    null = false
    type = text
  }
  column "targets" {
    null = false
    type = json
  }
  column "status" {
    null = false
    type = text
  }
  column "scheduled_after" {
    null = false
    type = datetime
  }
  column "last_run" {
    null = false
    type = datetime
  }
  column "tries_left" {
    null    = false
    type    = integer
    default = 0
  }
  column "logs" {
    null = false
    type = json
  }
  column "data" {
    null = false
    type = json
  }
  primary_key {
    columns = [column.id]
  }
  index "cron_task_id_idx" {
    columns = [column.id]
  }
  index "cron_task_reference_id_idx" {
    columns = [column.reference_id]
  }
  index "cron_task_status_last_run_idx" {
    columns = [column.status, column.last_run]
  }
  index "cron_task_status_created_at_idx" {
    columns = [column.status, column.created_at]
  }
  index "cron_task_status_scheduled_after_idx" {
    columns = [column.status, column.scheduled_after]
  }
}

