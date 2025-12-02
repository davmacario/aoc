BINARY_NAME := $(shell basename $(CURDIR))
BUILD_DIR := build
BINARY_PATH := ./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: init build run clean fmt

all: build run

init:
	go mod init $(BINARY_NAME)

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

build:
	@echo "Building $(BINARY_PATH)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BINARY_PATH) ./main.go

run:
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_PATH) $(input_path) $(output_path)

clean:
	@echo "Cleaning up ./$(BUILD_DIR)"
	go clean
	@rm -rf ./$(BUILD_DIR)

fmt:
	fmt:
	@echo "Formatting code..."
	go fmt ./...
