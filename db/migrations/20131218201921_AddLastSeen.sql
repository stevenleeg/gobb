-- +goose Up
ALTER TABLE users ADD COLUMN last_seen TIMESTAMP NOT NULL DEFAULT now();
ALTER TABLE users ADD COLUMN hide_online BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE users DROP COLUMN last_seen;
ALTER TABLE users DROP COLUMN hide_online;
