# This codebase desgin according to mozilla open source license.
# Redistribution , contribution and improve codebase under license
# convensions. @contact Ali Hassan AliMatrixCode@protonmail.com

# base image 
FROM golang:1.19beta1-alpine3.15

# environment params  
ENV CGO_ENABLED=0

ENV PORT=5000

ENV HOST=wizdwarfs


# app workspace

RUN mkdir /app

ADD . /app

WORKDIR /app

# set persist storage
ARG WIZ_DIR=/app_data

RUN mkdir -p ${WIZ_DIR}

# declaration persistance storage
ENV WIZ_VOLUME_DIR=/app${WIZ_DIR}/apps.txt

# app modules 
COPY go.mod go.sum ./

RUN go mod tidy

#RUN go mod vendor

RUN go mod download

# build app 
RUN go build -o wizdwarfs 

# testing
RUN go test ./...

# && go test -v ./... 

# publish app port
EXPOSE 5000

# peristance storage 
VOLUME [ ${WIZ_DIR} ]

# certs
RUN apk --no-cache add ca-certificates

LABEL designed="Wisdom-Enigma Inc"

# initialization container
CMD ["/app/wizdwarfs"]
