CREATE TABLE users(
    id  INTEGER PRIMARY KEY AUTOINCREMENT,
    created  INTEGER,
    username VARCHAR,
    password VARCHAR,
    salt     VARCHAR,
    sid      VARCHAR
);
