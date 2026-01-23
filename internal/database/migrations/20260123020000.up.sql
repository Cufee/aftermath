-- Add automod_verified column to user table
ALTER TABLE "user" ADD COLUMN IF NOT EXISTS "automod_verified" boolean NOT NULL DEFAULT false;

-- Backfill existing users as verified if they have activity
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
