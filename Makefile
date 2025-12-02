build:
	@echo "building...."
	@go build -o bin/auth main.go

run: build
	@./bin/auth

test:
	@go test -v ./internal/tests/...