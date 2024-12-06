-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Delete all snapshots due to a timestamp format and expiration logic change
DELETE FROM `vehicle_snapshot`;
DELETE FROM `account_snapshot`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
