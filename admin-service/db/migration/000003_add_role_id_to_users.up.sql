ALTER TABLE "users"
    ADD COLUMN "role_id" bigint NOT NULL DEFAULT 3; -- default roles is student

ALTER TABLE "users"
    ADD CONSTRAINT "fk_users_roles" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") NOT VALID;

ALTER TABLE "users"
    VALIDATE CONSTRAINT "fk_users_roles";

CREATE INDEX IF NOT EXISTS "idx_users_role_id" ON "users" ("role_id");

UPDATE
    "users"
SET
    role_id = 1
WHERE
    username = 'admin';