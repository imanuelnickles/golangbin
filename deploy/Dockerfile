## Builder
FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY . .

RUN apk add --no-cache

RUN GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/api app.go && chmod +x /usr/local/bin/api

## Deploy
FROM alpine:3.17

WORKDIR /

COPY --from=builder /usr/local/bin/api /usr/local/bin/api

ENTRYPOINT ["api"]
