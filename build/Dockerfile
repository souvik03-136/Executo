FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o executo ./cmd/main.go

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/executo .
CMD ["./executo"]
