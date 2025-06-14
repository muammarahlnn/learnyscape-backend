CREATE TABLE IF NOT EXISTS "roles" (
    "id" bigint PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "deleted_at" timestamptz
);

INSERT INTO "roles" ("id", "name")
VALUES
    (1, 'Admin'),
    (2, 'Lecturer'),
    (3, 'Student');