# --- STAGE 1: Build Stage ---
FROM golang:1.26.3-alpine AS builder

WORKDIR /app

# Install ca-certificates so we can copy them later
RUN apk add --no-cache ca-certificates

COPY go.mod main.go ./

COPY pkg pkg

# Create a non-root user (UID 10001 is a safe non-system choice)
RUN adduser -D -g '' -u 10001 appuser

# Build a static binary for the proxy
RUN CGO_ENABLED=0 go build -v -o /proxy .

# --- STAGE 2: Final Production Stage ---
FROM scratch AS final

# Copy the user definition from the builder stage
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy certificates from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy only the compiled binary from the builder stage
COPY --from=builder /proxy /proxy

EXPOSE 8080

# Switch to the non-root user
USER appuser

ENTRYPOINT ["/proxy"]
