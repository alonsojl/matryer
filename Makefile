##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: linter
linter: ## Lint source code.
	@golangci-lint run -c .golangci.yml > linter.txt
			 
.PHONY: live	
live:  ## Live reload for Go applications.
	@air -c .air.toml 

.PHONY: test
test: ## Run unit tests.
	@go test -v ./...

.PHONY: tools
tools: ## Install tools.
	@go install github.com/cosmtrek/air@latest
	@go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2

.PHONY: migrate-up
migrate-up: ## Up migrations.
	@migrate -path db/migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true" up

.PHONY: migrate-down
migrate-down: ## Down migrations.
	@migrate -path db/migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true" down