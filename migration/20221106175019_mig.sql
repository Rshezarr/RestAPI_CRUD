-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    data VARCHAR
);

-- +goose Down
DROP TABLE IF EXISTS users;