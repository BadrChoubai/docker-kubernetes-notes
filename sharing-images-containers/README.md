# Sharing Images and Containers 

## Docker Registry (Container Registry)

A container registry is a service used to manage the storage and distribution of
container images.

Inside of the Docker ecosystem, you would use DockerHub. But, you may also create
a private registry using the official distribution image:

[GitHub](https://github.com/distribution/distribution-library-image)

### Self-Hosting the Container Registry

Inside of this directory there is a folder called `registry`, which contains a
`docker-compose.yaml` file which may be used to start a container on your local
machine to be able to push images to.

```bash
cd registry

docker compose up -d

docker container logs docker-course-local-registry -n 1

# To stop the running container
docker compose down
```

- The Friendly Manual: [Deploy a Registry Server](https://distribution.github.io/distribution/about/deploying/)
