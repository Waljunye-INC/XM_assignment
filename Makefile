ENV_FILE := .env

include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

ifneq ("$(wildcard $(ENV_FILE))","")
	include $(ENV_FILE)
	export
endif

.PHONY: generate-mocks test generate-models
generate-mocks:
	@echo "Generating mocks..."
	@go generate ./...
	@echo "Mocks generated successfully!"

test:
	@echo "Starting tests"
	go test ./...

migrate-up:
	goose -dir migrations postgres "$(DB_URL)" up

generate-models:
	DB_NAME=$(DB_NAME) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) \
	envsubst < .sqlboiler.toml.tpl > .sqlboiler.toml && \
	sqlboiler psql --output internal/repository/models --wipe -c .sqlboiler.toml