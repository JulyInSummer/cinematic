CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd/app
DATABASE_URL="postgres://cinematic:password@localhost:5432/cinematic?sslmode=disable"


build:
	CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

run-local-docker:
	docker-compose -f docker-compose.local.yaml up -d --build

new-migration:
	migrate create -ext sql -dir ${CURRENT_DIR}/migrations -seq -digits 8 $(name)

migrate:
	migrate -source file://${CURRENT_DIR}/migrations -database ${DATABASE_URL} -verbose up

swag-init:
	swag init -g cmd/app/main.go -o internal/app/transport/http/docs

test:
	go test -p 1 ./...

run:
	make swag-init
	make build
	go run ${APP_CMD_DIR}/main.go

.PHONY: migrate test run build