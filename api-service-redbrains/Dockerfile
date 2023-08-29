# Use an official Go runtime as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o app

# Run the Go application when the container starts
CMD ["./app"]