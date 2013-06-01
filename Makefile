export GOPATH=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

.PHONY: all test

all: test

test:
	go test chewing
