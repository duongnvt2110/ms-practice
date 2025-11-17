
FROM golang:1.23.6-alpine

WORKDIR /app

RUN set -x && apk update && \
    apk add -y --no-cache \
    ca-certificates bash && \
    rm -rf /var/lib/apt/lists/*

# Hot Reload
RUN go install github.com/air-verse/air@v1.52.3  

COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

EXPOSE 3000

WORKDIR /app/booking-service

CMD ["air", "-c", ".air.toml"]