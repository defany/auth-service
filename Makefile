include .env

REGISTRY_HOST=docker.io
REGISTRY:=defany
CONTAINER_NAME:=auth-service:v0.0.1

protogen:
	buf generate proto

	go mod download

	go mod tidy

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
	goose -dir $(MIGRATIONS_DIR) postgres "user=${PG_USER} dbname=${PG_DATABASE_NAME} sslmode=disable password=${PG_PASSWORD}" up -v

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "user=${PG_USER} dbname=${PG_DATABASE_NAME} sslmode=disable password=${PG_PASSWORD}" down -v

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "user=${PG_USER} dbname=${PG_DATABASE_NAME} sslmode=disable password=${PG_PASSWORD}" status -v

up:
	docker compose up --build -d