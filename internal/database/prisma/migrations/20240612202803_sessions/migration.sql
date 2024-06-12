/*
  Warnings:

  - You are about to drop the column `updatedAt` on the `achievements_snapshots` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX "account_snapshots_accountId_lastBattleTime_idx";

-- DropIndex
DROP INDEX "vehicle_snapshots_accountId_vehicleId_lastBattleTime_idx";

-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_achievements_snapshots" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "accountId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "dataEncoded" BLOB NOT NULL
);
INSERT INTO "new_achievements_snapshots" ("accountId", "createdAt", "dataEncoded", "id", "referenceId") SELECT "accountId", "createdAt", "dataEncoded", "id", "referenceId" FROM "achievements_snapshots";
DROP TABLE "achievements_snapshots";
ALTER TABLE "new_achievements_snapshots" RENAME TO "achievements_snapshots";
CREATE INDEX "achievements_snapshots_createdAt_idx" ON "achievements_snapshots"("createdAt");
CREATE INDEX "achievements_snapshots_accountId_referenceId_idx" ON "achievements_snapshots"("accountId", "referenceId");
CREATE INDEX "achievements_snapshots_accountId_referenceId_createdAt_idx" ON "achievements_snapshots"("accountId", "referenceId", "createdAt");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;

-- CreateIndex
CREATE INDEX "account_snapshots_type_accountId_referenceId_idx" ON "account_snapshots"("type", "accountId", "referenceId");

-- CreateIndex
CREATE INDEX "account_snapshots_type_accountId_referenceId_createdAt_idx" ON "account_snapshots"("type", "accountId", "referenceId", "createdAt");

-- CreateIndex
CREATE INDEX "vehicle_snapshots_vehicleId_createdAt_idx" ON "vehicle_snapshots"("vehicleId", "createdAt");

-- CreateIndex
CREATE INDEX "vehicle_snapshots_accountId_referenceId_idx" ON "vehicle_snapshots"("accountId", "referenceId");

-- CreateIndex
CREATE INDEX "vehicle_snapshots_accountId_referenceId_vehicleId_idx" ON "vehicle_snapshots"("accountId", "referenceId", "vehicleId");

-- CreateIndex
CREATE INDEX "vehicle_snapshots_accountId_referenceId_type_idx" ON "vehicle_snapshots"("accountId", "referenceId", "type");

-- CreateIndex
CREATE INDEX "vehicle_snapshots_accountId_referenceId_createdAt_idx" ON "vehicle_snapshots"("accountId", "referenceId", "createdAt");
