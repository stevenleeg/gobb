-- +goose Up
CREATE TABLE IF NOT EXISTS settings(
    key VARCHAR PRIMARY KEY,
    value VARCHAR
);

-- +goose Down
DROP TABLE settings;
