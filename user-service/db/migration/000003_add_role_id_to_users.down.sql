DELETE FROM "users" WHERE username = 'admin';

DROP INDEX IF EXISTS "idx_users_role_id";

ALTER TABLE "users"
    DROP CONSTRAINT IF EXISTS "fk_users_roles";

ALTER TABLE "users"
    DROP COLUMN IF EXISTS "role_id";