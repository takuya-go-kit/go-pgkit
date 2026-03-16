-- +goose Up
CREATE TABLE IF NOT EXISTS pgkit_test (id SERIAL PRIMARY KEY, name TEXT);

-- +goose Down
DROP TABLE IF EXISTS pgkit_test;
