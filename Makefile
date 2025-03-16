CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd/app


build:
	CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go


run-local-docker:
	docker-compose -f docker-compose.local.yaml up -d --build

swag-init:
	swag init -g cmd/app/main.go -o internal/app/transport/http/docs

run:
	make swag-init
	make build
	go run ${APP_CMD_DIR}/main.go