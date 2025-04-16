#!/bin/bash
# This script is used to run a Docker container for a Go API application.
# It mounts the source code directory into the container and exposes port 8080.

# Linux command to run the Docker container
docker run -it --rm -v "$(pwd)/src:/app/go-api/src" -p 8080:8080 go-api-dev

# PowerShell command to run the Docker container
# docker run -it --rm -v "$(PWD)\src:/app/go-api/src" -p 8080:8080 go-api-dev

# CMD command to run the Docker container
# docker run -it --rm -v "%cd%/src:/app/go-api/src" -p 8080:8080 go-api-dev
