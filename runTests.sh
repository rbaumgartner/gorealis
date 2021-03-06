#!/bin/bash

docker-compose up -d

# If running docker-compose up gives any error, don't do anything.
if [ $? -ne 0 ]; then
    exit
fi

# Since we run our docker compose setup in bridge mode to be able to run on MacOS, we have to launch a Docker container within the bridge network in order to avoid any routing issues.
docker run --rm -t -v $(pwd):/go/src/github.com/paypal/gorealis --network gorealis_aurora_cluster golang:1.10-stretch go test -v github.com/paypal/gorealis $@

docker-compose down
