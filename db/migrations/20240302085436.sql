-- Create "status" table
CREATE TABLE "status" ("id" uuid NOT NULL, "name" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "board_status" table
CREATE TABLE "board_status" ("id" uuid NOT NULL, "status_id" uuid NOT NULL, "position" integer NOT NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "board_status_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "status" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "card" table
CREATE TABLE "card" ("id" uuid NOT NULL, "title" text NOT NULL, "content" text NOT NULL, "status_id" uuid NOT NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "card_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "status" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
