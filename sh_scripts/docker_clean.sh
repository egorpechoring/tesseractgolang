#!/bin/sh

# Stop and remove all running containers with a name starting with "golangtesseract_"
docker stop $(docker ps -q --filter "name=golangtesseract_" --no-trunc)
docker rm $(docker ps -aq --filter "name=golangtesseract_" --no-trunc)

# Remove any dangling images (not tagged and not used by any container)
docker image prune

# Remove all images not used by any container
docker image prune -a

# Remove all Docker volumes not used by any container
docker volume prune

# Remove all Docker networks not used by any container
docker network prune
