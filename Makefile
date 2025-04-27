.PHONY: run build test docker-up docker-down

# Run the application
run:
	go run ./cmd/api/...

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run tests
test:
	go test ./...

# Start Docker containers
docker-up:
	docker-compose -f docker-compose.local.yml up -d

# Stop Docker containers
docker-down:
	docker-compose -f docker-compose.local.yml down 