include ${CURDIR}/.env

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -v -o ${CURDIR}/bin/main ${CURDIR}/cmd/api/main.go

test:
	@go test -v -cover ${CURDIR}/...

test-cover:
	@go test ${CURDIR}/... -coverprofile=coverage.out
	@go tool cover -html=coverage.out && rm -f coverage.out

coverage:
	@go test ${CURDIR}/... -coverprofile=cover.out
	@go tool cover -html=cover.out && rm -rf cover.out

mockery:
	@mockery --all --case underscore --dir ${input} --output ${output}

migrate_create:
	@migrate create -ext sql -dir ${CURDIR}/db/migration/ -seq ${name}

migrate_force:
	@migrate -path ${CURDIR}/db/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose force 1

migrate_down:
	@migrate -path ${CURDIR}/db/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose down

migrate_down_one:
	@migrate -path ${CURDIR}/db/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose down 1

migrate_up:
	@migrate -path ${CURDIR}/db/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose up