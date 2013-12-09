-- +goose Up
ALTER TABLE users ADD COLUMN stylesheet_url varchar

-- +goose Down
ALTER TABLE users DROP COLUMN stylesheet_url
