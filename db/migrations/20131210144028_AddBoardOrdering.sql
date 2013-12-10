-- +goose Up
ALTER TABLE boards ADD COLUMN boardorder INTEGER NOT NULL DEFAULT 1;

-- +goose Down
ALTER TABLE boards DROP COLUMN IF EXISTS boardorder;
