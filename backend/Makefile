# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."


	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go
# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi
# Do database migrations
migrate:
	@if command -v tern > /dev/null; then \
           	set -a ; \
           	source ./.env; \
            echo ${DB_HOST}; \
            echo ${DB_DATABASE}; \
    		pushd db/migrations/; \
    		tern migrate; \
            popd; \
           	set +a ;\
		else \
		read -p "Tern is not installed on your machine. Would you like to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/jackc/tern@latest; \
               	source ./.env; \
                echo ${DB_HOST}; \
          		pushd db/migrations/; \
          		tern migrate; \
                popd; \
                echo "Migrated database successfully.";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi; \
       	DB_HOST= \
       	DB_USERNAME= \
       	DB_PASSWORD= \
       	DB_DATABASE= \
       	DB_PORT= \
       	DB_SCHEMA= \



.PHONY: all build run test clean watch docker-run docker-down itest migrate env
