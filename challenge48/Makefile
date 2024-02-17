# Define variables
IMAGE_NAME := data-privacy-vault
CONTAINER_NAME := data-privacy-vault-container
REDIS_CONTAINER_NAME := redis-container
REDIS_PORT := 6379
APP_PORT := 8080
NETWORK := data-privacy-vault-network

# Default target
.PHONY: start
start: network-create docker-build redis-start run

# Clean up and run the application
.PHONY: restart
restart: 
	@echo "Restarting..."
	clean start

# Clean up
.PHONY: clean
clean: stop redis-stop network-remove
	@echo "Cleaning up..."
	rm -f data-privacy-vault

#Create docker network
.PHONY: network-create
network-create:
	@echo "Creating network..."
	docker network create $(NETWORK)

#Remove docker network
.PHONY: network-remove
network-remove:
	@echo "Removing network..."
	docker network rm $(NETWORK)

# Build the Docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .

# Start the Redis container
.PHONY: redis-start
redis-start:
	@echo "Starting Redis container..."
	docker run --name $(REDIS_CONTAINER_NAME) -d --network $(NETWORK) -p $(REDIS_PORT):$(REDIS_PORT) redis

# Run the Docker container
.PHONY: run
run:
	@echo "Running Docker container..."
	docker run --name $(CONTAINER_NAME) -d --network $(NETWORK) -p $(APP_PORT):$(APP_PORT) $(IMAGE_NAME)

# Stop and remove the Docker container
.PHONY: stop
stop:
	@echo "Stopping and removing Docker container..."
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

# Stop and remove the Redis container
.PHONY: redis-stop
redis-stop:
	@echo "Stopping and removing Redis container..."
	docker stop $(REDIS_CONTAINER_NAME)
	docker rm $(REDIS_CONTAINER_NAME)

# Clean up
.PHONY: clean
clean: stop redis-stop
	@echo "Cleaning up..."
	rm -f data-privacy-vault
