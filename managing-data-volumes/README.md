# Managing Data and Volumes

[Volumes Documentation](https://docs.docker.com/storage/volumes/)

Volumes are folders on your host machine which are mounted ("made available")
into containers. Some benefits of volumes are that they will persist between
container runs and containers can both read and write to them.

In a docker container we can consider two types of data we may need to manage:

1. Temporary application data: [Stored with containers]
    - Fetched and produced in the running container
    - Stored in-memory or by temporary files
    -  Dynamic and changing, but cleared regularly
2. Permanent application data: [Stored with containers and volumes]
    - Fetched and produced in the running container
    - Stored in files or a database
    - Must not be lost if container stops or restarts

## Volumes Comparison

1. Anonymous Volumes:
    - Createed specifically for a single container
    - Survives container shutdown or restart unless `--rm` flag is used
    - Can not be shared across containers
    - Anonymous, hence cannot be reused
2. Named Volumes:
    - Created in general &mdash; not tied to any specific container
    - Survives container shutdown or restart
    - Can be shared across containers
    - Can be re-used for same container across restarts
3. Bind Mounts:
    - Location on host file system, not tied to any specific container
    - Survives container shutdown or restart and removal on hosts
    - Can be shared across containers
    - Can be re-used for same container across restarts

## Command-Line Management of Volumes

```bash
Usage:  docker volume COMMAND

Manage volumes

Commands:
  create      Create a volume
  inspect     Display detailed information on one or more volumes
  ls          List volumes
  prune       Remove unused local volumes
  rm          Remove one or more volumes

Run 'docker volume COMMAND --help' for more information on a command.
```
