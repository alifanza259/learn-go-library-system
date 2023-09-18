CREATE TYPE "gender" AS ENUM (
  'male',
  'female',
  'unknown'
);
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "admin" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "email" varchar NOT NULL UNIQUE,
  "first_name" varchar NOT NULL,
  "last_name" varchar,
  "password" varchar NOT NULL,
  "permission" varchar NOT NULL,
  "last_accessed_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "members" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "email" varchar NOT NULL UNIQUE,
  "first_name" varchar NOT NULL,
  "last_name" varchar,
  "dob" date NOT NULL,
  "gender" gender NOT NULL,
  "password" varchar NOT NULL,
  "password_changed_at" timestamptz,
  "last_accessed_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);