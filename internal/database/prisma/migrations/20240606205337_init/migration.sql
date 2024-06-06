-- CreateTable
CREATE TABLE "app_configurations" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "valueEncoded" TEXT NOT NULL,
    "metadataEncoded" TEXT NOT NULL DEFAULT ''
);

-- CreateTable
CREATE TABLE "auth_nonces" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "expiresAt" DATETIME NOT NULL,
    "referenceId" TEXT NOT NULL,
    "metadataEncoded" TEXT NOT NULL DEFAULT ''
);

-- CreateTable
CREATE TABLE "users" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "permissions" TEXT NOT NULL,
    "featureFlagsEncoded" TEXT NOT NULL DEFAULT ''
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
    "metadataEncoded" TEXT NOT NULL DEFAULT '',
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
    "valueEncoded" TEXT NOT NULL,
    "metadataEncoded" TEXT NOT NULL DEFAULT '',
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
CREATE TABLE "account_snapshots" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "updatedAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "accountId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "lastBattleTime" DATETIME NOT NULL,
    "dataEncoded" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "account_rating_season_snapshots" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "updatedAt" DATETIME NOT NULL,
    "seasonId" TEXT NOT NULL,
    "accountId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "lastBattleTime" DATETIME NOT NULL,
    "dataEncoded" TEXT NOT NULL
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
CREATE TABLE "glossary_averages" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "dataEncoded" TEXT NOT NULL
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
    "localizedNamesEncoded" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "glossary_achievements" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "section" TEXT NOT NULL,
    "dataEncoded" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "user_interactions" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "dataEncoded" TEXT NOT NULL,
    "metadataEncoded" TEXT NOT NULL DEFAULT ''
);

-- CreateTable
CREATE TABLE "discord_live_sessions" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "locale" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "guildId" TEXT,
    "channelId" TEXT NOT NULL,
    "messageId" TEXT NOT NULL,
    "lastUpdate" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "lastBattleTime" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "optionsEncoded" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "stats_request_options" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "accountId" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "dataEncoded" TEXT NOT NULL
);

-- CreateIndex
CREATE INDEX "auth_nonces_createdAt_idx" ON "auth_nonces"("createdAt");

-- CreateIndex
CREATE INDEX "auth_nonces_referenceId_idx" ON "auth_nonces"("referenceId");

-- CreateIndex
CREATE INDEX "auth_nonces_referenceId_expiresAt_idx" ON "auth_nonces"("referenceId", "expiresAt");

-- CreateIndex
CREATE INDEX "user_subscriptions_userId_idx" ON "user_subscriptions"("userId");

-- CreateIndex
CREATE INDEX "user_subscriptions_expiresAt_idx" ON "user_subscriptions"("expiresAt");

-- CreateIndex
CREATE INDEX "user_subscriptions_type_userId_expiresAt_idx" ON "user_subscriptions"("type", "userId", "expiresAt");

-- CreateIndex
CREATE INDEX "user_subscriptions_type_userId_idx" ON "user_subscriptions"("type", "userId");

-- CreateIndex
CREATE INDEX "user_subscriptions_referenceId_idx" ON "user_subscriptions"("referenceId");

-- CreateIndex
CREATE INDEX "user_subscriptions_type_referenceId_idx" ON "user_subscriptions"("type", "referenceId");

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
CREATE INDEX "accounts_clanId_idx" ON "accounts"("clanId");

-- CreateIndex
CREATE INDEX "accounts_nickname_idx" ON "accounts"("nickname");

-- CreateIndex
CREATE INDEX "accounts_lastBattleTime_idx" ON "accounts"("lastBattleTime");

-- CreateIndex
CREATE INDEX "account_snapshots_createdAt_idx" ON "account_snapshots"("createdAt");

-- CreateIndex
CREATE INDEX "account_snapshots_type_accountId_idx" ON "account_snapshots"("type", "accountId");

-- CreateIndex
CREATE INDEX "account_snapshots_type_accountId_createdAt_idx" ON "account_snapshots"("type", "accountId", "createdAt");

-- CreateIndex
CREATE INDEX "account_snapshots_type_accountId_lastBattleTime_idx" ON "account_snapshots"("type", "accountId", "lastBattleTime");

-- CreateIndex
CREATE INDEX "account_snapshots_type_referenceId_idx" ON "account_snapshots"("type", "referenceId");

-- CreateIndex
CREATE INDEX "account_snapshots_type_referenceId_createdAt_idx" ON "account_snapshots"("type", "referenceId", "createdAt");

-- CreateIndex
CREATE INDEX "account_rating_season_snapshots_createdAt_idx" ON "account_rating_season_snapshots"("createdAt");

-- CreateIndex
CREATE INDEX "account_rating_season_snapshots_seasonId_accountId_idx" ON "account_rating_season_snapshots"("seasonId", "accountId");

-- CreateIndex
CREATE INDEX "account_rating_season_snapshots_seasonId_referenceId_idx" ON "account_rating_season_snapshots"("seasonId", "referenceId");

-- CreateIndex
CREATE INDEX "account_clans_tag_idx" ON "account_clans"("tag");

-- CreateIndex
CREATE INDEX "user_interactions_type_idx" ON "user_interactions"("type");

-- CreateIndex
CREATE INDEX "user_interactions_userId_idx" ON "user_interactions"("userId");

-- CreateIndex
CREATE INDEX "user_interactions_userId_type_idx" ON "user_interactions"("userId", "type");

-- CreateIndex
CREATE INDEX "user_interactions_referenceId_idx" ON "user_interactions"("referenceId");

-- CreateIndex
CREATE INDEX "user_interactions_referenceId_type_idx" ON "user_interactions"("referenceId", "type");

-- CreateIndex
CREATE INDEX "discord_live_sessions_referenceId_idx" ON "discord_live_sessions"("referenceId");

-- CreateIndex
CREATE INDEX "stats_request_options_referenceId_idx" ON "stats_request_options"("referenceId");
