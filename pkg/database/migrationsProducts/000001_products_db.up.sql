BEGIN;

--Set timezone
SET TIME ZONE 'Asia/Bangkok';

--Create sequence
CREATE SEQUENCE products_id_seq START WITH 1 INCREMENT BY 1;

--Auto update
CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE TABLE "products" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('P', LPAD(NEXTVAL('products_id_seq')::TEXT, 6, '0')),
  "title" VARCHAR,
  "description" VARCHAR,
  "price" FLOAT NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at_timestamp_products_table BEFORE UPDATE ON "products" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();

COMMIT;