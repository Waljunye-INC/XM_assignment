# XM Assignment API

A Go-based API service for companies management with authentication using MySQL as a data store and Kafka for event processing.

## Features

- Authentication and authorization via JWT
- CRUD operations for companies
- MySQL database integration
- Kafka event processing for specific operations
- Configuration via environment variables or `.env` file
- OpenAPI specification

## Requirements

- Go 1.24+
- MySQL 8+
- Kafka (for advanced features)

## Getting Started

### 1. Environment Setup

Create an `.env` file in the project root based on the `.env.template`:

```sh
cp .env.template .env
```

Example content (already in .env.template):

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=mydb
DB_USER=user
DB_PASSWORD=password
JWT_KEY=mysecretkey
KAFKA_BROKERS=localhost:9092
KAFKA_TOPIC_COMPANIES=stark-topic
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

- `make test` - run tests (WIP)
- `make migrate-up` - apply database migrations
- `make generate-mocks` - generate test mocks (WIP)

## Project Structure

```
cmd/
    main.go                  # Entry point
    config/                  # Configuration
internal/
    contract/oapi/           # HTTP handlers
    repositories/            # Data access (repositories)
        auth/                # Authentication repository
        companies/           # Companies repository
    usecases/                # Business logic
        auth/                # Authentication use cases
        companies/           # Companies use cases
    events/                  # Kafka event processing
libs/                       # Common libraries (HTTP server, listeners)
migrations/                 # SQL migrations
utils/                      # Utility functions
.env.template               # Environment template file
```

## API Endpoints

### Authentication

- `POST /auth/login` — User login, returns JWT token
- `POST /auth/register` — Register a new user

### Companies

- `POST /companies` — Create a new company
- `GET /companies/{uuid}` — Get a company by UUID
- `PUT /companies` — Update a company
- `DELETE /companies/{uuid}` — Delete a company by UUID

## Testing

WIP

## Docker Support

The project includes Docker and docker-compose support for easy setup:

```sh
docker-compose up --build
```

This will start the API service with all necessary dependencies (MySQL, Kafka).

## API Documentation

OpenAPI specification is available at `/docs/oapi/openapi.yaml`.

_Test assignment for XM_
