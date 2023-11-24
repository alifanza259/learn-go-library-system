DROP TABLE IF EXISTS "email_verifications";

ALTER TABLE "members" DROP COLUMN "email_verified_at";
