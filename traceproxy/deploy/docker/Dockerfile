# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install ca-certificates for HTTPS/gRPC
RUN apk add --no-cache ca-certificates

# Copy source code
COPY . .

WORKDIR /app/traceproxy

RUN go mod download

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o encoredev-traceproxy .

# Final stage
FROM scratch

# Copy CA certificates for gRPC/HTTPS connections
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary
COPY --from=builder /app/traceproxy/encoredev-traceproxy /traceproxy

ENTRYPOINT ["/traceproxy"]
