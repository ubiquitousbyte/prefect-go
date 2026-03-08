.PHONY: help generate test build examples clean install-tools fetch-spec

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

generate: ## Generate client code from OpenAPI spec
	@echo "Generating client code..."
	@go generate ./...
	@echo "✓ Code generation complete"

test: ## Run tests
	@echo "Running tests..."
	@go test ./...

build: ## Build the prefect package
	@echo "Building prefect package..."
	@cd prefect && go build .
	@echo "✓ Build complete"

examples: ## Build all example programs
	@echo "Building examples..."
	@for dir in examples/*/; do \
		echo "Building $$dir..."; \
		(cd $$dir && go build -o $$(basename $$dir)) || exit 1; \
	done
	@echo "✓ All examples built"

clean: ## Clean generated files and build artifacts
	@echo "Cleaning..."
	@find examples -type f -executable -delete
	@echo "✓ Clean complete"

install-tools: ## Install development tools
	@echo "Installing oapi-codegen (experimental)..."
	@go get -tool github.com/oapi-codegen/oapi-codegen-exp/experimental/cmd/oapi-codegen@latest
	@echo "✓ Tools installed"

fetch-spec: ## Fetch OpenAPI spec from Prefect instance (requires VERSION and optionally API_URL)
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make fetch-spec VERSION=3.6.21 [API_URL=http://localhost:4200]"; \
		exit 1; \
	fi
	@./scripts/fetch-openapi.sh $(VERSION) $(API_URL)

all: generate build examples test ## Run all build steps

.DEFAULT_GOAL := help
