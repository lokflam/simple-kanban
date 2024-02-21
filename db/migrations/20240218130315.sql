-- Modify "cards" table
ALTER TABLE "cards" ADD COLUMN "status_id" uuid NOT NULL DEFAULT '018dbc48-4899-7aac-a1fa-0680a50c82a9', ADD CONSTRAINT "cards_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "statuses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
