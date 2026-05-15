.PHONY: all build test lint coverage clean

all: build

build:
	go build ./...

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./... && \
	go tool cover -func=coverage.out | \
	awk '/total:/ {pct=$$NF; sub(/%/, "", pct); if (pct+0 < 80) {print "coverage below 80%: " pct; exit 1} else {print "coverage: " pct "%"}}'

lint:
	golangci-lint run ./...

clean:
	@if [ -f hashutils ]; then rm hashutils; fi
