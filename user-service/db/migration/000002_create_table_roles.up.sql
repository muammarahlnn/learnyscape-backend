CREATE TABLE IF NOT EXISTS "roles" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

INSERT INTO "roles" ("name")
VALUES
    ('Admin'),
    ('Lecturer'),
    ('Student');