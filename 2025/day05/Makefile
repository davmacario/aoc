BINARY_NAME := $(shell basename $(CURDIR))
BUILD_DIR := build
BINARY_PATH := ./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: init deps build run all clean fmt help

init: ## Initialize go module and download dependencies
	go mod init $(BINARY_NAME)
	$(MAKE) deps

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

build: ## Build the binary
	@echo "Building $(BINARY_PATH)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BINARY_PATH) ./main.go

run: ## Run the binary
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_PATH)

all: build run ## Build and run the application

clean: ## Clean up build directory
	@echo "Cleaning up ./$(BUILD_DIR)"
	go clean
	@rm -rf ./$(BUILD_DIR)

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

help: ## Show this help
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## ' Makefile | awk 'BEGIN {FS="## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
