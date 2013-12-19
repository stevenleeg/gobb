-- +goose Up
ALTER TABLE users ADD COLUMN last_seen TIMESTAMP NOT NULL DEFAULT now();

-- +goose Down
ALTER TABLE users DROP COLUMN last_seen;
