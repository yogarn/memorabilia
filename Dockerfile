# Use the official Go image as the base image
FROM golang:1.22.0

# Set the working directory in the container
WORKDIR /cmd/app

# Copy the application files into the working directory
COPY . /cmd/app

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./main"]