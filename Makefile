.PHONY: all build test lint coverage clean yamllint rumdl ci help

all: build

build: ## Build all packages
	go build ./...

test: ## Run tests
	go test ./...

coverage: ## Run tests with coverage check (>=80%)
	go test -coverprofile=coverage.out ./... && \
	go tool cover -func=coverage.out | \
	awk '/total:/ {pct=$$NF; sub(/%/, "", pct); if (pct+0 < 80) {print "coverage below 80%: " pct; exit 1} else {print "coverage: " pct "%"}}'

lint: ## Run golangci-lint
	golangci-lint run ./...

yamllint: ## Lint YAML files
	yamllint .woodpecker/build.yaml renovate.json

rumdl: ## Format Markdown files
	rumdl fmt *.md

clean: ## Remove build artifacts
	@if [ -f hashutils ]; then rm hashutils; fi

ci: ## Run CI pipeline locally
	woodpecker exec --backend-engine docker --pipeline-event push --backend-docker-network bridge --repo-path "$$PWD" .woodpecker/build.yaml

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
