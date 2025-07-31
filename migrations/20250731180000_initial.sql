-- +goose Up
CREATE TABLE IF NOT EXISTS companies
(
    uuid                  CHAR(36)                                                                 NOT NULL,
    name                VARCHAR(15) UNIQUE                                                       NOT NULL,
    description         TEXT,
    employees_count INT                                                                      NOT NULL,
    is_registered       BOOLEAN                                                                  NOT NULL,
    type                ENUM ('corporation', 'non_profit', 'cooperative', 'sole_proprietorship') NOT NULL,
    PRIMARY KEY (uuid),
    CONSTRAINT chk_description_length CHECK (CHAR_LENGTH(description) <= 3000)
);

-- +goose Down
DROP TABLE IF EXISTS companies;