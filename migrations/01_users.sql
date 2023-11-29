CREATE TABLE IF NOT EXISTS users
(
    id         bigserial   NOT NULL PRIMARY KEY,
    username   text        NOT NULL UNIQUE,
    password   bytea       NOT NULL,
    first_name text,
    last_name  text,
    email      text,
    phone      text,
    status     bool        NOT NULL DEFAULT true,
    role       text        NOT NULL DEFAULT 'user' CHECK ( role IN ('admin', 'seller', 'user') ),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS users_username_index on users (username);