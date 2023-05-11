#!/bin/sh

clear
set -e # Stop the script on error

# Build the Docker image
docker build -t golangtesseract .

# Run the Docker container and map the port
docker run -it -p 8080:8080 --name golangtesseract_container golangtesseract