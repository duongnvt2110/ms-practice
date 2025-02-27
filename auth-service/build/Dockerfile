
FROM golang:1.23-alpine

WORKDIR /app/auth-service

RUN set -x && apk update && \
    apk add -y --no-cache \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Hot Reload
RUN go install github.com/air-verse/air@latest

# Copy go.work file and modules
COPY ../pkg ../pkg
COPY ../auth-service/go.* .

# Download dependencies
RUN go mod download

# Copy local code to the container image.
COPY ../auth-service/. ./

CMD ["air", "-c", ".air.toml"]