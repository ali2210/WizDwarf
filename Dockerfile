

FROM golang:1.14-alpine3.12

#ENV GO111MODULE=on

#ENV GOFLAGS= mod=vendor

#LABEL maintainer="Ali Hassan <Alideveloper95@protonmail.com>"

RUN mkdir /app

ADD . /app

WORKDIR /app

#COPY go.mod go.sum ./

#RUN go mod download

#COPY . .

#RUN go get -u -d ./...

   #EXPOSE 9101

RUN go build -o main

CMD ["/app/main"]
