CREATE TABLE "books" (
  "id" serial PRIMARY KEY,
  "isbn" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar,
  "author" varchar NOT NULL,
  "image_url" varchar,
  "genre" varchar NOT NULL,
  "quantity" int NOT NULL DEFAULT 0,
  "published_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);