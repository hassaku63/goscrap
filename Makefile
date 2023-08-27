GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")

.PHONY: clean
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf main

.PHONY: wire
wire:
	@echo "Generating wire files..."
	@wire ./...

.PHONY: build
build: $(GO_FILES) wire
	@echo "Building..."
	@go mod tidy
	@go build -o main

.PHONY: run
run: 
	@go run .