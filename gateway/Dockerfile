FROM golang:1.23.5 AS builder


WORKDIR /app
COPY . .

# Static build for Linux (required for distroless)
RUN cd cmd/gateway && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway

# --- Stage 2: Minimal runtime container ---
FROM gcr.io/distroless/static

# Copy binary and templates
COPY --from=builder /app/cmd/gateway/gateway /gateway
COPY internal/templates /templates
COPY .env /.env

ENTRYPOINT ["/gateway"]