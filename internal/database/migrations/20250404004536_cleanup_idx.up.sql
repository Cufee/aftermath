-- create index "account_snapshot_account_id_type_reference_id_created_at_idx" to table: "account_snapshot"
CREATE INDEX `account_snapshot_account_id_type_reference_id_created_at_idx` ON `account_snapshot` (`account_id`, `type`, `reference_id`, `created_at`);
-- create index "account_achievements_snapshot_account_id_type_reference_id_created_at_idx" to table: "account_achievements_snapshot"
CREATE INDEX `account_achievements_snapshot_account_id_type_reference_id_created_at_idx` ON `account_achievements_snapshot` (`account_id`, `type`, `reference_id`, `created_at`);
-- create index "vehicle_snapshot_vehicle_id_type_reference_id_created_at_idx" to table: "vehicle_snapshot"
CREATE INDEX `vehicle_snapshot_vehicle_id_type_reference_id_created_at_idx` ON `vehicle_snapshot` (`vehicle_id`, `type`, `reference_id`, `created_at`);
-- create index "vehicle_achievements_snapshot_vehicle_id_type_reference_id_created_at_idx" to table: "vehicle_achievements_snapshot"
CREATE INDEX `vehicle_achievements_snapshot_vehicle_id_type_reference_id_created_at_idx` ON `vehicle_achievements_snapshot` (`vehicle_id`, `type`, `reference_id`, `created_at`);
