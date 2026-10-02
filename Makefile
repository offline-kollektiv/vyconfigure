COMMIT_HASH  := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE         := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

test:
	go test -v ./...

lint:
	golangci-lint run

dist:
	mkdir dist

clean:
	rm -Rf dist

build: clean dist
	GO111MODULE=on go build -v  -ldflags="-s -w -X 'github.com/offline-kollektiv/vyconfigure/cmd.Commit=$(COMMIT_HASH)' -X 'github.com/offline-kollektiv/vyconfigure/cmd.Date=$(DATE)'" -trimpath -o ./dist/vyconfigure
