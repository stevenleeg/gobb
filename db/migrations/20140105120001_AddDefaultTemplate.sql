-- +goose Up
INSERT INTO SETTINGS (key, value) VALUES('template', 'default');
-- +goose Down
DELETE FROM SETTINGS WHERE key='template';
