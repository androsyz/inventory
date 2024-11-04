include .env

run:
	go run cmd/main.go

build: 
	docker compose up -d

test:
	go test ./internal/...

