#!/bin/bash

# Use environment var PET_STORE_PORT if set.
PORT=${PORT:-9000}
ACCESS_KEY=${PET_STORE_ACCESS_KEY:-"PetStoreAccessKey"}

if [[ "$1" == "build" ]]; then
    docker build -t pet-store .
else
    NAME="pet_store_container"
    docker rm -f ${NAME}
    docker run --name ${NAME} -p ${PORT}:${PORT} -e PET_STORE_ACCESS_KEY=${ACCESS_KEY} -e PORT=":${PORT}" -d pet-store
fi
