.PHONY: all build compose-up

dirs := ./auth-service ./gateway

all: build compose-up

run: compose-dev-up

build:
	@echo "Building all services..."
	@for dir in $(dirs); do \
		echo "Building $$dir..."; \
		cd $$dir && make build && cd -; \
	done

compose-up:
	@docker network create --driver bridge production
	@docker compose -f deployment/compose-api.yaml up -d --build

compose-dev-up:
	@docker network create --driver bridge development
	@docker compose -f deployment/compose-api.dev.yaml up -d --build

compose-down:
	@docker compose -f deployment/compose-pgpool.yaml \
		down -v
	@docker compose -f deployment/compose-api.yaml \
		-f deployment/compose-api.dev.yaml \
		down -v
	@docker network rm production
