# Makefile for Golang project

# Set the target binary name and directory
BINARY_NAME=./bin

# Main build target
all: fmt build build-win

# Install dependencies
deps:
	go mod download
	go get -v -t -d ./...

# Build the binary
build:
	mkdir -p ./bin
	go mod tidy
	go build -o $(BINARY_NAME) ./...

# Windows 64
build-win:
	mkdir -p $(BINARY_NAME)
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME) ./...

# Format the code using gofmt
fmt:
	go fmt ./...

# Help target to display available targets
help:
	@echo "Available targets:"
	@echo "  make all        - Build the binary (default target)"
	@echo "  make fmt        - Format the code using gofmt"
	@echo "  make help       - Show this help message"

.PHONY: all build clean deps run fmt test help
