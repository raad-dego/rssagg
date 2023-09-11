# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /rssagg

# Copy the Go project files into the container
COPY . .

# Build the Go application for a Linux target
RUN GOOS=linux GOARCH=amd64 go build -o rssagg

# Expose the port that your application listens on
EXPOSE 8080

# Define the command to run your application
CMD ["./rssagg"]
