# Start from the official Golang base image
FROM golang:1.21-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /app

# Install Git (for private module access, if needed) and necessary build tools
RUN apk add --no-cache git

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN go build -o finance-api ./cmd/server

# Create a minimal image for running
FROM alpine:latest

# Install ca-certificates for HTTPS support
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/finance-api .

# Expose port (change if needed)
EXPOSE 8080

# Command to run
CMD ["./finance-api"]
