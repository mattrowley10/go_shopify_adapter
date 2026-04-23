#!/bin/bash

# Load .env file
if [ -f .env ]; then 
    echo "Loading .env file..."
    set -o allexport
    source .env 
    set +o allexport
else 
    echo ".env file not found!"
fi 

go run ./cmd/api