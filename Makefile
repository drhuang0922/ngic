# Variables
APP_NAME = ngic
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DIR = build
MAIN_PATH = ./cmd/ngic

# Build info
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Go build flags
LDFLAGS = -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(GIT_COMMIT)
BUILD_FLAGS = -ldflags "$(LDFLAGS)" -trimpath

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building $(APP_NAME) v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# Install the application locally
.PHONY: install
install:
	@echo "Installing $(APP_NAME)..."
	go install $(BUILD_FLAGS) $(MAIN_PATH)

# Cross-compile for multiple platforms
.PHONY: build-all
build-all:
	@echo "Cross-compiling $(APP_NAME) v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	
	# macOS (Intel)
	GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 $(MAIN_PATH)
	
	# macOS (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 $(MAIN_PATH)
	
	# Linux (64-bit)
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 $(MAIN_PATH)
	
	# Linux (ARM64)
	GOOS=linux GOARCH=arm64 go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 $(MAIN_PATH)
	
	# Windows (64-bit)
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe $(MAIN_PATH)

# Create release archives
.PHONY: release
release: build-all
	@echo "Creating release archives..."
	@cd $(BUILD_DIR) && \
	tar -czf $(APP_NAME)-darwin-amd64.tar.gz $(APP_NAME)-darwin-amd64 && \
	tar -czf $(APP_NAME)-darwin-arm64.tar.gz $(APP_NAME)-darwin-arm64 && \
	tar -czf $(APP_NAME)-linux-amd64.tar.gz $(APP_NAME)-linux-amd64 && \
	tar -czf $(APP_NAME)-linux-arm64.tar.gz $(APP_NAME)-linux-arm64 && \
	zip $(APP_NAME)-windows-amd64.zip $(APP_NAME)-windows-amd64.exe

# Test the application
.PHONY: test
test:
	go test -v ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

# Run linter
.PHONY: lint
lint:
	golangci-lint run

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Tidy dependencies
.PHONY: tidy
tidy:
	go mod tidy

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  install    - Install the application locally"
	@echo "  build-all  - Cross-compile for multiple platforms"
	@echo "  release    - Create release archives"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  lint       - Run linter"
	@echo "  fmt        - Format code"
	@echo "  tidy       - Tidy dependencies"
	@echo "  help       - Show this help"
