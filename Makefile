run:
	@go build -o bin/main . && ./bin/main

create_migration:
	@read -p "Enter migration name: " name; \
		cd sql/schema && goose -s create $$name sql

migrate_up:
	@cd sql/schema && goose postgres postgres://postgres:password@localhost:5432/rssagg up

migrate_down:
	@cd sql/schema && goose postgres postgres://postgres:password@localhost:5432/rssagg down

db_build:
	@docker compose up -d

db_start:
	@docker compose start

db_stop:
	@docker compose stop

sqlc_generate:
	@sqlc generate

.PHONY: run migrate_up migrate_down create_migration db_build db_start db_stop sqlc_generate