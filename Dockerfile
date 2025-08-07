# Start from the official Golang base image
FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and migrations
COPY . .
COPY migrations ./migrations

# Build Go app
RUN go build -o finance-api ./cmd/server

# Final runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary and migrations from builder stage
COPY --from=builder /app/finance-api .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./finance-api"]

