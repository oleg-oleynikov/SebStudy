FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/certs ./certs/
COPY --from=builder /app/.env .

CMD ["sh", "-c",  "./main"]
