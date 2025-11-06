# -------- Builder Container --------
FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build the Go application with optimizations for a small binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /app/ordersystem

# -------- Runtime Container --------
FROM alpine:3.20
# Create user without home directory and no root privileges
RUN adduser -D -H app
USER app
WORKDIR /app
# Copy the compiled binary from the builder stage
COPY --from=builder /app/ordersystem /app/ordersystem
# Expose application port and define default command
EXPOSE 3000
CMD ["/app/ordersystem"]

# .env file is also present, but since it's in gitignore, it's not included in the repo.