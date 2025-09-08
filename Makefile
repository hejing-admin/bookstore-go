.PHONY: all test build vendor

.PHONY: tidy
tidy: ## Run go mod tidy against code.
	go mod tidy

.PHONY: vendor
vendor: ## Run go mod vendor against code.
	go mod vendor

doc:
	swag fmt && swag init -g cmd/bookstore/main.go --parseDependency