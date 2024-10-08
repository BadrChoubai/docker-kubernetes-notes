docker build -t docker-tutorial-volumes:prod $1
docker build -t docker-tutorial-volumes:dev --build-arg DEFAULT_PORT=8080 $1

