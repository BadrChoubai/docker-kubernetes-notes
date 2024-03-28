#!/bin/bash

echo "Testing node-app"

pushd node-app || exit
docker build -t node-app .
docker run -d --name node-app-local -p 8080:3000 node-app
popd || exit

echo "Testing python-app"

pushd python-app || exit
docker build -t python-app .
docker run -it --name python-app-local python-app
popd || exit

echo "Cleaning up..."

docker container stop node-app-local
docker container stop python-app-local
docker container rm node-app-local python-app-local
docker image rm node-app python-app