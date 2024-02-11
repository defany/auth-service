protogen:
	buf generate proto

	go mod tidy

run:
	go run ./app/cmd/main.go