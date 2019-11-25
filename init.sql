CREATE TABLE IF NOT EXISTS product(
    id      BIGSERIAL       PRIMARY KEY NOT NULL,
    name    varchar         NOT NULL,
    sku     varchar         NOT NULL,
    created timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated timestamp       NULL
);

CREATE TABLE IF NOT EXISTS categories (
    id          bigserial      PRIMARY KEY NOT NULL,
    name        varchar        NOT NULL,
    parent_id   bigint         NULL DEFAULT 0,
    created     timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated     timestamp      NULL
);

CREATE TABLE product_category (
    id      BIGSERIAL PRIMARY KEY NOT NULL,
    product_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS product_price (
    id          BIGSERIAL   PRIMARY KEY NOT NULL,
    amount      BIGINT      NOT NULL,
    price       DOUBLE PRECISION DEFAULT 0.0::double PRECISION NOT NULL,
    product_id  BIGINT      NOT NULL
);

CREATE TABLE IF NOT EXISTS order (
    id          BIGSERIAL   PRIMARY KEY NOT NULL,
    product_id  BIGINT      NOT NULL,
    amount      BIGINT      NOT NULL,
    price       DOUBLE PRECISION DEFAULT 0.0::double PRECISION NOT NULL,
    status      SMALLINT    NOT NULL,
    created     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);