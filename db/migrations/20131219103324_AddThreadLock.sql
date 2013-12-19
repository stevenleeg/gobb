-- +goose Up
ALTER TABLE posts ADD COLUMN locked BOOLEAN DEFAULT FALSE;

-- +goose Down
ALTER TABLE posts DROP COLUMN locked;
