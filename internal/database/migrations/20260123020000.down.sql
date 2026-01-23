-- Remove automod_verified column from user table
ALTER TABLE "user" DROP COLUMN IF EXISTS "automod_verified";
