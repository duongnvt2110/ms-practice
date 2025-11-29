
FROM golang:1.23-alpine

WORKDIR /app

RUN set -x && apk update && \
    apk add -y --no-cache \
    ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*

# Hot Reload
RUN go install github.com/air-verse/air@v1.52.3

# DB mate 
RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
RUN chmod +x /usr/local/bin/dbmate

# Copy local code to the container image.
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . ./

# Set working directory to build
WORKDIR /app/auth-service

CMD ["air", "-c", ".air.toml"]