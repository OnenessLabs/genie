VER_PACKAGE=github.com/OnenessLabs/genie/common

GIT_REVISION := $(shell git rev-parse --short HEAD)
COMMIT_DATE := $(shell git show -s --format=%cs HEAD)
BUILD_DATE := $(shell date -u +%d/%m/%Y@%H:%M:%S)

all: genied

.PHONY: genied
genied:
	go build -o build/genied -mod=readonly -ldflags "-X $(VER_PACKAGE).COMMIT=$(GIT_REVISION) -X $(VER_PACKAGE).BUILDDATE=$(BUILD_DATE)" ./cmd/genied

.PHONY: test
test:
	go test -failfast $(SHORTTEST) -race -v ./...

go.sum: go.mod
	@echo "Ensure dependencies have not been modified ..." >&2
	@go mod verify
	@go mod tidy
