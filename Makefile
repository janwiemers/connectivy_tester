.PHONY: default
default: test;

build:
	go mod tidy
	go build

test:
	go test

build_docker:
	docker build -t "connectivity_tester" .