-- +goose Up
ALTER TABLE boards RENAME COLUMN boardorder TO ordering;

-- +goose Down
ALTER TABLE boards RENAME COLUMN ordering TO boardorder;
