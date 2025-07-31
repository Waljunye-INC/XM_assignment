-- +goose Up
CREATE TABLE IF NOT EXISTS credentials
(
    username            VARCHAR(72) UNIQUE                                                            NOT NULL,
    password            TEXT                                                              NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS credentials;