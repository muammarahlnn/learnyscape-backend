CREATE TABLE "users" (
  "id" bigint PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hash_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

INSERT INTO users (id, username, email, hash_password)
VALUES (1, 'admin', 'admin@mail.com', '$2a$10$C4Y1IlZBxrChpGaaqEGCW.JBNqKB/CS02Ov6QzVlo2FU4e48ikkoq');