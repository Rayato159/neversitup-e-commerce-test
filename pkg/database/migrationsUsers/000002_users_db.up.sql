BEGIN;

INSERT INTO "users" (
    "username",
    "email",
    "password"
)
VALUES
    ('customer001', 'customer001@neversuitup.com', '$2a$10$8KzaNdKIMyOkASCH4QvSKuEMIY7Jc3vcHDuSJvXLii1rvBNgz60a6'),
    ('customer002', 'customer002@neversuitup.com', '$2a$10$8KzaNdKIMyOkASCH4QvSKuEMIY7Jc3vcHDuSJvXLii1rvBNgz60a6'); 

COMMIT;