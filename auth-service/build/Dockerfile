
FROM golang:1.23-alpine

WORKDIR /app

RUN set -x && apk update && \
    apk add -y --no-cache \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Hot Reload
RUN go install github.com/air-verse/air@latest

COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

CMD ["air", "-c", ".air.toml"]