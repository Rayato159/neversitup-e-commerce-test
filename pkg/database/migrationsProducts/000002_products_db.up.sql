BEGIN;

INSERT INTO "products" (
    "title",
    "description",
    "price"
)
VALUES
    ('Beer', 'Better than milk', 52),
    ('Cookie', 'This is for eat only not for any browser', 49),
    ('Coke', 'Better than Pepsi', 14),
    ('Chocolate', 'Sweetie pie', 25),
    ('F-35A', 'Prayusssssss plane', 80);

COMMIT;