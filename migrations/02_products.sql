CREATE TABLE IF NOT EXISTS products
(
    id           bigserial   NOT NULL PRIMARY KEY,
    name         text        NOT NULL,
    description  text        NOT NULL,
    review       text,
    rating       real        NOT NULL DEFAULT 0 CHECK ( 0 <= rating AND rating <= 5),
    rating_count bigint      NOT NULL DEFAULT 0 CHECK ( rating_count >= 0 ),
    price        bigint      NOT NULL DEFAULT 0 CHECK ( price >= 0 ),
    quantity     bigint      NOT NULL DEFAULT 0 CHECK ( quantity >= 0 ),
    status       bool        NOT NULL DEFAULT true,
    created_at   timestamptz NOT NULL DEFAULT now(),
    updated_at   timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS products_name_index on products (name);
CREATE INDEX IF NOT EXISTS products_rating_index on products (rating);
CREATE INDEX IF NOT EXISTS products_price_index on products (price);
CREATE INDEX IF NOT EXISTS products_status_index on products (status);