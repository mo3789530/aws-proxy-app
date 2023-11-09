BINARY=aws-app-proxy

defualt: build

build: 
	go build -o ${BINARY} .