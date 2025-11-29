FROM golang:1.23.6-alpine

WORKDIR /app

RUN set -x && apk update && \
    apk add --no-cache ca-certificates curl bash && \
    rm -rf /var/cache/apk/*

RUN go install github.com/air-verse/air@v1.52.3

# Install goose
RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh
RUN chmod +x /usr/local/bin/goose

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/event-service

CMD ["air", "-c", ".air.toml"]
