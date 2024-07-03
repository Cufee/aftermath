-- Add column "metadata" to table: "auth_nonces"
ALTER TABLE `auth_nonces` ADD COLUMN `metadata` json NOT NULL;
-- Add column "metadata" to table: "sessions"
ALTER TABLE `sessions` ADD COLUMN `metadata` json NOT NULL;
