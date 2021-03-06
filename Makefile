VERSION=$(shell bin/version)

.PHONY: all build

all: build

build:
	go build -ldflags "-X=main.Version=$(VERSION)" -i -o codestand

init:
	go get -u github.com/kardianos/govendor
	govendor sync
