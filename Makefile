export GOPATH=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

test:
	go test chewing
