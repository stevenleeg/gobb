-- +goose Up
CREATE TABLE IF NOT EXISTS views (
    id      VARCHAR(32) PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) NOT NULL,
    post_id INTEGER REFERENCES posts(id) NOT NULL,
    time    TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE views;
