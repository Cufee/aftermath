-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Alter "accounts" table
ALTER TABLE `accounts` RENAME TO `account`;
-- Alter "account_snapshots" table
ALTER TABLE `account_snapshots` RENAME TO `account_snapshot`;
-- Alter "clans" table
ALTER TABLE `clans` RENAME TO `clan`;
-- Alter "vehicle_snapshots" table
ALTER TABLE `vehicle_snapshots` RENAME TO `vehicle_snapshot`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
