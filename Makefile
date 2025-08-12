# Makefile for gin-template project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOFMT=gofmt

# Project parameters
BINARY_NAME=gin-template
MAIN_PATH=./cmd
BUILD_DIR=./build
TOOLS_DIR=./tools

.PHONY: all build clean test help run wire gen fmt tidy

# Default target
all: build

# Help target
help: ## Display available targets
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  run      - Run the application"
	@echo "  clean    - Clean build artifacts"
	@echo "  fmt      - Format Go code"
	@echo "  tidy     - Tidy Go modules"
	@echo "  wire     - Generate wire dependency injection"
	@echo "  gen      - Generate GORM models"
	@echo "  help     - Display this help"

# Build the binary
build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the application
run: tidy wire ## Run the application
	@echo "Running $(BINARY_NAME)..."
	$(GOCMD) run $(MAIN_PATH)

# Clean build artifacts
clean: ## Clean build artifacts
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean completed"

# Format code
fmt: ## Format Go code
	@echo "Formatting code..."
	$(GOFMT) -s -w .
	@echo "Code formatting completed"

# Tidy modules
tidy: ## Tidy Go modules
	@echo "Tidying modules..."
	$(GOMOD) tidy
	@echo "Module tidy completed"

# Generate wire dependency injection
wire: ## Generate wire dependency injection
	@echo "Generating wire..."
	@cd cmd/wire && wire
	@echo "Wire generation completed"

# Generate GORM models
gen: ## Generate GORM models
	@echo "Generating GORM models..."
	$(GOCMD) run $(TOOLS_DIR)/gorm_gen/gorm_gen.go
	@echo "GORM models generation completed"