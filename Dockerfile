# ---------- Stage 1: Build ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN go build -o main .

# ---------- Stage 2: Run ----------
FROM alpine:latest

WORKDIR /app

# Copy binary hasil build
COPY --from=builder /app/main .

# Expose port sesuai aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
