FROM golang:1.23.6-alpine

WORKDIR /app

RUN set -x && apk update && \
    apk add --no-cache ca-certificates curl bash && \
    rm -rf /var/cache/apk/*

# Copy local code to the container image.
COPY go.mod go.sum ./

# Hot Reload
RUN go install github.com/air-verse/air@v1.52.3

# DB mate 
RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
RUN chmod +x /usr/local/bin/dbmate

# Download dependencies
RUN go mod download

COPY . ./

WORKDIR /app/retry-management-service

CMD ["air", "-c", ".air.toml"]
