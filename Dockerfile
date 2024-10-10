# Use the official Golang image to build the Go application
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Expose necessary ports
# Port for API Gateway
EXPOSE 8080

# Port for auth-service
EXPOSE 8081

# Port for users-service
EXPOSE 8082

# Default environment variable for the service to run
ENV SERVICE=api-gateway

# Build and run the desired service based on the SERVICE environment variable
CMD go run util/cli/main.go ${SERVICE}
