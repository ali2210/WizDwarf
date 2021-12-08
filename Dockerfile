FROM golang:1.17.4-alpine3.14

ENV CGO_ENABLED=0

ENV PORT=5000

ENV HOST=wizdwarfs

RUN mkdir /app

ADD . /app

WORKDIR /app

ARG WIZ_DIR=/app_data

RUN mkdir -p ${WIZ_DIR}

ENV WIZ_VOLUME_DIR=/app${WIZ_DIR}/apps.txt

COPY go.mod go.sum ./

RUN go mod tidy

#RUN go mod vendor

RUN go mod download

RUN go build -o wizdwarfs 

RUN go test ./...

# && go test -v ./... 
EXPOSE 5000

VOLUME [ ${WIZ_DIR} ]

RUN apk --no-cache add ca-certificates

LABEL companyRelease="Wisdom-Enigma"

CMD ["/app/wizdwarfs"]

