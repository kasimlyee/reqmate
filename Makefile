# Build variables
BINARY_NAME ?= reqmate
VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_DATE ?= $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.buildDate=$(BUILD_DATE)"

# Tools
GOLANGCI_LINT = golangci-lint
GORELEASER = goreleaser

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/reqmate/main.go

.PHONY: install
install:
	go install $(LDFLAGS) ./cmd/reqmate

.PHONY: test
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: cover
cover: test
	go tool cover -html=coverage.out

.PHONY: lint
lint:
	@echo "Running linters..."
	$(GOLANGCI_LINT) run

.PHONY: release-dry-run
release-dry-run:
	$(GORELEASER) release --rm-dist --snapshot --skip-publish

.PHONY: release
release:
	$(GORELEASER) release --rm-dist

.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf bin/ coverage.out dist/

.PHONY: ci-verify
ci-verify: lint test