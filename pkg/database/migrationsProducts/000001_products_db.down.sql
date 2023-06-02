BEGIN;

DROP TRIGGER IF EXISTS set_updated_at_timestamp_products_table ON "products";

DROP FUNCTION IF EXISTS set_updated_at_column();

DROP TABLE IF EXISTS "products" CASCADE;

DROP SEQUENCE IF EXISTS products_id_seq;

COMMIT;