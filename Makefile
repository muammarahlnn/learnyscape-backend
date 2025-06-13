.PHONY: all build compose-up

dirs := ./user-service ./auth-service ./gateway

migrate_dirs := ./user-service ./auth-service ./admin-service

all: build compose-up

run: compose-dev-up

build:
	@echo "Building all services..."
	@for dir in $(dirs); do \
		echo "Building $$dir..."; \
		cd $$dir && make build && cd -; \
	done

migrate_up:
	@echo "Migrating all services..."
	@for dir in $(migrate_dirs); do\
		echo "Migrating $$dir..."; \
		DB_HOST=localhost; \
		if [ "$$dir" = "./user-service" ]; then \
			DB_PORT=5000; \
			DB_NAME=user_service_db; \
		elif [ "$$dir" = "./auth-service" ]; then \
			DB_PORT=5001; \
			DB_NAME=auth_service_db; \
		elif [ "$$dir" = "./admin-service" ]; then \
			DB_PORT=5002; \
			DB_NAME=admin_service_db; \
		fi;\
		migrate -path $$dir/db/migration/ -database "postgresql://postgres:postgres@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" -verbose up; \
	done

compose-up:
	@docker network create --driver bridge production
	@docker compose -f deployment/compose-pgpool.yaml up -d --build
	@make migrate_up
	@docker compose -f deployment/compose-api.yaml up -d --build

compose-dev-up:
	@docker network create --driver bridge production
	@docker compose -f deployment/compose-pgpool.yaml up -d --build
	@echo "Wait for 5 seconds to make sure pgpool is ready to accept connection..."
	@sleep 5
	@make migrate_up
	@docker compose -f deployment/compose-api.dev.yaml up -d --build

compose-down:
	@docker compose -f deployment/compose-pgpool.yaml \
		down -v
	@docker compose -f deployment/compose-api.yaml \
		-f deployment/compose-api.dev.yaml \
		down -v
	@docker network rm production
