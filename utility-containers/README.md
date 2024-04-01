# Utility Containers

Utility containers set up using Docker may be used for running commands that would 
otherwise not be available on the host machine.


## Using `docker run | build`

```bash
docker build -t npm-util .
docker run -it -v $(pwd):/app npm-util
```

## Using `docker compose`

Create a simple `docker-compose.yaml`

```yaml
services:
  npm:
    build: ./
    stdin_open: true
    tty: true
    volumes:
      - ./:/app
```

Use `docker compose run` to execute commands inside of the created container

```bash
docker compose run --rm npm init -y
```

