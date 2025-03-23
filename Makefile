.PHONY: build run clean restart

BIN_DIR = ./bin
BINARY_NAME = chat-server
PORT = 8080

# Detect OS (Linux, Darwin/macOS, Windows)
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := amd64

ifeq ($(OS), darwin)
    GOOS = darwin
    EXT =
else ifeq ($(OS), linux)
    GOOS = linux
    EXT =
else ifeq ($(OS), windows)
    GOOS = windows
    EXT = .exe
endif

BINARY = $(BIN_DIR)/$(BINARY_NAME)-$(GOOS)-$(ARCH)$(EXT)

# Build the Go application for the detected OS
build:
	@mkdir -p $(BIN_DIR)
	GOOS=$(GOOS) GOARCH=$(ARCH) go build -o $(BINARY) main.go
	@echo "Build completed: $(BINARY)"

# Build for all platforms
build-all: build-linux build-mac build-windows

build-linux:
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-linux-amd64 main.go
	@echo "Linux build completed."

build-mac:
	@mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-darwin-amd64 main.go
	@echo "macOS build completed."

build-windows:
	@mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go
	@echo "Windows build completed."

# Run the server after building
run: build
	$(BINARY) server

# Gracefully stop any process using the port and remove the binary
clean:
	@echo "Stopping process on port $(PORT)..."
	@-lsof -ti :$(PORT) | xargs -r kill -9
	@rm -rf $(BIN_DIR)
	@echo "Clean completed."

# Restart the server by cleaning and rerunning
restart: clean run
