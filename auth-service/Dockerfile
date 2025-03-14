FROM golang:1.24.1 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/auth-service .
COPY migrations /migrations

CMD ["./auth-service"]