# api-gateway Dockerfile

FROM golang:alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary, pointing to main.go in the cmd directory
RUN go build -o /api-gateway ./cmd/main.go

# Create a small final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /api-gateway .

# Expose the port on which the API Gateway listens
EXPOSE 8080

# Run the API Gateway binary
CMD ["./api-gateway"]
