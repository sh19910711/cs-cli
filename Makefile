VERSION=$(shell bin/version)

.PHONY: all build test

all: build

build:
	go build -ldflags "-X=main.Version=$(VERSION)" -i -o codestand

run:
	go run -ldflags "-X=main.Version=$(VERSION)" ./*.go ${ARGS}
