
FROM golang:1.23-alpine

WORKDIR /app

RUN set -x && apk update && \
    apk add -y --no-cache \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Hot Reload
RUN go install github.com/air-verse/air@latest

# Copy local code to the container image.
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . ./

# Set working directory to build
WORKDIR /app/auth-service

CMD ["air", "-c", ".air.toml"]