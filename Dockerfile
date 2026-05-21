# --- STAGE 1: Build Stage ---
FROM golang:1.26.3-alpine AS builder

WORKDIR /app

# Install ca-certificates so we can copy them later
RUN apk add --no-cache ca-certificates

COPY go.mod main.go ./

COPY pkg pkg

# Build a static binary for the proxy
RUN CGO_ENABLED=0 go build -v -o /proxy .

# --- STAGE 2: Final Production Stage ---
FROM scratch AS final

# Copy certificates from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy only the compiled binary from the builder stage
COPY --from=builder /proxy /proxy

EXPOSE 8080

ENTRYPOINT ["/proxy"]
