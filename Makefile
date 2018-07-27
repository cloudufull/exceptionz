.PHONY: all api bench install test

all: test

api:
	@protoc -I ./api/ api/*.proto --go_out=plugins=grpc:api

bench:
	@go test -bench=. ./...

install:
	@go get -u github.com/andygeiss/assert

test: install
	@go test -v ./...
