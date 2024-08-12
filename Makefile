# Makefile

# Project Name
PROJECT_NAME := wave-message

# Version
VERSION := 0.1.0

# Directories
CMD_DIR := cmd
PKG_DIR := pkg
TEST_DIR := tests

# Go commands
GO := go
GOFMT := gofmt
GOTEST := go test

# Default target
all: fmt build

# Format the code
fmt:
	$(GOFMT) -w $(CMD_DIR) $(PKG_DIR) $(TEST_DIR)

# Build the project
build:
	$(GO) build -o bin/$(PROJECT_NAME) $(CMD_DIR)/main.go

# Run tests
test:
	$(GOTEST) -v ./$(TEST_DIR)

# Clean up generated files
clean:
	rm -rf bin

# Show version
version:
	@echo $(VERSION)

.PHONY: all fmt build test clean version
