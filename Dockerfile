# Start from a base image with Golang already installed
FROM golang:latest

# Install Tesseract and its dependencies
RUN apt-get update && \
    apt-get install -y libleptonica-dev libtesseract-dev tesseract-ocr

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Install any needed packages specified in go.mod
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
