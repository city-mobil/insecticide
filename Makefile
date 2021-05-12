ROOT_APP_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
APP:=insecticide

.PHONY: all
all: build

.PHONY: build
build: clean ## Build application
	go generate ./...
	CGO_ENABLED=0 go build -o bin/${APP} cmd/${APP}/main.go

.PHONY: run
run: build ## Build and run
	bin/${APP}

.PHONY: fmt
fmt: ## Code format
	go fmt ./...
	goimports -w ./

.PHONY: lint
lint: ## Run linter
	golangci-lint run -v ./...

.PHONY: test
test: ## Run tests
	go test ./internal...

.PHONY: cover
cover: ## Run tests with coverage
	go test -coverprofile=coverage.out ./internal...

.PHONY: clean
clean: ## Cleanup binaries
	rm -f bin/*

.PHONY: help
help: ## List of commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
