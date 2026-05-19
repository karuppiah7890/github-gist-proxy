# --- STAGE 1: Build Stage ---
FROM golang:1.26.3-alpine AS builder

WORKDIR /app

COPY go.mod main.go ./

# Build a static binary for the proxy
RUN CGO_ENABLED=0 go build -v -o /proxy .

# --- STAGE 2: Final Production Stage ---
FROM scratch AS final

# Copy only the compiled binary from the builder stage
COPY --from=builder /proxy /proxy

EXPOSE 8080

ENTRYPOINT ["/proxy"]
