GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")
CMD_PACKAGES := $(shell go list ./cmd/... )

.PHONY: clean
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf main add

.PHONY: install-tools
install-tools:
	@echo "Installing tools..."
	go install golang.org/x/tools/cmd/godoc@latest

.PHONY: wire
wire:
	@echo "Generating wire files..."
	@wire ./...

.PHONY: build
build: $(GO_FILES) wire
	@echo "Building..."
	@go mod tidy
	@go build -o main

.PHONY: $(GO_FILES)
build-cmd:
	@echo "Building cmd..."
	@go build $(CMD_PACKAGES)

.PHONY: run
run: 
	@go run .