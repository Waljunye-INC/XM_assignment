-- +goose Up
ALTER TABLE apartments
    ALTER COLUMN sq_meters TYPE real USING sq_meters::real;

-- +goose Down
ALTER TABLE apartments
    ALTER COLUMN sq_meters TYPE double precision USING sq_meters::double precision;