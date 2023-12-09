# Get the latest commit branch, hash, and date
TAG=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
BRANCH=$(if $(TAG),$(TAG),$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null))
BRANCH_SAFE=$(shell echo $(BRANCH) | sed -e 's/[^a-zA-Z0-9-]/-/g')
HASH=$(shell git rev-parse --short=7 HEAD 2>/dev/null)
TIMESTAMP=$(shell git log -1 --format=%ct HEAD 2>/dev/null | xargs -I{} date -u -r {} +%Y%m%dT%H%M%S)
GIT_REV=$(shell printf "%s-%s-%s" "$(BRANCH_SAFE)" "$(HASH)" "$(TIMESTAMP)")
REV=$(if $(filter --,$(GIT_REV)),latest,$(GIT_REV)) # fallback to latest if not in git repo


all: lint test build release

build:
	cd cmd/gol && go build -ldflags "-X main.revision=$(REV) -s -w" -o ../../.bin/gol.$(REV)
	cp .bin/gol.$(REV) .bin/gol
	.bin/gol --version

release:
	@echo release to dist/
	goreleaser --snapshot --clean

test:
	go clean -testcache
	go test -v -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out
	rm coverage.out

lint:
	golangci-lint run

version:
	@echo "branch: $(BRANCH), hash: $(HASH), timestamp: $(TIMESTAMP)"
	@echo "revision: $(REV)"

.PHONY: build release test lint version