# ---- Stage 1: Builder ----
FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

RUN chmod +x scripts/build-application.sh
RUN ./scripts/build-application.sh

# ---- Stage 2: Runner ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ordersystem /app/ordersystem

EXPOSE 8080

CMD ["/app/ordersystem"]