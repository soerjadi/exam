-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product(
    id      BIGSERIAL       PRIMARY KEY NOT NULL,
    name    varchar         NOT NULL,
    sku     varchar         NOT NULL,
    created timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated timestamp       NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product;
-- +goose StatementEnd
