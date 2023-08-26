GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")

.PHONY: clean
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf main

.PHONY: build
build: $(GO_FILES)
	@echo "Building..."
	@go mod tidy
	@go build -o main main.go

.PHONY: run
run: 
	@go run .