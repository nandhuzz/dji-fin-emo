# Start from the official Golang base image
FROM golang:1.22

# Set working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]

# Expose the port used by the Go server
EXPOSE 8080
