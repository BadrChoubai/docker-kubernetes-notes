# Define variables
IMAGE_NAME := dockerfile-intro
CONTAINER_NAME := dockerfile-intro-container

# Display help message
help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  build       Build the Docker image"
	@echo "  run         Run the Docker container"
	@echo "  clean       Remove the Docker container and image"

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run:
	docker run --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

# Remove the Docker container and image
clean:
	docker rm -f $(CONTAINER_NAME) || true
	docker rmi -f $(IMAGE_NAME) || true

.PHONY: help build run clean
