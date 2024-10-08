# Docker Commands

## Create Network

```bash
docker network create goals-net
```

## Run MongoDB Container

```bash
docker run --name mongodb \
  -e MONGO_INITDB_ROOT_USERNAME=max \
  -e MONGO_INITDB_ROOT_PASSWORD=secret \
  -v data:/data/db \
  --rm \
  -d \
  --network goals-net \
  mongo
```

## Build Node API Image

```bash
docker build -t goals-node .
```

## Run Node API Container

```bash
docker run --name goals-backend \
  -e MONGODB_USERNAME=max \
  -e MONGODB_PASSWORD=secret \
  -v logs:/app/logs \
  -v /Users/maximilianschwarzmuller/development/teaching/udemy/docker-complete/backend:/app \
  -v /app/node_modules \
  --rm \
  -d \
  --network goals-net \
  -p 80:80 \
  goals-node
```

## Build React SPA Image

```bash
docker build -t goals-react .
```

## Run React SPA Container

```bash
docker run --name goals-frontend \
  -v /Users/maximilianschwarzmuller/development/teaching/udemy/docker-complete/frontend/src:/app/src \
  --rm \
  -d \
  -p 3000:3000 \
  -it \
  goals-react
```

## Stop all Containers

```bash
docker stop mongodb goals-backend goals-frontend
```
