.PHONY: all build clean install build-all build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-windows-amd64

# Binary name
BINARY_NAME=venvcleaner

# Version (read from version.go)
VERSION=$(shell grep 'const Version' version.go | cut -d'"' -f2)

# Build directory
BUILD_DIR=dist

# Default target - build for current platform
all: build

# Build for current platform
build:
	@echo "Building $(BINARY_NAME) v$(VERSION) for current platform..."
	@go build -o $(BINARY_NAME) .
	@echo "✅ Build complete: $(BINARY_NAME)"

# Install to GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME) v$(VERSION)..."
	@go install .
	@echo "✅ Installed to $(shell go env GOPATH)/bin/$(BINARY_NAME)"

# Build for all platforms
build-all: clean build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-windows-amd64
	@echo "✅ All builds complete! Check the $(BUILD_DIR) directory."

# Build for macOS Intel (amd64)
build-darwin-amd64:
	@echo "Building for macOS (Intel)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	@echo "✅ Built: $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64"

# Build for macOS Apple Silicon (arm64)
build-darwin-arm64:
	@echo "Building for macOS (Apple Silicon)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	@echo "✅ Built: $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64"

# Build for Linux (amd64)
build-linux-amd64:
	@echo "Building for Linux (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	@echo "✅ Built: $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64"

# Build for Windows (amd64)
build-windows-amd64:
	@echo "Building for Windows (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	@echo "✅ Built: $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"

# Show help
help:
	@echo "VenvCleaner v$(VERSION) - Makefile commands:"
	@echo ""
	@echo "  make build          - Build for current platform (default)"
	@echo "  make install        - Install to GOPATH/bin"
	@echo "  make build-all      - Build for all platforms (creates dist/ directory)"
	@echo "  make clean          - Remove build artifacts"
	@echo ""
	@echo "Platform-specific builds:"
	@echo "  make build-darwin-amd64   - Build for macOS (Intel)"
	@echo "  make build-darwin-arm64   - Build for macOS (Apple Silicon)"
	@echo "  make build-linux-amd64    - Build for Linux"
	@echo "  make build-windows-amd64  - Build for Windows"
	@echo ""
	@echo "For distribution:"
	@echo "  1. Run 'make build-all' to create binaries for all platforms"
	@echo "  2. Binaries will be in the dist/ directory"
	@echo "  3. Tag your release: git tag v$(VERSION)"
	@echo "  4. Push to GitHub: git push origin v$(VERSION)"
	@echo "  5. Create a GitHub release and upload the binaries from dist/"
