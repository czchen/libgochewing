# Introduction
[![Build Status](https://drone.io/github.com/czchen/libgochewing/status.png)](https://drone.io/github.com/czchen/libgochewing/latest)
[![Build Status](https://travis-ci.org/czchen/libgochewing.png)](https://travis-ci.org/czchen/libgochewing)

This project reimplement [libchewing](https://github.com/chewing/libchewing/) in [golang](http://golang.org/).

# Development
See [How to Write Go Code](http://golang.org/doc/code.html), or using command `go help gopath` to setup the development environment.

If your environment does not have golang, or the version of golang is too old, you can install it by [gvm](https://github.com/moovweb/gvm).

## Unit Test
The following cmomand run unit test of this project.

    go test

## Benchmark
The following command run benchmark of this project.

    go test -bench .

## Coverage
The following commands create summary coverage report for unit test.

    go get github.com/axw/gocov/gocov
    bin/gocov test github.com/czchen/libgochewing | bin/gocov report

The coverage report can also be generated as HTML with the following commands:

    go get github.com/axw/gocov/gocov
    go get github.com/matm/gocov-html
    bin/gocov test github.com/czchen/libgochewing | bin/gocov-html > coverage.html

# License
This project is licensed under [LGPL-2](http://www.gnu.org/licenses/old-licenses/lgpl-2.0.en.html).
