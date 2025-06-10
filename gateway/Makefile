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