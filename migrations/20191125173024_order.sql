-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order (
    id          BIGSERIAL   PRIMARY KEY NOT NULL,
    product_id  BIGINT      NOT NULL,
    amount      BIGINT      NOT NULL,
    price       DOUBLE PRECISION DEFAULT 0.0::double PRECISION NOT NULL,
    status      SMALLINT    NOT NULL,
    created     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order;
-- +goose StatementEnd
