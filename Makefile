.PHONY: all clean test build build-linux build-windows build-mac run-mac

all: clean build test

build: build-linux build-windows build-mac

build-linux:
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o bin/bees-in-the-trap-linux ./cmd/game

build-windows:
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/bees-in-the-trap-windows.exe ./cmd/game

build-mac:
	@echo "Building for macOS..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/bees-in-the-trap-mac ./cmd/game

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -rf bin/
	@mkdir -p bin/

run-mac: build-mac
	@echo "Running game on macOS..."
	@./bin/bees-in-the-trap-mac

run: build-linux
	@echo "Running game..."
	@./bin/bees-in-the-trap-linux 