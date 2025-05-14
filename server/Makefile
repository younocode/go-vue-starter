# Simple Makefile for a Go project

install_all: install_migrate install_sqlc install_wire
# migrate
install_migrate:
	@echo "Installing migrate..."
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# sqlc
install_sqlc:
	@echo "Installing sqlc..."
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# wire
install_wire:
	@echo "Installing wire..."
	@go install github.com/google/wire/cmd/wire@latest


# gen
gen_sqlc:
	@echo "gen sqlc..."
	@sqlc generate

# migrate
migrate_up:
	migrate -path="./database/migrate" -database=${databaseURL} up

migrate_drop:
	migrate -path="./database/migrate" -database=${databaseURL} drop -f

# Build the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o main.exe cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go
# Create DB container
docker-run:
	@docker compose up --build

# Shutdown DB container
docker-down:
	@docker compose down

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
	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command air -ErrorAction SilentlyContinue) { \
		air; \
		Write-Output 'Watching...'; \
	} else { \
		Write-Output 'Installing air...'; \
		go install github.com/air-verse/air@latest; \
		air; \
		Write-Output 'Watching...'; \
	}"

.PHONY: all build run test clean watch docker-run docker-down itest
