# Use an official Go runtime as a base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files to the container
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go

# Expose the port your application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]