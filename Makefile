.PHONY: all build test

VERSION=$(shell bin/version)

all: build

build:
	go build -ldflags "-X=main.Version=$(VERSION)" -i -o codestand

run:
	go run -ldflags "-X=main.Version=$(VERSION)" ./*.go ${ARGS}
