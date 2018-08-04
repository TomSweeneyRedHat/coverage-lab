PACKAGES := \
	github.com/coveragelab \

all: build test

build:
	go build -o bin/coveragelab hello.go

test:
	go test -v $(PACKAGES)
