CREATE TABLE "email_verifications" (
    "id" serial PRIMARY KEY,
    "member_id" uuid NOT NULL,
    "token" varchar NOT NULL,
    "is_used" boolean NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "email_verifications" ADD FOREIGN KEY ("member_id") REFERENCES "members" ("id") ON DELETE CASCADE;

ALTER TABLE "members" ADD "email_verified_at" timestamptz;