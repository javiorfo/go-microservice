# Variables
APP_NAME = microservice
VERSION = 0.0.1
MAIN_DIR = cmd
BIN_DIR = bin

# Detect OS
ifeq ($(OS),Windows_NT)
    RM = rmdir /s /q
    EXE = .exe
else
    RM = rm -rdf
    EXE =
endif

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BIN_DIR)
	@gofmt -w .
	@go build -o $(BIN_DIR)/$(APP_NAME)$(EXE) $(MAIN_DIR)/main.go
	@echo "$(APP_NAME)$(EXE) created!"

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./tests/*

.PHONY: clean
clean:
	@echo "Cleaning up '$(BIN_DIR)'"
	@$(RM) $(BIN_DIR)
	@echo "Done!"

.PHONY: run
run: build
	@echo "Running the application..."
	./$(BIN_DIR)/$(APP_NAME)$(EXE)

.PHONY: migrate
migrate:
	@echo "Running database schema migration..."
	@go run $(MAIN_DIR)/migrate/main.go

.PHONY: docker
docker: test build
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME):$(VERSION) .
	@make clean
	@echo "$(APP_NAME):$(VERSION) image created!"

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make build    - Build the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make run      - Run the application"
	@echo "  make migrate  - Migrate database schema"
	@echo "  make docker   - Create Docker image"
	@echo "  make help     - Show this help message"


