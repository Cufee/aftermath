-- Modify "user" table
ALTER TABLE "user" ADD COLUMN "automod_verified" boolean NOT NULL DEFAULT false;

-- Backfill data based on existence of related records
UPDATE "user" u
SET "automod_verified" = true
WHERE
    EXISTS (SELECT 1 FROM "user_connection" WHERE user_id = u.id)
    OR EXISTS (SELECT 1 FROM "user_subscription" WHERE user_id = u.id)
    OR EXISTS (SELECT 1 FROM "user_restriction" WHERE user_id = u.id)
    OR EXISTS (SELECT 1 FROM "user_content" WHERE user_id = u.id)
    OR EXISTS (SELECT 1 FROM "widget_settings" WHERE user_id = u.id)
    OR EXISTS (
        SELECT 1 FROM "moderation_request"
        WHERE requestor_id = u.id OR moderator_id = u.id
    );
