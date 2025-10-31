# ---------- Stage 1: Build ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd

# ---------- Stage 2: Run ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8001

CMD ["./main"]
