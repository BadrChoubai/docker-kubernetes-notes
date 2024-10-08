# Create network
docker network create networked-app--network

# Pull mongo image
docker pull mongo

# Run mongo image
docker run --network networked-app--network -d --name mongodb --rm mongo

# Build the Dockerfile
docker build -t networked-app-example .

# Run our app, exposing a port
export APP_PORT=8080
docker run --rm --network networked-app--network -p "${APP_PORT}:3000" --name networked-app networked-app-example
