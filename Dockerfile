

FROM golang:1.14-alpine3.12

ENV CGO_ENABLED=0

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main

CMD ["/app/main"]






#ENV GOFLAGS= mod=vendor

#LABEL maintainer="Ali Hassan <Alideveloper95@protonmail.com>"

#RUN mkdir /app

#ADD . /app

#ENV CGO_ENABLED=0

#ENV GO111MODULE=off

#COPY go.mod go.sum ./

#RUN go mod download

#COPY . .

#RUN go get -u -d ./...

   #EXPOSE 9101

#RUN go build -o main

