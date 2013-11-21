CREATE TABLE IF NOT EXISTS users (
    id       SERIAL PRIMARY KEY,
    created  INTEGER,
    username VARCHAR(20),
    password VARCHAR(75),
    salt     VARCHAR(25)
);

CREATE TABLE IF NOT EXISTS boards (
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(45),
    description  VARCHAR(140)
);

CREATE TABLE IF NOT EXISTS posts (
    id         SERIAL PRIMARY KEY,
    parent_id  INTEGER REFERENCES boards(id),
    author_id  INTEGER REFERENCES users(id) NOT NULL,
    title      VARCHAR(70),
    content    TEXT,
    created_on TIMESTAMP
);
