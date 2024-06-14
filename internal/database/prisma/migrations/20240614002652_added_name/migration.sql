/*
  Warnings:

  - Added the required column `name` to the `application_commands` table without a default value. This is not possible if the table is not empty.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_application_commands" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "name" TEXT NOT NULL,
    "version" TEXT NOT NULL,
    "optionsHash" TEXT NOT NULL
);
INSERT INTO "new_application_commands" ("createdAt", "id", "optionsHash", "updatedAt", "version") SELECT "createdAt", "id", "optionsHash", "updatedAt", "version" FROM "application_commands";
DROP TABLE "application_commands";
ALTER TABLE "new_application_commands" RENAME TO "application_commands";
CREATE INDEX "application_commands_optionsHash_idx" ON "application_commands"("optionsHash");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
