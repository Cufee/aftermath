-- Modify "user_content" table
ALTER TABLE "user_content" ALTER COLUMN "value" TYPE bytea USING "value"::TEXT::BYTEA;
