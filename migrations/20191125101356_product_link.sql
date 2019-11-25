-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_category (
    id      BIGSERIAL PRIMARY KEY NOT NULL,
    product_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product_category;
-- +goose StatementEnd
