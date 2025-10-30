#!/bin/bash

# Start db docker image
# I've created a volume called pgdata18 to persist the data
docker build . -t orderservice

# Remove old container if it exists
CONTAINER_NAME="postgres18"
if [ "$(docker ps -aq -f name=^${CONTAINER_NAME}$)" ]; then
  echo "Removing existing container: ${CONTAINER_NAME}"
  docker rm -f "${CONTAINER_NAME}"
fi

docker run -d \
  --name postgres18 \
  --env-file debug.env \
  -e PGDATA=/var/lib/postgresql/18/docker \
  -v pgdata18:/var/lib/postgresql/18/docker \
  -p 5432:5432 \
  --network host \
  postgres:18


# Wait for the database to start
sleep 5

# Remove old container if it exists
CONTAINER_NAME="orderservice"
if [ "$(docker ps -aq -f name=^${CONTAINER_NAME}$)" ]; then
  echo "Removing existing container: ${CONTAINER_NAME}"
  docker rm -f "${CONTAINER_NAME}"
fi

docker run -d \
  --name orderservice \
  --env-file debug.env \
  -p 3000:3000 \
  --network host \
  orderservice:latest
