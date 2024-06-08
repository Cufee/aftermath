-- CreateTable
CREATE TABLE "app_configurations" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "valueEncoded" BLOB NOT NULL,
    "metadataEncoded" BLOB NOT NULL
);

-- CreateTable
CREATE TABLE "cron_tasks" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "targetsEncoded" BLOB NOT NULL,
    "status" TEXT NOT NULL,
    "lastRun" DATETIME NOT NULL,
    "scheduledAfter" DATETIME NOT NULL,
    "logsEncoded" BLOB NOT NULL,
    "dataEncoded" BLOB NOT NULL
);

-- CreateTable
CREATE TABLE "users" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "permissions" TEXT NOT NULL,
    "featureFlags" INTEGER NOT NULL DEFAULT 0
);

-- CreateTable
CREATE TABLE "user_subscriptions" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "userId" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "expiresAt" DATETIME NOT NULL,
    "referenceId" TEXT NOT NULL,
    "permissions" TEXT NOT NULL,
    CONSTRAINT "user_subscriptions_userId_fkey" FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "user_connections" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "userId" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "permissions" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "metadataEncoded" BLOB NOT NULL,
    CONSTRAINT "user_connections_userId_fkey" FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "user_content" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "userId" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "valueEncoded" BLOB NOT NULL,
    "metadataEncoded" BLOB NOT NULL,
    CONSTRAINT "user_content_userId_fkey" FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "accounts" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "lastBattleTime" DATETIME NOT NULL,
    "accountCreatedAt" DATETIME NOT NULL,
    "realm" TEXT NOT NULL,
    "nickname" TEXT NOT NULL,
    "private" BOOLEAN NOT NULL DEFAULT false,
    "clanId" TEXT,
    CONSTRAINT "accounts_clanId_fkey" FOREIGN KEY ("clanId") REFERENCES "account_clans" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "account_clans" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "updatedAt" DATETIME NOT NULL,
    "tag" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "emblemId" TEXT NOT NULL DEFAULT '',
    "membersString" TEXT NOT NULL,
    "recordUpdatedAt" DATETIME NOT NULL
);

-- CreateTable
CREATE TABLE "account_snapshots" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "lastBattleTime" DATETIME NOT NULL,
    "accountId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "frameEncoded" BLOB NOT NULL
);

-- CreateTable
CREATE TABLE "vehicle_snapshots" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "lastBattleTime" DATETIME NOT NULL,
    "accountId" TEXT NOT NULL,
    "vehicleId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "frameEncoded" BLOB NOT NULL
);

-- CreateTable
CREATE TABLE "achievements_snapshots" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "accountId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "dataEncoded" BLOB NOT NULL
);

-- CreateTable
CREATE TABLE "glossary_averages" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "dataEncoded" BLOB NOT NULL
);

-- CreateTable
CREATE TABLE "glossary_vehicles" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "tier" INTEGER NOT NULL,
    "type" TEXT NOT NULL,
    "class" TEXT NOT NULL,
    "nation" TEXT NOT NULL,
    "localizedNamesEncoded" BLOB NOT NULL
);

-- CreateIndex
CREATE INDEX "cron_tasks_referenceId_idx" ON "cron_tasks"("referenceId");

-- CreateIndex
CREATE INDEX "cron_tasks_status_referenceId_scheduledAfter_idx" ON "cron_tasks"("status", "referenceId", "scheduledAfter");

-- CreateIndex
CREATE INDEX "cron_tasks_lastRun_status_idx" ON "cron_tasks"("lastRun", "status");

-- CreateIndex
CREATE INDEX "user_connections_userId_idx" ON "user_connections"("userId");

-- CreateIndex
CREATE INDEX "user_connections_type_userId_idx" ON "user_connections"("type", "userId");

-- CreateIndex
CREATE INDEX "user_connections_referenceId_idx" ON "user_connections"("referenceId");

-- CreateIndex
CREATE INDEX "user_connections_type_referenceId_idx" ON "user_connections"("type", "referenceId");

-- CreateIndex
CREATE INDEX "user_content_userId_idx" ON "user_content"("userId");

-- CreateIndex
CREATE INDEX "user_content_type_userId_idx" ON "user_content"("type", "userId");

-- CreateIndex
CREATE INDEX "user_content_referenceId_idx" ON "user_content"("referenceId");

-- CreateIndex
CREATE INDEX "user_content_type_referenceId_idx" ON "user_content"("type", "referenceId");

-- CreateIndex
CREATE INDEX "accounts_realm_idx" ON "accounts"("realm");

-- CreateIndex
CREATE INDEX "accounts_realm_lastBattleTime_idx" ON "accounts"("realm", "lastBattleTime");

-- CreateIndex
CREATE INDEX "accounts_id_lastBattleTime_idx" ON "accounts"("id", "lastBattleTime");

-- CreateIndex
CREATE INDEX "accounts_clanId_idx" ON "accounts"("clanId");

-- CreateIndex
CREATE INDEX "account_clans_tag_idx" ON "account_clans"("tag");

-- CreateIndex
CREATE INDEX "account_snapshots_createdAt_idx" ON "account_snapshots"("createdAt");

-- CreateIndex
CREATE INDEX "account_snapshots_accountId_lastBattleTime_idx" ON "account_snapshots"("accountId", "lastBattleTime");

-- CreateIndex
CREATE INDEX "vehicle_snapshots_createdAt_idx" ON "vehicle_snapshots"("createdAt");

-- CreateIndex
CREATE INDEX "vehicle_snapshots_accountId_vehicleId_lastBattleTime_idx" ON "vehicle_snapshots"("accountId", "vehicleId", "lastBattleTime");
