# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./
# Download dependencies (go.sum will be generated on the fly if missing, or we can ignore it for this phase)
RUN go mod tidy

# Copy source code
COPY . .

# Build the binary
RUN go build -o server cmd/valkey-inspector/main.go

# Run Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
