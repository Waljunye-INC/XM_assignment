ENV_FILE := .env

include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

DB_URL=$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

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
	goose -dir migrations mysql "$(DB_URL)" up

migrate-down:
	goose -dir migrations mysql "$(DB_URL)" down

