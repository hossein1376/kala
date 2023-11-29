CREATE TABLE IF NOT EXISTS products
(
    id           bigserial NOT NULL PRIMARY KEY,
    name         text      NOT NULL,
    description  text      NOT NULL,
    review       text,
    rating       real      NOT NULL CHECK ( 1 <= rating AND rating <= 5),
    rating_count bigint    NOT NULL CHECK ( rating_count >= 0 ),
    price        bigint    NOT NULL CHECK ( price >= 0 ),
    quantity     bigint    NOT NULL CHECK ( quantity >= 0 ),
    status       bool      NOT NULL DEFAULT true
);

CREATE INDEX IF NOT EXISTS products_name_index on products (name);
CREATE INDEX IF NOT EXISTS products_rating_index on products (rating);
CREATE INDEX IF NOT EXISTS products_price_index on products (price);
CREATE INDEX IF NOT EXISTS products_status_index on products (status);