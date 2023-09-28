CREATE TYPE "purpose" AS ENUM (
  'borrow',
  'return'
);

CREATE TYPE "status" AS ENUM (
  'pending',
  'declined',
  'approved'
);

CREATE TABLE "transactions" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "member_id" uuid NOT NULL,
  "admin_id" uuid,
  "borrow_id" uuid NOT NULL,
  "purpose" purpose NOT NULL,
  "status" status NOT NULL,
  "note" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "borrow_details" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "book_id" int NOT NULL,
  "borrowed_at" timestamptz NOT NULL,
  "returned_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("borrow_id") REFERENCES "borrow_details" ("id") ON DELETE CASCADE;
ALTER TABLE "transactions" ADD FOREIGN KEY ("member_id") REFERENCES "members" ("id") ON DELETE CASCADE;
ALTER TABLE "transactions" ADD FOREIGN KEY ("admin_id") REFERENCES "admin" ("id") ON DELETE CASCADE;
ALTER TABLE "borrow_details" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE;

CREATE INDEX IF NOT EXISTS idx_transactions_member_id ON "transactions" ("member_id");
CREATE INDEX IF NOT EXISTS idx_transactions_admin_id ON "transactions" ("admin_id");
CREATE INDEX IF NOT EXISTS idx_transactions_borrow_id ON "transactions" ("borrow_id");
CREATE INDEX IF NOT EXISTS idx_borrows_book_id ON "borrow_details" ("book_id");