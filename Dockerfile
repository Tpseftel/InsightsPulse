# Stage 1: Build the Go application
FROM golang:1.22.1-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Copy the built application into a minimal base image
FROM alpine:latest

# Install MariaDB client tools
RUN apk --no-cache add mariadb-client

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

# Copy the wait-for-it.sh script and set executable permissions
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Expose port if your application listens on a specific port
# INFO: I dont use this because 
EXPOSE 8080


# Command to run the executable
CMD ["/wait-for-it.sh", "mariadb:3306", "--", "./main"]
