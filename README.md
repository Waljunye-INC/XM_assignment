# Buildings & Apartments API

A Go-based API service for building and apartment management using PostgreSQL as a data store.

## Features

- CRUD operations for buildings and apartments
- Retrieve apartments by building ID
- PostgreSQL database integration
- Configuration via environment variables or `.env` file
- OpenAPI (Swagger) specification

## Requirements

- Go 1.24+
- PostgreSQL 15+

## Getting Started

### 1. Environment Setup

Create an `.env` file in the project root based on the `.env.template`:

```sh
cp .env.template .env
```

Example content (already in .env.template):

```env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=testdb
DB_USER=postgres
DB_PASSWORD=password
PUBLIC_API_PORT=8012
```

### 2. Run Migrations

The project uses goose for database migrations. To apply migrations:

```sh
make migrate-up
```

### 3. Start the Application

```sh
go run ./cmd/main.go -cfg-file=.env
```

The API server will start on the port specified in the `PUBLIC_API_PORT` environment variable (default is 8012).

## Make Commands

The project includes a Makefile with the following commands:

- `make test` - run tests
- `make migrate-up` - apply database migrations
- `make generate-models` - generate SQLBoiler models
- `make generate-mocks` - generate test mocks

## Project Structure

```
cmd/
    main.go                  # Entry point
    config/                  # Configuration
internal/
    contract/oapi/          # HTTP handlers
    repository/              # Data access (repositories and SQLBoiler models)
    usecases/               # Business logic
    domain/                 # Domain models
libs/                       # Common libraries (HTTP server, listeners)
migrations/                 # SQL migrations
utils/                      # Utility functions
.env.template               # Environment template file
```

## API Endpoints
[Postman Reference](https://oms999-6301.postman.co/workspace/OMS_assignment~251230bb-c018-4e54-b9ae-92431872fb81/collection/29502646-34be816b-2f58-4d02-911c-0ac97d310b5b?action=share&creator=29502646)

### Buildings

- `GET /buildings` — Get all buildings
- `POST /buildings` — Create or update a building
- `GET /buildings/{id}` — Get a building by ID
- `DELETE /buildings/{id}` — Delete a building by ID

### Apartments

- `GET /apartments` — Get all apartments
- `POST /apartments` — Create or update an apartment
- `GET /apartments/{id}` — Get an apartment by ID
- `DELETE /apartments/{id}` — Delete an apartment by ID
- `GET /apartments/building/{buildingId}` — Get all apartments in a specific building

## Testing

To run tests:

```sh
make test
```

> Note: Ensure the database is running and test database credentials are configured through environment variables.

## API Documentation

OpenAPI specification is available at `/docs/oapi/openapi.yaml`.

_Test assignment for OMS_
