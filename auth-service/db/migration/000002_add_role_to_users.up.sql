ALTER TABLE "users"
ADD COLUMN "role" varchar NOT NULL DEFAULT 'student';

UPDATE
    users
SET
    role = 'admin',
    updated_at = NOW()
WHERE
    username = 'admin'
    AND deleted_at IS NULL;