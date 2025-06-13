CREATE TABLE IF NOT EXISTS "users" (
  "id" bigint PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "profile_pic_url" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

INSERT INTO users (id, username, email, full_name)
VALUES (1, 'admin', 'admin@mail.com', 'Admin 1');