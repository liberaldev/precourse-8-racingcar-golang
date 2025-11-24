.PHONY: build clean test run help

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
	@echo "  make test   - Run tests"
	@echo "  make run    - Build and run"
	@echo "  make clean  - Remove build artifacts"
	@echo "  make help   - Show this help message"
	@echo ""
	@echo "Output: $(BINARY_PATH)"
