-- Modify "cards" table
ALTER TABLE "cards" ADD COLUMN "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, ADD COLUMN "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
