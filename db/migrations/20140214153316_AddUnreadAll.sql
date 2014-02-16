-- +goose Up
ALTER TABLE users ADD COLUMN last_unread_all TIMESTAMP;

-- +goose Down
ALTER TABLE users DROP COLUMN IF EXISTS last_unread_all;
