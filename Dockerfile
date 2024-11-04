# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd || { echo 'Build failed'; exit 1; }

# Run stage
FROM debian:bullseye-slim

WORKDIR /root/

COPY --from=builder /app/main .


COPY .env .env

EXPOSE 3000

CMD ["./main"]
