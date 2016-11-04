SHELL := /bin/bash
BINARY=dname

VERSION=0.0.1
BUILD_TIME=`date +%FT%T%z`

BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`

LDFLAGS="-X ${BINARY}.version ${VERSION} -X ${BINARY}.buildtime ${BUILD_TIME} -X ${BINARY}.branch ${BRANCH} -X ${BINARY}.commit ${COMMIT}"

GLIDE := $(shell glide --version 2>/dev/null)

default: build

deps:
ifdef GLIDE
	@glide install
else
	@echo "glide is not installed."
endif

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

pretest:
	@gofmt -d $$(find . -type f -name '*.go' -not -path "./vendor/*") 2>&1 | read; [ $$? == 1 ]

lint-test:
	@go get -v github.com/golang/lint/golint
	@golint ./... | grep -v vendor/ 2>&1 | read; [ $$? == 1 ]

vet:
	@go vet $(go list -f '{{ .ImportPath }}' ./... | grep -v vendor/)

test: pretest vet lint-test
	@go test -v $$(go list -f '{{ .ImportPath }}' ./... | grep -v vendor/) -p=1

build: test clean deps
	@go build -x -ldflags ${LDFLAGS} -o ${BINARY}

fmt:
	@gofmt -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

lint:
	@go get -v github.com/golang/lint/golint
	@golint ./... | grep -v vendor/ | true

list:
	@go list -f '{{ .ImportPath }}' ./... | grep -v vendor/

install:
	@go install github.com/umayr/dname/cmd/dname