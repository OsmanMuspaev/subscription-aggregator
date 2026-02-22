FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o subscription-service cmd/server/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/subscription-service .

COPY .env .

EXPOSE 8080

CMD ["./subscription-service"]