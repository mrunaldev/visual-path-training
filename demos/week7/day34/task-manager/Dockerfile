# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /task-manager ./cmd/server

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Create non-root user
RUN adduser -D appuser

# Set working directory
WORKDIR /app

# Copy binary and config files from builder
COPY --from=builder /task-manager .
COPY .env.example .env

# Set ownership and permissions
RUN chown -R appuser:appuser /app
USER appuser

# Expose port
EXPOSE 8080

# Command to run
CMD ["./task-manager"]
