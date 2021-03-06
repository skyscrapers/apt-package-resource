GOVERSION=$(shell go version | awk '{print $$3;}')
VERSION=$(shell git describe --tags | sed 's@^v@@' | sed 's@-@+@g')
TESTS?=
BINPATH?=$(GOPATH)/bin

all: test check install

dev:
	go get -u github.com/laher/goxc

prepare:
	go get -u github.com/mattn/goveralls
	go get -u github.com/axw/gocov/gocov
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

check:
	gometalinter --config=linter.json ./...

install:
	go install -v -ldflags "-X main.Version=$(VERSION)"

test:
	go test -v `go list ./... | grep -v vendor/`

goxc:
	rm -rf root/
	mkdir -p root/opt/concourse
	goxc -pv=$(VERSION) -max-processors=4 $(GOXC_OPTS)

version:
	@echo $(VERSION)

.PHONY: coverage.out version