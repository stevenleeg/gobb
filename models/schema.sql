CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    group_id    INTEGER DEFAULT 0,
    created_on  TIMESTAMP NOT NULL,
    username    VARCHAR(20),
    password    VARCHAR(75),
    avatar      VARCHAR,
    salt        VARCHAR(25)
);

CREATE TABLE IF NOT EXISTS boards (
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(45),
    description  VARCHAR(140)
);

CREATE TABLE IF NOT EXISTS posts (
    id           SERIAL PRIMARY KEY,
    board_id     INTEGER REFERENCES boards(id) NOT NULL,
    parent_id    INTEGER REFERENCES posts(id),
    author_id    INTEGER REFERENCES users(id) NOT NULL,
    title        VARCHAR(70) NOT NULL,
    content      TEXT NOT NULL,
    created_on   TIMESTAMP NOT NULL,
    latest_reply TIMESTAMP,
    last_edit    TIMESTAMP,
    sticky       BOOLEAN DEFAULT 'NO'
);
