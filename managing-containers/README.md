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
```

### `docker images`

```bash
Usage: docker images [OPTIONS]

Options:
    -a, --all
    -q, --quiet
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


### `docker run`

``` bash
Usage: docker run [OPTIONS] [CONTAINER_NAME]

Options:
    -d, --detached
    -i, --interactive
    -p, --publish
    -t, --tty
```
