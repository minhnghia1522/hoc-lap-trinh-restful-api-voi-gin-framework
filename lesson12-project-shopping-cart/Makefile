include .env
export

CONNECTION_STR = postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATION_DIRS = internal/db/migrations

# Import Database
importdb:
	docker exec -i postgres-db-go psql -U root -d master-golang < ./backupdb-master-golang.sql

# Export Database
exportdb:
	docker exec -i postgres-db-go pg_dump -U root -d master-golang > ./backupdb-master-golang.sql

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
	cd ./cmd/api && go run .

build:
	go build -o bin/app ./cmd/api/

build-exe:
	go build -o bin/app.exe ./cmd/api/


.PHONY: importdb exportdb migrate-create migrate-up migrate-down migrate-down-n migrate-force migration-drop migrate-goto sqlc-g build build-exe server
