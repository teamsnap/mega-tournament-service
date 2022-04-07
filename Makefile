.PHONY: lint test coverage

lint: ## Run linting over our code
	@golangci-lint run --path-prefix=server

test: ## Run all the tests
	@go test -v $$(go list ./... | grep -v "mocks") | grep -v "=== RUN"

coverage: ## Run test code coverage
	@go test $$(go list ./... | grep -v "mocks") -coverprofile coverage.out -covermode count && \
		go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+' | awk '{ printf "Code coverage: %s\n", $$1 }'