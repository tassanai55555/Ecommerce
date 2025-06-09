# Build stage (static binary)
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ecommerce-backend ./cmd

# Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ecommerce-backend .

EXPOSE 3000

CMD ["./ecommerce-backend"]