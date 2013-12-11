-- +goose Up
ALTER TABLE users ADD COLUMN user_title VARCHAR(40) DEFAULT '';

-- +goose Down
ALTER TABLE users DROP COLUMN user_title;
