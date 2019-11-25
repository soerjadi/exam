-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories (
    id          bigserial      PRIMARY KEY NOT NULL,
    name        varchar        NOT NULL,
    parent_id   bigint         NULL DEFAULT 0,
    created     timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated     timestamp      NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE category;
-- +goose StatementEnd
