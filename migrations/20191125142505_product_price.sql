-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product_price (
    id          BIGSERIAL   PRIMARY KEY NOT NULL,
    amount      BIGINT      NOT NULL,
    price       DOUBLE PRECISION DEFAULT 0.0::double PRECISION NOT NULL,
    product_id  BIGINT      NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product_price;
-- +goose StatementEnd
