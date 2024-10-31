CREATE TABLE users
(
    id            UUID default uuid_generate_v4() primary key,
    password      VARCHAR(255)        not null,
    email         VARCHAR(128) unique not null,
    refresh_token VARCHAR(255),
    created_at    TIMESTAMP default now()
);
