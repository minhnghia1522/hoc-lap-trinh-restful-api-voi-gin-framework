include .env
export

CONNECTION_STR = postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATION_DIRS = internal/db/migrations
PROD_COMPOSE = docker-compose.prod.yml
NOAPP_COMPOSE = docker-compose.noapp.yml
DEV_COMPOSE = docker-compose.dev.yml

ENV_FILE = .env

# Import Database
importdb:
	docker exec -i db psql -U root -d master-golang < ./backupdb-master-golang.sql

# Export Database
exportdb:
	docker exec -i db pg_dump -U root -d master-golang > ./backupdb-master-golang.sql

# Create new migration
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIRS) -seq $(NAME)

# Run all pending migrations (make migrate-up)
migrate-up:
	migrate -path $(MIGRATION_DIRS) -database "$(CONNECTION_STR)" up

# Rollback the last migration (make migrate-down)
migrate-down:
	migrate -path $(MIGRATION_DIRS) -database "$(CONNECTION_STR)" down 1

# Rollback N migrations
migrate-down-n:
	migrate -path $(MIGRATION_DIRS) -database "$(CONN_STRING)" down $(N)

# Force migration to the specified version (use with caution) (make migrate-force VERSION=1)
migrate-force:
	migrate -path $(MIGRATION_DIRS) -database "$(CONNECTION_STR)" force $(VERSION)

# Drop everything (include schema migration)
migration-drop:
	migrate -path $(MIGRATION_DIRS) -database "$(CONNECTION_STR)" drop

# Apply specific migration version (make migrate-goto VERSION=1)
migrate-goto:
	migrate -path $(MIGRATION_DIRS) -database "$(CONN_STRING)" goto $(VERSION)

sqlc:
	sqlc generate

server:
	go run ./cmd/api

build:
	go build -o bin/app ./cmd/api/

build-exe:
	go build -o bin/app.exe ./cmd/api/

run-binary:
	./bin/app

# Build container for dev
dev:
	docker compose -f $(DEV_COMPOSE) down
	docker compose -f $(DEV_COMPOSE) --env-file $(ENV_FILE) up --build

# Build container for noapp
noapp:
	docker compose -f $(NOAPP_COMPOSE) down
	docker compose -f $(NOAPP_COMPOSE) --env-file $(ENV_FILE) up --build -d

# Stop all containers in noapp
stop-noapp:
	docker compose -f $(NOAPP_COMPOSE) down

# Build and run the application in production mode using Docker Compose
prod:
	docker compose -f $(PROD_COMPOSE) down
	docker compose -f $(PROD_COMPOSE) --env-file $(ENV_FILE) up --build -d

# Stop all containers in production mode
stop-prod:
	docker compose -f $(PROD_COMPOSE) down

# View logs of the application in production mode
logs-prod:
	docker compose -f $(PROD_COMPOSE) logs -f --tail=100


.PHONY: importdb exportdb migrate-create migrate-up migrate-down migrate-down-n migrate-force migration-drop migrate-goto sqlc-g build build-exe server prod stop-prod logs-prod stop-noapp run-binary noapp dev
