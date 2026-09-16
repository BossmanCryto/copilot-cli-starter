.PHONY: build test clean help live-skills-test

# Variables
GO := go
GOFLAGS := -v
BINARY_NAME := copilot-cli

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build $(GOFLAGS) -o bin/$(BINARY_NAME) .
	@echo "Build complete: bin/$(BINARY_NAME)"

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Test with live API/network
live-skills-test:
	@echo "Running live skills tests..."
	$(GO) test -v -tags=live ./skills/...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	$(GO) clean

# Run linting
lint:
	@echo "Linting..."
	golangci-lint run ./...

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Help
help:
	@echo "Copilot CLI - Makefile targets"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
