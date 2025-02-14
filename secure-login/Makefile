# Makefile

# Variables
BINARY_NAME=webserver
SOURCE_FILES=main.go database.go utils.go middleware.go services.go

.PHONY: all build run clean

# Default target
all: build

# Build the binary
build:
	@echo "Building the application..."
	@go build -o $(BINARY_NAME) $(SOURCE_FILES)

# Run the binary
run: build
	@echo "Running the application..."
	@./$(BINARY_NAME)

# Clean the binary
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
