# Pomodoro Timer Makefile

APP_NAME = pomodoro
MAIN_PATH = ./cmd/pomodoro
BUILD_DIR = ./bin
INSTALL_PATH = /usr/local/bin

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get
GOMOD = $(GOCMD) mod

# Build flags
LDFLAGS = -ldflags "-s -w"

.PHONY: all build clean test deps install uninstall run help

all: clean deps test build

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "Build complete! Binary created at $(BUILD_DIR)/$(APP_NAME)"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Install the application system-wide
install: build
	@echo "Installing $(APP_NAME) to $(INSTALL_PATH)..."
	@sudo cp $(BUILD_DIR)/$(APP_NAME) $(INSTALL_PATH)
	@sudo chmod +x $(INSTALL_PATH)/$(APP_NAME)
	@echo "Installation complete! You can now use '$(APP_NAME)' from anywhere."

# Uninstall the application
uninstall:
	@echo "Uninstalling $(APP_NAME)..."
	@sudo rm -f $(INSTALL_PATH)/$(APP_NAME)
	@echo "Uninstall complete!"

# Run the application locally
run: build
	$(BUILD_DIR)/$(APP_NAME)

# Run with custom parameters
run-custom: build
	$(BUILD_DIR)/$(APP_NAME) -w 30 -b 10

# Show help
help:
	@echo "Pomodoro Timer - Available commands:"
	@echo ""
	@echo "  make build      - Build the application"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make deps       - Download dependencies"
	@echo "  make install    - Install system-wide (requires sudo)"
	@echo "  make uninstall  - Uninstall from system (requires sudo)"
	@echo "  make run        - Build and run locally"
	@echo "  make run-custom - Build and run with custom timing (30min work, 10min break)"
	@echo "  make all        - Clean, download deps, test, and build"
	@echo "  make help       - Show this help message"
	@echo ""
	@echo "After installation, you can use:"
	@echo "  pomodoro                 # Default: 25min work, 5min break"
	@echo "  pomodoro -w 30 -b 10     # 30min work, 10min break" 
	@echo "  pomodoro -h              # Show help"