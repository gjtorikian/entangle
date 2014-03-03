PACKAGES := \
	entangle/token \
	entangle/source \
	entangle/parser \
	entangle/lexer \
	entangle/declarations \
	entangle/utils \
	entangle/term \
	entangle/generators \
	entangle/generators/golang
SOURCE := $(wildcard $(addsuffix /*.go, $(addprefix src/, $(PACKAGES)))) src/entangle/data/assets.go
DATA_SOURCE := $(shell find data -type f ! -name '.*')

export GOPATH=$(shell pwd)

all: entangle

entangle: $(SOURCE)
	@go build -v -o bin/entangle cmds/entangle

src/entangle/data/assets.go: bin/go-bindata $(DATA_SOURCE)
	@./bin/go-bindata \
		-nocompress \
		-prefix="data/" \
		-o="src/entangle/data/assets.go" \
		-pkg="data" /
		data/...

bin/go-bindata: src/github.com/jteeuwen/go-bindata
	@go build -v -o bin/go-bindata github.com/jteeuwen/go-bindata/go-bindata

src/github.com/jteeuwen/go-bindata:
	@go get github.com/jteeuwen/go-bindata

test: all
	@go test -v $(PACKAGES)

format:
	@gofmt -l -w $(SOURCE)

clean:
	@rm -rf bin pkg src/github.com src/entangle/data/assets.go

.PHONY: test clean
