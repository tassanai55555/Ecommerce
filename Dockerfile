# Use official Golang image as the builder
FROM golang:1.24.3 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Use a minimal base image to run the app
FROM debian:bookworm-slim

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 3000

# Run the app
CMD ["./main"]
