# ==== STAGE 1: Build the Go app ====
FROM golang:1.26.3-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

# ==== STAGE 2: Build a small image ====
FROM alpine:3.22

COPY --from=builder /app/main ./

COPY --from=builder /app/.env ./

EXPOSE 8080

CMD ["/main"]