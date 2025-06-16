-- +goose Up
CREATE TABLE IF NOT EXISTS buildings(
    id bigserial primary key,
    name text NOT NULL,
    address text NOT NULL
);
CREATE TABLE IF NOT EXISTS apartments(
    id bigserial primary key,
    building_id bigint references buildings(id) NOT NULL,
    number text,
    floor int NOT NULL,
    sq_meters double precision NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS apartments;
DROP TABLE IF EXISTS buildings;