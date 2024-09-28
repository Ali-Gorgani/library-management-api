# books-service Dockerfile

FROM golang:alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary, pointing to main.go in the cmd directory
RUN go build -o /books-service ./cmd/main.go

# Create a small final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /books-service .

# Expose the port on which the Books Service listens
EXPOSE 8081

# Run the Books Service binary
CMD ["./books-service"]
