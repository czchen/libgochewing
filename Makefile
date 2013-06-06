export GOPATH=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

.PHONY: all build test

all: build test

build:
	go build chewing

test:
	go test chewing
