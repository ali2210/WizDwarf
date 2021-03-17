
FROM golang:1.15-alpine3.13

ENV CGO_ENABLED=0

ENV PORT=5000

ENV HOST=wizdwarfs

# ENV CoinbaseKey=uGJWOhYrm7X2njjC

# ENV CoinbaseSecret=U3D0pf9uwDGMAniaFyV17t2cd2ODHwVc

RUN mkdir /app

ADD . /app

WORKDIR /app

ARG WIZ_DIR=/seqDir

RUN mkdir -p ${WIZ_DIR}

ENV WIZ_VOLUME_DIR=/app${WIZ_DIR}/seqFile.txt

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o main && go test -v ./... 

EXPOSE 5000

VOLUME [ ${WIZ_DIR} ]

RUN apk --no-cache add ca-certificates

LABEL companyRelease="Wisdom-Enigma"

CMD ["/app/main"]





