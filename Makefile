include .env

REGISTRY_HOST=docker.io
REGISTRY:=defany
CONTAINER_NAME:=auth-service:v0.0.1

ifeq ($(OS),Windows_NT)
    path := $(shell echo %cd%)
else
    path := $(shell pwd)
endif

gen:
	make protogen

	make apigen

apigen:
	statik -src=app/pkg/gen/swagger/ -dest=app/pkg/gen/gen-swagger -include='*.css,*.html,*.js,*.json,*.png'

protogen:
	buf generate proto

	go mod download

run:
	go run ./app/cmd/main.go

build-push:
	docker buildx build --no-cache --platform linux/amd64 -t $(REGISTRY_HOST)/$(REGISTRY)/$(CONTAINER_NAME) .

	docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD) $(REGISTRY_HOST)/$(REGISTRY)

	docker push $(REGISTRY)/$(CONTAINER_NAME)

docker-run:
	docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)

	docker run -p 50001:50001 $(REGISTRY)/$(CONTAINER_NAME)

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "user=${PG_USER} port=${PG_PORT} dbname=${PG_DATABASE_NAME} sslmode=disable password=${PG_PASSWORD}" up -v

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "user=${PG_USER} port=${PG_PORT} dbname=${PG_DATABASE_NAME} sslmode=disable password=${PG_PASSWORD}" down -v

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "user=${PG_USER} port=${PG_PORT} dbname=${PG_DATABASE_NAME} sslmode=disable password=${PG_PASSWORD}" status -v

up-no-cache:
	docker compose down

	docker compose build --no-cache

	docker compose up -d

up:
	docker compose up --build -d

mockup:
	docker run --rm -v "$(path)":/src -w /src vektra/mockery --all

install:
	go mod download

	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
	github.com/rakyll/statik