.PHONY: all build compose-up

dirs := ./gateway

all: build compose-up

build:
	@echo "Building all services..."
	@for dir in $(dirs); do \
		echo "Building $$dir..."; \
		cd $$dir && make build && cd -; \
	done

compose-up:
	@docker network create --driver bridge production
	@docker compose -f deployment/compose-api.yaml up -d --build

compose-down:
	@docker compose -f deployment/compose-api.yaml down -v
	@docker network rm production
