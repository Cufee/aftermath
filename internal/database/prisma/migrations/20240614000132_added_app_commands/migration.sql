-- CreateTable
CREATE TABLE "application_commands" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "version" TEXT NOT NULL,
    "optionsHash" TEXT NOT NULL
);

-- CreateIndex
CREATE INDEX "account_snapshots_type_accountId_createdAt_idx" ON "account_snapshots"("type", "accountId", "createdAt");
