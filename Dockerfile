# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-stock .

# Final runtime stage
FROM alpine:latest

WORKDIR /app

# Install tzdata and other minimal packages in runtime container
RUN apk add --no-cache curl ca-certificates tzdata

# Set timezone to Asia/Jakarta
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone

# Copy binary from builder
COPY --from=builder /app/go-stock .

EXPOSE 3000

CMD ["./go-stock"]
