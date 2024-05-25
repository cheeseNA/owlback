.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install-modules: ## Install go modules
	@go mod tidy

generate: ## Generate code
	@go generate ./...

.PHONY: build
build: ## Build image
	docker compose -f build/docker-compose.yaml build

run: ## Run container
	docker compose -f build/docker-compose.yaml up -d

down: ## Down container
	docker compose -f build/docker-compose.yaml down
