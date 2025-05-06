# Tooling for Kubernetes

For working with Kubernetes locally, we first need to install the Kubernetes
tools: https://kubernetes.io/docs/tasks/tools/,
for the course I only set up `kubectl` and `minikube`.

**Jump To:**

- [Installing `kubectl`](#installing-kubectl)
- [Installing `minikube`](#installing-minikube)
- [Local Container Registry (Optional)](#local-container-registry-optional)

## Installing `kubectl`

Directions for installing and setting up `kubectl` for Linux can be found here:
[https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/).

Alternatively, you may use Snap or Homebrew to install the package as well:

**Using** `snap`:

```shell
snap install kubectl --classic
kubectl version --client
```

**Using** `brew`:

```shell
brew install kubectl
kubectl version --client
```

## Installing `minikube`

Directions for installing and setting up `minikube` for Linux can be found here:
[https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download](https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download)

Once you've got both `kubectl` and `minikube` installed, run the following commands to start and verify that your
development cluster is running:

```shell
minikube start
```

and

```shell
minikube status
```

## Local Container Registry (Optional)

[Docker Article](https://www.docker.com/blog/how-to-use-your-own-registry-2/)

```yaml
# docker-compose.yaml
version: "3"

services:
  registry:
    image: registry:2
    container_name: kubernetes-local-registry
    ports:
      - "5000:5000"
    environment:
      REGISTRY_STORAGE_FILESYSTEM_ROOTDIRECTORY: /data
    volumes:
      - ./data:/data
```
