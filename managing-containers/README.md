# Managing Images and Containers

Notes for the course on managing images and containers.

## Commands

Docker has a lot of useful commands which may be used to view information about
containers and images. Below are some that I've found useful in my workflow, but
running `docker help` in your terminal will give you all of the available options.

[Docker CLI Friendly Manual](https://docs.docker.com/reference/cli/docker/)

### `docker container`

```bash
Usage: docker container [OPTIONS]

Options:
    -a, --all
    -q, --quiet

Usage: docker container logs [OPTIONS]

Options:
    -f, --follow
    -n, --tail

Usage: docker container rm [OPTIONS] CONTAINER [CONTAINER...]

# Used to remove stopped containers

Options:
    -f, --force
```

### `docker images`

```bash
Usage: docker images [OPTIONS]

Options:
    -a, --all
    -q, --quiet
```

### `docker image inspect`

```bash
Usage: docker image inspect [OPTIONS] IMAGE [IMAGE...]

Options:
    -f, --format
```


### `docker image prune`

```bash
Usage: docker image prune [OPTIONS]

# Use to remove unused images

Options:
    -a, --all
```

### `docker build`

``` bash
Usage: docker build [OPTIONS] [Dockerfile] | [Path]

Options:
    -q, --quiet
    -t, --tag
```

### `docker push`

```bash
Usage: docker push [OPTIONS] REGISTRY[:TAG]
```

### `docker rmi`

``` bash
Usage: docker rmi [OPTIONS] IMAGE [IMAGE...]

# One or more image names may be passed to this command; `$(docker images -aq)` 
# outputs a list of all image IDs and can be used to remove all images not being
# used by a running container

Options:
    -f, --force
```

### `docker run`

``` bash
Usage: docker run [OPTIONS] [CONTAINER_NAME]

Options:
    --rm                automatically remove the container when it exits
    -d, --detached
    -i, --interactive
    -p, --publish
    -t, --tty
```
