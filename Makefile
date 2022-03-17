SHELL = /bin/sh

.PHONY: build
build:
	PWD=$$(PWD) user=$$(id -u) group=$$(id -g) entrypoint="" docker-compose --file ./docker-compose.yml up --build golang.test

.PHONY: install
install:
	PWD=$$(PWD) user=$$(id -u) group=$$(id -g) entrypoint="go mod tidy" docker-compose --file docker-compose.yml up golang.test

.PHONY: run
run:
	PWD=$$(PWD) user=$$(id -u) group=$$(id -g) entrypoint="go run ./cmd/main.go" docker-compose --file docker-compose.yml up golang.test