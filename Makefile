.PHONY: run test

VERSION=$(shell bin/version)

build: *.go
	go build -ldflags "-X=main.Version=$(VERSION)" -i -o codestand

run:
	go run -ldflags "-X=main.Version=$(VERSION)" ./*.go ${ARGS}
