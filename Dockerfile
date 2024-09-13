FROM golang:1.22.3-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/main.go

FROM alpine:latest
RUN apk add --no-cache nats-server
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
CMD ["sh", "-c", "nats-server & ./main"]