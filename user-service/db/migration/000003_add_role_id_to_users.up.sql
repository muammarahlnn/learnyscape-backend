ALTER TABLE "users"
    ADD COLUMN "role_id" bigint NOT NULL DEFAULT 3; -- default roles is student

ALTER TABLE "users"
    ADD CONSTRAINT "fk_users_roles" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") NOT VALID;

ALTER TABLE "users"
    VALIDATE CONSTRAINT "fk_users_roles";

CREATE INDEX IF NOT EXISTS "idx_users_role_id" ON "users" ("role_id");

INSERT INTO users (role_id, username, email, hash_password, full_name)
VALUES (1, 'admin', 'admin@mail.com', '$2a$10$C4Y1IlZBxrChpGaaqEGCW.JBNqKB/CS02Ov6QzVlo2FU4e48ikkoq', 'Admin 1');
