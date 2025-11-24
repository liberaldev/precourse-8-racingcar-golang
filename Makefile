.PHONY: build build-all clean test run help

BUILD_DIR=build
BINARY_PATH=$(BUILD_DIR)/$(BINARY_NAME)
BINARY_NAME=racingcar

# Default target
all: build

# Build the project
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BINARY_PATH) .
	@echo "Build complete: $(BINARY_PATH)"

# Build for all platforms (for GitHub releases)
build-all:
	@echo "Building for all platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/racingcar_linux_amd64 .
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/racingcar_mac_intel .
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/racingcar_mac_apple .
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/racingcar_windows_x64.exe .
	GOOS=windows GOARCH=386 go build -o $(BUILD_DIR)/racingcar_windows_x86.exe .
	@echo "Build complete for all platforms:"
	@ls -lh $(BUILD_DIR)/

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Build and run
run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BINARY_PATH)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# Show help
help:
	@echo "Available targets:"
	@echo "  make build  - Build the project (default)"
	@echo "  make build-all  - Build for all platforms (Linux, macOS, Windows)"
	@echo "  make test   - Run tests"
	@echo "  make run    - Build and run"
	@echo "  make clean  - Remove build artifacts"
	@echo "  make help   - Show this help message"
	@echo ""
	@echo "Output: $(BINARY_PATH)"
