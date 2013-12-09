-- +goose Up
ALTER TABLE users ADD COLUMN signature varchar

-- +goose Down
ALTER TABLE users DROP COLUMN signature
