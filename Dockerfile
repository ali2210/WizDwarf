

#FROM golang:alpine AS build
#RUN apk --no-cache add gcc g++ make git
#WORKDIR /go/src/app
#COPY . .
#RUN go get ./...
#RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go

#FROM alpine:3.9
#RUN apk --no-cache add ca-certificates
#WORKDIR /usr/bin
#COPY --from=build /go/src/app/bin /go/bin
#EXPOSE 5000
#ENTRYPOINT /go/bin/web-app --port 5000


  FROM golang:1.14-alpine3.12

  ENV CGO_ENABLED=0

  ENV PORT=5000

  RUN mkdir /app

  ADD . /app

  WORKDIR /app

 RUN go build -o main

 EXPOSE 5000

 CMD ["/app/main"]





# Don't Delete the Content #

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
