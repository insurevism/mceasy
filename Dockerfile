# Use official Go 1.19 Alpine image instead of AWS ECR custom image
FROM golang:1.19-alpine AS builder

# Set environment variables
ENV OPENAPI_ENTRY_POINT="cmd/main.go"
ENV OPENAPI_OUTPUT_DIR="cmd/docs"

WORKDIR /app

# Copy source code
COPY . /app

# Install dependencies
RUN apk add --no-cache make

# Install swag CLI for generating Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation
RUN make swagger

# Build the Go application
RUN make build

# ---- Runtime Stage ----
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates curl

# Set working directory
WORKDIR /app

# Copy built binary from builder stage
COPY --from=builder /app/main /app/main

# Copy environment files into the container (or mount at runtime)
COPY .env /app/

# Expose port
EXPOSE 8888

# Run the application
ENTRYPOINT [ "/app/main" ]
