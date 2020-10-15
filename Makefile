.PHONY: all test build fmt run

Binary="BinaryTree"

all: build fmt run

build:
	@go build -o ${Binary} .

fmt:
	@go fmt ./*

run:
	@./${Binary}

test:
	@go test ./test/Binary_test.go -v