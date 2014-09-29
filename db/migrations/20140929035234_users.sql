-- +goose Up
CREATE TABLE users (
    id           UUID PRIMARY KEY,
    stormpath_id varchar(255) NOT NULL,
    created_at   timestamp NOT NULL,
    updated_at   timestamp NOT NULL
);

-- +goose Down
DROP TABLE users;

