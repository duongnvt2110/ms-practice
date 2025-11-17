FROM golang:1.23.6-alpine

WORKDIR /app

RUN set -x && apk update && \
    apk add --no-cache ca-certificates bash && \
    rm -rf /var/cache/apk/*

RUN go install github.com/air-verse/air@v1.52.3

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/ticket-service

CMD ["air", "-c", ".air.toml"]
