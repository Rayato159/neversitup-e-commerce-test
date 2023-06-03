BEGIN;

DROP TRIGGER IF EXISTS set_updated_at_timestamp_oauth_table ON "oauth";

DROP FUNCTION IF EXISTS set_updated_at_column();

DROP TABLE IF EXISTS "oauth" CASCADE;

COMMIT;