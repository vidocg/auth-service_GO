ALTER TABLE users
    ADD first_name VARCHAR(128);

ALTER TABLE users
    ADD last_name VARCHAR(128);

ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(255);