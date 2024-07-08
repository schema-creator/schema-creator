CREATE TYPE "role" AS ENUM (
  'admin',
  'readwrite',
  'read'
);

CREATE TABLE "projects" (
  "project_id" varchar(16) PRIMARY KEY,
  "title" varchar(50) NOT NULL,
  "owner_id" varchar(16) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "is_deleted" boolean NOT NULL DEFAULT 'false'
);

CREATE TABLE "project_saves" (
  "project_save_id" varchar(16),
  "project_id" varchar(16) NOT NULL,
  "editor" jsonb,
  "object" jsonb,
  "saved_by" varchar(16),
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  PRIMARY KEY ("project_save_id", "saved_by")
);

CREATE TABLE "project_has_users" (
  "project_id" varchar(16) NOT NULL,
  "user_id" varchar(16) NOT NULL,
  "role" role NOT NULL
);

CREATE TABLE "users" (
  "user_id" varchar(16) PRIMARY KEY,
  "email" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "icon" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "is_deleted" boolean NOT NULL DEFAULT 'false'
);

CREATE TABLE "session" (
  "session_id" varchar(63) PRIMARY KEY,
  "user_agent" text NOT NULL,
  "user_id" varchar(63) NOT NULL,
  "token" text NOT NULL,
  "expiration_time" int NOT NULL
);

CREATE INDEX ON "projects" USING BTREE ("project_id");

CREATE INDEX ON "projects" USING BTREE ("title");

CREATE INDEX ON "project_saves" USING BTREE ("project_id");

CREATE INDEX ON "project_saves" USING BTREE ("updated_at");

CREATE INDEX ON "project_has_users" USING BTREE ("project_id");

CREATE INDEX ON "project_has_users" USING BTREE ("user_id");

CREATE INDEX ON "users" USING BTREE ("user_id");

CREATE INDEX ON "users" USING BTREE ("email");

ALTER TABLE "project_has_users" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("project_id");

ALTER TABLE "project_has_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "projects" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("user_id");

ALTER TABLE "project_saves" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("project_id");

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "project_saves" ADD FOREIGN KEY ("saved_by") REFERENCES "users" ("user_id");
