

run-local-docker:
	docker-compose -f docker-compose.local.yaml up -d --build

swag-init:
	swag init -g cmd/app/main.go -o internal/app/transport/http/docs