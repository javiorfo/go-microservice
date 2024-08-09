# Variables
APP_NAME = go-microservice
VERSION = 0.0.1
MAIN_DIR = cmd/main.go
BIN_DIR = bin

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BIN_DIR)
	@gofmt -w .
	@go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_DIR)

.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -drf $(BIN_DIR)

.PHONY: run
run: build
	@echo "Running the application..."
	./$(BIN_DIR)/$(APP_NAME)

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make build    - Build the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make run      - Run the application"
	@echo "  make help     - Show this help message"


