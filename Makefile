# Variables
APP_NAME = go-microservice
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
	@go build -o $(BIN_DIR)/$(APP_NAME)$(EXE) main.go
	@echo "$(APP_NAME)$(EXE)-$(VERSION) created!"

.PHONY: clean
clean:
	@echo "Cleaning up $(BIN_DIR)/$(APP_NAME)"
	@$(RM) $(BIN_DIR)
	@echo "Done!"

.PHONY: docker
docker: test build
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME):$(VERSION) .
	@make clean
	@echo "$(APP_NAME):$(VERSION) image created!"

.PHONY: info
info:
	@echo "PROJECT INFO:"
	@echo "+++++++++++++"
	@echo "Name: $(APP_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Binary folder: $(BIN_DIR)"

.PHONY: install
install:
	@echo "Downloading libraries..."
	@go mod download
	@echo "Installing swag..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Done!"

.PHONY: migrate
migrate:
	@echo "Running database schema migration..."
	@go run internal/migrator/main.go
	@echo "Done!"

.PHONY: run
run:
	@echo "Running the application $(APP_NAME)..."
	@go run main.go

.PHONY: swagger
swagger:
	@echo "Creating swagger api..."
	@swag init --parseDependency
	@swag fmt
	@go run docs/cmd/replace.go
	@echo "Done!"

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./tests/*

.PHONY: tidy
tidy:
	@echo "Running go mod tidy..."
	@go mod tidy
	@echo "Done!"

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make build    - Build the application"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make docker   - Create Docker image"
	@echo "  make help     - Show this help message"
	@echo "  make info     - Print Info"
	@echo "  make install  - Install libraries"
	@echo "  make migrate  - Migrate database schema"
	@echo "  make run      - Run the application"
	@echo "  make swagger  - Create swagger api"
	@echo "  make test     - Run tests"
	@echo "  make tidy     - Run go mod tidy"
