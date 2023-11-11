FROM golang:1.21-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update
RUN apk add make git
COPY go.mod go.sum ./
COPY . ${ROOT}

RUN go mod download

CMD ["go", "run", "main.go", "server"]

FROM golang:1.21-alpine as builder
ENV BINARY="aws-proxy-app"
ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update
RUN apk add make git

COPY go.mod go.sum ./
COPY . ${ROOT}

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux make build

FROM scratch as prod
ENV BINARY="aws-proxy-app"
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/${BINARY} ${ROOT}

EXPOSE 3000
CMD ["/go/src/app/aws-proxy-app"]