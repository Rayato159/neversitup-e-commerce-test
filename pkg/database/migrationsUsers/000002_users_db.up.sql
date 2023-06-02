BEGIN;

INSERT INTO "users" (
    "username",
    "password"
)
VALUES
    ('ruangyot', '$2a$10$8KzaNdKIMyOkASCH4QvSKuEMIY7Jc3vcHDuSJvXLii1rvBNgz60a6'),
    ('prayus', '$2a$10$8KzaNdKIMyOkASCH4QvSKuEMIY7Jc3vcHDuSJvXLii1rvBNgz60a6');

INSERT INTO "orders" (
    "user_id",
    "contact",
    "address",
    "status"
)
VALUES
    ('U000001', 'Ruangyot Nanchiang', '96/5 Chakrapatdipong Wat Sommanas Pom Prap Sattru Phai', 'completed'),
    ('U000001', 'Ruangyot Nanchiang', '96/5 Chakrapatdipong Wat Sommanas Pom Prap Sattru Phai', 'waiting'),
    ('U000002', 'Prayus Prayus', '919 Soi Phimtong, Ladprao 101 Rd., Bangkapi', 'completed');

INSERT INTO "products_orders" (
    "order_id",
    "qty",
    "product"
)
VALUES
    ('O000001', 1, '{"id": "P000001", "title": "Beer", "description": "Better than milk", "price": 52}'),
    ('O000001', 1, '{"id": "P000002", "title": "Cookie", "description": "This is for eat only not for any browser", "price": 49}'),
    ('O000001', 1, '{"id": "P000003", "title": "Coke", "description": "Better than Pepsi", "price": 14}'),
    ('O000002', 3, '{"id": "P000004", "title": "Chocolate", "description": "Sweetie pie", "price": 25}'),
    ('O000003', 2, '{"id": "P000005", "title": "F-35", "description": "Prayusssssss plane", "price": 80}');

COMMIT;