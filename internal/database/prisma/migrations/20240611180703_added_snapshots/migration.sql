/*
  Warnings:

  - You are about to drop the column `frameEncoded` on the `account_snapshots` table. All the data in the column will be lost.
  - Added the required column `battles` to the `vehicle_snapshots` table without a default value. This is not possible if the table is not empty.
  - Added the required column `ratingBattles` to the `account_snapshots` table without a default value. This is not possible if the table is not empty.
  - Added the required column `ratingFrameEncoded` to the `account_snapshots` table without a default value. This is not possible if the table is not empty.
  - Added the required column `regularBattles` to the `account_snapshots` table without a default value. This is not possible if the table is not empty.
  - Added the required column `regularFrameEncoded` to the `account_snapshots` table without a default value. This is not possible if the table is not empty.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_vehicle_snapshots" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "lastBattleTime" DATETIME NOT NULL,
    "accountId" TEXT NOT NULL,
    "vehicleId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "battles" INTEGER NOT NULL,
    "frameEncoded" BLOB NOT NULL
);
INSERT INTO "new_vehicle_snapshots" ("accountId", "createdAt", "frameEncoded", "id", "lastBattleTime", "referenceId", "type", "vehicleId") SELECT "accountId", "createdAt", "frameEncoded", "id", "lastBattleTime", "referenceId", "type", "vehicleId" FROM "vehicle_snapshots";
DROP TABLE "vehicle_snapshots";
ALTER TABLE "new_vehicle_snapshots" RENAME TO "vehicle_snapshots";
CREATE INDEX "vehicle_snapshots_createdAt_idx" ON "vehicle_snapshots"("createdAt");
CREATE INDEX "vehicle_snapshots_accountId_vehicleId_lastBattleTime_idx" ON "vehicle_snapshots"("accountId", "vehicleId", "lastBattleTime");
CREATE TABLE "new_account_snapshots" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "lastBattleTime" DATETIME NOT NULL,
    "accountId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "ratingBattles" INTEGER NOT NULL,
    "ratingFrameEncoded" BLOB NOT NULL,
    "regularBattles" INTEGER NOT NULL,
    "regularFrameEncoded" BLOB NOT NULL
);
INSERT INTO "new_account_snapshots" ("accountId", "createdAt", "id", "lastBattleTime", "referenceId", "type") SELECT "accountId", "createdAt", "id", "lastBattleTime", "referenceId", "type" FROM "account_snapshots";
DROP TABLE "account_snapshots";
ALTER TABLE "new_account_snapshots" RENAME TO "account_snapshots";
CREATE INDEX "account_snapshots_createdAt_idx" ON "account_snapshots"("createdAt");
CREATE INDEX "account_snapshots_accountId_lastBattleTime_idx" ON "account_snapshots"("accountId", "lastBattleTime");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
