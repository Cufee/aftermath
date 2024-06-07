-- CreateTable
CREATE TABLE "cron_tasks" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "type" TEXT NOT NULL,
    "referenceId" TEXT NOT NULL,
    "targetsEncoded" TEXT NOT NULL,
    "status" TEXT NOT NULL,
    "lastRun" DATETIME NOT NULL,
    "scheduledAfter" DATETIME NOT NULL,
    "logsEncoded" TEXT NOT NULL DEFAULT '',
    "dataEncoded" TEXT NOT NULL DEFAULT ''
);

-- CreateIndex
CREATE INDEX "cron_tasks_referenceId_idx" ON "cron_tasks"("referenceId");

-- CreateIndex
CREATE INDEX "cron_tasks_status_referenceId_scheduledAfter_idx" ON "cron_tasks"("status", "referenceId", "scheduledAfter");

-- CreateIndex
CREATE INDEX "cron_tasks_lastRun_status_idx" ON "cron_tasks"("lastRun", "status");
