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

## Command-Line Management of Volumes

```bash
docker volumes --help
```
