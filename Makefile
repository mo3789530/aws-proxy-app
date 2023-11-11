BINARY=aws-app-proxy

default: build


build:
	go build -o ${BINARY} .

run: build
	./${BINARY} server

build-docker-dev:
	docker build -t aws-proxy-app-dev --target=dev  .

run-docker: build-docker-dev
	dokcer run -d --name aws-proxy-app aws-proxy-app-dev:latest

install:
	go mod tidy


run-dev:
	go run main.go server