

# Docker Build ---[WIZDWARF]

FROM golang:latest as builder

VOLUME [${LOG_DIR}]

LABEL maintainer="Ali Hassan <Alideveloper95@protonmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get -u -d ./...

RUN apk add ca-certificates

RUN go build -o wiz

EXPOSE 9101

#	 ---End--- DockerFile 