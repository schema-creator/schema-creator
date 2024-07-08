-- Drop Foreign Keys
ALTER TABLE "project_has_users" DROP CONSTRAINT IF EXISTS "project_has_users_project_id_fkey";
ALTER TABLE "project_has_users" DROP CONSTRAINT IF EXISTS "project_has_users_user_id_fkey";
ALTER TABLE "projects" DROP CONSTRAINT IF EXISTS "projects_owner_id_fkey";
ALTER TABLE "project_saves" DROP CONSTRAINT IF EXISTS "project_saves_project_id_fkey";
ALTER TABLE "session" DROP CONSTRAINT IF EXISTS "session_user_id_fkey";
ALTER TABLE "project_saves" DROP CONSTRAINT IF EXISTS "project_saves_saved_by_fkey";

-- Drop Indexes
DROP INDEX IF EXISTS "projects_project_id_idx";
DROP INDEX IF EXISTS "projects_title_idx";
DROP INDEX IF EXISTS "project_saves_project_id_idx";
DROP INDEX IF EXISTS "project_saves_updated_at_idx";
DROP INDEX IF EXISTS "project_has_users_project_id_idx";
DROP INDEX IF EXISTS "project_has_users_user_id_idx";
DROP INDEX IF EXISTS "users_user_id_idx";
DROP INDEX IF EXISTS "users_email_idx";

-- Drop Tables
DROP TABLE IF EXISTS "projects";
DROP TABLE IF EXISTS "project_saves";
DROP TABLE IF EXISTS "project_has_users";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "session";

-- Drop Type
DROP TYPE IF EXISTS "role";