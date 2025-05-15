# Build stage
FROM golang:1.24.2 AS builder

WORKDIR /app

# Copy module files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/gaggle ./backend/cmd/Gaggle

# Runtime stage
FROM alpine:latest

# Install SSL certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -S gaggle && adduser -S gaggle -G gaggle

WORKDIR /home/gaggle

# Copy binary from builder
COPY --from=builder --chown=gaggle:gaggle /bin/gaggle ./

# Switch to non-root user
USER gaggle

EXPOSE 8080

CMD ["./gaggle"]
