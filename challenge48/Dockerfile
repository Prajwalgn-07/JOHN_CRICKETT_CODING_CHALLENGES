# Use the specified Golang version
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY src/go.mod src/go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy all source files to the working directory
COPY src/ ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /data-privacy-vault

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/data-privacy-vault"]
