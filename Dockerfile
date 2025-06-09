# Use official Golang image as the build stage
FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ecommerce-backend ./cmd

# Final image
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/ecommerce-backend .

EXPOSE 3000

CMD ["./ecommerce-backend"]