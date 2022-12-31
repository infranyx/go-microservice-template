PKG := github.com/infranyx/go-microservice-template
VERSION ?= $(shell git describe --match 'v[0-9]*' --dirty='.m' --always --tags)
BINARY_NAME=infranyx_go_grpc_template
BINARY_PATH=./out/bin/$(BINARY_NAME)
MAIN_PATH=./cmd/main.go

GOCMD=go

TEST_COVERAGE_FLAGS = -race -coverprofile=coverage.out -covermode=atomic
TEST_FLAGS?= -timeout 15m

# Set ENV
export PG_URL=postgres://admin:admin@localhost:5432/go_microservice_template?sslmode=disable ### DB Conn String For Migrations

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

## ---------- Usual ----------
.PHONY: vendor
vendor: ## go mod vendor
	$(GOCMD) mod vendor

.PHONY: tidy
tidy: ## go mod tidy
	$(GOCMD) mod tidy

.PHONY: vet
vet: ## go mod vet
	$(GOCMD) vet

.PHONY: dep
dep: ## go mod download
	$(GOCMD) mod download

.PHONY: run_dev
run_dev: ## go run cmd/main.go
	$(GOCMD) run $(MAIN_PATH)

.PHONY: watch
watch: ## Run the code with cosmtrek/air to have automatic reload on changes
	$(eval PACKAGE_NAME=$(shell head -n 1 go.mod | cut -d ' ' -f2))
	docker run -it --rm -w /go/src/$(PACKAGE_NAME) -v $(shell pwd):/go/src/$(PACKAGE_NAME) -p $(SERVICE_PORT):$(SERVICE_PORT) cosmtrek/air

## ---------- Build ----------
.PHONY: build
build: tidy vendor ## tidy , vendor , mkdir out/bin , build
	mkdir -p out/bin

	GOARCH=amd64 GOOS=darwin GO111MODULE=on $(GOCMD) build -mod vendor -o  ${BINARY_PATH}  ${MAIN_PATH}

.PHONY: run
run: ## run binary
	GOARCH=amd64 GOOS=darwin ./${BINARY_PATH}

.PHONY: clean
clean: ## Remove build related file
	go clean
	rm -fr ./bin
	rm -fr ./out
	rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./coverage.out ./profile.cov yamllint-checkstyle.xml

## ---------- Test ----------
.PHONY: test
test: ## go clean -testcache && go test ./...
	go clean -testcache && go test ./...

.PHONY: test_coverage
test_coverage: ## go test ./... -coverprofile=coverage.out
	go test ./... -coverprofile=coverage.out

## ---------- Lint ----------
.PHONY: lint
lint: lint-go lint-dockerfile  ## Run all available linters

.PHONY: lint-dockerfile
lint-dockerfile: ## Lint your Dockerfile
# If dockerfile is present we lint it.
ifeq ($(shell test -e ./Dockerfile && echo -n yes),yes)
	$(eval CONFIG_OPTION = $(shell [ -e $(shell pwd)/.hadolint.yaml ] && echo "-v $(shell pwd)/.hadolint.yaml:/root/.config/hadolint.yaml" || echo "" ))
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--format checkstyle" || echo "" ))
	$(eval OUTPUT_FILE = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "| tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run --rm -i $(CONFIG_OPTION) hadolint/hadolint hadolint $(OUTPUT_OPTIONS) - < ./Dockerfile $(OUTPUT_FILE)
endif

.PHONY: lint-go
lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s $(OUTPUT_OPTIONS)


## ---------- Help ----------
.PHONY: help
help: ## Show this help.
	@echo ''
	@echo ${CYAN}'PKG:' ${GREEN}$(PKG)
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)



# TODO : Fix & Complete migration commands
.PHONY: rollback
migrate-rollback: ### migration roll-back
	migrate -source db/migrations -database $(PG_URL) down

.PHONY: drop
migrate-drop: ### migration drop
	migrate -source db/migrations -database $(PG_URL)  drop

.PHONY: migrate-create
migrate-create:  ### create new migration
	migrate create -ext sql -dir db/migrations $(migrate_name)

.PHONY: migrate-up
migrate-up: ### migration up
	migrate -path db/migrations -database $(PG_URL) up

.PHONY: force
migrate-force: ### force
	migrate -path db/migrations -database $(PG_URL) force $(id)
